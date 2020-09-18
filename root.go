package main

import (
	"github.com/vugu/vugu"
	"log"
	_ "time"
)

type Root struct {
	Show     bool `vugu:"data"`
	ShowWasm bool `vugu:"data"`
	ShowGo   bool `vugu:"data"`
	ShowVugu bool `vugu:"data"`
}

func (c *Root) Toggle(e vugu.DOMEvent) {
	c.Show = !c.Show
	log.Printf("Toggled! Show is now %t", c.Show)
}
