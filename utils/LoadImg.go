package utils

import (
	"image"
	"os"

	"github.com/faiface/pixel"
)

// Utils es una estructura que contiene funciones de utilidad.
type Utils struct{}

func NewUtils() *Utils {
    return &Utils{}
}

// ImgCar es una estructura que representa una imagen de un coche.
type ImgCar struct {
	sprite     *pixel.Sprite
	ID         int
	entering   bool
	position   pixel.Vec
}

// NewImgCar crea una nueva instancia de ImgCar.
func NewImgCar(sprite *pixel.Sprite, ID int, state bool, position pixel.Vec) *ImgCar {
	return &ImgCar{
		sprite:   sprite,
		ID:       ID,
		entering: state,
		position: position,
	}
}

// GetSprite devuelve el sprite asociado a ImgCar.
func (ic *ImgCar) GetSprite() *pixel.Sprite {
	return ic.sprite
}

// GetPosition devuelve la posición del coche.
func (ic *ImgCar) GetPosition() pixel.Vec {
	return ic.position
}

// GetID devuelve el ID del coche.
func (ic *ImgCar) GetID() int {
	return ic.ID
}

// IsEntering devuelve si el coche está ingresando o no.
func (ic *ImgCar) IsEntering() bool {
	return ic.entering
}

// GetData devuelve una copia de la estructura ImgCar.
func (ic *ImgCar) GetData() *ImgCar {
	return ic
}

// LoadPicture carga una imagen desde el archivo especificado.
func (u *Utils) LoadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

// NewSprite crea un nuevo sprite con la imagen y forma especificadas.
func (u *Utils) NewSprite(picture pixel.Picture, form pixel.Rect) *pixel.Sprite {
	return pixel.NewSprite(picture, form)
}
