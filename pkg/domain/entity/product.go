package entity

type Product struct {
	Id		string
	Name	string
	Price	float64
	Stock	int
}

type FetchProductsResponse struct {
	Products []*Product
}

type FindProductByIdRequest struct {
	Id		string
}

type FindProductByIdResponse struct {
	Product *Product
}