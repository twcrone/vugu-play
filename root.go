package main

import (
	"encoding/json"
	"github.com/vugu/vugu"
	"log"
	"net/http"
)

type Root struct {
	Resp      resp `vugu:"data"`
	IsLoading bool `vugu:"data"`
}

type resp struct {
	Data []note `json:"data"`
}

type note struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
}

func (c *Root) HandleClick(event vugu.DOMEvent) {
	ee := event.EventEnv()

	go func() {

		ee.Lock()
		c.IsLoading = true
		ee.UnlockRender()

		res, err := http.Get("https://go-love-notes.herokuapp.com/notes")
		if err != nil {
			log.Printf("Error fetch()ing: %v", err)
			return
		}
		defer res.Body.Close()

		resp := resp{}
		err = json.NewDecoder(res.Body).Decode(&resp)
		if err != nil {
			log.Printf("Error JSON decoding: %v", err)
			return
		}

		ee.Lock()
		defer ee.UnlockRender()
		c.Resp = resp
		c.IsLoading = false

	}()
}
