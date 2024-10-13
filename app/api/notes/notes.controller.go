package notes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/quanganh247-qa/gorm-project/app/api/middleware"
	"github.com/quanganh247-qa/gorm-project/app/db"
)

type NoteControllerInterface interface {
	CreateNote(c *gin.Context)
	GetNoteByID(c *gin.Context)
	UpdateNote(c *gin.Context)
}

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
