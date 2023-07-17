package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prranavv/Backend_Project/database"
	"github.com/prranavv/Backend_Project/routes"
)

func setuproutes(app *fiber.App) {
	app.Post("/tasks", routes.CreateTask)
	app.Get("/tasks", routes.Gettasks)
	app.Get("/tasks/:task_id", routes.Gettask)
	app.Put("/tasks/:task_id", routes.ChangeStatus) //need to use query params eg localhost:3000/tasks/1/?status=finished
	app.Delete("/tasks/:task_id", routes.DeleteTask)
	app.Put("/tasks/update/:task_id", routes.UpdateTask)
}

func main() {
	database.ConnectDb()
	app := fiber.New()
	setuproutes(app)
	app.Listen(":3000")
}
