package pkg

import (
	"fmt"
	"sync"
)

type EntityRepository interface {
	Init()
	Add(entity *Entity) error
	Update(entity *Entity) error
	Get(id string) (*Entity, error)
	List(page, pageSize int) ([]*Entity, error)
	Delete(id string) error
}

type EntityMemoryRepository struct {
	mutex    sync.RWMutex
	entities []*Entity
}

func NewEntityMemoryRepository() *EntityMemoryRepository {
	return &EntityMemoryRepository{
		entities: []*Entity{},
	}
}

func (r *EntityMemoryRepository) Init() {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.entities = []*Entity{
		{
			ID:		"1",
			Type:	PersonEntityType,
			Name:	"John",
			Description: "John is a person",
		},
		{
			ID:		"2",
			Type:	CompanyEntityType,
			Name:	"Google",
			Description: "Google is a company",
		},
		{
			ID:		"3",
			Type:	PlaceEntityType,
			Name:	"seoul",
			Description: "seoul is a place",
		},
		{
			ID:		"4",
			Type:	BookEntityType,
			Name:	"The Little Prince",
			Description: "The Little Prince is a book",
		},
	}
}

func (r *EntityMemoryRepository) Add(entity *Entity) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if entity.ID == ""{
		entity.ID = fmt.Sprintf("%d", len(r.entities) + 1)
	}

	for _, e := range r.entities {
		if e.ID == entity.ID{
			return ErrEntityAlreadyExists
		}
	}

	r.entities = append(r.entities, entity)

	return nil
}