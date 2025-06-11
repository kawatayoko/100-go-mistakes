package main

import (
	"fmt"

	"github.com/kawatayoko/100-go-mistakes/src/02-code-project-organization/06-producer-side-if2/client"
	"github.com/kawatayoko/100-go-mistakes/src/02-code-project-organization/06-producer-side-if2/store"
)

func printCustomers(getter client.CustomersGetter) error {
	customers, err := getter.GetAllCustomers()
	if err != nil {
		fmt.Println("error getting customers:", err)
		return err
	}
	for _, c := range customers {
		fmt.Printf("Customer ID: %s, Name: %s\n", c.ID, c.Name)
	}
	return nil
}

func main() {
	s := store.Store{}
	printCustomers(s)
}
