package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

type Duration struct {
	Years  int
	Months int
}

func (d Duration) TotalMonths() int {
	return d.Years*12 + d.Months
}

var durationFormat = regexp.MustCompile(`(?i)^(?:(\d+)\s*(years?|y)(?:\s*(\d+)\s*(months?|m))?|(\d+)\s*(months?|m))$`)

func ParseDuration(str string) (Duration, error) {
	var years, months int
	var err error
	match := durationFormat.FindStringSubmatch(str)
	if match == nil {
		return Duration{}, fmt.Errorf("invalid string format")
	}
	if match[1] != "" {
		years, err = strconv.Atoi(match[1])
		if err != nil {
			return Duration{}, err
		}
	}
	if match[3] != "" {
		months, err = strconv.Atoi(match[3])
		if err != nil {
			return Duration{}, err
		}
	}
	if match[5] != "" {
		months, err = strconv.Atoi(match[5])
		if err != nil {
			return Duration{}, err
		}
	}
	return Duration{years, months}, nil
}

const floatPrecision = 128

func (d Duration) ToYearsDecimal() *big.Float {
	var iyears = big.NewInt(int64(d.Years))
	var imonths = big.NewInt(int64(d.Months))
	var years = new(big.Float).SetPrec(floatPrecision).SetInt(iyears)
	var months = new(big.Float).SetPrec(floatPrecision).SetInt(imonths)
	var monthsInYears = months.Quo(months, big.NewFloat(12.0))
	return years.Add(years, monthsInYears)
}

func SumDurations(ds []Duration) Duration {
	var years, months int
	for _, d := range ds {
		years += d.Years
		months += d.Months
	}
	years += months / 12
	months %= 12
	return Duration{years, months}
}

func (d Duration) FormatOutput() string {
	return fmt.Sprintf("%dy%dm", d.Years, d.Months)
}

func run(args []string, w io.Writer) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: timesum <file>")
	}
	path, err := filepath.Abs(args[0])
	if err != nil {
		return fmt.Errorf("could not read file %s: %v", path, err)
	}

	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("could not read file %s: %v", path, err)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatalf("error closing file: %v", err)
		}
	}()

	var durations []Duration

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		duration, err := ParseDuration(scanner.Text())
		if err != nil {
			return err
		}
		durations = append(durations, duration)
	}

	totalTime := SumDurations(durations)

	if err := scanner.Err(); err != nil {
		return err
	}

	_, err = fmt.Fprintln(w, totalTime.FormatOutput())
	if err != nil {
		return err
	}
	return nil
}

func main() {
	args := os.Args[1:]
	if err := run(args, os.Stdout); err != nil {
		log.Fatal(err)
	}
}
