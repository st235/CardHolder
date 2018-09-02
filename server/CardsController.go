package server

import (
	"net/http"
	"encoding/json"
	"card/core"
	"strconv"
)

func choose(ss []core.Card, size int, page int,
	test func(card core.Card) bool) (ret []core.Card) {
	if len(ss) < size*page {
		return
	}

	maxSize := size*page + page

	if maxSize > len(ss) {
		maxSize = len(ss)
	}

	for _, s := range ss[size*page:maxSize] {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func GetCards(w http.ResponseWriter, request *http.Request) {
	page, pe := strconv.ParseInt(request.URL.Query().Get("page"), 10, 32)
	size, se := strconv.ParseInt(request.URL.Query().Get("size"), 10, 32)

	if pe != nil || se != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"error\": 0}"))
		return
	}

	cards := choose(core.Cards.CardsArray, int(page), int(size), func(card core.Card)bool {
		return true
	})

	if len(cards) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("{\"error\": 404}"))
		return
	}

	b, _ := json.Marshal(cards)
	w.Write(b)
}

func GetCard(w http.ResponseWriter, request *http.Request) {
	id, pe := strconv.ParseInt(request.URL.Query().Get("id"), 10, 32)

	if pe != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"error\": 0}"))
		return
	}

	cards := choose(core.Cards.CardsArray, 0, len(core.Cards.CardsArray), func(card core.Card)bool {
		return card.Id == int(id)
	})

	if len(cards) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("{\"error\": 404}"))
		return
	}

	b, _ := json.Marshal(cards[0])
	w.Write(b)
}
