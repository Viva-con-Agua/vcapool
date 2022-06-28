package vcapool

import (
	"github.com/Viva-con-Agua/vcago"
	"github.com/google/uuid"
)

type ParticipationCreate struct {
	EventID string `json:"event_id" bson:"event_id"`
	Comment string `json:"comment" bson:"comment"`
}

type ParticipationUpdate struct {
	ID      string `json:"id" bson:"_id"`
	Status  string `json:"status" bson:"status"`
	Comment string `json:"comment" bson:"comment"`
	//Confirmer UserInternal `json:"confirmer" bson:"confirmer"`
}

type Participation struct {
	ID      string       `json:"id" bson:"_id"`
	User    UserInternal `json:"user" bson:"user"`
	EventID string       `json:"event_id" bson:"event_id"`
	Comment string       `json:"comment" bson:"comment"`
	Status  string       `json:"status" bson:"status"`
	Crew    CrewSimple   `json:"crew" bson:"crew"`
	//Confirmer UserInternal   `json:"confirmer" bson:"confirmer"`
	Modified vcago.Modified `json:"modified" bson:"modified"`
}

type ParticipationList []Participation

type ParticipationParam struct {
	ID string `param:"id"`
}

type ParticipationQuery struct {
	ID       []string `query:"id" qs:"id"`
	EventID  []string `query:"event_id" qs:"event_id"`
	Comment  []string `query:"comment" bson:"comment"`
	Status   []string `query:"status" bson:"status"`
	UserId   []string `query:"user_id" bson:"user_id"`
	CrewName []string `query:"crew_name" bson:"crew_name"`
	CrewId   []string `query:"crew_id" qs:"crew_id"`
}

func (i *ParticipationQuery) Match() (r *vcago.MongoMatch) {
	r = vcago.NewMongoMatch()
	r.StringList("_id", i.ID)
	r.StringList("event_id", i.EventID)
	r.StringList("status", i.Status)
	r.StringList("comment", i.Comment)
	r.StringList("user.id", i.UserId)
	r.StringList("crew.name", i.CrewName)
	r.StringList("crew.id", i.CrewId)
	return
}

func (i *ParticipationCreate) NewParticipation(token *AccessToken) *Participation {
	return &Participation{
		ID:       uuid.NewString(),
		User:     *token.UserInternal(),
		EventID:  i.EventID,
		Comment:  i.Comment,
		Status:   "requested",
		Crew:     *token.CrewSimple(),
		Modified: vcago.NewModified(),
	}
}
