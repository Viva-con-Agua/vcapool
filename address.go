package vcapool

import (
	"encoding/json"

	"github.com/Viva-con-Agua/vcago"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type AddressCreate struct {
	Street      string `json:"street" bson:"street"`
	Number      string `json:"number" bson:"number"`
	Zip         string `json:"zip" bson:"zip"`
	City        string `json:"city" bson:"city"`
	Country     string `json:"country" bson:"country"`
	Additionals string `json:"additionals" bson:"additionals"`
}

func (i *AddressCreate) Address(token *AccessToken) (r *Address) {
	bytes, _ := json.Marshal(&i)
	r = new(Address)
	_ = json.Unmarshal(bytes, &r)
	r.ID = uuid.NewString()
	r.UserID = token.ID
	r.Modified = vcago.NewModified()
	return
}

type AddressUpdate struct {
	ID          string `json:"id" bson:"_id"`
	Street      string `json:"street" bson:"street"`
	Number      string `json:"number" bson:"number"`
	Zip         string `json:"zip" bson:"zip"`
	City        string `json:"city" bson:"city"`
	Country     string `json:"country" bson:"country"`
	Additionals string `json:"additionals" bson:"additionals"`
	UserID      string `json:"user_id" bson:"user_id"`
}

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

type AddressParam struct {
	ID     string `param:"id"`
	UserID string
}

func (i *AddressParam) Filter() bson.M {
	filter := bson.M{"_id": i.ID}
	if i.UserID != "" {
		filter["user_id"] = i.UserID
	}
	return filter
}

type AddressQuery struct {
	ID     string `query:"id"`
	CrewID string `query:"crew_id"`
	UserID string `query:"user_id"`
}

func (i *AddressQuery) Filter() bson.M {
	f := vcago.NewMongoFilter()
	f.Equal("_id", i.ID)
	return f.Filter
}
