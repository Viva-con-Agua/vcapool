package vcapool

import (
	"github.com/Viva-con-Agua/vcago"
	"github.com/google/uuid"
)

type (
	UserCrew struct {
		ID       string         `bson:"_id" json:"id"`
		UserID   string         `bson:"user_id" json:"user_id"`
		Name     string         `bson:"name" json:"name"`
		Email    string         `bson:"email" json:"email"`
		Roles    []vcago.Role   `bson:"roles" json:"roles"`
		CrewID   string         `bson:"crew_id" json:"crew_id"`
		Modified vcago.Modified `bson:"modified" json:"modified"`
	}
)

func NewUserCrew(userID string, crewID string, crewName string, email string) *UserCrew {
	return &UserCrew{
		ID:       uuid.NewString(),
		UserID:   userID,
		Name:     crewName,
		Email:    email,
		CrewID:   crewID,
		Modified: vcago.NewModified(),
	}
}
