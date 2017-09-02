package errors

import (
	"testing"
)

func TestErrorf(t *testing.T) {
	err := Errorf(nil, "this is a error message")
	t.Log(err)

	f1 := func() {
		err = Errorf(nil, "error message in closure")
	}
	f1()
	t.Log(err)

	err = Errorf(err, "wrap error")
	t.Log(err)
}

func TestExample(t *testing.T) {
	func1()
}

func TestExample2(t *testing.T) {
	func11()
}
