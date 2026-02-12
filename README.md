# timesum

A CLI tool written in Go that sums durations (years and months) from a text file.

## Usage

```bash
timesum <file>
```

Where `<file>` is a plain text file with one duration per line.

## Input Format

Each line should contain a duration in years and/or months. The parser is case-insensitive and accepts flexible formatting:

```
1y2m
2years3months
3 Years 2 Months
4Years2Months
7months
```

## Output

The tool outputs the total duration normalized into years and months:

```bash
$ timesum durations.txt
18y0m
```

## Building

```bash
go build -o timesum .
```

## Running Tests

```bash
go test -v ./...
```

## License

MIT
