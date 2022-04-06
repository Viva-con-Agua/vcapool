package vcapool

import (
	"github.com/Viva-con-Agua/vcago"
	"github.com/google/uuid"
)

type AvatarCreate struct {
	URL    string `bson:"url" json:"url"`
	Type   string `bson:"type" json:"type"`
	UserID string `bson:"user_id" json:"user_id"`
}

func (i *AvatarCreate) Avatar() *Avatar {
	return &Avatar{
		ID:       uuid.NewString(),
		URL:      i.URL,
		Type:     i.Type,
		UserID:   i.UserID,
		Modified: vcago.NewModified(),
	}
}

type AvatarUpdate struct {
	ID     string `bson:"_id" json:"id"`
	URL    string `bson:"url" json:"url"`
	Type   string `bson:"type" json:"type"`
	UserID string `bson:"user_id" json:"user_id"`
}
type AvatarDelete struct {
	ID string `bson:"_id" json:"id"`
}

type Avatar struct {
	ID       string         `bson:"_id" json:"id"`
	URL      string         `bson:"url" json:"url"`
	Type     string         `bson:"type" json:"type"`
	UserID   string         `bson:"user_id" json:"user_id"`
	Modified vcago.Modified `bson:"modified" json:"modified"`
}
