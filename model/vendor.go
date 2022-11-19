package model

type Vendor struct {
	Name         string   `json:"VendorName"`
	Email        string   `json:"email"`
	Address      string   `json:"address"`
	CeoName      string   `json:"ceoName"`
	ListProducts []string `json:"listTProducts"`
}
