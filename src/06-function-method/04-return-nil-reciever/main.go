package main

import (
	"errors"
	"log"
	"strings"
)

type MultiError struct {
	errs []string
}

func (m *MultiError) Add(err error) {
	m.errs = append(m.errs, err.Error())
}

// MultiErrorは、このError() 関数があることで、errorインタフェースを満たしている
func (m *MultiError) Error() string {
	return strings.Join(m.errs, ";")
}

type Customer struct {
	Age  int
	Name string
}

func (c Customer) Validate() error {
	var m *MultiError
	// この段階ではmはnil

	if c.Age < 0 {
		m = &MultiError{} //新たなMultiErrorを割り当てる。nilでなくなる。
		m.Add(errors.New("age is negative"))
	}
	if c.Name == "" {
		if m == nil {
			m = &MultiError{} //新たなMultiErrorを割り当てる。nilでなくなる。
		}
		m.Add(errors.New("name is nil"))
	}
	return m
}

type Foo struct{}

func (foo *Foo) Bar() string {
	return "bar"
}

func main() {
	customer := Customer{
		Age:  33,
		Name: "John",
	}
	if err := customer.Validate(); err != nil {
		log.Fatalf("customer is invalid: %v", err)
	}

	// var foo *Foo
	// fmt.Println(foo.Bar())
}
