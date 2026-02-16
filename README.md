# openCalc

Fast, minimal, terminal-first calculator written in Go.
Type an expression, get an answer. No UI, no waiting.

## Highlights

- Zero-friction CLI usage
- Exponent support (example: `5e5`)
- Clean, small Go codebase

## Quick Start

Build the binary:

```bash
go build -o openCalc ./cmd
```

Run a calculation:

```bash
./openCalc 8+5
./openCalc 5e5
```

Optional: move it to your PATH so you can call it anywhere.

## Examples

```bash
./openCalc 12/3
./openCalc 2e8
./openCalc (3+4)*5
```

## Project Layout

```text
cmd/         CLI entrypoint and parsing
operations/  Math operations implementation
```

## Development

```bash
go run ./cmd 9*(6-2)
```

## License

See [LICENSE](LICENSE).
