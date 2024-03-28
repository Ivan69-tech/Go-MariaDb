package database

import (
	"database/sql"
	"fmt"
)

func CreateDbTable(user string, mdp string, host string) (*sql.DB, error) {

	dataSourceName := user + ":" + mdp + "@tcp(" + host + ":3306)/"

	// Connexion à la base de données
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println("Erreur lors de la connexion à la base de données:", err)
		return nil, err
	}

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + "test2")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Base de données créée avec succès.")

	dataSourceName = dataSourceName + "test2"

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

	// Création des bonnes colonnes

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS temp3 (
		Heure TIMESTAMP NOT NULL,
		PCSState INT UNSIGNED,
		PSetpoint INT UNSIGNED,
		QSetpoint SMALLINT
	)`)

	if err != nil {
		panic(err.Error())
	}

	return db, nil

}
