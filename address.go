package vcapool

import "github.com/Viva-con-Agua/vcago"

type Address struct {
	ID          string         `json:"id" bson:"_id"`
	Street      string         `json:"street" bson:"street"`
	Number      string         `json:"number" bson:"number"`
	Zip         string         `json:"zip" bson:"zip"`
	City        string         `json:"city" bson:"city"`
	Country     string         `json:"country" bson:"country"`
	Additionals string         `json:"additionals" bson:"additionals"`
	UserID      string         `json:"user_id" bson:"user_id"`
	Modified    vcago.Modified `json:"modified" bson:"modified"`
}
