package domain

import (
	"errors"
	"github.com/google/uuid"
)

var (
	ErrCatalogItemAlreadyExists = errors.New("catalog item already exists")
)

type Catalog struct {
	Id    string
	Title string
	Items map[string]CatalogItem
}

func CreateCatalog(id, title string, items []CatalogItem) *Catalog {
	if id == "" {
		id = uuid.NewString()
	}

	var itemsMap map[string]CatalogItem
	if items != nil {
		for _, i := range items {
			if i.Id == "" {
				i.Id = uuid.NewString()
			}
			itemsMap[i.Id] = i
		}
	}

	return &Catalog{
		Id:    id,
		Title: title,
		Items: itemsMap,
	}
}

func (c *Catalog) AddItem(id, value string) error {
	if id == "" {
		id = uuid.NewString()
	}

	if _, ok := c.Items[id]; ok {
		return ErrCatalogItemAlreadyExists
	}

	c.Items[id] = CatalogItem{
		Id:    id,
		Value: value,
	}

	return nil
}
