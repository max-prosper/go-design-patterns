package creational

import "testing"

func TestBuilderPatternCar(t *testing.T) {
	manufactoringComplex := ManufactoringDirector{}

	carBuilder := &CarBuilder{}
	manufactoringComplex.SetBuilder(carBuilder)
	manufactoringComplex.Construct()

	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Errorf("Wheels on a car must be 4 but they were %d", car.Wheels)
	}

	if car.Seats != 5 {
		t.Errorf("Seats on a car must be 5 but they were %d", car.Seats)
	}

	if car.Structure != "Car" {
		t.Errorf("Structure on a car must be 'Car' but it was %s", car.Structure)
	}
}

func TestBuildPatternBike(t *testing.T) {
	manufactoringComplex := ManufactoringDirector{}

	bikeBuilder := &BikeBuilder{}
	manufactoringComplex.SetBuilder(bikeBuilder)
	manufactoringComplex.Construct()

	bike := bikeBuilder.GetVehicle()

	if bike.Wheels != 2 {
		t.Errorf("Wheels on a bike must be 2 but they were %d", bike.Wheels)
	}

	if bike.Seats != 2 {
		t.Errorf("Seats on a bike must be 2 but they were %d", bike.Seats)
	}

	if bike.Structure != "Motorbike" {
		t.Errorf("Structure on a bike must be 'Motorbike' but it was %s", bike.Structure)
	}
}
