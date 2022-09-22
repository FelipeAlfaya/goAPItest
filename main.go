package main

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

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
	e.POST("/cars", createCar)
	e.Logger.Fatal(e.Start(":8080"))
}

func getCars(c echo.Context) error {
	return c.JSON(200, cars)
}

func createCar(c echo.Context) error {
	car := new(Car)
	if err := c.Bind(car); err != nil {
		return err
	}
	cars = append(cars, *car)
	saveCar(*car)
	return c.JSON(200, cars)
}

func saveCar(car Car) error {
	db, err := sql.Open("mysql", "cars.db")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO cars (name, price) VALUES (?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(car.Name, car.Price)
	if err != nil {
		return err
	}
	return nil
}
