package vcapool

import (
	"time"

	"github.com/Viva-con-Agua/vcago"
	"github.com/google/uuid"
)

type UserNVM struct {
	ID        string          `bson:"_id" json:"id"`
	Status    string          `bson:"status" json:"status"`
	Since     int64           `bson:"since" json:"since"`
	Expired   int64           `bson:"expired" json:"expired"`
	ExpiresAt int64           `bson:"expires_at" json:"expires_at"`
	History   HistoryDateList `bson:"history" json:"history"`
	UserID    string          `bson:"user_id" json:"user_id"`
	Modified  vcago.Modified  `bson:"modified" json:"modified"`
}

func NewUserNVM(userID string) *UserNVM {
	return &UserNVM{
		ID:        uuid.NewString(),
		Status:    "requested",
		Since:     0,
		Expired:   0,
		ExpiresAt: 0,
		UserID:    userID,
		Modified:  vcago.NewModified(),
	}
}

func (i *UserNVM) Confirm() {
	i.Status = "confirmed"
	i.Since = time.Now().Unix()
	i.Modified.Update()
}

func (i *UserNVM) Reject() {
	now := time.Now().Unix()
	if i.Status == "confirmed" {
		i.History.Add(i.Since, now)
	}
	i.Status = "rejected"
	i.Expired = now
	i.Modified.Update()
}

func (i *UserActive) Request() {
	i.Status = "requested"
	i.Since = 0
	i.Expired = 0
	i.Modified.Update()
}
