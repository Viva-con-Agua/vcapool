package vcapool

/*
import (
	"net/http"

	"github.com/Viva-con-Agua/vcago"
)


var authCookie = vcago.NewCookieConfig()

//AuthToken represents the authentication tokens for handling access via jwt.
type AuthToken struct {
	AccessToken  string `json:"access_token" bson:"access_token"`
	RefreshToken string `json:"refresh_token" bson:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at" bson:"expires_at"`
}
*/
//NewAuthToken creates an new access and refresh token for the given user.
/*func NewAuthToken(user *User) (r *AuthToken, err error) {
	r = new(AuthToken)
	if r.AccessToken, err = NewAccessToken(user).SignedString(JWTSecret); err != nil {
		return
	}
	r.RefreshToken, err = NewRefreshToken(user.ID).SignedString(JWTSecret)
	return
}*/
/*
//AccessCookie return an cookie conains the access_token.
func (i *AuthToken) AccessCookie() (r *http.Cookie) {
	return authCookie.Cookie("access_token", i.AccessToken)
}

//ResetAccessCookie returns an cookie for reset the access_token.
func ResetAccessCookie() *http.Cookie {
	return authCookie.Cookie("access_token", "")
}

//RefreshCookie returns an cookie conains the refresh_token.
func (i *AuthToken) RefreshCookie() *http.Cookie {
	return authCookie.Cookie("refresh_token", i.RefreshToken)
}

//ResetRefreshCookie returns an cookie for reset the refresh_token.
func ResetRefreshCookie() *http.Cookie {
	return authCookie.Cookie("refresh_token", "")
}*/
