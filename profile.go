package vcapool

import "github.com/Viva-con-Agua/vcago"

type (
	Profile struct {
		ID          string         `bson:"_id" json:"id"`
		FirstName   string         `bson:"first_name" json:"first_name" validate:"required"`
		LastName    string         `bson:"last_name" json:"last_name" validate:"required"`
		FullName    string         `bson:"full_name" json:"full_name"`
		DisplayName string         `bson:"display_name" json:"display_name"`
		Gender      string         `bson:"gender" json:"gender"`
		Avatar      Avatar         `bson:"avatar" json:"avatar"`
		Phone       string         `bson:"phone" json:"phone"`
		UserID      string         `bson:"user_id" json:"user_id"`
		Modified    vcago.Modified `bson:"modified" json:"modified"`
	}
	Avatar struct {
		URL  string `bson:"url" json:"url"`
		Type string `bson:"type" json:"type"`
	}
)

func (i *Profile) NewFromVcaGo(profile *vcago.Profile) {
	i.FirstName = profile.FirstName
	i.LastName = profile.LastName
	i.FullName = profile.FullName
}
