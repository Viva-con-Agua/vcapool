package vcapool

import "github.com/Viva-con-Agua/vcago/vmod"

type (
	Profile struct {
		ID        string        `bson:"_id" json:"id"`
		Gender    string        `bson:"gender" json:"gender"`
		Phone     string        `bson:"phone" json:"phone"`
		Birthdate int64         `bson:"birthdate" json:"birthdate"`
		UserID    string        `bson:"user_id" json:"user_id"`
		Modified  vmod.Modified `bson:"modified" json:"modified"`
	}
)
