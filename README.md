# hsize [![GitHub license](https://img.shields.io/github/license/gonejack/hsize.svg?color=blue)](LICENSE.md)
command line tool for human readable size

## Installation
```
go get -u github.com/gonejack/hsize
```

## Usage

by arguments
```
> hsize 123 45678
123 => 123B
45678 => 44.61KB
```

by stdin
```
> echo 19129219219129119 | hsize
19129219219129119 => 16.99PB
```
