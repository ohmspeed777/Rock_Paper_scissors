package model

import "github.com/kamva/mgm/v3"

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Username             string  `json:"username" bson:"username"`
}
