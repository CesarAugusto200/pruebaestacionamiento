package controllers

import (
	"SimuladorEstacionamiento/models"
	"SimuladorEstacionamiento/utils"
	"SimuladorEstacionamiento/views"
	"sync"

	"github.com/faiface/pixel/pixelgl"
)

type EntranceController struct {
	model *models.Entrada 
	view  *views.ViewEntrada
	mu    *sync.Mutex
}

func NewEntranceController(win *pixelgl.Window, mu *sync.Mutex, utilsInstance *utils.Utils) *EntranceController {
	viewEntrada := views.NewViewEntrada(win, utilsInstance) // Declara y asigna viewEntrada
	return &EntranceController{
		model: models.NewEntrada(),
		view:  viewEntrada, // Asigna viewEntrada a ec.view
		mu:    mu,
	}
}


func (ec *EntranceController) LoadStates() {
	imgs := ec.view.LoadStatesImages()
	ec.view.SetStateImages(imgs)
}

func (ec *EntranceController) PaintEntrance(pos int) {
	ec.view.PaintEntrance(pos)
}
