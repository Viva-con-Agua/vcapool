package vcapool

type Location struct {
	Street      string   `json:"street" bson:"street"`
	City        string   `json:"city" bson:"city"`
	Country     string   `json:"country" bson:"country"`
	CountryCode string   `json:"country_code" bson:"country_code"`
	Number      string   `json:"number" bson:"number"`
	Position    Position `json:"position" bson:"position"`
	PlaceID     string   `json:"place_id" bson:"place_id"`
	Sublocality string   `json:"sublocality" bson:"sublocality"`
}

type Position struct {
	Lat float64 `json:"lat" bson:"lat"`
	Lng float64 `json:"lng" bson:"lng"`
}

type Meeting struct {
	MeetAt   int64    `json:"meet_at" bson:"meet_at"`
	Location Location `json:"location" bson:"location"`
	Type     string   `json:"type" bson:"type"`
	Title    string   `json:"title" bson:"title"`
	Subtitle string   `json:"subtitle" bson:"subtitle"`
}

type MeetingList []Meeting
