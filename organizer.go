package vcapool

type Organizer struct {
	ID              string `json:"id" bson:"_id"`
	Name            string `json:"name" bson:"name"`
	TypeOfOrganizer string `json:"type_of_organizer" bson:"type_of_organizer"`
	Modified        string `json:"modified" bson:"modified"`
}
