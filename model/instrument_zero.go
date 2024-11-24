package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type OptionInstrument struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	InstrumentToken int64              `bson:"instrument_token" json:"instrument_token"`
	ExchangeToken   string             `bson:"exchange_token" json:"exchange_token"`
	TradingSymbol   string             `bson:"tradingsymbol" json:"tradingsymbol"`
	Name            string             `bson:"name" json:"name"`
	LastPrice       float64            `bson:"last_price" json:"last_price"`
	Expiry          string             `bson:"expiry" json:"expiry"`
	Strike          int64              `bson:"strike" json:"strike"`
	TickSize        float64            `bson:"tick_size" json:"tick_size"`
	LotSize         int                `bson:"lot_size" json:"lot_size"`
	InstrumentType  string             `bson:"instrument_type" json:"instrument_type"`
	Segment         string             `bson:"segment" json:"segment"`
	Exchange        string             `bson:"exchange" json:"exchange"`
}
