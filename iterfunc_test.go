package iterfunc_test

import (
	"bufio"
	"fmt"
	"log"
	"strings"

	"github.com/phelmkamp/iterfunc"
)

type MyIter struct{ i int }

func (m *MyIter) Next() bool {
	m.i++
	return m.i < 4
}

func (m *MyIter) Value() int {
	return m.i
}

func Example() {
	var m MyIter
	for i := range iterfunc.New(m.Next, m.Value) {
		fmt.Println(i)
	}
	// Output:
	// 1
	// 2
	// 3
}

func Example_bufio_Scanner() {
	txt := `Line 1
Line 2
Line 3
`
	sc := bufio.NewScanner(strings.NewReader(txt))
	for v := range iterfunc.New(sc.Scan, sc.Bytes) {
		fmt.Println(string(v))
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	// Output:
	// Line 1
	// Line 2
	// Line 3
}
