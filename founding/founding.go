package founding

import (
	"bytes"

	"github.com/shinji62/gwenn-playground/utils"
)

type ConsecutiveFounds struct {
	ListConsectiveFound []ConsecutiveFound
}

type ConsecutiveFound struct {
	Lenght int
	Index  int
}

func NewConsecutiveFounds(genericSlice interface{}) (*ConsecutiveFounds, error) {
	p := ReturnLonguestFounds(genericSlice)
	return &ConsecutiveFounds{ListConsectiveFound: p.ListConsectiveFound}, nil
}

func FindAllLonguestBytes(slide [][]byte) *ConsecutiveFounds {
	p := &ConsecutiveFounds{}
	c := ConsecutiveFound{Index: 0, Lenght: 1}

	//Edge case if only one element
	if len(slide) <= 1 {
		p.ListConsectiveFound = append(p.ListConsectiveFound,
			ConsecutiveFound{Index: 0, Lenght: len(slide)})
		return p
	}

	for x := 1; x < len(slide); x++ {
		switch bytes.Compare(slide[x-1], slide[x]) {
		case -1:
			c.Lenght++
		case +1:
			p.ListConsectiveFound = append(p.ListConsectiveFound, c)
			c = ConsecutiveFound{Index: x, Lenght: 1}
		}
		//Edge Case if we are at the end of the slide
		if x == len(slide)-1 {
			p.ListConsectiveFound = append(p.ListConsectiveFound, c)
		}

	}
	return p
}

func FindAllLonguestBinBytes(slide [][]byte) *ConsecutiveFounds {
	p := &ConsecutiveFounds{}
	c := ConsecutiveFound{Index: 0, Lenght: 1}

	//Edge case if only one element
	if len(slide) <= 1 {
		p.ListConsectiveFound = append(p.ListConsectiveFound,
			ConsecutiveFound{Index: 0, Lenght: len(slide)})
		return p
	}

	for x := 1; x < len(slide); x++ {
		switch bytes.Compare(slide[x-1], slide[x]) {
		case 0:
			c.Lenght++
		default:
			p.ListConsectiveFound = append(p.ListConsectiveFound, c)
			c = ConsecutiveFound{Index: x, Lenght: 1}
		}
		//Edge Case if we are at the end of the slide
		if x == len(slide)-1 {
			p.ListConsectiveFound = append(p.ListConsectiveFound, c)
		}

	}
	return p
}

func FindAllLonguestBinaryInt(slide []int) *ConsecutiveFounds {
	p := &ConsecutiveFounds{}
	c := ConsecutiveFound{Index: 0, Lenght: 1}
	for x := 1; x <= len(slide); x++ {
		switch {
		case slide[x-1] == slide[x]:
			c.Lenght++
		case slide[x-1] != slide[x]:
			p.ListConsectiveFound = append(p.ListConsectiveFound, c)
			c = ConsecutiveFound{Index: x, Lenght: 1}
		}

	}
	return p
}

//Return the longuest one that we found
func (cs *ConsecutiveFounds) LonguestFoundings() (int, int) {
	lenght := 0
	index := 0
	for _, v := range cs.ListConsectiveFound {
		if v.Lenght > lenght {
			lenght = v.Lenght
			index = v.Index
		}
	}
	return lenght, index
}

// the smallestConsecutive we found
func (cs *ConsecutiveFounds) SmallestFoundings() (int, int) {
	lenght := 0
	index := 0
	for _, v := range cs.ListConsectiveFound {
		if v.Lenght < lenght {
			lenght = v.Lenght
			index = v.Index
		}
	}
	return lenght, index
}

// Support multiple Input
// we use reflection to get the right type and create a array of byte array
// that we use to parse after wards.

func ReturnLonguestFounds(genericSlice interface{}) *ConsecutiveFounds {
	p := &ConsecutiveFounds{}
	switch genericSlice.(type) {
	case []int:
		p = FindAllLonguestBytes(utils.IntSliceToByteSlice(genericSlice.([]int)))
	case []byte:
		p = FindAllLonguestBytes(utils.ByteTobyteArray(genericSlice.([]byte)))
	case string:
		if utils.IsBinaryNumber(genericSlice.(string)) {
			p = FindAllLonguestBinBytes(utils.StringsToByteSlice(genericSlice.(string)))
		} else {
			p = FindAllLonguestBytes(utils.StringsToByteSlice(genericSlice.(string)))

		}
	}
	return p
}
