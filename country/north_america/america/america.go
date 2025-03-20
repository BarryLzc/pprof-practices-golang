package america

import (
	"github.com/go-pprof-practices/constant"
	"log"
	"time"
)

type America struct {
	buffer     [][constant.Mi]byte
	slowBuffer [][constant.Mi]byte
}

func (*America) Name() string {
	return "america"
}

func (a *America) Exist() {
	a.Trade()
	a.Develop()
	a.War()
	a.Legislate()
	a.Lead()
	a.Rebel()
}

func (a *America) Trade() {
	log.Println(a.Name(), "trade")
}

func (a *America) Develop() {
	log.Println(a.Name(), "develop")
}

func (a *America) War() {
	log.Println(a.Name(), "war")
}

func (a *America) Legislate() {
	log.Println(a.Name(), "legislate")
	go func() {
		time.Sleep(time.Second * 30)
		max := constant.Gi
		for len(a.slowBuffer)*constant.Mi < max {
			a.slowBuffer = append(a.slowBuffer, [constant.Mi]byte{})
			time.Sleep(time.Millisecond * 500)
		}
	}()
}

func (a *America) Lead() {
	log.Println(a.Name(), "lead")
}

func (a *America) Rebel() {
	log.Println(a.Name(), "rebel")
	max := constant.Gi
	for len(a.buffer)*constant.Mi < max {
		a.buffer = append(a.buffer, [constant.Mi]byte{})
	}
}
