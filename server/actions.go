package server

type Action struct {
	Name     string
	Function func() map[int]Action
}

type ActionFunc func() map[int]Action

func GenerateActions(game *Game) map[int]Action {
	return nil
}
