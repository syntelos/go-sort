Comparable GOST for GOPL


  type Ordered interface 

  func Sort[T Ordered](array []T) ([]T) {


Review

  A "traditional" abstraction

    type Comparable interface

    func Sort(array []Comparable) ([]Comparable)

  fails syntactically.

    go test
    # github.com/syntelos/go-sort [github.com/syntelos/go-sort.test]
    ./sort_test.go:61:7: cannot use vector (variable of type TestList) as []Comparable value in argument to Sort
    FAIL	github.com/syntelos/go-sort [build failed]

  A syntactic dissonance which remains unclear and
  unresolved.

  The dynamical semantics required are available via
  parametric type abstraction, represented by the "Ordered"
  interface found in [MERO].


Note Bene

  The definition provided currently is a test of conception.

    type Ordered interface {
	    ~int | ~uint | ~string
    }


References

  [GOST] http://pkg.go.dev/sort
  [MERO] https://blog.merovius.de/posts/2024-01-05_constraining_complexity/
