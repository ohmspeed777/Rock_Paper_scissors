package controllers

import (
	"Rock_Paper_Scossors/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

type ChallengeController interface {
	Challenge(ctx *gin.Context)
}

type challengeController struct {
	scoreController ScoreController
}

func NewChallengeController(score ScoreController) ChallengeController{
	return &challengeController{
		scoreController: score,
	}
}

func (controller *challengeController) Challenge(ctx *gin.Context) {

	// Check header from jwt
	id := ctx.MustGet("user").(string)
	if id == "" {
			return
	}

	// challenge form json api
	challenger := service.SingleChallenge()
	ctx.ShouldBindJSON(challenger)

	// check opponent still exist
	opponent := service.SingleUser()
	opponentErr := mgm.Coll(opponent).First(bson.M{"username": challenger.Opponent}, opponent)


	if opponent.ID.Hex() == id {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "You can not challenge yourself",
		})
		return
	}

	

	// not found opponent username in users collection
	if opponentErr != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"msg": "Not found that opponent",
		})
		return
	}

	// check action form api
	// Not found action is list
	// actionList := []string{"rock", "paper", "scissor"}
	if !service.ContainsAction(service.GetListAction(), challenger.Action) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "Your action incorrect",
		})
		return
	}

	// modified challenger before insert
	challenger.Challenger = id
	challenger.Opponent = opponent.ID.Hex()


	// check challenge is exist, if it exist  then cerate score 
	// check form opponent id
	challengeSuccess := service.SingleChallenge()
	_ = mgm.Coll(challengeSuccess).First(bson.M{"opponent": challenger.Challenger, "challenger": opponent.ID.Hex()}, challengeSuccess)


	// check already challenge if not create
	alreadyChallenge := service.SingleChallenge()
	mgm.Coll(alreadyChallenge).First(bson.M{"challenger": challenger.Challenger, "opponent": challenger.Opponent}, alreadyChallenge)

	if alreadyChallenge.Opponent != "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "You already challenge that one !!, Please try again.",
		})
		return
	}


	// challenge is not exist So crate a new challenge
	if challengeSuccess.Opponent == "" {
		challengerErr := mgm.Coll(challenger).Create(challenger)
		if challengerErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg": "Can not create a challenge",
			})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"msg": "success",
			"data": challenger,
		})
		return
	}

	score, scoreErr := controller.scoreController.InsertScore(challengeSuccess, challenger.Action)

	if scoreErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "Your action incorrect",
		})
		return
	}


	ctx.JSON(http.StatusCreated, gin.H{
		"msg": "success",
		"data": score,
	})

	
}