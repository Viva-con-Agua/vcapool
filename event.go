package vcapool

import (
	"encoding/json"

	"github.com/Viva-con-Agua/vcago"
	"github.com/google/uuid"
)

//EventApplication represents the application type of an event.
type EventApplication struct {
	StartDate      int64 `json:"start_date" bson:"start_date"`
	EndDate        int64 `json:"end_date" bson:"end_date"`
	SupporterCount int   `json:"supporter_count" bson:"supporter_count"`
}

//InternalASP represents the model for an asp with pool account. Used for event_asp and internal_asp.
type InternalASP struct {
	UserID      string `json:"user_id" bson:"user_id"`
	FullName    string `json:"full_name" bson:"full_name"`
	DisplayName string `json:"display_name" bson:"display_name"`
	Email       string `json:"email" bson:"email"`
	Phone       string `json:"phone" bson:"phone"`
}

//EventTools contains tools
type EventTools struct {
	Tools   []string `json:"tools" bson:"tools"`
	Special string   `json:"special" bson:"special"`
}

//ExternalASP represents an external asp without user_id.
type ExternalASP struct {
	FullName    string `json:"full_name" bson:"full_name"`
	DisplayName string `json:"display_name" bson:"display_name"`
	Email       string `json:"email" bson:"email"`
	Phone       string `json:"phone" bson:"phone"`
}

//EventState represents the state of an event.
type EventState struct {
	State                string `json:"state" bson:"state"`
	CrewConfirmation     string `json:"crew_confirmation" bson:"crew_confirmation"`
	InternalConfirmation string `json:"internal_confirmation" bson:"internal_confirmation"`
	TakingID             string `json:"taking_id" bson:"taking_id"`
	DepositID            string `json:"deposit_id" bson:"deposit_id"`
}

type Event struct {
	ID                    string           `json:"id" bson:"_id"`
	Name                  string           `json:"name" bson:"name"`
	TypeOfEvent           string           `json:"type_of_event" bson:"type_of_event"`
	AdditionalInformation string           `json:"additional_information" bson:"additional_information"`
	Website               string           `json:"website" bson:"website"`
	TourID                string           `json:"tour_id" bson:"tour_id"`
	Location              Location         `json:"location" bson:"location"`
	ArtistIDs             []string         `json:"artist_ids" bson:"artist_ids"`
	Artists               ArtistList       `json:"artists" bson:"artists"`
	Organizer             Organizer        `json:"organizer" bson:"organizer"`
	StartAt               int64            `json:"start_at" bson:"start_at"`
	EndAt                 int64            `json:"end_at" bson:"end_at"`
	Crew                  CrewSimple       `json:"crew" bson:"crew"`
	EventASP              InternalASP      `json:"event_asp" bson:"event_asp"`
	InteralASP            InternalASP      `json:"internal_asp" bson:"internal_asp"`
	ExternalASP           ExternalASP      `json:"external_asp" bson:"external_asp"`
	Application           EventApplication `json:"application" bson:"application"`
	EventTools            EventTools       `json:"event_tools" bson:"event_tools"`
	Creator               InternalASP      `json:"creator" bson:"creator"`
	EventState            EventState       `json:"event_state" bson:"event_state"`
	Modified              vcago.Modified   `json:"modified" bson:"modified"`
}

type EventCreate struct {
	Name                  string           `json:"name" bson:"name"`
	TypeOfEvent           string           `json:"type_of_event" bson:"type_of_event"`
	AdditionalInformation string           `json:"additional_information" bson:"additional_information"`
	Website               string           `json:"website" bson:"website"`
	TourID                string           `json:"tour_id" bson:"tour_id"`
	Location              Location         `json:"location" bson:"location"`
	ArtistIDs             []string         `json:"artist_ids" bson:"artist_ids"`
	Organizer             Organizer        `json:"organizer" bson:"organizer"`
	StartAt               int64            `json:"start_at" bson:"start_at"`
	EndAt                 int64            `json:"end_at" bson:"end_at"`
	Crew                  CrewSimple       `json:"crew" bson:"crew"`
	EventASP              InternalASP      `json:"event_asp" bson:"event_asp"`
	InternalASP           InternalASP      `json:"internal_asp" bson:"internal_asp"`
	ExternalASP           ExternalASP      `json:"external_asp" bson:"external_asp"`
	Application           EventApplication `json:"application" bson:"application"`
	EventTools            EventTools       `json:"event_tools" bson:"event_tools"`
	CreatorID             string           `json:"creator_id" bson:"creator_id"`
}

