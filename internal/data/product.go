package data

// Product is the DTO for the product domain
type Product struct {
	ID    string `json:"id"`
	Type  string `json:"product_type"`
	Brand string `json:"brand"`
	Name  string `json:"name"`
}
