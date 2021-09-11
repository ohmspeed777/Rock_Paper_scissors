package controllers

import (
	"Rock_Paper_Scossors/model"
	"Rock_Paper_Scossors/service"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/operator"
	"go.mongodb.org/mongo-driver/bson"
)

type ScoreController interface {
	InsertScore(challenge *model.Challenge, op_act string) (*model.Score, error)
	// Ranking(ctx *gin.Context)
	ViewStat(ctx *gin.Context)
}

type scoreController struct{}

func NewScoreController() ScoreController {
	return &scoreController{}
}

func (controller *scoreController) InsertScore(challenge *model.Challenge, op_act string) (*model.Score, error) {

	score := service.SingleScore()
	score.Challenger = challenge.Challenger
	score.Opponent = challenge.Opponent
	score.Challenger_action = challenge.Action
	score.Opponent_action = op_act
	score.Results = service.DecideResult(challenge.Action, op_act)

	err := mgm.Coll(score).Create(score)
	if err != nil {
		return nil , err 
	}
	
	mgm.Coll(challenge).Delete(challenge)
	return score, nil
}

// func (controller *scoreController) Ranking(ctx *gin.Context) {
// 	return := 
// }

func (controller *scoreController) ViewStat(ctx *gin.Context) {
	id := ctx.MustGet("user").(string)
	if id == "" {
			return
	}

	opponent_id := ctx.Param("id")

	result := []model.Stat{}

	_ = mgm.Coll(&model.Score{}).SimpleAggregate(&result,
		bson.M{
			operator.Match: bson.M{
				"challenge": bson.M{operator.In: []string{id, opponent_id}},
				"opponent": bson.M{operator.In: []string{id, opponent_id}},
			},
		},
		bson.M{
			operator.Group: bson.M{
				"_id": bson.M{},
				"history": bson.M{operator.Push: "$results"},
			},
		},
	)

	fmt.Println(result)
}