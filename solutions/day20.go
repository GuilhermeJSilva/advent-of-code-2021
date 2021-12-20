package solutions

import (
	"fmt"
)

func readImage() []string {
	image := make([]string, 0)
	for {
		var line string
		fmt.Scanln(&line)
		if len(line) == 0 {
			break
		}
		image = append(image, line)
	}

	return image
}

func readEnhancement() string {
	var enhancement string
	for {
		var line string
		fmt.Scanln(&line)
		if len(line) == 0 {
			break
		}
		enhancement += line
	}
	return enhancement
}

func solvePixel(image []string, enhancement string, i int, j int, infinity bool) byte {
	height := len(image)
	width := len(image[0])
	position := 0

	for i_off := i - 1; i_off <= i+1; i_off++ {
		for j_off := j - 1; j_off <= j+1; j_off++ {
			position = position << 1
			if i_off < 0 || i_off >= height || j_off < 0 || j_off >= width {
				if infinity {
					position |= 1
				}
			} else if image[i_off][j_off] == '#' {
				position |= 1
			}
		}
	}
	return enhancement[position]
}

func enhance(image []string, enhancement string, infinity bool) []string {
	new_image := make([]string, 0)
	height := len(image)
	width := len(image[0])
	for i := -1; i <= width; i++ {
		line := make([]byte, 0)
		for j := -1; j <= height; j++ {
			pixel := solvePixel(image, enhancement, i, j, infinity)
			line = append(line, pixel)
		}
		new_image = append(new_image, string(line))
	}
	return new_image
}

func countLight(image []string) int {
	count := 0
	for _, line := range image {
		for _, value := range line {
			if value == '#' {
				count++
			}
		}
	}

	return count
}

func enhanceN(image []string, enhancement string, times int) []string {
	alternating := enhancement[0] == '#'
	infinity := false
	new_image := enhance(image, enhancement, infinity)
	infinity = alternating
	for i := 1; i < times; i++ {
		new_image = enhance(new_image, enhancement, infinity)
		if alternating {
			infinity = !infinity
		}
	}
	return new_image
}
func SolveDay20() {
	enhancement := readEnhancement()
	image := readImage()
	image2 := enhanceN(image, enhancement, 2)
	fmt.Println(countLight(image2))
	image50 := enhanceN(image, enhancement, 50)
	fmt.Println(countLight(image50))
}
