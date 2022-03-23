package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Car struct {
	carID string `json:"car_id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price string `json:"price"`
}

var CarDatas = []Car{}

func CreateCar(ctx *gin.Context) {
	var newCar Car

	if err := ctx.ShouldBindJSON(&newCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newCar.carID = fmt.Sprintf("c%d", len(CarDatas)+1)
	CarDatas = append(CarDatas, newCar)

	ctx.JSON(http.StatusCreated, gin.H{
		"car": newCar,
	})
}

func UpdateCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false
	var updatedCar Car

	if err := ctx.ShouldBindJSON(&updatedCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, car := range CarDatas {
		if car.carID == carID {
			condition = true
			CarDatas[i] = updatedCar
			CarDatas[i].carID = carID
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("Car with ID %v not found", carID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Car with ID %v has been updated", carID),
	})
}

func GetCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false
	var carData Car

	for i, car := range CarDatas {
		if carID == car.carID {
			condition = true
			carData = CarDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("Car with ID %v not found", carID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"car": carData,
	})
}

func DeleteCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false
	var carIndex int

	for i, car := range CarDatas {
		if car.carID == carID {
			condition = true
			carIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("Car with ID %v not found", carID),
		})
		return
	}

	copy(CarDatas[carIndex:], CarDatas[carIndex+1:])
	CarDatas[len(CarDatas)-1] = Car{}
	CarDatas = CarDatas[:len(CarDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Car with ID %v has been deleted", carID),
	})
}
