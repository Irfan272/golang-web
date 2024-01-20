package categorymodel

import (
	"log"
	"web-crud/config"
	"web-crud/entities"
)

func GetAll() []entities.Categories {
	db := config.ConnectDB()
	rows, err := db.Query("SELECT * FROM categories")
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	var categories []entities.Categories

	for rows.Next() {
		var category entities.Categories
		err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt)

		if err != nil {
			panic(err)
		}

		categories = append(categories, category)
	}
	return categories

}

func Create(categories entities.Categories) bool {
	db := config.ConnectDB()

	result, err := db.QueryRow(`
        INSERT INTO categories (name, created_at, updated_at) VALUES (?, ?, ?);`,
		categories.Name, categories.CreatedAt, categories.UpdatedAt,
	)

	if err != nil {
		log.Println("Error in query", err)
		log.Println("Result :", result)
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()

	if err != nil {
		log.Println("Error in lastinsertid", err)
		panic(err)
	}

	return lastInsertId > 0
}
