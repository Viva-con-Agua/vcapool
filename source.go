package vcapool

import (
	"encoding/json"

	"github.com/Viva-con-Agua/vcago"
	"github.com/google/uuid"
)

type SourceCategory struct {
	Name  string `json:"name" bson:"name"`
	Norms string `json:"norms" bson:"norms"`
}

type SourceExternal struct {
	Location      string `json:"location" bson:"location"`
	ContactPerson string `json:"contact_person" bson:"contact_person"`
	Email         string `json:"email" bson:"email"`
	Address       string `json:"address" bson:"address"`
	Receipt       bool   `json:"receipt" bson:"receipt"`
	Purpose       string `json:"purpose" bson:"purpose"`
}

type SourceCreate struct {
	Category    SourceCategory `json:"category" bson:"category"`
	Description string         `json:"description" bson:"description"`
	HasExternal bool           `json:"has_external" bson:"has_external"`
	Purpose     string         `json:"purpose" bson:"purpose"`
	Confirmed   bool           `json:"confirmed" bson:"confirmed"`
	ConfirmDate int64          `json:"confirm_date" bson:"confirm_date"`
	External    SourceExternal `json:"external" bson:"external"`
	Creator     UserInternal   `json:"creator" bson:"creator"`
	Confirmer   UserInternal   `json:"confirmer" bson:"confirmer"`
}

func (i *SourceCreate) Source(id string) (r *Source) {
	bytes, _ := json.Marshal(&i)
	r = new(Source)
	_ = json.Unmarshal(bytes, &r)
	r.ID = uuid.NewString()
	r.Modified = vcago.NewModified()
	r.TakingID = id
	return
}

type SourceCreateList []SourceCreate

func (i *SourceCreateList) SourceList(id string) (r *SourceList) {
	r = new(SourceList)
	for n, _ := range *i {
		temp := (*i)[n].Source(id)
		*r = append(*r, *temp)
	}
	return
}

type Source struct {
	ID          string         `json:"id" bson:"_id"`
	Category    SourceCategory `json:"category" bson:"category"`
	Description string         `json:"description" bson:"description"`
	HasExternal bool           `json:"has_external" bson:"has_external"`
	Purpose     string         `json:"purpose" bson:"purpose"`
	Confirmed   bool           `json:"confirmed" bson:"confirmed"`
	ConfirmDate int64          `json:"confirm_date" bson:"confirm_date"`
	External    SourceExternal `json:"external" bson:"external"`
	Creator     UserInternal   `json:"creator" bson:"creator"`
	Confirmer   UserInternal   `json:"confirmer" bson:"confirmer"`
	TakingID    string         `json:"taking_id" bson:"taking_id"`
	Modified    vcago.Modified `json:"modified" bson:"modified"`
}

type SourceList []Source

func (i *SourceList) Insert() (r []interface{}) {
	for n, _ := range *i {
		r = append(r, (*i)[n])
	}
	return
}
