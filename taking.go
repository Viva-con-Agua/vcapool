package vcapool

type TakingEventSettings struct {
	TypeOfEvent string `json:"type_of_event" bson:"type_of_event"`
	EventID     string `json:"event_id" bson:"event_id"`
	EventName   string `json:"event_name" bson:"event_name"`
	TourID      string `json:"tour_id" bson:"tour_id"`
	TourName    string `json:"tour_name" bson:"tour_name"`
}

type Taking struct {
	ID            string      `json:"id" bson:"_id"`
	EventSettings string      `json:"event_settings" bson:"event_settings"`
	Sources       SourceList  `json:"sources" bson:"sources"`
	Crew          CrewSimple  `json:"crew" bson:"crew"`
	Author        InternalASP `json:"author" bson:"author"`
}

type TakingList []Taking
