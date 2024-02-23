package managerController

import (
	managerModel "github.com/chronicler-org/core/src/manager/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)


func HandleGetAll(c *fiber.Ctx) error {
  var db = managerModel.InitHandler()

  managers, err := db.FindAll()
  if err != nil {
    return c.SendStatus(fiber.StatusInternalServerError)
  }

  return c.Status(fiber.StatusOK).JSON(managers)
}

func HandleGetById(c *fiber.Ctx) error {
  db := managerModel.InitHandler()

  manager, err := db.FindByID(c.Params("id"))
  if err != nil {
    return c.SendStatus(fiber.StatusInternalServerError)
  }
  if manager.ID == uuid.Nil {
    return c.SendStatus(fiber.StatusNotFound)
  }

  return c.Status(fiber.StatusOK).JSON(manager)
}

func HandleCreateManager(c *fiber.Ctx) error {
  var db = managerModel.InitHandler()

  var manager managerModel.CreateManagerDTO
  err := c.BodyParser(&manager)
  if err != nil {
    return c.SendStatus(fiber.StatusBadRequest)
  }

  newUserID, err := db.Create(&manager)
  if err != nil {
    return c.SendStatus(fiber.StatusInternalServerError)
  }

  return c.Status(fiber.StatusOK).JSON(fiber.Map{
    "location": newUserID,
  })
}

func HandleUpdateManager(c *fiber.Ctx) error {
  var db = managerModel.InitHandler()

  id := c.Params("id")

  var updateData managerModel.CreateManagerDTO
  err := c.BodyParser(&updateData)

  if err != nil {
    return c.SendStatus(fiber.StatusBadRequest)
  }
  
  managerToUpdate, err := db.FindByID(id)
  if managerToUpdate.ID == uuid.Nil {
    return c.SendStatus(fiber.StatusNotFound)
  }

  err = db.Update(&managerToUpdate, &updateData)

  if err != nil {
    return c.SendStatus(fiber.StatusInternalServerError)
  }

  return c.Status(fiber.StatusOK).JSON(managerToUpdate)
}

func HandleDeleteManager(c *fiber.Ctx) error {
  db := managerModel.InitHandler()
  id := c.Params("id")

  manager, err := db.FindByID(id)
  if err != nil {
    return c.SendStatus(fiber.StatusInternalServerError)
  }
  if manager.ID == uuid.Nil {
    return c.SendStatus(fiber.StatusNotFound)
  }

  err = db.Delete(id)
  if err != nil {
    return c.SendStatus(fiber.StatusBadRequest)
  }
  
  return c.Status(fiber.StatusOK).JSON(manager)
}
