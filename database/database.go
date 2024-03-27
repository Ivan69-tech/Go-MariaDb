package database

import (
	"database/sql"
	"fmt"
)

func DbConnection() (*sql.DB, error) {

	dataSourceName := "root:malivajo@tcp(localhost:3306)/"

	// Connexion à la base de données
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println("Erreur lors de la connexion à la base de données:", err)
		return nil, err
	}

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + "test")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Base de données créée avec succès.")

	dataSourceName = "root:malivajo@tcp(localhost:3306)/test"

	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println("Erreur lors de la connexion à la base de données:", err)
		return nil, err
	}

	// Vérification de la connexion à la base de données
	err = db.Ping()
	if err != nil {
		fmt.Println("Erreur lors de la vérification de la connexion à la base de données:", err)
		return nil, err
	}
	fmt.Println("Connexion à la base de données réussie !")
	return db, nil
}
