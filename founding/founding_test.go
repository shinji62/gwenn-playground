package founding_test

import (
	. "github.com/shinji62/gwenn-playground/founding"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Founding", func() {
	Describe("With the right INT array input", func() {
		It("Should give expected result when array is not empty", func() {
			test := []int{4, 1, 3, 4, 2, 3, 5, 8, 1, 9, 8, 6, 4, 6}
			result := []int{2, 3, 5, 8}
			p, err := NewConsecutiveFounds(test)
			lenght, index := p.LonguestFoundings()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(test[index:(index + lenght)]).Should(Equal(result))

		})
		It("Should give expected result when result is at the end", func() {
			test := []int{4, 1, 3, 4, 2, 3, 5, 8, 1, 9, 8, 6, 4, 6, 2, 3, 5, 8, 9}
			result := []int{2, 3, 5, 8, 9}
			p, err := NewConsecutiveFounds(test)
			lenght, index := p.LonguestFoundings()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(test[index:(index + lenght)]).Should(Equal(result))

		})
		It("Should give expected result when array is empty", func() {
			test := []int{}
			result := []int{}
			p, err := NewConsecutiveFounds(test)
			lenght, index := p.LonguestFoundings()
			Expect(index).To(Equal(0))
			Expect(lenght).To(Equal(0))
			Expect(err).ShouldNot(HaveOccurred())
			Expect(test[index:(index + lenght)]).Should(Equal(result))

		})
	})

	Describe("With the right byte String array input", func() {
		It("Should give expected result when array is not empty", func() {
			test := []byte{'n', 'b', 'w', 'x', 'z', 'q', 'd', 'n', 'r', 'e', 'v', 'k', 'd', 'n'}
			result := []byte{'b', 'w', 'x', 'z'}
			p, err := NewConsecutiveFounds(test)
			lenght, index := p.LonguestFoundings()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(test[index:(index + lenght)]).Should(Equal(result))

		})
		It("Should give expected result when array is empty", func() {
			test := []byte{}
			result := []byte{}
			p, err := NewConsecutiveFounds(test)
			lenght, index := p.LonguestFoundings()
			Expect(index).To(Equal(0))
			Expect(lenght).To(Equal(0))
			Expect(err).ShouldNot(HaveOccurred())
			Expect(test[index:(index + lenght)]).Should(Equal(result))

		})

	})
	Describe("With the right byte 0 or 1 array input", func() {
		It("Should give expected result when array is not empty", func() {
			test := "01100100001001"
			result := "0000"
			p, err := NewConsecutiveFounds(test)
			lenght, index := p.LonguestFoundings()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(test[index:(index + lenght)]).Should(Equal(result))

		})
		It("Should give expected result when array is empty", func() {
			test := ""
			result := ""
			p, err := NewConsecutiveFounds(test)
			lenght, index := p.LonguestFoundings()
			Expect(index).To(Equal(0))
			Expect(lenght).To(Equal(0))
			Expect(err).ShouldNot(HaveOccurred())
			Expect(test[index:(index + lenght)]).Should(Equal(result))

		})

	})
	Describe("Given a string", func() {
		It("Should give expected result when string is not empty and at the end", func() {
			test_r := "nbwxzqdnrevkdnabcdefghi"
			result := "abcdefghi"
			p, err := NewConsecutiveFounds(test_r)
			lenght, index := p.LonguestFoundings()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(test_r[index:(index + lenght)]).Should(Equal(result))

		})
		It("Should give expected result when string is not empty and at the beginning", func() {
			test_r := "abcdefghinbwxzqdnrevkdnabcdefghi"
			result := "abcdefghin"
			p, err := NewConsecutiveFounds(test_r)
			lenght, index := p.LonguestFoundings()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(test_r[index:(index + lenght)]).Should(Equal(result))

		})
		It("Should give expected result when empty", func() {
			test := ""
			result := ""
			p, err := NewConsecutiveFounds(test)
			lenght, index := p.LonguestFoundings()
			Expect(index).To(Equal(0))
			Expect(lenght).To(Equal(0))
			Expect(err).ShouldNot(HaveOccurred())
			Expect(test[index:(index + lenght)]).Should(Equal(result))

		})
		It("Should give expected result when one elem", func() {
			test := "b"
			result := "b"
			p, err := NewConsecutiveFounds(test)
			lenght, index := p.LonguestFoundings()
			Expect(index).To(Equal(0))
			Expect(lenght).To(Equal(1))
			Expect(err).ShouldNot(HaveOccurred())
			Expect(test[index:(index + lenght)]).Should(Equal(result))

		})

	})

})
