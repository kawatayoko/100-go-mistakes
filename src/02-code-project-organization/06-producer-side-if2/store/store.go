package store

type CustomerStorage interface {
	StoreCustomer(customer Customer) error
	GetCustomer(id string) (Customer, error)
	UpdateCustomer(customer Customer) error
	GetAllCustomers() ([]Customer, error)
	GetCustomersWithoutContract() ([]Customer, error)
	GetCustomersWithNegativeBalance() ([]Customer, error)
}

type Customer struct {
	ID   string
	Name string
}

func NewCutomer(id, name string) Customer {
	return Customer{
		ID:   id,
		Name: name,
	}
}

type Store struct{}

func (s Store) GetAllCustomers() ([]Customer, error) {
	// ここでデータベースや外部APIから顧客情報を取得するロジックを実装
	return []Customer{
		NewCutomer("1", "yoko"),
		NewCutomer("2", "taro"),
	}, nil
}
