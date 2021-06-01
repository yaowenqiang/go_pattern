package creational

import (
    "testing"
)

func TestBuilderPattern(t *testing.T) {
    manufacturingComplex := ManufacturingDirector{}
    carBuilder := &CarBuilder{}
    manufacturingComplex.SetBuilder(carBuilder)
    manufacturingComplex.Construct()
    car := carBuilder.GetVehicle()

    if car.Wheels != 4 {
        t.Errorf("Wheels on a car must be 4 but they were %d\n", car.Wheels)
    }

    if car.Structure != "Car" {
        t.Errorf("Structure on a car must be 'Car' but was %s\n", car.Structure)
    }

    if car.Seats != 5 {
        t.Errorf("Seats on a car must be 5 but they were %d\n", car.Seats)
    }

    bikeBuilder := &BikeBuilder{}
    manufacturingComplex.SetBuilder(bikeBuilder)
    manufacturingComplex.Construct()

    motobike := bikeBuilder.GetVehicle()
    motobike.Seats = 1

    if motobike.Wheels != 2 {
        t.Errorf("Wheels on a motobike  must be 2 but they were %d\n", motobike.Wheels)
    }
    if motobike.Structure != "Motobike" {
        t.Errorf("Structure on a motobike  must be 'Motobike' but was %s\n", motobike.Structure)
    }

}
