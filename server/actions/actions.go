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
		ActionLinks:   map[ActionType]ActionInfo{},
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
	Token  int
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
	ActionLinks   map[ActionType]ActionInfo
	randGenerator *rand.Rand
	baseUrl       string
}

func (m *GameManager) generateActions() {
	m.ActionLinks = map[ActionType]ActionInfo{}
	m.ActiveActions = Actions{}
	m.generateNewGameAction()
	if m.ActiveGame == nil {
		return
	}
	if len(m.ActiveGame.ActiveTurn.LastRoll) == 0 {
		m.generateRollAction()
		return
	}
	if len(m.ActiveGame.ActiveTurn.LastRoll) > 0 && m.ActiveGame.ActiveTurn.Score == 0 {
		m.generateNewTurnAction()
		return
	}
	m.generateRollAction()
	m.generateNewTurnAction()
}

func (m *GameManager) generateNewGameAction() {
	id := m.randGenerator.Int()
	m.ActiveActions.NewGameAction = func(winningScore int) {
		m.ActiveGame = game.NewGame(winningScore)
		m.generateActions()
	}
	m.ActionLinks[NewGameAction] = ActionInfo{
		Token:  id,
		Url:    fmt.Sprintf(`%v/NewGame?token=%v&winningScore={WinningScore}`, m.baseUrl, id),
		Method: "POST",
	}
}

func (m *GameManager) generateRollAction() {
	id := m.randGenerator.Int()
	m.ActiveActions.RollAction = func() {
		m.ActiveGame.Roll()
		m.generateActions()
	}
	m.ActionLinks[RollAction] = ActionInfo{
		Token:  id,
		Url:    fmt.Sprintf(`%v/Roll?token=%v`, m.baseUrl, id),
		Method: "POST",
	}
}

func (m *GameManager) generateNewTurnAction() {
	id := m.randGenerator.Int()
	m.ActiveActions.NewTurnAction = func() {
		m.ActiveGame.NewTurn()
		m.generateActions()
	}
	m.ActionLinks[NewTurnAction] = ActionInfo{
		Token:  id,
		Url:    fmt.Sprintf(`%v/NewTurn?token=%v`, m.baseUrl, id),
		Method: "POST",
	}
}
