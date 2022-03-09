package vcapool

import (
	"time"

	"github.com/Viva-con-Agua/vcago"
	"github.com/google/uuid"
)

type UserNVM struct {
	ID       string          `bson:"_id" json:"id"`
	Status   string          `bson:"status" json:"status"`
	Since    int64           `bson:"since" json:"since"`
	Expired  int64           `bson:"expired" json:"expired"`
	History  HistoryDateList `bson:"history" json:"history"`
	UserID   string          `bson:"user_id" json:"user_id"`
	Modified vcago.Modified  `bson:"modified" json:"modified"`
}

func NewUserNVM(userID string) *UserNVM {
	return &UserNVM{
		ID:       uuid.NewString(),
		Status:   "confirmed",
		Since:    time.Now().Unix(),
		Expired:  0,
		UserID:   userID,
		Modified: vcago.NewModified(),
	}
}

func (i *UserNVM) Confirmed() {
	i.Status = "confirmed"
	i.Since = time.Now().Unix()
	i.Modified.Update()
}

func (i *UserNVM) Rejected() {
	now := time.Now().Unix()
	if i.Status == "confirmed" {
		i.History.Add(i.Since, now)
	}
	i.Status = "rejected"
	i.Expired = now
	i.Modified.Update()
}

func (i *UserNVM) Withdraw() {
	now := time.Now().Unix()
	if i.Status == "confirmed" {
		i.History.Add(i.Since, now)
	}
	i.Status = "withdrawn"
	i.Since = 0
	i.Expired = now
	i.Modified.Update()
}

func (i *UserNVM) IsConfirmed() bool {
	if i.Status == "confirmed" {
		return true
	}
	return false
}
func (i *UserNVM) IsWithdrawn() bool {
	if i.Status == "withdrawn" {
		return true
	}
	return false
}
func (i *UserNVM) IsRejected() bool {
	if i.Status == "rejected" {
		return true
	}
	return false
}
