package models

import (
	"fmt"
	"math/rand"
	"time"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Vehicle struct {
    parking *Parking
	I               int
	ParkingSpace int
	skin			*canvas.Image
}

func NewVehicle(p *Parking, s *canvas.Image) *Vehicle {
	return &Vehicle{
        parking: p,
		skin: s,
	}
}


func (v *Vehicle) RunVehicle() {
	v.parking.Space <- true
	v.parking.mutex.Lock();	
		for i := 0; i < len(v.parking.ParkingSpaces); i++ {
			if !v.parking.ParkingSpaces[i].occupied{
				v.skin.Move(fyne.NewPos(v.parking.ParkingSpaces[i].x, v.parking.ParkingSpaces[i].y))
				v.skin.Refresh()
				v.ParkingSpace = i
				v.parking.ParkingSpaces[i].occupied = true
				break
			}
		}
		fmt.Println("vehicle ", v.I, " enters")
	v.parking.mutex.Unlock()

		randomSleepSeconds := rand.Intn(30) + 5
		time.Sleep(time.Duration(randomSleepSeconds) * time.Second) // Tiempo que espera en el parking

	v.parking.mutex.Lock()
		<- v.parking.Space
		v.parking.ParkingSpaces[v.ParkingSpace].occupied = false
		v.skin.Move(fyne.NewPos( 460,45 ))
		fmt.Println("vehicle ", v.I, "exit")
		time.Sleep(300 *time.Millisecond) //Tiempo que espera en la salida
    v.parking.mutex.Unlock()
		v.skin.Move(fyne.NewPos( 460000,45000 ))
}



func GenerateVehicle(n int, parking *Parking) {
	parking.Space <- true
	for i := 0; i < n; i++ {
		randomVehicleNumber := rand.Intn(8) + 1
		vehicleImageName := fmt.Sprintf("./assets/vehicle%d.png", randomVehicleNumber)

        VehicleImg := canvas.NewImageFromURI(storage.NewFileURI(vehicleImageName))
		VehicleImg.Resize(fyne.NewSize(70,120))
		VehicleImg.Move(fyne.NewPos(460, 650))

		NewVehicle := NewVehicle(parking, VehicleImg)
		NewVehicle.I = i + 1

		parking.DrawVehicle <- VehicleImg
		time.Sleep(time.Millisecond*400)
		go NewVehicle.RunVehicle()
		time.Sleep(time.Duration(rand.ExpFloat64() * float64(time.Second)))
	}
}