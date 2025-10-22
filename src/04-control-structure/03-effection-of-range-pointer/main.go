package main

import "fmt"

type Foo struct {
	value string
}

type Store struct {
	m map[string]*Customer
}

type Customer struct {
	ID      string
	Balance float64
}

func (s *Store) StoreCustomers(customers []Customer) {
	for _, customer := range customers {
		// mapにcustomerへのポインタを保存する
		s.m[customer.ID] = &customer
	}
}

func main() {
	s := Store{m: map[string]*Customer{}}
	s.StoreCustomers([]Customer{
		{
			ID:      "1",
			Balance: 10,
		},
		{
			ID:      "2",
			Balance: -10,
		},
		{
			ID:      "3",
			Balance: 0,
		},
	})
	for k, v := range s.m {
		fmt.Printf("key: %s, value: %+v, address: %p\n", k, v, v)
		// fmt.Printf("key: %s", k)
		// fmt.Printf("value: %v", v)
	}
}
