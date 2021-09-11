package model

import "github.com/kamva/mgm/v3"

type Challenge struct {
	mgm.DefaultModel 						`bson:",inline"`
	Challenger 				string 		`bson:"challenger"`
	Opponent 					string 		`bson:"opponent" json:"opponent"`
	Action 						string 		`bson:"action" json:"action"`
}

