/*
 * Comparable GOST for GOPL
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package sort

import (
	"fmt"
	"testing"
)

type TestObject string

type TestList []TestObject

func (this TestObject) Compare(that TestObject) int {

	var bi, bj byte
	var x, y, z int = 0, len(this), len(that)
	var d, c int = 0, 0 

	if y == z {

		c = z
	} else {
		if y < z {
			c = y
			d = -1
		} else {
			c = z
			d = 1
		}
	}

	for ; x < c; x++ {
		bi = this[x]
		bj = that[x]
		if bi != bj {

			if bi < bj {

				return -1
			} else {

				return 1
			}
		}
	}
	return d
}

func (this TestList) Print() {
	for _, el := range this {
		fmt.Println(el)
	}
}

func TestSort(t *testing.T) {
	var vector TestList = TestList{ TestObject("20231219192613"), TestObject("20231221074246"), TestObject("20240102214104"), TestObject("20231222063428"), TestObject("20240104112200"), TestObject("20231217190339"), TestObject("20231213155157"), TestObject("20231219065525"), TestObject("20231231120412"), TestObject("20231221152849"), TestObject("20240102073948"), TestObject("20240101083455") }

	Sort(vector)

	vector.Print()
}
