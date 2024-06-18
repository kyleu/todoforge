# Installation

## Pre-built binaries
Download any package from the [release page](https://github.com/kyleu/todoforge/releases).

### Homebrew
```shell
brew install kyleu/kyleu/todoforge
```

### deb, rpm and apk packages
Download the .deb, .rpm or .apk packages from the [release page](https://github.com/kyleu/todoforge/releases) and install them with the appropriate tools.

## Running with Docker
```shell
docker run -p 12000:12000 ghcr.io/kyleu/todoforge:latest
docker run -p 12000:12000 ghcr.io/kyleu/todoforge:latest-debug
```

## Built from source

### go install
```shell
go install github.com/kyleu/todoforge@latest
```

### Source code

If you want to contribute to the project, please follow the steps on our [contributing guide](contributing).

If you just want to build from source for whatever reason, follow these steps:

```shell
git clone https://github.com/kyleu/todoforge
cd todoforge
make build
./build/debug/todoforge --help
```

A script has been provided at `./bin/dev.sh` that will auto-reload when the project changes.

Note that you may need to run `./bin/bootstrap.sh` before building the project for the first time.