type EventCreateList []EventCreate

func (i *EventCreateList) Database(tourID string) (r *EventDatabaseList) {
	r = new(EventDatabaseList)
	for n, _ := range *i {
		temp := (*i)[n].Database()
		temp.TourID = tourID
		*r = append(*r, *temp)
	}
	return
}

func (i *EventCreate) Database() (r *EventDatabase) {
	bytes, _ := json.Marshal(&i)
	r = new(EventDatabase)
	_ = json.Unmarshal(bytes, &r)
	r.ID = uuid.NewString()
	r.EventState = EventState{
		State: "created",
	}
	r.Modified = vcago.NewModified()
	return
}

func (i *Event) Database(r *EventDatabase) {
	bytes, _ := json.Marshal(&i)
	r = new(EventDatabase)
	_ = json.Unmarshal(bytes, &r)
	return
}

func (i *EventDatabase) Event() (r *Event) {
	bytes, _ := json.Marshal(&i)
	r = new(Event)
	_ = json.Unmarshal(bytes, &r)
	return
}

type EventDatabase struct {
	ID                    string           `json:"id" bson:"_id"`
	Name                  string           `json:"name" bson:"name"`
	TypeOfEvent           string           `json:"type_of_event" bson:"type_of_event"`
	AdditionalInformation string           `json:"additional_information" bson:"additional_information"`
	Website               string           `json:"website" bson:"website"`
	TourID                string           `json:"tour_id" bson:"tour_id"`
	Location              Location         `json:"location" bson:"location"`
	ArtistIDs             []string         `json:"artist_ids" bson:"artist_ids"`
	Organizer             Organizer        `json:"organizer" bson:"organizer"`
	StartAt               int64            `json:"start_at" bson:"start_at"`
	EndAt                 int64            `json:"end_at" bson:"end_at"`
	Crew                  CrewSimple       `json:"crew" bson:"crew"`
	EventASP              InternalASP      `json:"event_asp" bson:"event_asp"`
	InternalASP           InternalASP      `json:"internal_asp" bson:"internal_asp"`
	ExternalASP           ExternalASP      `json:"external_asp" bson:"external_asp"`
	Application           EventApplication `json:"application" bson:"application"`
	EventTools            EventTools       `json:"event_tools" bson:"event_tools"`
	CreatorID             string           `json:"creator_id" bson:"creator_id"`
	EventState            EventState       `json:"event_state" bson:"event_state"`
	Modified              vcago.Modified   `json:"modified" bson:"modified"`
}

type EventDatabaseList []EventDatabase

func (i *EventDatabaseList) Insert() (r []interface{}) {
	for n, _ := range *i {
		r = append(r, (*i)[n])
	}
	return
}

type EventList []Event

type EventQuery struct {
	ID          string `query:"id" qs:"id"`
	Name        string `query:"name" qs:"name"`
	UpdatedTo   string `query:"updated_to" qs:"updated_to"`
	UpdatedFrom string `query:"updated_from" qs:"updated_from"`
	CreatedTo   string `query:"created_to" qs:"created_to"`
	CreatedFrom string `query:"created_from" qs:"created_from"`
}

func (i *EventQuery) Match() *vcago.MongoMatch {
	match := vcago.NewMongoMatch()
	match.EqualString("_id", i.ID)
	match.LikeString("name", i.Name)
	match.GteInt64("modified.updated", i.UpdatedFrom)
	match.GteInt64("modified.created", i.CreatedFrom)
	match.LteInt64("modified.updated", i.UpdatedTo)
	match.LteInt64("modified.created", i.CreatedTo)
	return match
}
