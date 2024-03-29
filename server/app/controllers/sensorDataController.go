package controller

import (
	repo "smhome/pkg/repository"
	service "smhome/pkg/services"
	"smhome/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func SensorStats(c *fiber.Ctx) error {
	typ := c.Query("type", "none")
	if typ == "none" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "require ?type = ...",
			"success": false,
		})
	}
	date := c.Query("date", "none")
	if date == "none" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "require ?date = ...",
			"success": false,
		})
	}
	var sensor struct {
		service string
		data    string
	}
	switch typ {
	case "temp":
		sensor.service = repo.DTemp
		sensor.data = repo.TEMPERATURE
	case "humid":
		sensor.service = repo.DHumid
		sensor.data = repo.HUMIDITY
	case "light":
		sensor.service = repo.DLight
		sensor.data = repo.LIGHT
	default:
		return c.SendStatus(fiber.StatusBadRequest)
	}
	dataService := service.NewDataService(sensor.service)
	data, err := dataService.GetSensorData(sensor.data, date)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   err.Error(),
			"success": false,
		})
	}
	res, _ := utils.SensorDataStat(*data, date)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": res,
		"success": true,
	})
}
