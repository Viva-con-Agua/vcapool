package vcapool

import (
	"errors"

	"github.com/Viva-con-Agua/vcago"
)

type User struct {
	ID       string         `json:"id,omitempty" bson:"_id"`
	Email    string         `json:"email" bson:"email" validate:"required,email"`
	Profile  Profile        `json:"profile" bson:"profile"`
	Crew     UserCrew       `json:"crew" bson:"crew"`
	Address  Address        `json:"address" bson:"address"`
	Roles    vcago.RoleList `json:"roles" bson:"roles"`
	Active   UserActive     `json:"active" bson:"active"`
	NWM      UserNWM        `json:"nwm" bson:"nwm"`
	Country  string         `bson:"country" json:"country"`
	Modified vcago.Modified `json:"modified" bson:"modified"`
}

func (i *User) ToAuthToken() (r *AuthToken, err error) {
	r = new(AuthToken)
	return NewAuthToken(i)
}

func UserFromInterface(i interface{}) (r *User, err error) {
	var ok bool
	if r, ok = i.(*User); !ok {
		return nil, vcago.NewStatusInternal(errors.New("no user models"))
	}
	return
}
