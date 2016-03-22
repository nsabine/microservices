package controllerlib

import (
	"math"
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
        for i := range GameState {
                for j := range GameState[i] {
			//fmt.Println("Comparing: ", me.MyName, GameState[i][j].MyName)
			if me.MyName == GameState[i][j].MyName {
				return GameState[i][j].XPos, GameState[i][j].YPos
			}
                }
        }
	return -1, -1
}

func WhereNearest(GameState [][]UpdateRequest, me UpdateRequest, mytype string) (UpdateRequest) {
	nearest := UpdateRequest{"Empty", mytype, 0, 0}
	nearestDistance := 100
	var distance int
        for i := range GameState {
                for j := range GameState[i] {
			if mytype == GameState[i][j].Type {
				distance = CalculateDistance(me, GameState[i][j])
				if distance < nearestDistance {
					nearestDistance = distance
					nearest = GameState[i][j]
				}
			}
                }
        }
	return nearest
}

func CalculateDistance(me UpdateRequest, them UpdateRequest) (int) {
	xDistance := int(math.Abs(float64(them.XPos - me.XPos)))
	yDistance := int(math.Abs(float64(them.YPos - me.YPos)))
	distance := math.Sqrt(math.Pow(float64(xDistance), 2) + math.Pow(float64(yDistance), 2))
	return int(distance)
}
