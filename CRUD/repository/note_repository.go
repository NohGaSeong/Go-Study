package repository

import "golang-fiber-crud/model"

type NoteRepository interface {
	Save(note model.Note)
	Update(note model.Note)
	Delete(note int)
	FindById(noteId int) (model.Note, error)
	FindAll() []model.Note
}
