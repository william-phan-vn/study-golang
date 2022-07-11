package main

import "fmt"

// TODO: define the 'Car' type struct
type Car struct {
    battery int
    batteryDrain int
    speed int
    distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car {
        battery: 100,
        speed: speed,
        batteryDrain: batteryDrain,
        distance: 0,
    }
}

// TODO: define the 'Track' type struct
type Track struct {
    distance int
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	return Track {
        distance: distance,
    }
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	return Car {
        speed: car.speed,
        batteryDrain: car.batteryDrain,
        battery: car.battery - car.batteryDrain,
        distance: car.distance + car.speed,
    }
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	fmt.Println("track distance: ", track.distance)
	fmt.Println("remain distance: ", car.speed * (car.battery / car.batteryDrain))
    if car.speed * (car.battery / car.batteryDrain) >= track.distance {
        return true
    }
	return false
}

func main() {
	car := Car {
        battery: 3,
        speed: 5,
        batteryDrain: 7,
        distance: 0,
    }

	distance := 100
	track := NewTrack(distance)

	fmt.Println(CanFinish(car, track))
}