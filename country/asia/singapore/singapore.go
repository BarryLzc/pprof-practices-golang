package singapore

import (
	"log"
	"sync"
	"time"
)

type Singapore struct {
	buffer []byte
}

func (s *Singapore) Name() string {
	return "singapore"
}

func (s *Singapore) Exist() {

}

func (s *Singapore) Trade() {
	log.Println(s.Name(), "trade")
}

func (s *Singapore) Develop() {
	log.Println(s.Name(), "develop")
	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(30 * time.Second)
		}()
	}
}

func (s *Singapore) War() {
	log.Println(s.Name(), "war")
}

func (s *Singapore) Legislate() {
	log.Println(s.Name(), "legislate")
}

func (s *Singapore) PreserveCulture() {
	log.Println(s.Name(), "preserveCulture")
}

func (s *Singapore) Expand() {
	log.Println(s.Name(), "expand")

	m := &sync.Mutex{}
	m.Lock()
	go func() {
		time.Sleep(time.Second)
		m.Unlock()
	}()
	m.Lock()
}
