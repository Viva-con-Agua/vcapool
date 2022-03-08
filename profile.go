package vcapool

import "github.com/Viva-con-Agua/vcago"

type (
	Profile struct {
		ID        string         `bson:"_id" json:"id"`
		Gender    string         `bson:"gender" json:"gender"`
		Phone     string         `bson:"phone" json:"phone"`
		Birthdate int64          `bson:"birthdate" json:"birthdate"`
		UserID    string         `bson:"user_id" json:"user_id"`
		Modified  vcago.Modified `bson:"modified" json:"modified"`
	}
	Avatar struct {
		ID       string         `bson:"_id" json:"id"`
		URL      string         `bson:"url" json:"url"`
		Type     string         `bson:"type" json:"type"`
		Modified vcago.Modified `bson:"modified" json:"modified"`
	}
)
