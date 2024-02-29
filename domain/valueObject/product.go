package valueObject

type Product struct {
	ProductId string `json:"productId"`
	Type      string `json:"type"`
	Quantity  int    `json:"quantity"`
}
