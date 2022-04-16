package vcapool

import (
	"encoding/json"

	"github.com/Viva-con-Agua/vcago"
	"github.com/google/uuid"
)

type (
	TakingCreate struct {
		TypeOfEvent string     `json:"type_of_event" bson:"type_of_event"`
		EventID     string     `json:"event_id" bson:"event_id"`
		EventName   string     `json:"event_name" bson:"event_name"`
		TourID      string     `json:"tour_id" bson:"tour_id"`
		TourName    string     `json:"tour_name" bson:"tour_name"`
		Crew        CrewSimple `json:"crew" bson:"crew"`
	}

	TakingUpdate struct {
		ID          string       `json:"id" bson:"_id"`
		TypeOfEvent string       `json:"type_of_event" bson:"type_of_event"`
		EventID     string       `json:"event_id" bson:"event_id"`
		EventName   string       `json:"event_name" bson:"event_name"`
		TourID      string       `json:"tour_id" bson:"tour_id"`
		TourName    string       `json:"tour_name" bson:"tour_name"`
		Crew        CrewSimple   `json:"crew" bson:"crew"`
		Author      UserInternal `json:"author" bson:"author"`
	}

	TakingDatabase struct {
		ID          string         `json:"id" bson:"_id"`
		TypeOfEvent string         `json:"type_of_event" bson:"type_of_event"`
		EventID     string         `json:"event_id" bson:"event_id"`
		EventName   string         `json:"event_name" bson:"event_name"`
		TourID      string         `json:"tour_id" bson:"tour_id"`
		TourName    string         `json:"tour_name" bson:"tour_name"`
		Crew        CrewSimple     `json:"crew" bson:"crew"`
		Author      UserInternal   `json:"author" bson:"author"`
		Modified    vcago.Modified `json:"modified" bson:"modified"`
	}

	Taking struct {
		ID          string         `json:"id" bson:"_id"`
		TypeOfEvent string         `json:"type_of_event" bson:"type_of_event"`
		EventID     string         `json:"event_id" bson:"event_id"`
		EventName   string         `json:"event_name" bson:"event_name"`
		TourID      string         `json:"tour_id" bson:"tour_id"`
		TourName    string         `json:"tour_name" bson:"tour_name"`
		Crew        CrewSimple     `json:"crew" bson:"crew"`
		Sources     SourceList     `json:"sources" bson:"sources"`
		Author      UserInternal   `json:"author" bson:"author"`
		Modified    vcago.Modified `json:"modified" bson:"modified"`
	}

	TakingParam struct {
		ID string `param:"id" bson:"_id"`
	}
	TakingQuery struct {
		ID []string `query:"id" qs:"id"`
	}
	TakingList []Taking
)

func (i *TakingCreate) Database(token *AccessToken) (r *TakingDatabase) {
	bytes, _ := json.Marshal(&i)
	r = new(TakingDatabase)
	_ = json.Unmarshal(bytes, &r)
	r.ID = uuid.NewString()
	r.Author = *token.UserInternal()
	r.Crew = *token.CrewSimple()
	r.Modified = vcago.NewModified()
	return
}

func (i *TakingDatabase) Taking() (r *Taking) {
	bytes, _ := json.Marshal(&i)
	r = new(Taking)
	_ = json.Unmarshal(bytes, &r)
	return
}

func (i *TakingQuery) Match() (r *vcago.MongoMatch) {
	r = vcago.NewMongoMatch()
	r.StringList("_id", i.ID)
	return
}
