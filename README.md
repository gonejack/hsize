# human readable size


## installation
```
go get github.com/gonejack/hsize
```

## usage

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
