package vcapool

//UserDatabase represents the database model of the users collection.
type (
	UserSimple struct {
		ID          string `json:"id,omitempty" bson:"_id"`
		Email       string `json:"email" bson:"email" validate:"required,email"`
		FirstName   string `bson:"first_name" json:"first_name" validate:"required"`
		LastName    string `bson:"last_name" json:"last_name" validate:"required"`
		FullName    string `bson:"full_name" json:"full_name"`
		DisplayName string `bson:"display_name" json:"display_name"`
	}
	UserInternal struct {
		ID          string  `json:"id,omitempty" bson:"_id"`
		Email       string  `json:"email" bson:"email"`
		FirstName   string  `bson:"first_name" json:"first_name"`
		LastName    string  `bson:"last_name" json:"last_name"`
		FullName    string  `bson:"full_name" json:"full_name"`
		DisplayName string  `bson:"display_name" json:"display_name"`
		Profile     Profile `json:"profile" bson:"profile"`
	}
)
