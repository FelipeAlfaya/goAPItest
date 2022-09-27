package main

import (
	"database/sql"
	"goAPItest/diretorio2"

	"github.com/labstack/echo/v4"
)

type Dog struct {
	Name  string
	Price float64
	Breed string
}

type Cat struct {
	Name  string
	Price float64
	Breed string
}

var dogs []Dog
var cats []Cat

func GenDogs() {
	dogs = []Dog{
		{Name: "Wilson", Price: 19.99, Breed: "Golden Retriever"},
		{Name: "Roberto", Price: 19.99, Breed: "Shi tzu"},
		{Name: "Matilda", Price: 19.99, Breed: "Mixed"},
	}

	// 	cars = append(cars, Car{Name: "Ford", Price: 19.99})
	// 	cars = append(cars, Car{Name: "BMW", Price: 19.99})
	// 	cars = append(cars, Car{Name: "Chevrolet", Price: 19.99})
}

func GenCats() {
	cats = []Cat{
		{Name: "Rodisney", Price: 15.99, Breed: "Ragdoll"},
		{Name: "Clovis", Price: 9.99, Breed: "British Shorthair"},
		{Name: "", Price: 17.99, Breed: "Sphynx"},
	}
}

func main() {
	GenDogs()
	GenCats()
	e := echo.New()
	e.GET("/dogs", getDogs)
	e.POST("/dogs", createDog)
	e.GET("/cats", getCats)
	e.POST("/cats", createCat)
	e.Logger.Fatal(e.Start(":8080"))
	diretorio2.Test()
}

func getCats(c echo.Context) error {
	return c.JSON(200, cats)
}

func createCat(c echo.Context) error {
	cat := new(Cat)
	if err := c.Bind(cat); err != nil {
		return err
	}
	cats = append(cats, *cat)
	saveCat(*cat)
	return c.JSON(200, cats)
}

func saveCat(cat Cat) error {
	db, err := sql.Open("mysql", "dogs.db")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO cats (name, price, breed) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(cat.Name, cat.Price, cat.Breed)
	if err != nil {
		return err
	}
	return nil
}

func getDogs(c echo.Context) error {
	return c.JSON(200, dogs)
}

func createDog(c echo.Context) error {
	dog := new(Dog)
	if err := c.Bind(dog); err != nil {
		return err
	}
	dogs = append(dogs, *dog)
	saveDog(*dog)
	return c.JSON(200, dogs)
}

func saveDog(dog Dog) error {
	db, err := sql.Open("mysql", "dogs.db")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO dogs (name, price, breed) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(dog.Name, dog.Price, dog.Breed)
	if err != nil {
		return err
	}
	return nil
}
