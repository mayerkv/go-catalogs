package domain

type CatalogRepository interface {
	Save(catalog *Catalog) error
	FindById(id string) (*Catalog, error)
}
