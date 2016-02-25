package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

type member struct {
	name string
	isBoke bool
}

func main() {
	wg := &sync.WaitGroup{}
	for _, m := range members() {
		wg.Add(1)
		go func(m member, wg *sync.WaitGroup) {
			defer wg.Done()
			if false == m.isBoke {
				machi(4)
				fmt.Printf("%s: 私がやります！\n", m.name)
			}
		}(m, wg)
	}
	wg.Wait()

	for _, m := range members() {
		wg.Add(1)
		go func(m member, wg *sync.WaitGroup) {
			defer wg.Done()
			if m.isBoke {
				machi(4)
				fmt.Printf("%s: じゃあ、私がやります・・・\n", m.name)
			}
		}(m, wg)
	}
	wg.Wait()

	for _, m := range members() {
		wg.Add(1)
		go func(m member, wg *sync.WaitGroup) {
			defer wg.Done()
			if false == m.isBoke {
				machi(2)
				fmt.Printf("%s: どうぞどうぞ！\n", m.name)
			}
		}(m, wg)
	}
	wg.Wait()
}

func machi(s int) {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Intn(s)) * time.Second)
}

func members() []member {
	return []member{
		member{"肥後", false},
		member{"寺門", false},
		member{"上島", true},
	}
}
