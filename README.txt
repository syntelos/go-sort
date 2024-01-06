Comparable GOST for GOPL


  type Comparable interface {

	  Compare(Comparable) int
  }

  func Sort(array []Comparable) ([]Comparable)


Review


  The test case fails syntactically.

    go test
    # github.com/syntelos/go-sort [github.com/syntelos/go-sort.test]
    ./sort_test.go:61:7: cannot use vector (variable of type TestList) as []Comparable value in argument to Sort
    FAIL	github.com/syntelos/go-sort [build failed]

  A syntactic dissonance which remains unclear and
  unresolved.


References

  [GOST] http://pkg.go.dev/sort

