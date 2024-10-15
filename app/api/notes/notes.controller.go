package notes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/quanganh247-qa/gorm-project/app/api/middleware"
	"github.com/quanganh247-qa/gorm-project/app/db"
	"github.com/quanganh247-qa/gorm-project/app/util"
)

type NoteControllerInterface interface {
	CreateNote(c *gin.Context)
	GetNoteByID(c *gin.Context)
	UpdateNote(c *gin.Context)
	DeleteNote(c *gin.Context)
	GetNotes(c *gin.Context)
	GetNotesOfUser(c *gin.Context)
}

// CreateNote godoc
// @Summary Create a new note
// @Description Create a new note
// @Tags notes
// @Accept json
// @Produce json
// @Param note body CreateNoteRequest true "Note info"
// @Success 201 {object} db.Notes
// @Failure 401 {string} string "unauthorized"
// @Router /notes/create [post]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func (controller *NoteController) CreateNote(c *gin.Context) {
	var req CreateNoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	authPayload, err := middleware.GetAuthorizationPayload(c)
	fmt.Println(authPayload.Username)
	note, err := controller.service.CreateNote(c, authPayload.Username, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Create note successfully": note})
}

// GetNoteByID godoc
// @Summary Get note by ID
// @Description Get note by ID
// @Tags notes
// @Accept json
// @Produce json
// @Param note_id path int true "Note ID"
// @Success 200 {object} db.Notes
// @Router /notes/{note_id} [get]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func (controller *NoteController) GetNoteByID(c *gin.Context) {
	noteID := c.Param("note_id")
	idNum, err := strconv.ParseInt(noteID, 10, 16)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	note, err := controller.service.GetNoteByID(c, idNum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Note": note})
}

// UpdateNote godoc
// @Summary Update note
// @Description Update note
// @Tags notes
// @Accept json
// @Produce json
// @Param note_id path int true "Note ID"
// @Param note body db.UpdateNoteRequest true "Note info"
// @Success 200 {object} db.Notes
// @Router /notes/{note_id} [put]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func (controller *NoteController) UpdateNote(c *gin.Context) {
	noteID := c.Param("note_id")
	idNum, err := strconv.ParseInt(noteID, 10, 16)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var req db.UpdateNoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(req)
	note, err := controller.service.UpdateNote(c, req, idNum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Update note successfully": note})
}

// DeleteNote godoc
// @Summary Delete note
// @Description Delete note
// @Tags notes
// @Accept json
// @Produce json
// @Param note_id path int true "Note ID"
// @Success 200 {string} string "Delete note successfully"
// @Router /notes/{note_id} [delete]

func (controller *NoteController) DeleteNote(c *gin.Context) {
	noteID := c.Param("note_id")
	idNum, err := strconv.ParseInt(noteID, 10, 16)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = controller.service.DeleteNote(c, idNum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Delete note successfully": noteID})
}

// GetNotes godoc
// @Summary Get all notes
// @Description Get all notes
// @Tags notes
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param pageSize query int false "Page size"
// @Success 200 {object} []db.Notes
// @Router /notes [get]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func (controller *NoteController) GetNotes(c *gin.Context) {
	pagination, err := util.GetPageInQuery(c.Request.URL.Query())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(pagination)
	notes, err := controller.service.GetNotes(c, pagination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Notes": notes})
}

// GetNotesOfUser godoc
// @Summary Get all notes of user
// @Description Get all notes of user
// @Tags notes
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Success 200 {object} []db.Notes
// @Router /notes/user [get]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func (controller *NoteController) GetNotesOfUser(c *gin.Context) {
	pagination, err := util.GetPageInQuery(c.Request.URL.Query())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	authPayload, err := middleware.GetAuthorizationPayload(c)
	fmt.Println(authPayload.Username)
	notes, err := controller.service.GetNotesOfUser(c, authPayload.Username, pagination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Notes": notes})
}
