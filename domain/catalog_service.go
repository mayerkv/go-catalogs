package domain

import "errors"

var (
	ErrCatalogNotExists     = errors.New("catalog not exists")
	ErrCatalogAlreadyExists = errors.New("catalog already exists")
)

type CatalogService struct {
	catalogRepository CatalogRepository
}

func NewCatalogService(catalogRepository CatalogRepository) *CatalogService {
	return &CatalogService{catalogRepository: catalogRepository}
}

func (s *CatalogService) CreateCatalog(id, title string, items []CatalogItem) (*Catalog, error) {
	c, err := s.catalogRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	if c != nil {
		return nil, ErrCatalogAlreadyExists
	}

	catalog := CreateCatalog(id, title, items)
	if err := s.catalogRepository.Save(catalog); err != nil {
		return nil, err
	}

	return catalog, nil
}

func (s *CatalogService) AddCatalogItem(catalogId, itemId, value string) error {
	catalog, err := s.getCatalog(catalogId)
	if err != nil {
		return err
	}

	if err := catalog.AddItem(itemId, value); err != nil {
		return err
	}

	return s.catalogRepository.Save(catalog)
}

func (s *CatalogService) GetCatalog(id string) (*Catalog, error) {
	return s.getCatalog(id)
}

func (s *CatalogService) getCatalog(id string) (*Catalog, error) {
	catalog, err := s.catalogRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	if catalog == nil {
		return nil, ErrCatalogNotExists
	}

	return catalog, nil
}
