package model

type Product struct {
	ProductName  string `json:"productName"`
	VendorName   string `json:"vendorName"`
	Available    int    `json:"available"`
	IsBestSeller bool   `json:"isBestSeller"`
	Minimum      int    `json:"minimum"`
	BuyNow       bool   `json:"buyNow"`
	Sku          string `json:"sku"`
}
