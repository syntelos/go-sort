/*
 * Comparable GOST for GOPL
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package sort

type Ordered interface {
	~int | ~uint | ~string
}
/*
 * Zero-positive array index order is ascending.
 */
func Ascending[T Ordered](array []T) ([]T) {

	var a, b int = 0, len(array)
	var x, y T
	/*
	 * Partition sort adapted from GOST
	 */
	for i := a + 1; i < b; i++ {

		for j := i; j > a; j-- {

			x = array[j]
			y = array[j-1]

			if x < y {

				array[j] = y
				array[j-1] = x
			} else {
				break
			}
		}
	}
	return array
}
/*
 * Zero-positive array index order is descending.
 */
func Descending[T Ordered](array []T) ([]T) {

	var a, b int = 0, len(array)
	var x, y T
	/*
	 * Partition sort adapted from GOST
	 */
	for i := a + 1; i < b; i++ {

		for j := i; j > a; j-- {

			x = array[j]
			y = array[j-1]

			if x > y {

				array[j] = y
				array[j-1] = x
			} else {
				break
			}
		}
	}
	return array
}
