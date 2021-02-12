# Chaos Client

Simple app used in chaos testing as a potential target. It can ping specified server at a fixed rate using /counter endpoint.

- [Chaos Client](#chaos-client)
  - [Command line arguments](#command-line-arguments)
  - [Development](#development)

## Command line arguments

Service supports several command line arguments set (example values are provided in parentheses):

- `host` — server to ping (`server:80`);
- `verb` — request verb to use (`get`); `get` if not specified;
- `rate` — requests per second (`1`); 0 to send 1 request and close.

## Development

To build project:

```shell
go build ./...
```
