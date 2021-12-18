package solutions

import (
	"fmt"
	"math"
	"strconv"
)

type BinaryNode struct {
	value  int
	pair   bool
	left   *BinaryNode
	right  *BinaryNode
	parent *BinaryNode
}

func binaryNumberMiddle(number string) int {
	open_param := 0

	for i, char := range number {
		if char == '[' {
			open_param++
		} else if char == ']' {
			open_param--
		} else if char == ',' && open_param == 1 {
			return i
		}
	}
	return -1
}

func readBinaryNumber(number string) *BinaryNode {
	node := &BinaryNode{}
	if number[0] != '[' {
		node.value, _ = strconv.Atoi(number)
		node.pair = false
	} else {
		node.pair = true
		middle := binaryNumberMiddle(number)
		left := readBinaryNumber(number[1:middle])
		node.left = left
		left.parent = node
		right := readBinaryNumber(number[middle+1 : len(number)-1])
		node.right = right
		right.parent = node
	}

	return node
}

func readBinaryNumbers() []*BinaryNode {
	numbers := make([]*BinaryNode, 0)
	for {
		var number_ln string
		_, err := fmt.Scanln(&number_ln)
		if err != nil {
			break
		}
		numbers = append(numbers, readBinaryNumber(number_ln))
	}

	return numbers
}

func addBinaryNumbers(lhs *BinaryNode, rhs *BinaryNode) *BinaryNode {
	node := &BinaryNode{0, true, lhs, rhs, nil}
	lhs.parent = node
	rhs.parent = node
	node.Reduce()
	return node
}

func (node *BinaryNode) Reduce() {
	for {
		exploded := node.explode(0)
		if exploded {
			continue
		}
		splitted := node.split()
		if !splitted {
			break
		}
	}
}

func (node *BinaryNode) ToString() string {
	if node.pair {
		return "[" + node.left.ToString() + "," + node.right.ToString() + "]"
	}
	return strconv.Itoa(node.value)

}

func (node *BinaryNode) explodeNode() {
	curr := node.parent
	prev := node
	for curr != nil && curr.right != prev {
		prev = curr
		curr = curr.parent
	}

	if curr != nil {
		curr = curr.left
		for curr.pair {
			curr = curr.right
		}
		curr.value += node.left.value
	}
	curr = node.parent
	prev = node
	for curr != nil && curr.left != prev {
		prev = curr
		curr = curr.parent
	}
	if curr != nil {
		curr = curr.right
		for curr.pair {
			curr = curr.left
		}
		curr.value += node.right.value
	}
}

func (node *BinaryNode) explode(level int) bool {
	if level != 4 {
		exploded := node.left != nil && node.left.explode(level+1)
		if exploded {
			return true
		}
		exploded = node.right != nil && node.right.explode(level+1)
		return exploded
	}

	if !node.pair {
		return false
	}
	node.explodeNode()
	node.value = 0
	node.pair = false
	node.left = nil
	node.right = nil
	return true
}

func (node *BinaryNode) split() bool {
	if node.pair {
		splitted := node.left != nil && node.left.split()
		if splitted {
			return true
		}
		splitted = node.right != nil && node.right.split()
		return splitted
	}

	if node.value < 10 {
		return false
	}

	node.pair = true
	node.left = &BinaryNode{int(math.Floor(float64(node.value) / 2.0)), false, nil, nil, node}
	node.right = &BinaryNode{int(math.Ceil(float64(node.value) / 2.0)), false, nil, nil, node}
	return true
}

func addAllNumbers(numbers []*BinaryNode) *BinaryNode {
	number := numbers[0]
	for _, other := range numbers[1:] {
		number = addBinaryNumbers(number, other)
	}
	return number
}

func (node *BinaryNode) Magnitude() int {
	if node.pair {
		return node.left.Magnitude()*3 + node.right.Magnitude()*2
	}
	return node.value
}

func (orig *BinaryNode) Copy() *BinaryNode {
	node := &BinaryNode{}
	if !orig.pair {
		node.value = orig.value
		node.pair = false
	} else {
		node.pair = true
		left := orig.left.Copy()
		node.left = left
		left.parent = node
		right := orig.right.Copy()
		node.right = right
		right.parent = node
	}

	return node
}

func largestMagnitude(numbers []*BinaryNode) int {
	max := 0

	for i, number := range numbers {
		for j, other := range numbers {
			if i == j {
				continue
			}
			result := addBinaryNumbers(number.Copy(), other.Copy()).Magnitude()
			if max < result {
				max = result
			}
		}
	}
	return max

}

func SolveDay18() {
	numbers := readBinaryNumbers()
	fmt.Println("Part 2: ", largestMagnitude(numbers))
	result := addAllNumbers(numbers)
	fmt.Println("Part 1: ", result.Magnitude())

}
