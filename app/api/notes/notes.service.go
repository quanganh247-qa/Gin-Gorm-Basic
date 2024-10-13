package notes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/quanganh247-qa/gorm-project/app/db"
)

type NoteServiceInterface interface {
	CreateNote(c *gin.Context, username string, req CreateNoteRequest) (*db.Notes, error)
	GetNoteByID(c *gin.Context, noteID int64) (*db.Notes, error)
	UpdateNote(c *gin.Context, req db.UpdateNoteRequest, nodeId int64) (*db.Notes, error)
}

func (s *NoteService) CreateNote(c *gin.Context, username string, req CreateNoteRequest) (*db.Notes, error) {
	note, err := s.store.CreateNote(c, db.CreateNoteRequest{
		Username: username,
		Title:    req.Title,
		Content:  req.Content,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create note: %w", err)
	}
	return &note, nil
}

func (s *NoteService) GetNoteByID(c *gin.Context, noteID int64) (*db.Notes, error) {
	note, err := s.store.GetNoteByID(c, noteID)
	if err != nil {
		return nil, fmt.Errorf("failed to get note: %w", err)
	}
	return &note, nil
}

func (s *NoteService) UpdateNote(c *gin.Context, req db.UpdateNoteRequest, nodeId int64) (*db.Notes, error) {
	note, err := s.store.UpdateNote(c, nodeId, db.UpdateNoteRequest{
		Title:   req.Title,
		Content: req.Content,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update note: %w", err)
	}
	return &note, nil
}
