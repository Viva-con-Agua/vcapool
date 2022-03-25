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
	NVM       UserNVM        `json:"nvm" bson:"nvm,omitempty"`
	Modified  vcago.Modified `json:"modified" bson:"modified"`
}

type UserList []User

type UserQuery struct {
	FirstName     string   `query:"first_name" qs:"first_name"`
	LastName      string   `query:"last_name" qs:"last_name"`
	FullName      string   `query:"full_name" qs:"full_name"`
	DisplayName   string   `query:"display_name" qs:"display_name"`
	ActiveState   []string `query:"active_state" qs:"active_state"`
	SystemRoles   []string `query:"system_roles" qs:"system_roles"`
	PoolRoles     []string `query:"pool_roles" qs:"pool_roles"`
	PrivacyPolicy string   `query:"privacy_policy" qs:"privacy_policy"`
	NVMState      []string `query:"nvm_state" qs:"nvm_state"`
	CrewID        string   `query:"crew_id" qs:"crew_id"`
	Country       string   `query:"country" qs:"country"`
	Confirmed     string   `query:"confirmed" qs:"confirmed"`
	UpdatedTo     string   `query:"updated_to" qs:"updated_to"`
	UpdatedFrom   string   `query:"updated_from" qs:"updated_from"`
	CreatedTo     string   `query:"created_to" qs:"created_to"`
	CreatedFrom   string   `query:"created_from" qs:"created_from"`
}

type UserSimple struct {
	ID          string `json:"id,omitempty" bson:"_id"`
	Email       string `json:"email" bson:"email" validate:"required,email"`
	FirstName   string `bson:"first_name" json:"first_name" validate:"required"`
	LastName    string `bson:"last_name" json:"last_name" validate:"required"`
	FullName    string `bson:"full_name" json:"full_name"`
	DisplayName string `bson:"display_name" json:"display_name"`
}

func (i *UserList) GetEmailList() []string {
	list := []string{}
	for n, _ := range *i {
		list = append(list, (*i)[n].Email)
	}
	return list
}

func (i *User) ToAuthToken() (r *AuthToken, err error) {
	r = new(AuthToken)
	return NewAuthToken(i)
}

func UserFromInterface(i interface{}) (r *User, err error) {
	var ok bool
	if r, ok = i.(*User); !ok {
		return nil, errors.New("no user models")
	}
	return
}
