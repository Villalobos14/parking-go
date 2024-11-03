package models

import (
	"sync"
	"fyne.io/fyne/v2/canvas"
)

type Slot struct{
	x float32
	y float32
	occupied bool
}

type Parking struct {
	Space chan bool
	DrawVehicle chan *canvas.Image
	mutex sync.Mutex
	ParkingSpaces []Slot
}


func NewParking(nS int) *Parking {
	return &Parking{
		Space: make(chan bool, nS+1),
		DrawVehicle: make(chan *canvas.Image,1),
		ParkingSpaces: []Slot{
			{x: 40, y: 200, occupied:false},
			{x: 120, y: 200, occupied:false},
			{x: 200, y: 200, occupied:false},
			{x: 280, y: 200, occupied:false},
			{x: 360, y: 200, occupied:false},
			{x: 560, y: 200, occupied:false},
			{x: 640, y: 200, occupied:false},
			{x: 720, y: 200, occupied:false},
			{x: 800, y: 200, occupied:false},
			{x: 890, y: 200, occupied:false},

			{x: 40, y: 480, occupied:false},
			{x: 120, y: 480, occupied:false},
			{x: 200, y: 480, occupied:false},
			{x: 280, y: 480, occupied:false},
			{x: 360, y: 480, occupied:false},
			{x: 560, y: 480, occupied:false},
			{x: 640, y: 480, occupied:false},
			{x: 720, y: 480, occupied:false},
			{x: 800, y: 480, occupied:false},
			{x: 890, y: 480, occupied:false},
		},
	}
}