# Advent of code 2021 in Go

An opportunity to learn the Go language during this year's Advent of code challenge.

## Logbook

### Setting everything up

Basically following this tutorial: https://spf13.com/presentation/building-an-awesome-cli-app-in-go-oscon/

Install Go package: `pacman -S go`

Create base Go path: `mkdir ~/go`

Add Go's bin path to `~/.bashrc`:

```
export GOPATH="$HOME/go"
export PATH="$PATH:$GOPATH/bin"
```

Install Cobra (CLI application framework and library): `go get -v github.com/spf13/cobra/cobra`

```
mkdir -p ~/go/src/github.com/sirzerator/advent2021
cobra init github.com/sirzerator/advent2021 -a "Émile Plourde-Lavoie" --pkg-name github.com/sirzerator/advent2021
cd ~/go/src/github.com/sirzerator/advent2021
go mod init
go get github.com/spf13/cobra
go get github.com/spf13/viper
```
