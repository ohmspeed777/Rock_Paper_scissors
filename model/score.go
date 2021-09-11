package model

import (
	"github.com/kamva/mgm/v3"
)

type Score struct {
	mgm.DefaultModel 						`bson:",inline"`
	Challenger 				string 		`bson:"challenger"`
	Opponent 					string 		`bson:"opponent"`
	Challenger_action string 		`bson:"Challenger_action"`
	Opponent_action 	string 		`bson:"Opponent_action"`
	Results 					string 		`bson:"results"`
}


// type Ranking struct {
// 	mgm.DefaultModel 						`bson:",inline"`
// 	user string	 `bson:"_id"`
// }

type Stat struct {
	History []string `bson:"history"`
	LastMatch string `bson:"lastMatch"`
	Status string `bson:"status"`
}

