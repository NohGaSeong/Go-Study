package repository

import (
	"errors"
	"golang-fiber-crud/data/request"
	"golang-fiber-crud/helper"
	"golang-fiber-crud/model"

	"gorm.io/gorm"
)

type NoteRepositoryImpl struct {
	Db *gorm.DB
}

func NewNoteRepositoryImpl(Db *gorm.DB) NoteRepository {
	return &NoteRepositoryImpl{Db: Db}
}

// Delete implements NoteRepository.
func (n *NoteRepositoryImpl) Delete(noteId int) {
	var note model.Note
	result := n.Db.Where("id = ?", noteId).Delete(&note)
	helper.ErrorPanic(result.Error)
}

// FindAll implements NoteRepository.
func (n *NoteRepositoryImpl) FindAll() []model.Note {
	var note []model.Note
	result := n.Db.Find(&note)
	helper.ErrorPanic(result.Error)
	return note
}

// FindById implements NoteRepository.
func (n *NoteRepositoryImpl) FindById(noteId int) (model.Note, error) {
	var note model.Note
	result := n.Db.Find(&note, noteId)
	if result != nil {
		return note, nil
	} else {
		return note, errors.New("note is not found")
	}
}

// Save implements NoteRepository.
func (n *NoteRepositoryImpl) Save(note model.Note) {
	result := n.Db.Create(&note)
	helper.ErrorPanic(result.Error)
}

// Update implements NoteRepository.
func (n *NoteRepositoryImpl) Update(note model.Note) {
	var updateNote = request.UpdateNoteRequest{
		Id:      note.Id,
		Content: note.Content,
	}
	result := n.Db.Model(&note).Updates(updateNote)
	helper.ErrorPanic(result.Error)
}
