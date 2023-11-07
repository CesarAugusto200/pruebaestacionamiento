package Logica

import (
	"SimuladorEstacionamiento/controllers"
	"SimuladorEstacionamiento/models"
	"SimuladorEstacionamiento/utils"
	"SimuladorEstacionamiento/views"

	"sync"

	"github.com/faiface/pixel/pixelgl"
)

type Simulation struct {
    carChannel      chan models.Car
    entranceChannel chan int
    winChannel      chan utils.ImgCar
    mu              *sync.Mutex
    parkingCtrl     *controllers.ControllerParking
    entranceCtrl    *controllers.EntranceController
    carCtrl         *controllers.ControllerCar
    carSprites      []utils.ImgCar
    win             *pixelgl.Window // Nuevo campo para la ventana
    viewCar         *views.ViewCar  // Nuevo campo para ViewCar
}

func NewSimulation(win *pixelgl.Window, u *utils.Utils, viewCar *views.ViewCar) *Simulation {
    sim := &Simulation{
        carChannel:      make(chan models.Car, 100),
        entranceChannel: make(chan int),
        winChannel:      make(chan utils.ImgCar),
        mu:              &sync.Mutex{},
        win:             win,
        viewCar:         viewCar, // Asignar viewCar al campo viewCar de la estructura
    }

    sim.parkingCtrl = controllers.NewControllerParking(win, sim.mu)
    sim.entranceCtrl = controllers.NewEntranceController(win, sim.mu, u)
    sim.carCtrl = controllers.NewControllerCar(win, sim.mu, u)

    return sim
}


func (s *Simulation) Init(win *pixelgl.Window, u *utils.Utils) {
    s.carChannel = make(chan models.Car, 100)
    s.entranceChannel = make(chan int)
    s.winChannel = make(chan utils.ImgCar)
    s.mu = &sync.Mutex{}

    s.parkingCtrl = controllers.NewControllerParking(win, s.mu)
    s.entranceCtrl = controllers.NewEntranceController(win, s.mu, u) // Agregar el objeto utils como argumento

    s.carCtrl = controllers.NewControllerCar(win, s.mu, u) // Agregar el objeto utils como argumento
    s.win = win
}


func (s *Simulation) Run() {
    s.viewCar.SetSprite() // Llamar a SetSprite en viewCar en lugar de carCtrl

    go s.parkingCtrl.Park(&s.carChannel, s.entranceCtrl, s.carCtrl, &s.entranceChannel, s.winChannel)
    s.entranceCtrl.LoadStates()
    go s.carCtrl.GenerateCars(100, &s.carChannel)

    for {
        select {
        case val := <-s.winChannel:
            if val.IsEntering() {
                s.carSprites = append(s.carSprites, val)
            } else {
                var arrAux []utils.ImgCar
                for _, value := range s.carSprites {
                    if value.GetID() != val.GetID() {
                        arrAux = append(arrAux, value)
                    }
                }
                s.carSprites = s.carSprites[:0]
                s.carSprites = append(s.carSprites, arrAux...)
            }
        }
    }
}
