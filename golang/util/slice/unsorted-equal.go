// given 2 slices which are not sorted, find if they both have same elements (need not be in the same order)

//TODO: below logic need not be the most performant one. it works good for small size of the slices
package main

import "fmt"

func main() {
	a := []string{"a", "b", "c"}
	b := []string{"a", "b", "c"} // same data case
	c := []string{"c", "b", "a"} // same data different order case
	d := []string{"b", "c", "d"} // different data, same length case
	e := []string{"b", "c"} // different data, different lenght case
	f := []string{} // empty case

	a1 := []string{"a", "a", "b"}
	a2 := []string{"a", "a", "b"} // same data case with duplicates
	a3 := []string{"a", "b", "a"} //
	g := []string{"a", "b", "b"}
	h := []string{"a", "b"}

	a4 := []string{}
	a5 := []string{} // both empty case

	fmt.Println(sliceComp(a, b))   // true
	fmt.Println(sliceComp(a, c))   // true
	fmt.Println(sliceComp(a, d))   // false
	fmt.Println(sliceComp(a, e))   // false
	fmt.Println(sliceComp(a, f))   // false
	fmt.Println(sliceComp(a1, g))  // false
	fmt.Println(sliceComp(a1, h))  // false
	fmt.Println(sliceComp(a1, a2)) // true
	fmt.Println(sliceComp(a1, a3)) // true
	fmt.Println(sliceComp(a4, a5)) // true
}

func sliceComp(aa []string, bb []string) bool {
	if len(aa) != len(bb) {
		return false
	}

	ma := slice2map(aa)
	mb := slice2map(bb)

	for a, acnt := range ma {
		if bcnt, exists := mb[a]; !exists || acnt != bcnt {
			return false
		}

	}
	return true
}

func slice2map(aa []string) map[string]int {
	m := map[string]int{}
	for _, a := range aa {
		m[a]++
	}
	return m
}
