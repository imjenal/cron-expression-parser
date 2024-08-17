# Cron Expression Parser

`cron-expression-parser` is a command-line utility written in Go that parses a cron string and expands each field to show the times at which it will run. It follows the standard cron format with five time fields (minute, hour, day of month, month, and day of week) plus a command.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Testing](#testing)
- [Cron Expression Format](#cron-expression-format)
- [Project Structure](#project-structure)

## Installation

To use Cron Expression Parser, you need to have Go installed on your system.

1. Clone the repository:

   ```shell
   git clone https://github.com/imjenal/cron-expression-parser.git
   cd cron-expression-parser
   ```
2. Build the application:

    ```bash 
    go build -o cron-expression-parser cmd/main.go
    ```

3. Run the application:
    ```bash 
    ./cron-expression-parser "<cron_expression>"
    ```
   Replace the example cron expression with your own.

## Usage

Cron Expression Parser accepts a single argument, which is the cron expression you want to parse. It then expands each field and displays the resulting schedule.

```bash 
./cron-expression-parser "*/15 3 1,5 * 1-4 /usr/bin/find"
 ```

This command will output the expanded cron fields, such as minute, hour, day of the month, month, and day of the week, along with the provided command.

## Testing

### Running Tests

To run the tests for this project, follow these steps:

- Run all tests
```bash
go test ./...
```

- Run all tests in a package
```bash
go test ./internal/fields
```

### Test Coverage
```bash
go test -cover ./...
```
This will display the percentage of code covered by the tests.

## Cron Expression Format

Cron Expression Parser follows the standard cron format with five fields:

1. Minute (0 - 59)
2. Hour (0 - 23)
3. Day of the Month (1 - 31)
4. Month (1 - 12)
5. Day of the Week (0 - 6, where Sunday is 0 and Saturday is 6)

You can use `*` to indicate all possible values in a field. Additionally, you can specify step values like */15 to indicate every 15 minutes or ranges like 1-5 to specify a range of values.

### Supported features:
- Expressions must consist of five "parts" followed by a command.
- Asterisk (`*`)
- Basic slashes, for some time interval `x` (e.g. `*/15`)
    - note that this includes the smallest value, and you are expected to provide intervals that evenly divide
- Basic commas (e.g. `1,2`)
- Hyphens (e.g. `1-7`)

### Unsupported features:
- Other forms of slashes
- [special time strings](https://en.wikipedia.org/wiki/Cron#Nonstandard_predefined_scheduling_definitions)
- Words in place of Month or Day of week
- Question marks

## Project Structure

- `cmd/`:
    - `main.go`: Entry point for the CLI application.

- `internal/`:
    - `config/`: Configuration settings (`config.go`).
    - `fields/`: Field types for cron parsing (e.g., `minute_field.go`).
    - `utils/`: Shared utility functions (`utils.go`).
    - `common/`: interface `field.go`
    - Core logic files: `parser.go`, `print.go`.

- `go.mod` and `go.sum`: Go module and dependency management files.