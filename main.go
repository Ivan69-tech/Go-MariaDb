package main

import (
	"Grafana/database"
	"fmt"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql" // Import du pilote MySQL/MariaDB
)

func main() {

	db, err := database.DbConnection()

	if err != nil {
		fmt.Println("Erreur lors de la connection à la database:", err)
		return
	}
	defer db.Close()

	for {
		Heure1 := time.Now()
		Heure := Heure1.Format("2006-01-02 15:04:05")
		rand.Seed(time.Now().UnixNano())
		Temperature := rand.Float64()*9 + 1

		_, err = db.Exec("INSERT INTO temp2 (Heure, Temperature) VALUES (?, ?)", Heure, Temperature)
		if err != nil {
			fmt.Println("Erreur lors de l'insertion des données:", err)
			return
		}
		fmt.Println("Données insérées avec succès !")

		time.Sleep(2 * time.Second)

	}

}
