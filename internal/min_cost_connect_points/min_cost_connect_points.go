package min_cost_connect_points

import (
	"math"
	"sort"
)

// 1584. Min Cost to Connect All Points
// https://leetcode.com/problems/min-cost-to-connect-all-points/submissions/891487080/
func minCostConnectPointsUnionFind(points [][]int) int {
	// full graph - build edges
	var edges Edges
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			edges = append(edges, Edge{
				point1: Coord{points[i][0], points[i][1]},
				point2: Coord{points[j][0], points[j][1]},
				weight: int(math.Abs(float64(points[i][0]-points[j][0]))) + int(math.Abs(float64(points[i][1]-points[j][1])))})
		}
	}
	roots := make(map[Coord]Coord, len(points))
	for i := 0; i < len(points); i++ {
		p := points[i]
		roots[Coord{p[0], p[1]}] = Coord{p[0], p[1]}
	}

	sort.Sort(edges)
	result := 0
	for _, e := range edges {
		if find(e.point1, roots) != find(e.point2, roots) {
			result += e.weight
			union(e.point1, e.point2, &roots)
		}
	}

	return result
}

type Coord struct {
	x, y int
}

type Edge struct {
	point1 Coord
	point2 Coord
	weight int
}

type Edges []Edge

func (w Edges) Len() int           { return len(w) }
func (w Edges) Swap(i, j int)      { w[i], w[j] = w[j], w[i] }
func (w Edges) Less(i, j int) bool { return w[i].weight < w[j].weight }

func find(point Coord, roots map[Coord]Coord) Coord {
	for roots[point] != point {
		point = roots[point]
	}
	return point
}

func union(point1, point2 Coord, roots *map[Coord]Coord) {
	rp1 := find(point1, *roots)
	rp2 := find(point2, *roots)
	(*roots)[rp1] = rp2
}

// Prim algorithm to find minimum spanning tree
func minCostConnectPointsPrim(points [][]int) (ans int) {
	n := len(points)

	dis := make([]int, n)
	x, y := points[0][0], points[0][1]
	for i := range dis {
		dis[i] = abs(points[i][0]-x) + abs(points[i][1]-y)
	}

	mark := make([]bool, n)
	for {
		chosen, min := -1, 1_000_000_000
		for i := range dis {
			if !mark[i] && dis[i] < min {
				min, chosen = dis[i], i
			}
		}

		if chosen == -1 {
			break
		}

		mark[chosen] = true
		ans += dis[chosen]

		for i := range dis {
			if mark[i] {
				continue
			}

			d := abs(points[i][0]-points[chosen][0]) + abs(points[i][1]-points[chosen][1])
			if dis[i] > d {
				dis[i] = d
			}
		}
	}

	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
