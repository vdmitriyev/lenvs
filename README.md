## About

Loads environmental variables from `.env` file. Runs a sub-process within an operating system terminal, which will get environment variables from the provided `.env` file.

## Install

**Dependencies:** `go` must be installed first. Latest version can be obtained and installed from [https://go.dev/doc/install](https://go.dev/doc/install).

Install
```
go install github.com/vdmitriyev/lenvs@latest
```

## Usage

Command should load environment variables from a `.env` file:
```
lenvs .env
```

## License

[MIT](LICENSE)