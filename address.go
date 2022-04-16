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
	CountryCode string `json:"country_code" bson:"country_code"`
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
	CountryCode string `json:"country_code" bson:"country_code"`
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
	CountryCode string         `json:"country_code" bson:"country_code"`
	Additionals string         `json:"additionals" bson:"additionals"`
	UserID      string         `json:"user_id" bson:"user_id"`
	Modified    vcago.Modified `json:"modified" bson:"modified"`
}

type AddressList []Address

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
	ID          []string `query:"id" qs:"id"`
	CrewID      []string `query:"crew_id" qs:"crew_id"`
	UserID      []string `query:"user_id" qs:"user_id"`
	UpdatedTo   string   `query:"updated_to" qs:"updated_to"`
	UpdatedFrom string   `query:"updated_from" qs:"updated_from"`
	CreatedTo   string   `query:"created_to" qs:"created_to"`
	CreatedFrom string   `query:"created_from" qs:"created_from"`
}

func (i *AddressQuery) Match() (r *vcago.MongoMatch) {
	r = vcago.NewMongoMatch()
	r.StringList("_id", i.ID)
	r.StringList("crew_id", i.CrewID)
	r.StringList("user_id", i.UserID)
	r.GteInt64("modified.updated", i.UpdatedFrom)
	r.GteInt64("modified.created", i.CreatedFrom)
	r.LteInt64("modified.updated", i.UpdatedTo)
	r.LteInt64("modified.created", i.CreatedTo)
	return
}
