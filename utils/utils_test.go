package utils_test

import (
	"strconv"

	. "github.com/shinji62/gwenn-playground/utils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Utils", func() {
	Describe("Giving a string", func() {
		It("Should return expected byte array", func() {
			arr := StringsToByteSlice("m")
			result := make([][]byte, 1, 1)
			result[0] = []byte("m")
			Expect(arr).To(Equal(result))
		})
	})
	Describe("Giving a Int Array", func() {
		It("Should return expected byte array", func() {
			arr := IntSliceToByteSlice([]int{413})
			result := make([][]byte, 1, 1)
			result[0] = []byte(strconv.Itoa(413))
			Expect(arr).To(Equal(result))
		})
	})

	Describe("Giving String", func() {
		Context("with valid binary number", func() {
			It("Should detect a 0-1 number", func() {
				test := "010101110111010101"
				Expect(IsBinaryNumber(test)).To(BeTrue())
			})
		})
		Context("with invalid binary number", func() {
			It("Should not detect a 0-1 number", func() {
				test := "d010d1011f10111010101"
				Expect(IsBinaryNumber(test)).To(BeFalse())
			})
		})

	})
})
