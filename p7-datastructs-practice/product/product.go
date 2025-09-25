package product

type Product struct {
	title string
	price float64
	id int
}

func New(title string, price float64, id int) Product {
	return Product{
		title: title,
		price: price,
		id: id,
	}
}