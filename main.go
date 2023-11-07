package main

import (
    "SimuladorEstacionamiento/Logic"
    "SimuladorEstacionamiento/utils"
    "SimuladorEstacionamiento/views"
    "github.com/faiface/pixel"
    "github.com/faiface/pixel/pixelgl"
)

func main() {
    pixelgl.Run(func() {
        win, err := pixelgl.NewWindow(pixelgl.WindowConfig{
            Title:  "Simulator Chiapas Parking",
            Bounds: pixel.R(0, 0, 1024, 768), // Ajusta el tamaño de la ventana según tus necesidades
        })
        if err != nil {
            panic(err)
        }

        utilsInstance := utils.NewUtils() // Crear una instancia de utils.Utils

        // Crear una nueva instancia de ViewCar
        viewCar := views.NewViewCar(win, utilsInstance)

        // Asegurarse de que el sprite se establece antes de intentar pintarlo
        viewCar.SetSprite()

        simulation := Logica.NewSimulation(win, utilsInstance, viewCar)
        simulation.Init(win, utilsInstance) 
        simulation.Run()
    })
}
