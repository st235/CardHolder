package server

import (
	"encoding/json"
	"net/http"
	"time"
	"math/rand"
	"card/utils"
	"fmt"
)

var authToken string

type Error struct {
	Code int
}

func Register() {
	rand.Seed(time.Now().UnixNano())
	authToken = utils.RandStringRunes(32)

	fmt.Printf("token: %s", authToken)

	NewServer().
		AddRoute("/cards/get", GetCards).
		AddRoute("/card/get", GetCard).
		AddRoute("/auth", CheckClmTry).
		Listen(determineListenAddress())
}

func CheckAuth(token string) bool {
	return token == authToken
}

func WriteError(code int, w http.ResponseWriter) {
	error := new(Error)
	error.Code = code

	jerror, e := json.Marshal(error)

	if e != nil {
	}

	w.WriteHeader(code)
	w.Write(jerror)
}
