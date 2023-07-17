package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/prranavv/Backend_Project/database"
	"github.com/prranavv/Backend_Project/models"
)

type TaskDTO struct {
	TaskID    int    `json:"task_id"`
	Task_Name string `json:"task_name"`
	Priority  string `json:"priority"`
	Status    string `json:"status"`
}

func CreateTaskDTO(task models.Task) TaskDTO {
	return TaskDTO{
		TaskID:    int(task.ID),
		Task_Name: task.Task_Name,
		Priority:  task.Priority,
		Status:    task.Status,
	}
}

// POST
func CreateTask(c *fiber.Ctx) error {
	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	flag := 0
	if task.Priority == "low" || task.Priority == "mid" || task.Priority == "high" {
		flag = flag + 1
	}
	if flag == 0 {
		return c.Status(400).SendString("Enter valid priority")
	}
	task.Status = "Pending"
	database.Database.Db.Create(&task)
	responsetask := CreateTaskDTO(task)
	return c.Status(200).JSON(responsetask)
}

// GET
func Gettasks(c *fiber.Ctx) error {
	tasks := []models.Task{}
	database.Database.Db.Find(&tasks)
	responsetasks := []TaskDTO{}
	for _, task := range tasks {
		responsetask := CreateTaskDTO(task)
		responsetasks = append(responsetasks, responsetask)
	}
	return c.Status(200).JSON(responsetasks)
}

func findtaskbyid(id int, task *models.Task) error {
	database.Database.Db.Find(&task, "id=?", id)
	if task.ID == 0 {
		return errors.New("task does not exist")
	}
	return nil
}

// GET individual Task
func Gettask(c *fiber.Ctx) error {
	id, err := c.ParamsInt("task_id")
	if err != nil {
		return c.Status(400).SendString("enter a number to search for")
	}
	var task models.Task
	if err := findtaskbyid(id, &task); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	responsetask := CreateTaskDTO(task)
	return c.Status(200).JSON(responsetask)
}

type ChangeStatusQeury struct {
	Status string `json:"status"`
}

func ChangeStatus(c *fiber.Ctx) error {
	id, err := c.ParamsInt("task_id")
	if err != nil {
		return c.Status(400).SendString("enter a number to search for")
	}
	q := new(ChangeStatusQeury)
	if err := c.QueryParser(q); err != nil {
		return c.SendString("query not parsed")
	}
	var task models.Task
	if err := findtaskbyid(id, &task); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	task.Status = q.Status
	database.Database.Db.Save(&task)
	responsetask := CreateTaskDTO(task)
	return c.JSON(responsetask)

}

//Delete Tasks

func DeleteTask(c *fiber.Ctx) error {
	id, err := c.ParamsInt("task_id")
	if err != nil {
		return c.Status(400).JSON("Enter a number")
	}
	var task models.Task
	if err := findtaskbyid(id, &task); err != nil {
		return c.Status(400).SendString("task not found")
	}
	if err := database.Database.Db.Delete(&task, "id=?", id).Error; err != nil {
		return c.Status(404).JSON(err)
	}
	return c.Status(200).JSON("Successfully deleted")
}
