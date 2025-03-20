package germany

import "log"

type Germany struct {
}

func (g *Germany) Name() string {
	return "tiger"
}

func (g *Germany) Exist() {
	g.Trade()
	g.Develop()
	g.War()
	g.Legislate()
	g.Integrate()
	g.Innovate()
}

func (g *Germany) Trade() {
	log.Println(g.Name(), "trade")
	loop := 10000000000
	for i := 0; i < loop; i++ {
		// do nothing
	}
}

func (g *Germany) Develop() {
	log.Println(g.Name(), "develop")
}

func (g *Germany) War() {
	log.Println(g.Name(), "war")
}

func (g *Germany) Legislate() {
	log.Println(g.Name(), "legislate")
}

func (g *Germany) Integrate() {
	log.Println(g.Name(), "integrate")
}

func (g *Germany) Innovate() {
	log.Println(g.Name(), "innovate")
}
