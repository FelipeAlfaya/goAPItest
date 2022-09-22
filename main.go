package main

import "github.com/labstack/echo/v4"

type Car struct {
	Name  string
	Price float64
}

var cars []Car

func CreateCars() {
	cars = []Car{
		{Name: "Ford", Price: 19.99},
		{Name: "BMW", Price: 19.99},
		{Name: "Chevrolet", Price: 19.99},
	}

	// 	cars = append(cars, Car{Name: "Ford", Price: 19.99})
	// 	cars = append(cars, Car{Name: "BMW", Price: 19.99})
	// 	cars = append(cars, Car{Name: "Chevrolet", Price: 19.99})
}

func main() {
	CreateCars()
	e := echo.New()
	e.GET("/cars", getCars)
	e.Logger.Fatal(e.Start(":8080"))
}

func getCars(c echo.Context) error {
	return c.JSON(200, cars)
}
