package lead

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/maitnngo2002/go_lang_projects/go-fiber-crm-basic/database"
)

type Lead struct {
	gorm.Model
	Name    string	`json:"name"`
	Company string	`json: "company"`
	Email   string	`json: "email"`
	Phone   int		`json: "phone"`
}

// context enables the program to use the data from the user
func GetLeads(c *fiber.Ctx) {
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
}

func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead)
}

func NewLead(c *fiber.Ctx) {

	db := database.DBConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&lead)
	c.JSON(lead)

}

func DeleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.First(&lead, id) // find the first Lead with the provided id
	if lead.Name == "" {
		c.Status(500).Send("No Lead found with ID")
	}
	db.Delete(&lead)
	c.Send("Lead successfully deleted")
}