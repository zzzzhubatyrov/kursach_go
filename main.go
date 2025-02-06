package main

import (
	"kurs/intrenal/models"
	storage "kurs/intrenal/storage/sqlite"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// devices := []struct {
	// 	GPONDevice models.GPONDevice
	// 	OLTDevice  models.OLT
	// }{
	// 	GPONDevice: models.GPONDevice{},
	// 	OLTDevice:  models.OLT{},
	// }

	mod := []interface{}{
		&models.Router{},
		&models.Route{},
		&models.NetworkInterface{},
		&models.NAT{},
		&models.NATRule{},
		&models.Switch{},
		&models.Port{},
		&models.VLAN{},
		&models.AccessPoint{},
		&models.ConnectedDevice{},
		&models.Firewall{},
		&models.FirewallRule{},
		&models.FirewallLog{},
		&models.OLT{},
		&models.ONU{},
		&models.PONPort{},
		&models.NetworkDevice{},
		&models.GPONDevice{},
	}

	db, err := storage.SQLiteStorageInit()
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(mod...)
	// Вывод информации о устройствах
	// for _, device := range devices {
	// fmt.Printf("Device ID: %s, Type: %s, Status: %s\n", device.DeviceID, device.DeviceType, device.Status)
	// }

	if err := app.Listen(":5000"); err != nil {
		log.Fatal(err)
	}
}
