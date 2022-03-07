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

func (i *UserActive) Confirmed() {
	i.Status = "confirmed"
	i.Since = time.Now().Unix()
	i.Modified.Update()
}

func (i *UserActive) Rejected() {
	now := time.Now().Unix()
	if i.Status == "confirmed" {
		i.History.Add(i.Since, now)
	}
	i.Status = "rejected"
	i.Expired = now
	i.Since = 0
	i.Modified.Update()
}

func (i *UserActive) Requested() {
	i.Status = "requested"
	i.Since = 0
	i.Expired = 0
	i.Modified.Update()
}

func (i *UserActive) Withdraw() {
	now := time.Now().Unix()
	if i.Status == "confirmed" {
		i.History.Add(i.Since, now)
	}
	i.Status = "withdrawn"
	i.Since = 0
	i.Expired = now
	i.Modified.Update()
}

func (i *UserActive) IsRequested() bool {
	if i.Status == "requested" {
		return true
	}
	return false
}
func (i *UserActive) IsConfirmed() bool {
	if i.Status == "confirmed" {
		return true
	}
	return false
}
func (i *UserActive) IsWithdrawn() bool {
	if i.Status == "withdrawn" {
		return true
	}
	return false
}
func (i *UserActive) IsRejected() bool {
	if i.Status == "rejected" {
		return true
	}
	return false
}
