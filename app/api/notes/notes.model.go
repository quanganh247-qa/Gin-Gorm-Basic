package notes

import "github.com/quanganh247-qa/gorm-project/app/db"

type NoteService struct {
	store db.Store
}

type NoteAPI struct {
	controller NoteControllerInterface
}

type NoteController struct {
	service NoteServiceInterface
}

type CreateNoteRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type GetNoteByTitleRequest struct {
	Sub_title string `json:"sub_title"`
}
