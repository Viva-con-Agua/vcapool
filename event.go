package vcapool

import "github.com/Viva-con-Agua/vcago"

type Event struct {
	ID           string                `json:"id" bson:"_id"`
	Name         string                `json:"name" bson:"name"`
	TypeOfEvent  string                `json:"type_of_event" bson:"type_of_event"`
	TourneeID    string                `json:"tournee_id" bson:"tournee_id"`
	Location     Location              `json:"location" bson:"location"`
	Artists      ArtistList            `json:"artists" bson:"artists"`
	Organizer    Organizer             `json:"organizer" bson:"organizer"`
	StartAt      int64                 `json:"start_at" bson:"start_at"`
	EndAt        int64                 `json:"end_at" bson:"end_at"`
	Application  EventApplication      `json:"application" bson:"application"`
	Organisation EventOrganisationList `json:"organistion"`
	Tools        []string              `json:"tools" bson:"tools"`
	CreatorID    string                `json:"creator_id" bson:"creator_id"`
	Modified     string                `json:"modified" bson:"modified"`
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

type EventLocation struct {
	ID       string   `json:"id" bson:"_id"`
	Name     string   `json:"name" bson:"name"`
	Location Location `json:"location" bson:"location"`
}

type EventDate struct {
	StartAt int64 `json:"start_at" bson:"start_at"`
	EndAt   int64 `json:"end_at" bson:"end_at"`
}

type EventOrganisation struct {
	ID         string          `json:"id" bson:"_id"`
	CrewID     string          `json:"crew_id" bson:"crew_id"`
	CrewName   string          `json:"crew_name" bson:"crew_name"`
	Meetings   MeetingList     `json:"meetings" bson:"meetings"`
	EventASP   EventASP        `json:"event_asp" bson:"event_asp"`
	Contingent EventContingent `json:"contingent" bson:"contingent"`
	EventID    string          `json:"event_id" bson:"event_id"`
	Modified   vcago.Modified  `json:"modified" bson:"modified"`
}

type EventOrganisationList []EventOrganisation

type EventApplication struct {
	StartDate      int64 `json:"start_date" bson:"start_date"`
	EndDate        int64 `json:"end_date" bson:"end_date"`
	SupporterCount int   `json:"supporter_count" bson:"supporter_count"`
}

type EventASP struct {
	UserID      string `json:"user_id" bson:"user_id"`
	FullName    string `json:"full_name" bson:"full_name"`
	DisplayName string `json:"display_name" bson:"display_name"`
	Email       string `json:"email" bson:"email"`
	Phone       string `json:"phone" bson:"phone"`
}

type EventContingent struct {
	CrewID string `json:"crew_id" bson:"crew_id"`
	Count  int    `json:"count" bson:"count"`
}

type EventContingentList []EventContingent

func (i *EventQuery) Match() *vcago.MongoMatch {
	match := vcago.NewMongoMatch()
	match.EqualString("id", i.ID)
	match.LikeString("name", i.Name)
	match.GteInt64("modified.updated", i.UpdatedFrom)
	match.GteInt64("modified.created", i.CreatedFrom)
	match.LteInt64("modified.updated", i.UpdatedTo)
	match.LteInt64("modified.created", i.CreatedTo)
	return match
}
