package solutions

import (
	"encoding/hex"
	"fmt"
	"math"
)

type HexaReader struct {
	hexa    []byte
	digit_i uint
}

type Literal struct {
	version uint
	type_id uint
	value   uint
}

type Operator struct {
	version    uint
	type_id    uint
	length_id  uint
	length     uint
	subpackets []interface{}
}

func readLargeHexadecimal() HexaReader {
	var hexa string
	fmt.Scanln(&hexa)
	bytes, _ := hex.DecodeString(hexa)
	return HexaReader{bytes, 0}
}

func getBits(value byte, bit_start uint, bit_end uint) byte {
	extracted := value >> (7 - bit_end)
	length := bit_end - bit_start + 1
	extracted &= ^(0xff << length)
	return extracted
}

func (hexa *HexaReader) readBits(n_bits uint) uint {
	var value uint = 0
	size_to_read := uint(n_bits)
	for bits_read := uint(0); bits_read < n_bits; {
		letter_i := hexa.digit_i / 8
		bit_i := hexa.digit_i % 8
		size_to_read = n_bits - bits_read
		if size_to_read > 8-bit_i {
			size_to_read = 8 - bit_i
		}
		bit_value := getBits(hexa.hexa[letter_i], bit_i, bit_i+size_to_read-1)
		value = (value << size_to_read) | uint(bit_value)
		hexa.digit_i += size_to_read
		bits_read += size_to_read
	}
	return value
}

func (hexa *HexaReader) readLiteral(version uint) (Literal, uint) {
	packet_size := uint(6)
	last := uint(1)
	var value uint = 0
	for last != 0 {
		value = value << 4
		next := hexa.readBits(5)
		last = next >> 4
		value += next & 0b1111
		packet_size += 5
	}
	return Literal{version, 4, value}, packet_size
}

func (hexa *HexaReader) readPacket() (interface{}, uint) {
	version := hexa.readBits(3)
	type_id := hexa.readBits(3)
	if type_id == 4 {
		return hexa.readLiteral(version)
	}
	return hexa.readOperator(version, type_id)
}

func (hexa *HexaReader) readOperator(version uint, type_id uint) (Operator, uint) {
	packet_size := uint(7)
	length_id := hexa.readBits(1)
	subpackets := make([]interface{}, 0)
	lenght_to_read := uint(15)
	if length_id == 1 {
		lenght_to_read = 11
	}
	packet_size += lenght_to_read
	size := hexa.readBits(lenght_to_read)
	for curr_size := size; curr_size > 0; {
		subpacket, subpacket_size := hexa.readPacket()
		subpackets = append(subpackets, subpacket)
		packet_size += subpacket_size
		if length_id == 1 {
			curr_size--
		} else {
			curr_size -= subpacket_size
		}
	}
	return Operator{version, type_id, length_id, size, subpackets}, packet_size
}

func (op Operator) sumVersions() uint {
	sum := op.version
	for _, sub := range op.subpackets {
		switch sub := sub.(type) {
		case Operator:
			sum += sub.sumVersions()
		case Literal:
			sum += sub.version
		}
	}
	return sum
}

func sumMembers(values []uint) uint {
	sum := uint(0)
	for _, value := range values {
		sum += value
	}
	return sum
}

func multMembers(values []uint) uint {
	sum := uint(1)
	for _, value := range values {
		sum *= value
	}
	return sum
}

func minMembers(values []uint) uint {
	var minimum uint = math.MaxUint
	for _, value := range values {
		if minimum > value {
			minimum = value
		}
	}
	return minimum
}

func maxMembers(values []uint) uint {
	var maximum uint = 0
	for _, value := range values {
		if maximum < value {
			maximum = value
		}
	}
	return maximum
}

func greaterMembers(values []uint) uint {
	if values[0] > values[1] {
		return 1
	}
	return 0
}
func lesserMembers(values []uint) uint {
	if values[0] < values[1] {
		return 1
	}
	return 0
}
func equalMembers(values []uint) uint {
	if values[0] == values[1] {
		return 1
	}
	return 0
}

var applyFunctions []func([]uint) uint = []func([]uint) uint{
	sumMembers,
	multMembers,
	minMembers,
	maxMembers,
	nil,
	greaterMembers,
	lesserMembers,
	equalMembers,
}

func (op Operator) result() uint {
	results := make([]uint, len(op.subpackets))
	for i, subpacket := range op.subpackets {
		results[i] = result(subpacket)
	}
	return applyFunctions[op.type_id](results)
}

func result(cmd interface{}) uint {
	switch cmd := cmd.(type) {
	case Operator:
		return cmd.result()
	case Literal:
		return cmd.value
	}
	return 0
}

func SolveDay16() {
	hexa := readLargeHexadecimal()
	version := hexa.readBits(3)
	type_id := hexa.readBits(3)
	operator, _ := hexa.readOperator(version, type_id)
	fmt.Println(operator.sumVersions())
	fmt.Println(operator.result())
}
