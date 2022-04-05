package vcapool

import (
	"github.com/Viva-con-Agua/vcago"
	"github.com/google/uuid"
)

type Application struct {
	ID       string         `json:"id" bson:"_id"`
	User     UserSimple     `json:"user" bson:"user"`
	EventID  string         `json:"event_id" bson:"event_id"`
	Status   string         `json:"status" bson:"status"`
	Crew     CrewSimple     `json:"crew" bson:"crew"`
	Modified vcago.Modified `json:"modified" bson:"modified"`
}

type ApplicationCreate struct {
	EventID string `json:"event_id"`
}

func (i *ApplicationCreate) Application(token AccessToken) *Application {
	return &Application{
		ID: uuid.NewString(),
		User: UserSimple{
			ID:          token.ID,
			Email:       token.Email,
			DisplayName: token.DisplayName,
		},
	}
}
