package views

import (
	"SimuladorEstacionamiento/utils"
	"fmt"
	"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type ViewCar struct {
    win    *pixelgl.Window
    utils  *utils.Utils //se Agrega un campo para los utilidades
    sprite *pixel.Sprite
}


type SpriteCar struct {
	img *pixel.Sprite
	Id  int
}

func NewViewCar(win *pixelgl.Window, utils *utils.Utils) *ViewCar {
    return &ViewCar{
        win:   win,
        utils: utils, // Pasar la instancia de utils.Utils
    }
}


func (cv *ViewCar) SetSprite() {
    fmt.Println("cv:", cv)
    if cv == nil {
        log.Println("Error: La instancia de ViewCar es nula en SetSprite.")
    }
    carSprite := cv.loadCarSprite()
    cv.sprite = carSprite
}


func (cv *ViewCar) PaintCar(pos pixel.Vec) *pixel.Sprite {
	fmt.Println("cv:", cv)
	cv.sprite.Draw(cv.win, pixel.IM.Moved(pos))
	return cv.sprite
}

func (cv *ViewCar) loadCarSprite() *pixel.Sprite {
    if cv != nil {
        picCar, _ := cv.utils.LoadPicture("./assets/Car.png")
        return cv.utils.NewSprite(picCar, picCar.Bounds())
    } else {
        // Manejar el caso en el que cv es nulo
        log.Println("Error: La instancia de ViewCar es nula.")
        // Devolver un sprite en blanco u otro valor predeterminado
        return pixel.NewSprite(nil, pixel.Rect{})
    }
}



func NewImgCar(spr *pixel.Sprite, Id int) *SpriteCar {
	return &SpriteCar{
		img: spr,
		Id:  Id,
	}
}
