package utils

import (
	"regexp"
	"strconv"
)

func ByteTobyteArray(inputByte []byte) [][]byte {
	slice := make([][]byte, len(inputByte))
	for i, v := range inputByte {
		slice[i] = []byte{v}
	}
	return slice
}

func StringsToByteSlice(inputString string) [][]byte {
	slice := make([][]byte, len(inputString))
	for i, v := range inputString {
		slice[i] = []byte(string(v))
	}
	return slice
}

func IntSliceToByteSlice(input []int) [][]byte {
	slice := make([][]byte, cap(input), len(input))
	for i, v := range input {
		slice[i] = []byte(strconv.Itoa(v))
	}
	return slice
}

func IsBinaryNumber(s string) bool {
	binaryPattern := regexp.MustCompile(`^[0-1]+$`)
	return binaryPattern.MatchString(s)
}
