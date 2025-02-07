package main

import (
	"kurs/intrenal/models"
	storage "kurs/intrenal/storage/sqlite"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db, err := storage.SQLiteStorageInit()
	if err != nil {
		log.Fatal(err)
	}

	mod := []interface{}{
		&models.Route{},
		&models.NetworkInterface{},
		&models.NATRule{},
		&models.Port{},
		&models.VLAN{},
		&models.ConnectedDevice{},
		&models.FirewallRule{},
		&models.FirewallLog{},
		&models.ONU{},
		&models.NetworkDevice{},
		&models.GPONDevice{},
		&models.Router{},
		&models.NAT{},
		&models.Switch{},
		&models.AccessPoint{},
		&models.Firewall{},
		&models.OLT{},
		&models.PONPort{},
	}

	db.Debug().AutoMigrate(mod...)
	// Вывод информации о устройствах
	// for _, device := range devices {
	// fmt.Printf("Device ID: %s, Type: %s, Status: %s\n", device.DeviceID, device.DeviceType, device.Status)
	// }

	if err := app.Listen(":5000"); err != nil {
		log.Fatal(err)
	}
}
