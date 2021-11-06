package repository

import (
	"github.com/mayerkv/go-catalogs/domain"
	"sync"
)

type InMemoryCatalogRepository struct {
	sync.Mutex
	items map[string]domain.Catalog
}

func NewInMemoryCatalogRepository() domain.CatalogRepository {
	return &InMemoryCatalogRepository{}
}

func (r *InMemoryCatalogRepository) Save(catalog *domain.Catalog) error {
	r.Lock()
	defer r.Unlock()

	r.items[catalog.Id] = *catalog

	return nil
}

func (r *InMemoryCatalogRepository) FindById(id string) (*domain.Catalog, error) {
	r.Lock()
	defer r.Unlock()

	if catalog, ok := r.items[id]; ok {
		return &catalog, nil
	}

	return nil, nil
}
