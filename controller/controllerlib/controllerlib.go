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
