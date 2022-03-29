package vcapool

import "github.com/Viva-con-Agua/vcago"

type (
	Crew struct {
		ID       string         `json:"id,omitempty" bson:"_id"`
		Name     string         `json:"name" bson:"name"`
		Email    string         `json:"email" bson:"email"`
		Cities   []City         `json:"cities" bson:"cities"`
		Modified vcago.Modified `json:"modified" bson:"modified"`
	}
	City struct {
		City        string   `json:"city" bson:"city"`
		Country     string   `json:"country" bson:"country"`
		CountryCode string   `json:"country_code" bson:"country_code"`
		PlaceID     string   `json:"place_id" bson:"place_id"`
		Position    Position `json:"position" bson:"position"`
	}
	CrewList  []Crew
	CrewQuery struct {
		ID     string   `query:"id,omitempty" qs:"id"`
		Name   string   `query:"name" qs:"name"`
		Email  string   `query:"email" qs:"email"`
		Cities []string `query:"cities" qs:"cities"`
	}
)
