# hsize
command line tool for human readable size

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gonejack/hsize)
![Build](https://github.com/gonejack/hsize/actions/workflows/go.yml/badge.svg)
[![GitHub license](https://img.shields.io/github/license/gonejack/hsize.svg?color=blue)](LICENSE)


## Installation
```
go get -u github.com/gonejack/hsize
```

## Usage

by arguments
```
> hsize 123 383764
123B
374.76KB
```

by stdin
```
> echo 19129219219129119 | hsize
16.99PB
```

## Options

#### -p precision
```
> hsize 1025
1KB
> hsize -p 5 1025
1.00097KB
```
