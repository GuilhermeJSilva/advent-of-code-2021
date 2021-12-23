package solutions

import (
	"fmt"
)

type Point3D struct {
	x int
	y int
	z int
}

type Area3D struct {
	start Point3D
	end   Point3D
}

func readAreas() ([]Area3D, []string) {
	areas := make([]Area3D, 0)
	commands := make([]string, 0)

	for {
		var area Area3D
		var command string
		_, err := fmt.Scanf("%s x=%d..%d,y=%d..%d,z=%d..%d", &command, &area.start.x, &area.end.x, &area.start.y, &area.end.y, &area.start.z, &area.end.z)
		if err != nil {
			break
		}
		area.end.AddConstant(1)
		areas = append(areas, area)
		commands = append(commands, command)
	}

	return areas, commands
}

func (p *Point3D) AddConstant(value int) {
	p.x += value
	p.y += value
	p.z += value
}

func (area Area3D) IsEmpty() bool {
	return area.start.x >= area.end.x || area.start.y >= area.end.y || area.start.z >= area.end.z
}

func (area Area3D) Contains(other Area3D) bool {
	return other.start.x >= area.start.x && other.end.x <= area.end.x && other.start.y >= area.start.y && other.end.y <= area.end.y && other.start.z >= area.start.z && other.end.z <= area.end.z
}

func maxInt(lhs int, rhs int) int {
	if lhs > rhs {
		return lhs
	}
	return rhs
}

func minInt(lhs int, rhs int) int {
	if lhs < rhs {
		return lhs
	}
	return rhs
}

func (lhs Area3D) Intersection(rhs Area3D) Area3D {
	xmin, xmax := maxInt(lhs.start.x, rhs.start.x), minInt(lhs.end.x, rhs.end.x)
	ymin, ymax := maxInt(lhs.start.y, rhs.start.y), minInt(lhs.end.y, rhs.end.y)
	zmin, zmax := maxInt(lhs.start.z, rhs.start.z), minInt(lhs.end.z, rhs.end.z)
	return Area3D{Point3D{xmin, ymin, zmin}, Point3D{xmax, ymax, zmax}}
}

func (lhs Area3D) Overlap(rhs Area3D) []Area3D {
	if lhs.Intersection(rhs).IsEmpty() {
		return []Area3D{lhs, rhs}
	}
	areas := make([]Area3D, 0)

	x_values, y_values, z_values := lhs.orderCoords(rhs)
	for x_i := 1; x_i < len(x_values); x_i++ {
		for y_i := 1; y_i < len(y_values); y_i++ {
			for z_i := 1; z_i < len(z_values); z_i++ {
				start := Point3D{x_values[x_i-1], y_values[y_i-1], z_values[z_i-1]}
				end := Point3D{x_values[x_i], y_values[y_i], z_values[z_i]}
				area := Area3D{start, end}
				if !area.IsEmpty() && (lhs.Contains(area) || rhs.Contains(area)) {
					areas = append(areas, area)

				}
			}
		}
	}
	return areas
}

func (lhs Area3D) orderCoords(rhs Area3D) ([4]int, [4]int, [4]int) {
	x_values := [4]int{
		minInt(lhs.start.x, rhs.start.x),
		maxInt(lhs.start.x, rhs.start.x),
		minInt(lhs.end.x, rhs.end.x),
		maxInt(lhs.end.x, rhs.end.x),
	}
	y_values := [4]int{
		minInt(lhs.start.y, rhs.start.y),
		maxInt(lhs.start.y, rhs.start.y),
		minInt(lhs.end.y, rhs.end.y),
		maxInt(lhs.end.y, rhs.end.y),
	}
	z_values := [4]int{
		minInt(lhs.start.z, rhs.start.z),
		maxInt(lhs.start.z, rhs.start.z),
		minInt(lhs.end.z, rhs.end.z),
		maxInt(lhs.end.z, rhs.end.z),
	}
	return x_values, y_values, z_values
}

func (a Area3D) size() int {
	res := (a.start.x - a.end.x) * (a.start.y - a.end.y) * (a.start.z - a.end.z)
	if res < 0 {
		return -res
	}
	return res
}

func sumSizes(areas map[Area3D]int, value int) int {
	sum := 0
	for area, state := range areas {
		if state == value {
			sum += area.size()
		}
	}
	return sum
}

func lampsLit(areas []Area3D, commands []string) int {
	areaToState := map[Area3D]int{}
	for i, area := range areas {
		new_state := 1
		if commands[i] == "off" {
			new_state = 0
		}

		noop := false
		for key_area, area_state := range areaToState {
			if !key_area.Intersection(area).IsEmpty() {
				if new_state == area_state && key_area.Contains(area) {
					noop = true
					break
				}
				delete(areaToState, key_area)
				for _, overlap_area := range area.Overlap(key_area) {
					if !area.Contains(overlap_area) {
						areaToState[overlap_area] = area_state
					}
				}
			}

		}
		if !noop {
			areaToState[area] = new_state
		}
	}

	return sumSizes(areaToState, 1)
}

func SolveDay22() {
	areas, commands := readAreas()
	fmt.Println(lampsLit(areas, commands))

	// area1 := Area3D{Point3D{1.9, 0.9, 0.9}, Point3D{6.1, 5.1, 2.1}}
	// area2 := Area3D{Point3D{2.9, 1.9, 0.9}, Point3D{7.1, 4.1, 2.1}}
	// res := area1.decomposeAreas(area2)
	// for _, r := range res {
	// 	fmt.Println(r)
	// 	fmt.Println(sumSizes(r))
	// }

}
