package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"server/actions"
	"strconv"
	"sync"
)

type Config struct {
	BaseUrl string
}

var manager *actions.GameManager
var m *sync.Mutex

func main() {
	m = &sync.Mutex{}
	configBytes, err := ioutil.ReadFile(`config.json`)
	if err != nil {
		log.Printf(err.Error())
		return
	}
	var config Config
	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		log.Printf(err.Error())
	}

	manager = actions.NewGameManager(config.BaseUrl)

	http.HandleFunc("/GameStatus", handleGameStatus)
	http.HandleFunc("/NewGame", handleNewGame)
	http.HandleFunc("/Roll", handleRoll)
	http.HandleFunc("/NewTurn", handleNewTurn)
	log.Printf("listening on %v", config.BaseUrl)
	log.Println(http.ListenAndServe(config.BaseUrl, nil))
}

func handleGameStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != `GET` {
		_, _ = w.Write(generateResponse("GameStatus endpoint requires GET method"))
		return
	}
	_, _ = w.Write(generateResponse(""))
}

func handleNewGame(w http.ResponseWriter, r *http.Request) {
	if r.Method != `POST` {
		_, _ = w.Write(generateResponse("NewGame endpoint requires POST method"))
		return
	}
	params, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		_, _ = w.Write(generateResponse(err.Error()))
		return
	}
	token, err := getToken(params)
	if err != nil {
		_, _ = w.Write(generateResponse(err.Error()))
		return
	}
	winningScore, err := getWinningScore(params)
	if err != nil {
		_, _ = w.Write(generateResponse(err.Error()))
		return
	}

	m.Lock()
	defer m.Unlock()
	action, ok := manager.ActionLinks[actions.NewGameAction]
	if !ok {
		_, _ = w.Write(generateResponse("NewGame is not valid at this time"))
		return
	}
	if action.Token != token {
		_, _ = w.Write(generateResponse("the provided token is invalid"))
		return
	}

	manager.ActiveActions.NewGameAction(winningScore)
	_, _ = w.Write(generateResponse(""))
}

func handleRoll(w http.ResponseWriter, r *http.Request) {
	if r.Method != `POST` {
		_, _ = w.Write(generateResponse("Roll endpoint requires POST method"))
		return
	}
	params, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		_, _ = w.Write(generateResponse(err.Error()))
		return
	}
	token, err := getToken(params)
	if err != nil {
		_, _ = w.Write(generateResponse(err.Error()))
		return
	}

	m.Lock()
	defer m.Unlock()
	action, ok := manager.ActionLinks[actions.RollAction]
	if !ok {
		_, _ = w.Write(generateResponse("Roll is not valid at this time"))
		return
	}
	if action.Token != token {
		_, _ = w.Write(generateResponse("the provided token is invalid"))
		return
	}

	manager.ActiveActions.RollAction()
	_, _ = w.Write(generateResponse(""))
}

func handleNewTurn(w http.ResponseWriter, r *http.Request) {
	if r.Method != `POST` {
		_, _ = w.Write(generateResponse("NewTurn endpoint requires POST method"))
		return
	}
	params, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		_, _ = w.Write(generateResponse(err.Error()))
		return
	}
	token, err := getToken(params)
	if err != nil {
		_, _ = w.Write(generateResponse(err.Error()))
		return
	}

	m.Lock()
	defer m.Unlock()
	action, ok := manager.ActionLinks[actions.NewTurnAction]
	if !ok {
		_, _ = w.Write(generateResponse("NewTurn is not valid at this time"))
		return
	}
	if action.Token != token {
		_, _ = w.Write(generateResponse("the provided token is invalid"))
		return
	}

	manager.ActiveActions.NewTurnAction()
	_, _ = w.Write(generateResponse(""))
}

type response struct {
	Error     string
	GameState *actions.GameManager
}

func generateResponse(error string) []byte {
	r := response{
		Error:     error,
		GameState: manager,
	}
	responseBytes, err := json.Marshal(r)
	if err != nil {
		errorBytes, _ := json.Marshal(map[string]string{"Error": err.Error()})
		return errorBytes
	}
	return responseBytes
}

func getToken(params url.Values) (int, error) {
	token, ok := params["token"]
	if !ok || len(token) == 0 {
		return 0, errors.New("token was not provided")
	}
	tokenInt, err := strconv.Atoi(token[0])
	if err != nil {
		return 0, errors.New("token was not a valid integer")
	}
	return tokenInt, nil
}

func getWinningScore(params url.Values) (int, error) {
	winningScore, ok := params["winningScore"]
	if !ok || len(winningScore) == 0 {
		return 0, errors.New("winningScore was not provided")
	}
	winningScoreInt, err := strconv.Atoi(winningScore[0])
	if err != nil {
		return 0, errors.New("winningScore was not a valid integer")
	}
	return winningScoreInt, nil
}
