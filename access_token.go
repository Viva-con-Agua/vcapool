package vcapool

import (
	"errors"
	"time"

	"github.com/Viva-con-Agua/vcago"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type AccessToken struct {
	ID            string               `json:"id,omitempty" bson:"_id"`
	Email         string               `json:"email" bson:"email" validate:"required,email"`
	FirstName     string               `bson:"first_name" json:"first_name" validate:"required"`
	LastName      string               `bson:"last_name" json:"last_name" validate:"required"`
	FullName      string               `bson:"full_name" json:"full_name"`
	DisplayName   string               `bson:"display_name" json:"display_name"`
	Roles         vcago.RoleListCookie `json:"system_roles" bson:"system_roles"`
	Country       string               `bson:"country" json:"country"`
	PrivacyPolicy bool                 `bson:"privacy_policy" json:"privacy_policy"`
	Confirmd      bool                 `bson:"confirmed" json:"confirmed"`
	LastUpdate    string               `bson:"last_update" json:"last_update"`
	Profile       Profile              `json:"profile" bson:"profile,truncate"`
	CrewName      string               `json:"crew_name"`
	CrewID        string               `json:"crew_id"`
	CrewEmail     string               `json:"crew_email"`
	AddressID     string               `json:"address_id"`
	PoolRoles     vcago.RoleListCookie `json:"pool_roles"`
	ActiveState   string               `json:"active_state"`
	NVMState      string               `json:"nvm_state"`
	AvatarID      string               `json:"avatar_id"`
	Modified      vcago.Modified       `json:"modified"`
	jwt.StandardClaims
}

func NewAccessToken(user *User) *AccessToken {
	return &AccessToken{
		user.ID,
		user.Email,
		user.FirstName,
		user.LastName,
		user.FullName,
		user.DisplayName,
		*user.Roles.Cookie(),
		user.Country,
		user.PrivacyPolicy,
		user.Confirmd,
		user.LastUpdate,
		user.Profile,
		user.Crew.Name,
		user.Crew.CrewID,
		user.Crew.Email,
		user.Address.ID,
		*user.PoolRoles.Cookie(),
		user.Active.Status,
		user.NVM.Status,
		user.Avatar.ID,
		user.Modified,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}
}

func (i *AccessToken) SignedString(secret string) (string, error) {
	temp := jwt.NewWithClaims(jwt.SigningMethodHS256, i)
	return temp.SignedString([]byte(secret))
}

func (i *AccessToken) UserInternal() *UserInternal {
	return &UserInternal{
		ID:          i.ID,
		Email:       i.Email,
		FirstName:   i.FirstName,
		LastName:    i.LastName,
		FullName:    i.FullName,
		DisplayName: i.DisplayName,
		Profile:     i.Profile,
	}
}

func (i *AccessToken) UserSimple() *UserSimple {
	return &UserSimple{
		ID:          i.ID,
		Email:       i.Email,
		FirstName:   i.FirstName,
		LastName:    i.LastName,
		FullName:    i.FullName,
		DisplayName: i.DisplayName,
	}
}

func (i *AccessToken) CrewSimple() *CrewSimple {
	return &CrewSimple{
		ID:    i.CrewID,
		Name:  i.CrewName,
		Email: i.CrewEmail,
	}
}

//AccessCookieConfig can with echo for middleware.JWTWithConfig(vmod.AccessConfig) to handling access controll
//The token is reachable with c.Get("token")
func AccessCookieConfig() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(
		middleware.JWTConfig{
			Claims:      &AccessToken{},
			ContextKey:  "token",
			TokenLookup: "cookie:access_token",
			SigningKey:  []byte(jwtSecret),
		})
}

func AccessCookieUser(c echo.Context) (r *AccessToken, err error) {
	token := c.Get("token").(*jwt.Token)
	if token == nil {
		return nil, errors.New("No user in Conext")
	}
	r = token.Claims.(*AccessToken)
	return
}
