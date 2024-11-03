package screens

import (
	"parking/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
)

type GameScene struct {
	window fyne.Window
	content        *fyne.Container
}

func NewScene(window fyne.Window) *GameScene {
	scene := &GameScene{window: window}
    scene.Render()
    return scene
}
func (s *GameScene) StartGame() {
	e := models.NewParking(20)
	go models.GenerateVehicle(100, e)
	go s.DrawVehicles(e)
}

func (s *GameScene) Render() {
	backgroundImage := canvas.NewImageFromURI( storage.NewFileURI("./assets/parkingbackground.png") )
    backgroundImage.Resize(fyne.NewSize(1000,800))
	backgroundImage.Move( fyne.NewPos(0,0) )

	s.content = container.NewWithoutLayout(
        backgroundImage,
    )
    s.window.SetContent(s.content) 
    s.StartGame()
}



func (s *GameScene) DrawVehicles(e *models.Parking) {
	for {
		imagen := <- e.DrawVehicle
		s.content.Add(imagen)
        s.window.Canvas().Refresh(s.content)
	}
}