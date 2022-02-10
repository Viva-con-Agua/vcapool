package vcapool

import "github.com/Viva-con-Agua/vcago"

type Crew struct {
	ID       string         `json:"id,omitempty" bson:"_id"`
	Name     string         `json:"name" bson:"name"`
	City     string         `json:"city" bson:"city"`
	Country  string         `json:"country" bson:"country"`
	Modified vcago.Modified `json:"modified" bson:"modified"`
}
