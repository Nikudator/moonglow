package handlers

import (
	"moonglow/database"
	"moonglow/models"

	"github.com/gofiber/fiber/v2"
)

// получение списка всех продуктов
func GetPosts(c *fiber.Ctx) error {
	rows, err := database.DB.Query("SELECT id, name, description, price, stock, image_url FROM products")
	if err != nil {
		return c.Status(500).SendString("Ошибка выполнения запроса к базе данных")
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.Name, &post.Description, &post.Price, &post.Stock, &post.ImageURL)
		if err != nil {
			return c.Status(500).SendString("Ошибка сканирования данных")
		}
		posts = append(posts, post)
	}

	return c.JSON(posts)
}

// создание нового продукта
func CreatePost(c *fiber.Ctx) error {
	post := new(models.Post)
	if err := c.BodyParser(post); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}

	_, err := database.DB.Exec("INSERT INTO products (name, description, price, stock, image_url) VALUES ($1, $2, $3, $4, $5)",
		post.Name, post.Description, post.Price, post.Stock, post.ImageURL)
	if err != nil {
		return c.Status(500).SendString("Ошибка вставки данных в базу")
	}

	return c.Status(201).SendString("Продукт успешно создан")
}

// получение продукта по ID
func GetPost(c *fiber.Ctx) error {
	id := c.Params("id")
	row := database.DB.QueryRow("SELECT id, name, description, price, stock, image_url FROM products WHERE id = $1", id)

	var post models.Post
	err := row.Scan(&post.ID, &post.Name, &post.Description, &post.Price, &post.Stock, &post.ImageURL)
	if err != nil {
		return c.Status(404).SendString("Продукт не найден")
	}

	return c.JSON(post)
}

// обновление продукта
func UpdatePost(c *fiber.Ctx) error {
	id := c.Params("id")
	post := new(models.Post)

	if err := c.BodyParser(post); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}

	_, err := database.DB.Exec("UPDATE products SET name = $1, description = $2, price = $3, stock = $4, image_url = $5 WHERE id = $6",
		post.Name, post.Description, post.Price, post.Stock, post.ImageURL, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления данных")
	}

	return c.SendString("Продукт успешно обновлён")
}

// удаление продукта
func DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	_, err := database.DB.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return c.Status(500).SendString("Ошибка удаления продукта")
	}

	return c.SendString("Продукт успешно удалён")
}
