package vcapool

import (
	"time"

	"github.com/Viva-con-Agua/vcago"
	"github.com/google/uuid"
)

type UserActive struct {
	ID       string          `bson:"_id" json:"id"`
	Status   string          `bson:"status" json:"status"`
	Since    int64           `bson:"since" json:"since"`
	Expired  int64           `bson:"expired" json:"expired"`
	History  HistoryDateList `bson:"history" json:"history"`
	UserID   string          `bson:"user_id" json:"user_id"`
	Modified vcago.Modified  `bson:"modified" json:"modified"`
}

func NewUserActive(userID string) *UserActive {
	return &UserActive{
		ID:       uuid.NewString(),
		Status:   "requested",
		Since:    0,
		Expired:  0,
		UserID:   userID,
		Modified: vcago.NewModified(),
	}
}

func (i *UserActive) Active() {
	i.Status = "active"
	i.Since = time.Now().Unix()
	i.Modified.Update()
}

func (i *UserActive) Rejected() {
	now := time.Now().Unix()
	if i.Status == "active" {
		i.History.Add(i.Since, now)
	}
	i.Status = "rejected"
	i.Expired = now
	i.Modified.Update()
}

func (i *UserActive) Requested() {
	i.Status = "requested"
	i.Since = 0
	i.Expired = 0
	i.Modified.Update()
}

func (i *UserActive) IsRequested() bool {
	if i.Status == "requested" {
		return true
	}
	return false
}
