package main

import (
	"kurs/intrenal/models"
	storage "kurs/intrenal/storage/sqlite"
	"log"

	"github.com/gofiber/fiber/v2"
	webview "github.com/webview/webview_go"
)

func main() {
	app := fiber.New()

	db, err := storage.NewSqlite()
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

	db.AutoMigrate(mod...)

	// customLogger := db.Logger.(*pkg.CustomLogger)
	// customLogger.LogMode(logger.Info)
	// Вывод информации о устройствах
	// for _, device := range devices {
	// fmt.Printf("Device ID: %s, Type: %s, Status: %s\n", device.DeviceID, device.DeviceType, device.Status)
	// }

	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Fatal(err)
		}
	}()

	wv := webview.New(true)
	// defer wv.Destroy()
	wv.SetTitle("zhubatyrov-kp")
	wv.SetSize(800, 600, webview.HintNone)
	wv.Navigate("http://localhost:3000")
	wv.Run()
}
