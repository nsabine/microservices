package controllerlib

import (
	"fmt"
)

type UpdateRequest struct {
	MyName 	string
	Type	string
	XPos	int
	YPos	int
}

var XSize = 25
var YSize = 25


func GetGridCode(r UpdateRequest) string {
	switch r.Type {
		case "Empty":
			return " "
		case "Robber":
			return "R"
		case "Cop":
			return "C"
		default:
			return "?"
	}
}

func WhereAmI(GameState [][]UpdateRequest, me UpdateRequest) (int, int) {
	fmt.Println("In controllerlib.WhereAmI")
        for i := range GameState {
                for j := range GameState[i] {
			fmt.Println("Comparing: ", me.MyName, GameState[i][jj].MyName)
			if me.MyName == GameState[i][j].MyName {
				return GameState[i][j].XPos, GameState[i][j].YPos
			}
                }
        }
	return -1, -1
}
