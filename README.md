# weather-cli

A simple weather forecasting command line tool by go.

## Usage

**Native**

```bash
go build -o weather-cli utils.go types.go main.go && ./weather-cli
```

**Use Lib**

```bash
go build -o weather-cli utils.go types.go cli_main.go && ./weather-cli
```

## Package

```bash
go build -ldflags "-s -w" -o weather-cli utils.go types.go cli_main.go && upx ./weather-cli
```

**Linux Arch**

```bash
GOOS=linux GOARCH=amd64 go build ...
```

**Windows Arch**

```bash
GOOS=windows GOARCH=amd64 go build ...
```

**MacOSX Arch**

```bash
GOOS=darwin GOARCH=amd64 go build ...
```

## Compression

```bash
go build -ldflags "-s -w" -o weather-cli utils.go types.go cli_main.go && upx ./weather-cli
```
