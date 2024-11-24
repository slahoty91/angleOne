package model

type TokenData struct {
	Token          string `json:"token" bson:"token"`
	Symbol         string `json:"symbol" bson:"symbol"`
	Name           string `json:"name" bson:"name"`
	Expiry         string `json:"expiry" bson:"expiry"`
	Expiry_Zrd_Fmt string `json:"expiry_zer_fmt" bson:"expiry_zer_fmt"`
	Strike         string `json:"strike" bson:"strike"`
	Strike_Int     int64  `json:"strike_int" bson:"strike_int"`
	LotSize        string `json:"lotsize" bson:"lotsize"`
	InstrumentType string `json:"instrumenttype" bson:"instrument_type"`
	ExchangeSeg    string `json:"exch_seg" bson:"exch_seg"`
	TickSize       string `json:"tick_size" bson:"tick_size"`
	InstType       string `json:"inst_type" bson:"inst_type"`
}
