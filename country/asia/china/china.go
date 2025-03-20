package china

import (
	"github.com/go-pprof-practices/constant"
	"log"
)

type China struct {
}

func (c *China) Name() string {
	return "china"
}

func (c *China) Exist() {
	c.Trade()
	c.Develop()
	c.War()
	c.Legislate()
}

func (c *China) Trade() {
	log.Println(c.Name(), "trade")
}

func (c *China) Develop() {
	log.Println(c.Name(), "develop")
}

func (c *China) War() {
	log.Println(c.Name(), "war")
}

func (c *China) Legislate() {
	log.Println(c.Name(), "legislate")
}

func (c *China) Expand() {
	log.Println(c.Name(), "expand")
	_ = make([]byte, 16*constant.Mi)
}

func (c *China) PreserveCulture() {
	log.Println(c.Name(), "preserveCulture")
}
