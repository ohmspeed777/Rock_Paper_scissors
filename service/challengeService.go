package service

import "Rock_Paper_Scossors/model"

func SingleChallenge() *model.Challenge{
	return&model.Challenge{}
}

func ContainsAction(actionList []string, action string) bool {
	for _, val := range actionList {
			if action == val {
					return true
			}
	}
	return false
}