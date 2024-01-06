/*
 * Comparable GOST for GOPL
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package sort

type Comparable interface {

	Compare(Comparable) int
}

func Sort(array []Comparable) ([]Comparable) {

	var a, b int = 0, len(array)
	var x, y Comparable
	/*
	 * Partition sort adapted from GOST
	 */
	for i := a + 1; i < b; i++ {

		for j := i; j > a; j-- {

			x = array[j]
			y = array[j-1]

			if 0 > x.Compare(y) {

				array[j] = y
				array[j-1] = x
			} else {
				break
			}
		}
	}
	return array
}
