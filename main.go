package main

import (
	"flag"
	"fmt"

	"github.com/shinji62/gwenn-playground/founding"
)

var testSuite []interface{}

func main() {
	flag.Parse()
	founder := &founding.ConsecutiveFounds{}
	var i, l int
	var err error
	fmt.Println("-------------- Test From String in Command Line --------------")
	for _, commandLineArg := range flag.Args() {
		founder, err = founding.NewConsecutiveFounds(commandLineArg)
		if err != nil {
			fmt.Println(err)
		}
		l, i = founder.LonguestFoundings()
		fmt.Printf("\nTest number %v\n", commandLineArg)
		fmt.Printf("Lenght: %d, Index: %d ", l, i)
		fmt.Printf("Result: %v\n", commandLineArg[i:(i+l)])

	}

	fmt.Println("-------------- Test in source Code --------------")
	test := []int{4, 1, 3, 4, 2, 3, 5, 8, 1, 9, 8, 6, 4, 6}
	testStringByte := []byte{'n', 'b', 'w', 'x', 'z', 'q', 'd', 'n', 'r', 'e', 'v', 'k', 'd', 'n'}
	testString := "nbwxzqdnrevkdn"
	testOther := "01100100001001"

	founder, err = founding.NewConsecutiveFounds(test)
	if err != nil {
		fmt.Println(err)
	}
	l, i = founder.LonguestFoundings()
	fmt.Printf("\nTest number %v\n", test)
	fmt.Printf("Lenght: %d, Index: %d ", l, i)
	fmt.Printf("Result: %v\n", test[i:(i+l)])

	founder, err = founding.NewConsecutiveFounds(testStringByte)
	if err != nil {
		fmt.Println(err)
	}
	l, i = founder.LonguestFoundings()
	fmt.Printf("\nTest String Byte %s\n", testStringByte)
	fmt.Printf("Lenght: %d, Index: %d ", l, i)
	fmt.Printf("Result:  %s\n", testStringByte[i:(i+l)])

	founder, err = founding.NewConsecutiveFounds(testString)
	if err != nil {
		fmt.Println(err)
	}
	l, i = founder.LonguestFoundings()
	fmt.Printf("\nTest String from string %s\n", testString)
	fmt.Printf("Lenght: %d, Index: %d ", l, i)
	fmt.Printf("Result:  %s\n", testString[i:(i+l)])

	founder, err = founding.NewConsecutiveFounds(testOther)
	if err != nil {
		fmt.Println(err)
	}
	l, i = founder.LonguestFoundings()
	fmt.Printf("\nTest Binary 0 or 1 from string %s\n", testOther)
	fmt.Printf("Lenght: %d, Index: %d ", l, i)
	fmt.Printf("Result:  %s\n", testOther[i:(i+l)])
}
