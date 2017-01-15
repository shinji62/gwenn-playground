## Description
Some test with Longuest Consecutive

### Requirement
**Go** >= 1.7.X

Except for the test we use the std library
For the test we use [gingko](https://onsi.github.io/ginkgo/) and [gomega](https://onsi.github.io/gomega/)

```
$ go get github.com/onsi/ginkgo/ginkgo
$ go get github.com/onsi/gomega
```
### Compile
This will compile the binary
```
make compile
```

## Running
Use your corresponding binary
```shell
cd ./dist/....
```
#### Running included test in the code
```shell
./gwenn-playground
```
#### Running with command line Argument
You can give some string to be parsed using command line.
```
./gwenn-playground arg1 arg2 ...
```
example
```shell
./gwenn-playground 01010101000110000101 fkdsjfksldjfkljsdjlfk
```

Should give
```
-------------- Test From String in Command Line --------------

Test number 01010101000110000101
Lenght: 4, Index: 13 Result: 0000

Test number fkdsjfksldjfkljsdjlfk
Lenght: 3, Index: 5 Result: fks
-------------- Test in source Code --------------

Test number [4 1 3 4 2 3 5 8 1 9 8 6 4 6]
Lenght: 4, Index: 4 Result: [2 3 5 8]

Test String Byte nbwxzqdnrevkdn
Lenght: 4, Index: 1 Result:  bwxz

Test String from string nbwxzqdnrevkdn
Lenght: 4, Index: 1 Result:  bwxz

Test Binary 0 or 1 from string 01100100001001
Lenght: 4, Index: 6 Result:  0000
```
### Run Unitesting
```shell
make test
```
