package main

import "fmt"

type point struct {
	x, y int
}

func main() {
	var m map[string]point
	m = make(map[string]point)
	fmt.Printf("empty m: %#v\n", m)

	m["start"] = point{0, 0}
	m["end"] = point{1, 1}

	fmt.Printf("m: %#v\nm[\"middle\"]: %#v\n", m, m["middle"])

	m = map[string]point{
		"different start": point{-1, -5},
		"middleearth":     {1, 1},
		"other end":       {2, 3},
	}

	type namedPoint struct {
		name string
		val  point
	}

	ordered := []namedPoint{
		{name: "different start", val: point{-1, -5}},
	}
	// sort.Sort() can now be used
	fmt.Println(ordered)

	// iterate
	for k, v := range m {
		fmt.Printf("item %v at key %s\n", v, k)
	}

	for _, v := range m {
		fmt.Printf("item %v\n", v)
	}

	for k := range m {
		fmt.Printf("key %s\n", k)
	}

	// whoa
	rev := map[point]string{}
	for k, v := range m {
		rev[v] = k
	}
	fmt.Printf("reversed %v\n", rev)
	v, ok := rev[point{1, 1}]
	fmt.Printf("does it have a middle (%v)? %v\n", v, ok)
	delete(rev, point{1, 1})
}
