package service

import "Rock_Paper_Scossors/model"

func SingleScore() *model.Score{
	return &model.Score{}
}

func GetStat() *model.Stat{
	return &model.Stat{}
}

func DecideResult(action1, action2 string) string {
	index1 := findIndexInListAction(action1)
	index2 := findIndexInListAction(action2)
	return matchResultWithNumber(index1 - index2)
}


func matchResultWithNumber(num int) string {
	if num == -2 || num == 1 {
		return "win"
	} else if num == -1 || num == 2 {
		return "lose"
	} else if num == 0 {
		return "draw"
	}

	return "unknown"
}

func findIndexInListAction(action string) int {
	listAction := GetListAction()
	for index, act := range listAction {
		if act == action {
			return index
		}
	}
	// default value
	return 0
}

func GetListAction() []string{
	return []string{"rock", "paper", "scissor"}
}