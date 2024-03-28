package database

import (
	"fmt"
	"time"

	"github.com/simonvetter/modbus"
)

func FillDbModbus() {

	var (
		PCSState  uint16
		PSetpoint uint16
		QSetpoint uint16
		client    *modbus.ModbusClient
		err       error
	)

	client, err = modbus.NewClient(&modbus.ClientConfiguration{
		URL:     "tcp://localhost:5502",
		Timeout: 1 * time.Second,
	})

	if err != nil {
		fmt.Println("Erreur lors de la connection au serveur modbus:", err)
		return
	}

	err = client.Open()
	if err != nil {
		fmt.Println("impossible d'ouvrir la connexion", err)
	}

	db, err := CreateDbTable("root", "malivajo", "localhost")

	if err != nil {
		fmt.Println("Erreur lors de la connection à la database:", err)
		return
	}
	defer db.Close()

	client.WriteRegister(101, 2)
	client.WriteRegister(102, 25)
	client.WriteRegister(103, 50)

	for {
		PCSState, _ = client.ReadRegister(101, modbus.HOLDING_REGISTER)
		PSetpoint, _ = client.ReadRegister(102, modbus.HOLDING_REGISTER)
		QSetpoint, _ = client.ReadRegister(103, modbus.HOLDING_REGISTER)
		Heure1 := time.Now()
		Heure := Heure1.Format("2006-01-02 15:04:05")

		_, err = db.Exec("INSERT INTO temp3 (Heure, PCSState, PSetpoint, QSetpoint) VALUES (?, ?, ?, ?)",
			Heure,
			PCSState,
			PSetpoint,
			QSetpoint)

		if err != nil {
			fmt.Println("Erreur lors de l'insertion des données:", err)
			return
		}
		fmt.Println("Données insérées avec succès !")
		fmt.Println(Heure, PCSState, PSetpoint, QSetpoint)

		time.Sleep(5 * time.Second)

	}
}
