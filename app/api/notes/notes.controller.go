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
