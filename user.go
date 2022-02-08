package vpool

import (
	"errors"

	"github.com/Viva-con-Agua/vcago"
)

type User struct {
	ID       string         `json:"id,omitempty" bson:"_id"`
	Email    string         `json:"email" bson:"email" validate:"required,email"`
	Profile  vcago.Profile  `json:"profile" bson:"profile"`
	CrewID   string         `json:"crew_id" bson:"crew_id"`
	Crew     Crew           `json:"crew" bson:"-"`
	Address  Address        `json:"address" bson:"address"`
	Roles    string         `json:"roles" bson:"roles"`
	Modified vcago.Modified `json:"modified" bson:"modified"`
}

type Address struct {
	Street      string `json:"street" bson:"street"`
	Number      string `json:"number" bson:"number"`
	Zip         string `json:"zip" bson:"zip"`
	City        string `json:"city" bson:"city"`
	Country     string `json:"country" bson:"country"`
	Additionals string `json:"additionals" bson:"additionals"`
}

func NewUser(user *vcago.User) (r *User) {
	return &User{
		ID:       user.ID,
		Email:    user.Email,
		Profile:  user.Profile,
		Roles:    user.Roles,
		Modified: vcago.NewModified(),
	}
}

func UserFromInterface(i interface{}) (r *User, err error) {
	var ok bool
	if r, ok = i.(*User); !ok {
		return nil, vcago.NewStatusInternal(errors.New("no user models"))
	}
	return
}
