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

func (this TestList) Print() {
	for _, el := range this {
		fmt.Println(el)
	}
}

func TestAscending(t *testing.T) {
	var vector TestList = TestList{ TestObject("20231219192613"), TestObject("20231221074246"), TestObject("20240102214104"), TestObject("20231222063428"), TestObject("20240104112200"), TestObject("20231217190339"), TestObject("20231213155157"), TestObject("20231219065525"), TestObject("20231231120412"), TestObject("20231221152849"), TestObject("20240102073948"), TestObject("20240101083455") }


	Ascending(vector)

	vector.Print()

	fmt.Println()
}

func TestDescending(t *testing.T) {
	var vector TestList = TestList{ TestObject("20231219192613"), TestObject("20231221074246"), TestObject("20240102214104"), TestObject("20231222063428"), TestObject("20240104112200"), TestObject("20231217190339"), TestObject("20231213155157"), TestObject("20231219065525"), TestObject("20231231120412"), TestObject("20231221152849"), TestObject("20240102073948"), TestObject("20240101083455") }


	Descending(vector)

	vector.Print()

	fmt.Println()
}
