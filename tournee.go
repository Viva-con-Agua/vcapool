package vcapool

import (
	"encoding/json"

	"github.com/Viva-con-Agua/vcago"
	"github.com/google/uuid"
)

type Tour struct {
	ID        string         `json:"id" bson:"_id"`
	Name      string         `json:"name" bson:"name"`
	ArtistIDs []string       `json:"artist_ids" bson:"artist_ids"`
	Artists   ArtistList     `json:"artists" bson:"artists"`
	Events    EventList      `json:"events" bson:"events"`
	Creator   UserInternal   `json:"creator" bson:"creator"`
	Modified  vcago.Modified `json:"modified" bson:"modified"`
}

type TourList []Tour

type TourCreate struct {
	Name      string          `json:"name" bson:"name"`
	ArtistIDs []string        `json:"artist_ids" bson:"artist_ids"`
	Creator   UserInternal    `json:"creator" bson:"creator"`
	Events    EventCreateList `json:"events" bson:"events"`
}

type TourUpdate struct {
	ID        string         `json:"id" bson:"_id"`
	Name      string         `json:"name" bson:"name"`
	ArtistIDs []string       `json:"artist_ids" bson:"artist_ids"`
	Modified  vcago.Modified `json:"modified" bson:"modified"`
}

func (i *TourDatabase) Tour() (r *Tour) {
	bytes, _ := json.Marshal(&i)
	r = new(Tour)
	_ = json.Unmarshal(bytes, &r)
	return
}

type TourDatabase struct {
	ID        string         `json:"id" bson:"_id"`
	Name      string         `json:"name" bson:"name"`
	ArtistIDs []string       `json:"artist_ids" bson:"artist_ids"`
	Creator   UserInternal   `json:"creator" bson:"creator"`
	Modified  vcago.Modified `json:"modified" bson:"modified"`
}

func (i *TourCreate) Database(token *AccessToken) *TourDatabase {
	return &TourDatabase{
		ID:        uuid.NewString(),
		Name:      i.Name,
		ArtistIDs: i.ArtistIDs,
		Creator: UserInternal{
			UserID:      token.ID,
			FullName:    token.FullName,
			DisplayName: token.DisplayName,
			Email:       token.Email,
			Phone:       token.Profile.Phone,
		},
		Modified: vcago.NewModified(),
	}
}

type TourQuery struct {
	ID          string `query:"id" qs:"id"`
	Name        string `query:"name" qs:"name"`
	UpdatedTo   string `query:"updated_to" qs:"updated_to"`
	UpdatedFrom string `query:"updated_from" qs:"updated_from"`
	CreatedTo   string `query:"created_to" qs:"created_to"`
	CreatedFrom string `query:"created_from" qs:"created_from"`
}

func (i *TourQuery) Match() *vcago.MongoMatch {
	match := vcago.NewMongoMatch()
	match.EqualString("_id", i.ID)
	match.LikeString("name", i.Name)
	match.GteInt64("modified.updated", i.UpdatedFrom)
	match.GteInt64("modified.created", i.CreatedFrom)
	match.LteInt64("modified.updated", i.UpdatedTo)
	match.LteInt64("modified.created", i.CreatedTo)
	return match
}
