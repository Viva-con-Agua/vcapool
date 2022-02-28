package vcapool

import "github.com/Viva-con-Agua/vcago"

type UserNWM struct {
	Status    string         `bson:"status" json:"status"`
	Since     int64          `bson:"since" json:"since"`
	Expired   int64          `bson:"expired" json:"expired"`
	ExpiresAt int64          `bson:"expires_at" json:"expires_at"`
	Modified  vcago.Modified `bson:"modified" json:"modified"`
}
