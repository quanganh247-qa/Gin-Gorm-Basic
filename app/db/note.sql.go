package db

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type CreateNoteRequest struct {
	Username string `json:"username"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

type UpdateNoteRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (s *Store) CreateNote(c context.Context, arg CreateNoteRequest) (Notes, error) {
	note := Notes{
		Username: arg.Username,
		Title:    arg.Title,
		Content:  arg.Content,
	}
	result := s.db.WithContext(c).Create(&note)
	if result.Error != nil {
		return Notes{}, result.Error
	}
	return note, nil
}

func (s *Store) GetNoteByID(c context.Context, noteID int64) (Notes, error) {
	var note Notes
	result := s.db.WithContext(c).Where("note_id = ?", noteID).First(&note)
	if result.Error != nil {
		return Notes{}, fmt.Errorf("error getting note: %w", result.Error)
	}
	return note, nil
}

func (s *Store) UpdateNote(c context.Context, noteID int64, arg UpdateNoteRequest) (Notes, error) {

	var note Notes
	if err := s.db.WithContext(c).First(&note, "note_id = ?", noteID).Error; err != nil {
		return Notes{}, fmt.Errorf("note not found: %w", err)
	}

	err := s.ExecTx(c, func(tx *gorm.DB) error {

		updatedFields := map[string]interface{}{}
		if arg.Title != "" {
			updatedFields["title"] = arg.Title
		}
		if arg.Content != "" {
			updatedFields["content"] = arg.Content
		}

		fmt.Println("Updated fields: ", arg.Title, arg.Content)

		if len(updatedFields) > 0 {
			result := s.db.WithContext(c).Model(&note).Updates(updatedFields)
			if result.Error != nil {
				return fmt.Errorf("error updating note: %w", result.Error)
			}
			fmt.Println("Updated note: ", result.RowsAffected)
		}

		if err := s.db.WithContext(c).First(&note, "note_id = ?", noteID).Error; err != nil {
			return fmt.Errorf("error fetching updated note: %w", err)
		}

		return nil

	})
	if err != nil {
		return Notes{}, err
	}

	return note, nil
}

func (s *Store) DeleteNoteByID(c context.Context, noteID int64) error {
	if err := s.ExecTx(c, func(tx *gorm.DB) error {
		var note Notes
		if err := s.db.WithContext(c).First(&note, "note_id = ?", noteID).Error; err != nil {
			return fmt.Errorf("note not found: %w", err)
		}

		result := s.db.WithContext(c).Delete(&note)
		if result.Error != nil {
			return fmt.Errorf("error deleting note: %w", result.Error)
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *Store) GetNotes(c context.Context, limit, offset int) ([]Notes, error) {
	var notes []Notes
	result := s.db.WithContext(c).Limit(limit).Offset(offset).Find(&notes)
	if result.Error != nil {
		return nil, fmt.Errorf("error getting notes: %w", result.Error)
	}
	return notes, nil
}

func (s *Store) GetNotesOfUser(c context.Context, username string, limit, offset int) ([]Notes, error) {
	var notes []Notes
	result := s.db.WithContext(c).Where("username = ?", username).Limit(limit).Offset(offset).Find(&notes)
	if result.Error != nil {
		return nil, fmt.Errorf("error getting notes: %w", result.Error)
	}
	return notes, nil

}
