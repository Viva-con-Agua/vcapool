package vcapool

import (
	"net/http"

	"github.com/Viva-con-Agua/vcago"
)

var jwtSecret = vcago.Config.GetEnvString("JWT_SECRET", "w", "secret")

//AuthCookie represents the default authentication cookie
var AuthCookie = AuthCookieDefault()

//AuthCookieDefault returns an http.Cookie with the default set parameters.
func AuthCookieDefault() (r *http.Cookie) {
	r = new(http.Cookie)
	var sameSite string
	sameSite = vcago.Config.GetEnvString("COOKIE_SAME_SITE", "w", "strict")
	if sameSite == "lax" {
		r.SameSite = http.SameSiteLaxMode
	}
	if sameSite == "strict" {
		r.SameSite = http.SameSiteStrictMode
	}
	if sameSite == "none" {
		r.SameSite = http.SameSiteNoneMode
	}
	r.Secure = vcago.Config.GetEnvBool("COOKIE_SECURE", "w", true)
	r.HttpOnly = vcago.Config.GetEnvBool("COOKIE_HTTP_ONLY", "w", true)
	r.Path = "/"
	r.Domain = vcago.Config.GetEnvString("COOKIE_DOMAIN", "w", "localhost")
	return
}

type AuthToken struct {
	AccessToken  string `json:"access_token" bson:"access_token"`
	RefreshToken string `json:"refresh_token" bson:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at" bson:"expires_at"`
}

func NewAuthToken(user *User) (r *AuthToken, err error) {
	r = new(AuthToken)
	if r.AccessToken, err = NewAccessToken(user).SignedString(jwtSecret); err != nil {
		return
	}
	r.RefreshToken, err = NewRefreshToken(user.ID).SignedString(jwtSecret)
	return
}

func (i *AuthToken) AccessCookie() (r *http.Cookie) {
	r = AuthCookie
	r.Name = "access_token"
	r.Value = i.AccessToken
	return
}

func (i *AuthToken) RefreshCookie() (r *http.Cookie) {
	r = AuthCookie
	r.Name = "refresh_token"
	r.Value = i.RefreshToken
	return
}
