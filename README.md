# Vedran

> Polkadot chain loadbalancer.

## Installation
You can download already built binaries for your platform from `builds` folder or get `vedran` golang package and build it locally. Find detailed instructions below.

### Get `vedran` package
1. Install [Golang](https://golang.org/doc/install) **1.13 or greater**
2. Run the command below
```
go get github.com/NodeFactoryIo/vedran
```
3. Run vedran from your Go bin directory. For linux systems it will likely be:
```
~/go/bin/vedran
```
Note that if you need to do this, you probably want to add your Go bin directory to your $PATH to make things easier!

## Usage

```
$ vedran -h

vedran is a command line interface for ....

Usage:
  vedran [command]

Available Commands:
  version     Show the current version of Vedran deamon app

Use "vedran [command] --help" for more information about a command.
```

## Development
Run daemon app with `go run main.go [command]`.

Expected name of the configuration file depends on `ENV` variable. For example, if you run a daemon app with `ENV=test go run main.go start`, expected config file name is `config-test.yaml`

## License

This project is licensed under Apache 2.0:
- Apache License, Version 2.0, ([LICENSE-APACHE](http://www.apache.org/licenses/LICENSE-2.0))