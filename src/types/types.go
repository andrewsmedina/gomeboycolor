package types

import "fmt"

type Register byte
type Word uint16

func (w Word) String() string {
	var zeroes string
	switch {
	case w < 0x0010:
		zeroes += "000"
	case w < 0x0100:
		zeroes += "00"
	case w < 0x1000:
		zeroes += "0"
	}
	return fmt.Sprintf("0x%s%X", zeroes, uint16(w))
}

type Words []Word

func (w Words) Len() int {
	return len(w)
}

func (w Words) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func (w Words) Less(i, j int) bool {
	return w[i] < w[j]
}