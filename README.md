# Hex Tool

A command line tool that converts number between different number systems.

## Installation

Make sure your go version is greater than `1.13`. 
And make sure Go Module is on, by setting environment variable:

```shell
$ export GO111MODULE=on
```

And simply run:

```shell
$ go install github.com/gg-tools/hex
```

Make sure your `PATH` includes the `$GOPATH/bin` directory so commands can be easily used:

`export PATH=$PATH:$GOPATH/bin`

## Usage


```shell
# convert to binary system
hex -b 1988

# convert to octal system
hex -o 1988

# convert to hexadecimal system
hex -H 1988

# convert number from binary system to hexadecimal system
hex 0b110

# convert number from octal system to hexadecimal system
hex 06

# convert number from hexadecimal system to hexadecimal system
hex -hex 0x1e
hex -H 0x1e
```

