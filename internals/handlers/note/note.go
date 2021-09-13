package noteHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/percoguru/notes-api-fiber/database"
	"github.com/percoguru/notes-api-fiber/internals/model"
)

// GetNotes func gets all existing notes
// @Description Get all existing notes
// @Tags Notes
// @Accept json
// @Produce json
// @Success 200 {array} model.Note
// @router /api/note [get]
func GetNotes(c *fiber.Ctx) error {
	db := database.DB
	var notes []model.Note

	// find all notes in the database
	db.Find(&notes)

	// If no note is present return an error
	if len(notes) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No notes present", "data": nil})
	}

	// Else return notes
	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": notes})
}

// CreateNote func create a note
// @Description Create a Note
// @Tags Notes
// @Accept json
// @Produce json
// @Param title body string true "Title"
// @Param sub_title body string true "SubTitle"
// @Param text body string true "Text"
// @Success 200 {object} model.Note
// @router /api/note [post]
func CreateNote(c *fiber.Ctx) error {
	db := database.DB
	note := new(model.Note)

	// Store the body in the note and return error if encountered
	err := c.BodyParser(note)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	// Add a uuid to the note
	note.ID = uuid.New()
	// Create the Note and return error if encountered
	err = db.Create(&note).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create note", "data": err})
	}

	// Return the created note
	return c.JSON(fiber.Map{"status": "success", "message": "Created Note", "data": note})
}

// GetNote func one note by ID
// @Description Get one note by ID
// @Tags Note
// @Accept json
// @Produce json
// @Success 200 {object} model.Note
// @router /api/note/{id} [get]
func GetNote(c *fiber.Ctx) error {
	db := database.DB
	var note model.Note

	// Read the param noteId
	id := c.Params("noteId")

	// Find the note with the given Id
	db.Find(&note, "id = ?", id)

	// If no such note present return an error
	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	}

	// Return the note with the Id
	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": note})
}

// UpdateNote update a note by ID
// @Description Update a Note by ID
// @Tags Notes
// @Accept json
// @Produce json
// @Param title body string true "Title"
// @Param sub_title body string true "SubTitle"
// @Param text body string true "Text"
// @Success 200 {object} model.Note
// @router /api/note/{id} [post]
func UpdateNote(c *fiber.Ctx) error {
	type updateNote struct {
		Title    string `json:"title"`
		SubTitle string `json:"sub_title"`
		Text     string `json:"Text"`
	}
	db := database.DB
	var note model.Note

	// Read the param noteId
	id := c.Params("noteId")

	// Find the note with the given Id
	db.Find(&note, "id = ?", id)

	// If no such note present return an error
	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	}

	// Store the body containing the updated data and return error if encountered
	var updateNoteData updateNote
	err := c.BodyParser(&updateNoteData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Edit the note
	note.Title = updateNoteData.Title
	note.SubTitle = updateNoteData.SubTitle
	note.Text = updateNoteData.Text

	// Save the Changes
	db.Save(&note)

	// Return the updated note
	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": note})
}

// DeleteNote delete a Note by ID
// @Description Delete a note by ID
// @Tags Note
// @Accept json
// @Produce json
// @Success 200
// @router /api/note/{id} [delete]
func DeleteNote(c *fiber.Ctx) error {
	db := database.DB
	var note model.Note

	// Read the param noteId
	id := c.Params("noteId")

	// Find the note with the given Id
	db.Find(&note, "id = ?", id)

	// If no such note present return an error
	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	}

	// Delete the note and return error if encountered
	err := db.Delete(&note, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete note", "data": nil})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Deleted Note"})
}
