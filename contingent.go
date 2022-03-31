package vcapool

import "github.com/Viva-con-Agua/vcago"

type Contingent struct {
	ID       string         `json:"id" bson:"_id"`
	CrewID   string         `json:"crew_id" bson:"crew_id"`
	CrewName string         `json:"crew_name" bson:"crew_name"`
	Meetings MeetingList    `json:"meetings" bson:"meetings"`
	EventID  string         `json:"event_id" bson:"event_id"`
	Count    int            `json:"count" bson:"count"`
	Modified vcago.Modified `json:"modified" bson:"modified"`
}

type ContingentList []Contingent

type ContingentQuery struct {
	ID          string `query:"id" qs:"id"`
	CrewID      string `query:"crew_id" qs:"crew_id"`
	CrewName    string `query:"crew_name" qs:"crew_name"`
	UpdatedTo   string `query:"updated_to" qs:"updated_to"`
	UpdatedFrom string `query:"updated_from" qs:"updated_from"`
	CreatedTo   string `query:"created_to" qs:"created_to"`
	CreatedFrom string `query:"created_from" qs:"created_from"`
}

func (i *ContingentQuery) Match() *vcago.MongoMatch {
	match := vcago.NewMongoMatch()
	match.EqualString("_id", i.ID)
	match.EqualString("crew_id", i.CrewID)
	match.LikeString("crew_name", i.CrewName)
	match.GteInt64("modified.updated", i.UpdatedFrom)
	match.GteInt64("modified.created", i.CreatedFrom)
	match.LteInt64("modified.updated", i.UpdatedTo)
	match.LteInt64("modified.created", i.CreatedTo)
	return match
}
