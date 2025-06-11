package client

import "github.com/kawatayoko/100-go-mistakes/src/02-code-project-organization/06-producer-side-if2/store"

type CustomersGetter interface {
	GetAllCustomers() ([]store.Customer, error)
}
