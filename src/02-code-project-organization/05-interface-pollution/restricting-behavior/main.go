package main

import "fmt"

type IntConfig struct {
	Value int32
}

func (c *IntConfig) Get() int32 {
	return c.Value
}

func (c *IntConfig) Set(value int32) {
	c.Value = value
}

type intConfigGetter interface {
	Get() int32
}

type Foo struct {
	thresholder intConfigGetter
}

func NewFoo(thresholder intConfigGetter) Foo {
	return Foo{thresholder: thresholder}
}

func (f Foo) Bar() {
	threshold := f.thresholder.Get()
	fmt.Printf("threadshold is %v\n", threshold)
	// _ = threshold
}
func main() {
	f := NewFoo(&IntConfig{
		Value: 11,
	})
	f.Bar()
}
