package entity

type Product struct {
	ID    int
	Name  string
	Price int
	Stock int
}

func (p Product) StockStatus() string {
	var status string
	if p.Stock < 10 {
		status = "Almost out of stock"
	} else if p.Stock < 100 {
		status = "Limited stock"
	}

	return status
}
