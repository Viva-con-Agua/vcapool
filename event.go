package vcapool

import "github.com/Viva-con-Agua/vcago"

type Event struct {
	ID                    string           `json:"id" bson:"_id"`
	Name                  string           `json:"name" bson:"name"`
	TypeOfEvent           string           `json:"type_of_event" bson:"type_of_event"`
	AdditionalInformation string           `json:"additional_information" bson:"additional_information"`
	Website               string           `json:"website" bson:"website"`
	TourneeID             string           `json:"tournee_id" bson:"tournee_id"`
	Location              Location         `json:"location" bson:"location"`
	Artists               ArtistList       `json:"artists" bson:"artists"`
	Organizer             Organizer        `json:"organizer" bson:"organizer"`
	StartAt               int64            `json:"start_at" bson:"start_at"`
	EndAt                 int64            `json:"end_at" bson:"end_at"`
	Crew                  CrewSimple       `json:"crew" bson:"crew"`
	EventASP              EventASP         `json:"event_asp" bson:"event_asp"`
	InteralASP            EventASP         `json:"interal_asp" bson:"internal_asp"`
	ExternalASP           EventASPExternal `json:"external_asp" bson:"external_asp"`
	Application           EventApplication `json:"application" bson:"application"`
	EventTools            EventTools       `json:"event_tools" bson:"event_tools"`
	CreatorID             string           `json:"creator_id" bson:"creator_id"`
	EventState            EventState       `json:"event_state" bson:"event_state"`
	Modified              string           `json:"modified" bson:"modified"`
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

type EventTools struct {
	Tools   []string `json:"tools" bson:"tools"`
	Special string   `json:"special" bson:"special"`
}

type EventASPExternal struct {
	FullName    string `json:"full_name" bson:"full_name"`
	DisplayName string `json:"display_name" bson:"display_name"`
	Email       string `json:"email" bson:"email"`
	Phone       string `json:"phone" bson:"phone"`
}

type EventState struct {
	State                string `json:"state" bson:"state"`
	CrewConfirmation     string `json:"crew_confirmation" bson:"crew_confirmation"`
	InternalConfirmation string `json:"internal_confirmation" bson:"internal_confirmation"`
	TakingID             string `json:"taking_id" bson:"taking_id"`
	DepositID            string `json:"deposit_id" bson:"deposit_id"`
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
