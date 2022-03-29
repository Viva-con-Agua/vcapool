package vcapool

import "github.com/Viva-con-Agua/vcago"

type Artist struct {
	ID           string         `json:"id" bson:"_id"`
	Name         string         `json:"name" bson:"name"`
	TypeOfArtist string         `json:"type_of_artist" bson:"type_of_artist"`
	Modified     vcago.Modified `json:"modified" bson:"modified"`
}

type ArtistList []Artist

type ArtistQuery struct {
	ID          string `query:"id" qs:"id"`
	Name        string `query:"name" qs:"name"`
	UpdatedTo   string `query:"updated_to" qs:"updated_to"`
	UpdatedFrom string `query:"updated_from" qs:"updated_from"`
	CreatedTo   string `query:"created_to" qs:"created_to"`
	CreatedFrom string `query:"created_from" qs:"created_from"`
}

func (i *ArtistQuery) Match() *vcago.MongoMatch {
	match := vcago.NewMongoMatch()
	match.EqualString("id", i.ID)
	match.LikeString("name", i.Name)
	match.GteInt64("modified.updated", i.UpdatedFrom)
	match.GteInt64("modified.created", i.CreatedFrom)
	match.LteInt64("modified.updated", i.UpdatedTo)
	match.LteInt64("modified.created", i.CreatedTo)
	return match
}
