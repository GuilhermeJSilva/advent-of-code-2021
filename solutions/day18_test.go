package solutions

import (
	"fmt"
	"strings"
	"testing"
)

func TestExplosion(t *testing.T) {
	numbers := []string{
		"[[[[[9,8],1],2],3],4]|[[[[0,9],2],3],4]",
		"[7,[6,[5,[4,[3,2]]]]]|[7,[6,[5,[7,0]]]]",
		"[[6,[5,[4,[3,2]]]],1]|[[6,[5,[7,0]]],3]",
		"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]|[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
		"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]|[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
		"[[[[7,7],[7,8]],[[9,5],[8,0]]],[[[9,0],[0,[5,5]]],[8,[9,0]]]]|[[[[7,7],[7,8]],[[9,5],[8,0]]],[[[9,0],[5,0]],[13,[9,0]]]]",
	}
	for i, number := range numbers {
		split := strings.Split(number, "|")
		bn := readBinaryNumber(split[0])
		bn.explode(0)
		if bn.ToString() != split[1] {
			t.Errorf("Explosion error int test %v: %v\nResult:   %v\nExpected: %v\n", i, strings.Split(numbers[i], "|")[0], bn.ToString(), split[1])
		}

	}
}

func TestAdd(t *testing.T) {
	tests := []string{
		"[[[[4,3],4],4],[7,[[8,4],9]]]|[1,1]|[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]|[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]|[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]",
		"[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]|[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]|[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]]",
		"[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]]|[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]|[[[[7,0],[7,7]],[[7,7],[7,8]]],[[[7,7],[8,8]],[[7,7],[8,7]]]]",
		"[[[[7,0],[7,7]],[[7,7],[7,8]]],[[[7,7],[8,8]],[[7,7],[8,7]]]]|[7,[5,[[3,8],[1,4]]]]|[[[[7,7],[7,8]],[[9,5],[8,7]]],[[[6,8],[0,8]],[[9,9],[9,0]]]]",
	}

	for i, test := range tests {
		split := strings.Split(test, "|")
		bn1 := readBinaryNumber(split[0])
		bn2 := readBinaryNumber(split[1])
		result := addBinaryNumbers(bn1, bn2)
		if result.ToString() != split[2] {
			t.Errorf("Add error %v:\nResult:   %v\nExpected: %v", i, result.ToString(), split[2])
		}
	}
}

func TestCopy(t *testing.T) {

	tests := []string{
		"[[[[4,3],4],4],[7,[[8,4],9]]]",
	}
	for i, test := range tests {
		bn := readBinaryNumber(test)
		copy := bn.Copy()
		fmt.Println(test)
		fmt.Println(copy.ToString())
		if bn.ToString() != copy.ToString() {
			t.Errorf("Copy error %v:\nResult:   %v\nExpected: %v", i, copy.ToString(), bn.ToString())
		}

	}
}
