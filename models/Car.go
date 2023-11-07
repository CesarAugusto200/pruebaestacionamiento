package models

import (
	"fmt"
	"math/rand"
	"SimuladorEstacionamiento/utils"
	"sync"
	"time"

	"github.com/faiface/pixel"
)

type Car struct {
	ParkingTime int 
	Id int
}

func NewCar() *Car {
	rand.Seed(time.Now().UnixNano()) 
	parkingTime := rand.Intn(10) + 15
	return &Car{ParkingTime: parkingTime}
}

func (c *Car) GenerateCars(n int, ch chan Car) {
	for i := 1; i <= n; i++ {
		car := NewCar()
		car.Id = i
		ch<- *car
		rand.Seed(time.Now().UnixNano()) 
		newTime := rand.Intn(2) + 1
		time.Sleep(time.Second * time.Duration(newTime))
	}
	close(ch)
	fmt.Println("Ha Terminado El tiempo de la generacion de Autos")
}

func (c *Car) Timer(pos int, pc *Parking, mu *sync.Mutex, spaces *[20]bool, chEntrance *chan int, sprite *pixel.Sprite, chWin chan utils.ImgCar, coo pixel.Vec) {
	mu.Lock()
	data := utils.NewImgCar(sprite, pos, true, coo)
	chWin<-*data
	*chEntrance<-0
	mu.Unlock()

	mu.Lock()
	pc.nSpaces--
	fmt.Println("The Car", c, "He just parked and is parked in spot number:", pos)
	fmt.Println("They remain", pc.nSpaces, "spaces available")
	mu.Unlock()

	time.Sleep(time.Second * time.Duration(c.ParkingTime))

	fmt.Println("The Car", c, c.Id,"It was parked in the place with the number:", pos)
	
	mu.Lock()
	data = utils.NewImgCar(sprite, pos, false, coo)
	chWin<-*data
	pc.nSpaces = pc.nSpaces + 1
	spaces[pos] = true
	fmt.Println("There is", pc.nSpaces, "Availability of spaces, after which the car left")
	mu.Unlock()

	mu.Lock()
	*chEntrance<-1
	mu.Unlock()
}