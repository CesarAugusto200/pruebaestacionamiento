package controllers

import (
	"SimuladorEstacionamiento/models"
	"SimuladorEstacionamiento/utils"
	"SimuladorEstacionamiento/views"
	"sync"

	"github.com/faiface/pixel/pixelgl"
)

type ControllerParking struct {
	model *models.Parking
	view  *views.ViewParking
	mu    *sync.Mutex
}

func NewControllerParking(win *pixelgl.Window, mu *sync.Mutex) *ControllerParking {
	return &ControllerParking{
		model: models.NewParking(),
		view:  views.NewViewParking(win),
		mu:    mu,
	}
}

func (pc *ControllerParking) PaintParking() {
	pc.view.PaintParking()
}

func (pc *ControllerParking) PaintStreet() {
	pc.view.PaintStreet()
}

func (pc *ControllerParking)Park(chCar *chan models.Car, entranceController *EntranceController, ControllerCar *ControllerCar, chEntrance *chan int, chWin chan utils.ImgCar) {
	go pc.ChangingState(chEntrance, entranceController)
	for car := range *chCar {
		pos := pc.model.FindSpaces()
		if pos != -1 {
			coo := pc.view.GetCoordinates(pos)
			ControllerCar.view.SetSprite()
			sprite := ControllerCar.view.PaintCar(coo)
			if entranceController.model.GetState() == "Idle" || entranceController.model.GetState() == "Entering" {
				go car.Timer(pos, pc.model, pc.mu, pc.model.GetAllSpaces(), chEntrance, sprite, chWin, coo)
			} else {
				*chEntrance <- 0
				go car.Timer(pos, pc.model, pc.mu, pc.model.GetAllSpaces(), chEntrance, sprite, chWin, coo)
			}
		}
	}
}

func (pc *ControllerParking) ChangingState(chEntrance *chan int, entranceController *EntranceController) {
	for change := range *chEntrance {
		entranceController.model.SetState(change)
	}
}
