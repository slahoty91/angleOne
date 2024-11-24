package model

type AngleOneOrder struct {
	Script      string `json:"script"`
	OrderID     string `json:"orderid"`
	Symbol      string `json:"symbol"`
	Qty         string `json:"qty"`
	LotSize     int    `json:"lot_size" bson:"lot_size"`
	SymbolToken string `json:"token"`
	Status      string `json:"status" bson:"statuss"`
	TransType   string `json:"trans_type" bson:"trans_type"`
	ZeroOrderID string `json:"zero_order_id" bson:"zero_order_id"`
}
