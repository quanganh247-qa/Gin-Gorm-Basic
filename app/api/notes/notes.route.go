package notes

import (
	"github.com/quanganh247-qa/gorm-project/app/api/middleware"
	"github.com/quanganh247-qa/gorm-project/app/db"
)

func Routes(routerGroup middleware.RouterGroup) {
	notes := routerGroup.RouterDefault.Group("/notes")
	authRoute := routerGroup.RouteAuth(notes)

	notesAPI := &NoteAPI{
		&NoteController{
			service: &NoteService{
				store: db.StoreDB,
			},
		},
	}
	{
		authRoute.POST("/create", notesAPI.controller.CreateNote)
		authRoute.GET("/:note_id", notesAPI.controller.GetNoteByID)
		authRoute.PUT("/:note_id", notesAPI.controller.UpdateNote)
		authRoute.DELETE("/:note_id", notesAPI.controller.DeleteNote)
		authRoute.GET("/all-notes", notesAPI.controller.GetNotes)
		authRoute.GET("/", notesAPI.controller.GetNotesByUser)
	}
}
