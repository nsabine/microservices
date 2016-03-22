package controllerlib

type UpdateRequest struct {
	MyName 	string
	Type	string
	XPos	int
	YPos	int
}

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
			if me.MyName == GameState[i][j].MyName {
				return GameState[i][j].XPos, GameState[i][j].YPos
			}
                }
        }
	return -1, -1
}
