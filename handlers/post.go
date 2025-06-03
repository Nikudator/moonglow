package handlers

import (
	"moonglow/database"
	"moonglow/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

// получение списка всех постов
func GetPosts(c *fiber.Ctx) error {
	rows, err := database.DB.Query("SELECT id, title, lead, body, created, udated, author FROM posts")
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

// создание нового поста
func CreatePost(c *fiber.Ctx) error {
	post := new(models.Post)
	if err := c.BodyParser(post); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}

	_, err := database.DB.Exec("INSERT INTO posts (id, title, lead, body, created, udated, author) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		uuid.new, post.Title, post.Lead, post.Body, time.Now().Unix(), time.Now().Unix(), post.Author)
	if err != nil {
		return c.Status(500).SendString("Ошибка вставки данных в базу")
	}

	return c.Status(201).SendString("Пост успешно создан")
}

// получение поста по ID
func GetPost(c *fiber.Ctx) error {
	id := c.Params("id")
	row := database.DB.QueryRow("SELECT id, title, lead, body, created, udated, author FROM posts WHERE id = $1", id)

	var post models.Post
	err := row.Scan(&post.ID, &post.Title, &post.Lead, &post.Body, &post.Created, &post.Updated, &post.Author)
	if err != nil {
		return c.Status(404).SendString("Пост не найден")
	}

	return c.JSON(post)
}

// обновление поста
func UpdatePost(c *fiber.Ctx) error {
	id := c.Params("id")
	post := new(models.Post)

	if err := c.BodyParser(post); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}

	_, err := database.DB.Exec("UPDATE posts SET title = $1, lead = $2, body = $3, udated = $4, author = $5 WHERE id = $6",
		post.Title, post.Lead, post.Body, time.Now().Unix(), post.Author, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления данных")
	}

	return c.SendString("Пост успешно обновлён")
}

// удаление поста
func DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	_, err := database.DB.Exec("DELETE FROM posts WHERE id = $1", id)
	if err != nil {
		return c.Status(500).SendString("Ошибка удаления поста")
	}

	return c.SendString("Пост успешно удалён")
}
