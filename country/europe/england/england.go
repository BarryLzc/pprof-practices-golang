package england

import (
	"log"
	"time"
)

type England struct {
}

func (e *England) Name() string {
	return "cat"
}

func (e *England) Exist() {
	e.Trade()
	e.Develop()
	e.War()
	e.Legislate()
	e.Integrate()
	e.Innovate()
}

func (e *England) Trade() {
	log.Println(e.Name(), "trade")
}

func (e *England) Develop() {
	log.Println(e.Name(), "develop")
}

func (e *England) War() {
	log.Println(e.Name(), "war")
}

func (e *England) Legislate() {
	log.Println(e.Name(), "legislate")

	<-time.After(time.Second)
}

func (e *England) Integrate() {
	log.Println(e.Name(), "integrate")
}

func (e *England) Innovate() {
	log.Println(e.Name(), "innovate")
}
