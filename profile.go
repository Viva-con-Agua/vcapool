package vcapool

import (
	"github.com/Viva-con-Agua/vcago"
	"github.com/google/uuid"
)

type (
	ProfileCreate struct {
		Gender    string `bson:"gender" json:"gender"`
		Phone     string `bson:"phone" json:"phone"`
		Birthdate int64  `bson:"birthdate" json:"birthdate"`
	}
	Profile struct {
		ID        string         `bson:"_id" json:"id"`
		Gender    string         `bson:"gender" json:"gender"`
		Phone     string         `bson:"phone" json:"phone"`
		Birthdate int64          `bson:"birthdate" json:"birthdate"`
		UserID    string         `bson:"user_id" json:"user_id"`
		Modified  vcago.Modified `bson:"modified" json:"modified"`
	}
)

func (i *ProfileCreate) Profile(userID string) *Profile {
	return &Profile{
		ID:        uuid.NewString(),
		Gender:    i.Gender,
		Phone:     i.Phone,
		Birthdate: i.Birthdate,
		UserID:    userID,
		Modified:  vcago.NewModified(),
	}
}
