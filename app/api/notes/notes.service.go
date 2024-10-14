package notes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/quanganh247-qa/gorm-project/app/db"
	"github.com/quanganh247-qa/gorm-project/app/util"
)

type NoteServiceInterface interface {
	CreateNote(c *gin.Context, username string, req CreateNoteRequest) (*db.Notes, error)
	GetNoteByID(c *gin.Context, noteID int64) (*db.Notes, error)
	UpdateNote(c *gin.Context, req db.UpdateNoteRequest, nodeId int64) (*db.Notes, error)
	DeleteNote(c *gin.Context, noteID int64) error
	GetNotes(c *gin.Context, pagination *util.Pagination) ([]db.Notes, error)
	GetNotesOfUser(c *gin.Context, username string, pagination *util.Pagination) ([]db.Notes, error)
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

func (s *NoteService) DeleteNote(c *gin.Context, noteID int64) error {
	err := s.store.DeleteNoteByID(c, noteID)
	if err != nil {
		return fmt.Errorf("failed to delete note: %w", err)
	}
	return nil

}

func (s *NoteService) GetNotes(c *gin.Context, pagination *util.Pagination) ([]db.Notes, error) {
	limit := pagination.PageSize
	offset := (pagination.Page - 1) * pagination.PageSize
	fmt.Println(limit, offset)
	notes, err := s.store.GetNotes(c, int(limit), int(offset))
	if err != nil {
		return nil, fmt.Errorf("failed to get notes: %w", err)
	}
	return notes, nil
}

func (s *NoteService) GetNotesOfUser(c *gin.Context, username string, pagination *util.Pagination) ([]db.Notes, error) {
	limit := pagination.PageSize
	offset := (pagination.Page - 1) * pagination.PageSize

	notes, err := s.store.GetNotesOfUser(c, username, int(limit), int(offset))
	if err != nil {
		return nil, fmt.Errorf("failed to get notes: %w", err)
	}
	return notes, nil
}
