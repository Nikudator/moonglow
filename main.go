package main

import (
	"database"
	"log"
	"routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// инициализируем базу данных
	if err := database.Connect(); err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	// создаём новое приложение Fiber
	app := fiber.New(fiber.Config{
		Prefork: false, // используем предварительное форкование для увеличения производительности
	})

	// Подключаем middleware
	app.Use(logger.New())   // Логирование запросов
	app.Use(compress.New()) // Сжатие ответов
	app.Use(recover.New())  // Восстановление после паники
	app.Use(limiter.New())  // Лимит запросов для предотвращения DDOS атак

	// Регистрация маршрутов
	routes.RegisterProductRoutes(app)

	// Запускаем сервер
	log.Fatal(app.Listen(":8000"))
}
