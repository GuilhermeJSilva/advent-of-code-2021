package solutions

import (
	"fmt"
	"sort"

	"gonum.org/v1/gonum/mat"
)

var option_order [][]int = [][]int{
	{1, 2, 3},
	{1, -2, -3},
	{-1, 2, -3},
	{-1, -2, 3},
	{1, 3, -2},
	{1, -3, 2},
	{-1, 3, 2},
	{-1, -3, -2},
	{2, 3, 1},
	{2, -3, -1},
	{-2, 3, -1},
	{-2, -3, 1},
	{2, 1, -3},
	{2, -1, 3},
	{-2, 1, 3},
	{-2, -1, -3},
	{3, 1, 2},
	{3, -1, -2},
	{-3, 1, -2},
	{-3, -1, 2},
	{3, 2, -1},
	{3, -2, 1},
	{-3, 2, 1},
	{-3, -2, -1},
}

func createTransformations() []*mat.Dense {
	options := make([]*mat.Dense, 0)
	for _, order := range option_order {
		mat_values := make([]float64, 0)
		for _, value := range order {
			for i := 1; i <= 3; i++ {
				if value == i || value == -i {
					mat_values = append(mat_values, float64(value/i))
				} else {
					mat_values = append(mat_values, 0)
				}
			}
		}
		options = append(options, mat.NewDense(3, 3, mat_values))
	}

	return options
}

func readScannerPoints() (*mat.Dense, error) {
	var scanner_id int
	_, err := fmt.Scanf("--- scanner %d ---\n", &scanner_id)
	if err != nil {
		return nil, err
	}
	var n_points int = 0
	points := make([]float64, 0)
	for {
		var x, y, z int
		_, err := fmt.Scanf("%d,%d,%d\n", &x, &y, &z)
		if err != nil {
			break
		}
		points = append(points, float64(x))
		points = append(points, float64(y))
		points = append(points, float64(z))
		n_points++
	}

	return mat.NewDense(n_points, 3, points), nil
}

func readAllScanners() []*mat.Dense {
	scanners := make([]*mat.Dense, 0)
	for {
		points, err := readScannerPoints()
		if err != nil {
			break
		}
		scanners = append(scanners, points)
	}
	return scanners
}

func matchPoints(points *mat.Dense, ref_points *mat.Dense, new_ref *mat.Dense) mat.Dense {
	rows, _ := points.Dims()
	var center2center mat.Dense
	center2center.Sub(new_ref, ref_points)
	replicated := center2center.Grow(rows-1, 0).(*mat.Dense)
	replicated.Apply(func(i, j int, v float64) float64 { return replicated.At(0, j) }, replicated)
	replicated.Add(replicated, points)
	return *replicated
}

func flatUniqueMatrix(matrix [][]float64) []float64 {
	flat := make([]float64, 0)
	flat = append(flat, matrix[0]...)
	for i, row := range matrix[1:] {
		if matrix[i][0] != row[0] || matrix[i][1] != row[1] || matrix[i][2] != row[2] {
			flat = append(flat, row...)
		}
	}
	return flat
}

func mergePoints(lhs *mat.Dense, rhs *mat.Dense) *mat.Dense {
	points := make([][]float64, 0)
	data := lhs.RawMatrix().Data
	for i := 0; i < len(data); i += 3 {
		points = append(points, []float64{data[i], data[i+1], data[i+2]})
	}
	data = rhs.RawMatrix().Data
	for i := 0; i < len(data); i += 3 {
		points = append(points, []float64{data[i], data[i+1], data[i+2]})
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0] || (points[i][0] == points[j][0] && points[i][1] < points[j][1]) || (points[i][0] == points[j][0] && points[i][1] == points[j][1] && points[i][2] < points[j][2])
	})

	unique_points := flatUniqueMatrix(points)
	return mat.NewDense(len(unique_points)/3, 3, unique_points)
}

func matchScanner(known *mat.Dense, unknown_raw *mat.Dense, transformations []*mat.Dense) (*mat.Dense, mat.Dense, bool) {
	known_rows, _ := known.Dims()
	for i := 0; i < known_rows; i++ {
		match_known := mat.NewDense(1, 3, known.RawRowView(i))
		for _, trans := range transformations {
			var unknown mat.Dense
			unknown.Mul(unknown_raw, trans)
			unknown_rows, _ := unknown.Dims()
			for j := 0; j < unknown_rows; j++ {
				match_unknown := mat.NewDense(1, 3, unknown.RawRowView(j))
				points := matchPoints(&unknown, match_unknown, match_known)
				all_points := mergePoints(known, &points)
				n_points, _ := all_points.Dims()
				if n_points <= unknown_rows+known_rows-12 {
					var center mat.Dense
					center.Sub(match_known, match_unknown)
					fmt.Println(mat.Formatted(&center))
					return all_points, center, true
				}
			}
		}
	}
	return known, *mat.NewDense(1, 1, []float64{0}), false
}

func matchAll(scanners []*mat.Dense, transformations []*mat.Dense) (*mat.Dense, []mat.Dense) {

	points := scanners[0]
	to_search := make([]int, 0)
	centers := make([]mat.Dense, 0)
	for i, scanner := range scanners[1:] {
		var matched bool
		var center mat.Dense
		points, center, matched = matchScanner(points, scanner, transformations)
		if !matched {
			to_search = append(to_search, i+1)
		} else {
			centers = append(centers, center)
		}
	}

	for len(to_search) > 0 {
		to_search_again := make([]int, 0)
		fmt.Println(to_search)
		for _, i := range to_search {
			var matched bool
			var center mat.Dense
			points, center, matched = matchScanner(points, scanners[i], transformations)
			if !matched {
				to_search_again = append(to_search_again, i)
			} else {
				centers = append(centers, center)
			}
		}
		to_search = to_search_again
	}
	return points, centers
}

func SolveDay19() {
	scanners := readAllScanners()
	transformations := createTransformations()
	points, centers := matchAll(scanners, transformations)
	fmt.Println(points.Dims())
	max := 0
	for i, center := range centers {
		for j, other_center := range centers {
			if i == j {
				continue
			}
			dst := 0
			for i, val := range center.RawMatrix().Data {
				diff := val - other_center.RawMatrix().Data[i]
				if diff < 0 {
					dst -= int(diff)
				} else {
					dst += int(diff)
				}
			}
			if max < dst {
				max = dst
			}
		}
	}
	fmt.Println(max)
}
