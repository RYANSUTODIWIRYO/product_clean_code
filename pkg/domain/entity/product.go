package entity

type Product struct {
	ID		string
	Name	string
	Price	float64
	Stok	int
}


type FetchProductsResponse struct {
	Products []*Product
}
