package main

import "fmt"

// Car represents a car object
type Car struct {
	Make     string
	Model    string
	Year     int
	Color    string
	Engine   string
	Transmission string
	Sunroof    bool
	GPS      bool
	Leather   bool
}

// CarBuilder is a builder for the Car object
type CarBuilder struct {
	car *Car
}

// NewCarBuilder creates a new CarBuilder
func NewCarBuilder() *CarBuilder {
	return &CarBuilder{car: &Car{}}
}

// WithMake sets the make of the car
func (b *CarBuilder) WithMake(make string) *CarBuilder {
	b.car.Make = make
	return b
}

// WithModel sets the model of the car
func (b *CarBuilder) WithModel(model string) *CarBuilder {
	b.car.Model = model
	return b
}

// WithYear sets the year of the car
func (b *CarBuilder) WithYear(year int) *CarBuilder {
	b.car.Year = year
	return b
}

// WithColor sets the color of the car
func (b *CarBuilder) WithColor(color string) *CarBuilder {
	b.car.Color = color
	return b
}

// WithEngine sets the engine type of the car
func (b *CarBuilder) WithEngine(engine string) *CarBuilder {
	b.car.Engine = engine
	return b
}

// WithTransmission sets the transmission type of the car
func (b *CarBuilder) WithTransmission(transmission string) *CarBuilder {
	b.car.Transmission = transmission
	return b
}

// WithSunroof sets the sunroof option of the car
func (b *CarBuilder) WithSunroof(sunroof bool) *CarBuilder {
	b.car.Sunroof = sunroof
	return b
}

// WithGPS sets the GPS option of the car
func (b *CarBuilder) WithGPS(GPS bool) *CarBuilder {
	b.car.GPS = GPS
	return b
}

// WithLeather sets the leather seats option of the car
func (b *CarBuilder) WithLeather(leather bool) *CarBuilder {
	b.car.Leather = leather
	return b
}

// Build constructs and returns the Car object
func (b *CarBuilder) Build() *Car {
	return b.car
}

func main() {
	// Creating a car object using a builder
	car := NewCarBuilder()
		.WithMake("Toyota")
		.WithModel("Camry")
		.WithYear(2023)
		.WithColor("Blue")
		.WithEngine("2.5L I4")
		.WithTransmission("Automatic")
		.WithSunroof(true)
		.WithGPS(true)
		.WithLeather(true)
		.Build()

	fmt.Printf("Car: %+v\n", car)
}