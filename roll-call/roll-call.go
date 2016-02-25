package main

import (
	"fmt"
	"time"
	"math/rand"
)

const MaxClassPeopleNumber = 6

type class struct {
	name string
	peopleNumber int
}

func main() {
	start := time.Now();

	execute()

	end := time.Now();
	fmt.Printf("経過時間%d秒\n", int((end.Sub(start)).Seconds()))
}

func execute() {
	cls := classes()

	ch := make(chan class)
	for _, cl := range cls {
		go rollCall(cl, ch)
	}

	var total int
	for i := 0; i < len(cls); i++ {
		tmpCl := <- ch
		total = total + tmpCl.peopleNumber
		fmt.Printf("%sクラスは%d人でした\n", tmpCl.name, tmpCl.peopleNumber)
	}

	fmt.Printf("合計は%d人です\n", total)
}

func rollCall(cl class, ch chan class) {
	cl.peopleNumber = generateRandomNumber()
	wait(cl.peopleNumber)
	ch <- cl
}

func generateRandomNumber() int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(MaxClassPeopleNumber)
}

func wait(s int) {
	time.Sleep(time.Duration(s) * time.Second)
}

func classes() []class {
	return []class{
		class{"A", 0},
		class{"B", 0},
		class{"C", 0},
	}
}
