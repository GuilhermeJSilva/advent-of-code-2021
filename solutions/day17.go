package solutions

import "fmt"

type Target struct {
	x_start int
	x_end   int
	y_start int
	y_end   int
}

func readTarget() Target {
	var target Target
	fmt.Scanf("target area: x=%d..%d, y=%d..%d", &target.x_start, &target.x_end, &target.y_start, &target.y_end)
	return target
}

func searchVelocities(target Target) int {
	total := 0
	for start_vx := 1; start_vx <= target.x_end; start_vx++ {
		for start_vy := target.y_start; start_vy != -target.y_start; start_vy++ {
			x, y := 0, 0
			vx, vy := start_vx, start_vy

			for x <= target.x_end && y >= target.y_start {

				if x >= target.x_start && y <= target.y_end {
					total += 1
					break
				}

				x, y = x+vx, y+vy
				vy -= 1

				if vx > 0 {
					vx -= 1
				}
			}
		}
	}

	return total
}
func SolveDay17() {
	target := readTarget()
	fmt.Println(target.y_start * (target.y_start + 1) / 2)
	fmt.Println(searchVelocities(target))
}
