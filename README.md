# hsize [![GitHub license](https://img.shields.io/github/license/gonejack/hsize.svg?color=blue)](LICENSE)
command line tool for human readable size

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

## Arguments

#### -p precision
```
> hsize 1025
1KB
> hsize -p 5 1025
1.00097KB
```
