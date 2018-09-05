package server

import (
	"net/http"
	"math/rand"
	"bytes"
	"encoding/json"
	"strings"
	"fmt"
)

const maxTryCount = 7
var currentCount int = maxTryCount

var codename string

type Clm struct {
	Bulls int
	Cows int
}

func CheckClmTry(w http.ResponseWriter, request *http.Request) {
	try := request.URL.Query().Get("try")

	if currentCount == maxTryCount {
		currentCount = 0
		codename = string(GenerateString(4))

		fmt.Printf("Generated: %s\n", codename)
	}

	currentCount++

	if len(codename) != len(try) {
		WriteResponse(w, -1, -1)
		return
	}

	bulls := 0
	cows := 0
	for i := 0; i < len(codename); i++ {
		if codename[i] == try[i] {
			bulls++
			continue
		}

		index := strings.Index(codename, string(try[i]))
		if index != -1 {
			cows++
		}
	}

	WriteResponse(w, bulls, cows)

	if bulls == len(codename) {
		w.Header().Add("token", authToken)
	}
}

func WriteResponse(w http.ResponseWriter, b int, c int) {
	res := new(Clm)
	res.Bulls = b
	res.Cows = c
	ba, _ := json.Marshal(res)
	w.Write(ba)
}

func GenerateString(n int) string {
	var buffer bytes.Buffer
	ur := new(UniqueRand)
	ur.generated = make(map[int]bool, 10)

	for i := 0; i < n; i++ {
		buffer.WriteString(string(ur.Int() + '0'))
	}

	return buffer.String()
}

type UniqueRand struct {
	generated map[int]bool
}

func (u *UniqueRand) Int() int {
	for {
		i := rand.Int() % 10
		if !u.generated[i] {
			u.generated[i] = true
			return i
		}
	}
}
