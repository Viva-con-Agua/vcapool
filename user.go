package vcapool

import (
	"errors"

	"github.com/Viva-con-Agua/vcago"
)

type User struct {
	ID            string         `json:"id,omitempty" bson:"_id"`
	Email         string         `json:"email" bson:"email" validate:"required,email"`
	FirstName     string         `bson:"first_name" json:"first_name" validate:"required"`
	LastName      string         `bson:"last_name" json:"last_name" validate:"required"`
	FullName      string         `bson:"full_name" json:"full_name"`
	DisplayName   string         `bson:"display_name" json:"display_name"`
	Roles         vcago.RoleList `json:"system_roles" bson:"system_roles"`
	Country       string         `bson:"country" json:"country"`
	PrivacyPolicy bool           `bson:"privacy_policy" json:"privacy_policy"`
	Confirmd      bool           `bson:"confirmed" json:"confirmed"`
	LastUpdate    string         `bson:"last_update" json:"last_update"`
	//extends the vcago.User
	Profile   Profile        `json:"profile" bson:"profile,truncate"`
	Crew      UserCrew       `json:"crew" bson:"crew,omitempty"`
	Avatar    Avatar         `bson:"avatar,omitempty" json:"avatar"`
	Address   Address        `json:"address" bson:"address,omitempty"`
	PoolRoles vcago.RoleList `json:"pool_roles" bson:"pool_roles,omitempty"`
	Active    UserActive     `json:"active" bson:"active,omitempty"`
	NWM       UserNWM        `json:"nwm" bson:"nwm,omitempty"`
	Modified  vcago.Modified `json:"modified" bson:"modified"`
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
