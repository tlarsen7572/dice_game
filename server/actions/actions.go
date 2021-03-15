package actions

import (
	"fmt"
	"math/rand"
	"server/game"
	"strings"
	"time"
)

func NewGameManager(baseUrl string) *GameManager {
	baseUrl = strings.TrimSuffix(baseUrl, `/`)
	manager := &GameManager{
		ActiveGame:    nil,
		ActiveActions: Actions{},
		ActionLinks:   map[int]ActionInfo{},
		randGenerator: rand.New(rand.NewSource(time.Now().UTC().UnixNano())),
		baseUrl:       baseUrl,
	}
	manager.generateActions()
	return manager
}

type ActionType int

const (
	NewGameAction ActionType = 0
	RollAction    ActionType = 1
	NewTurnAction ActionType = 2
)

type ActionInfo struct {
	Type   ActionType
	Url    string
	Method string
}

type Actions struct {
	NewGameAction func(int)
	RollAction    func()
	NewTurnAction func()
}

type GameManager struct {
	ActiveGame    *game.Game
	ActiveActions Actions `json:"-"`
	ActionLinks   map[int]ActionInfo
	randGenerator *rand.Rand
	baseUrl       string
}

func (m *GameManager) generateActions() {
	m.ActionLinks = map[int]ActionInfo{}
	m.ActiveActions = Actions{}
	if m.ActiveGame == nil {
		m.generateNewGameAction()
		return
	}
}

func (m *GameManager) generateNewGameAction() {
	id := m.randGenerator.Int()
	m.ActiveActions.NewGameAction = func(winningScore int) {
		m.ActiveGame = game.NewGame(winningScore)
		m.generateActions()
	}
	m.ActionLinks[id] = ActionInfo{
		Type:   NewGameAction,
		Url:    fmt.Sprintf(`%v/%v?winningScore={WinningScore}`, m.baseUrl, id),
		Method: "POST",
	}
}
