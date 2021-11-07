package grpc_service

import (
	context "context"
	"github.com/mayerkv/go-catalogs/domain"
)

type CatalogsServiceServerImpl struct {
	catalogService *domain.CatalogService
}

func NewCatalogsServiceServerImpl(catalogService *domain.CatalogService) CatalogsServiceServer {
	return &CatalogsServiceServerImpl{catalogService: catalogService}
}

func (s *CatalogsServiceServerImpl) CreateCatalog(ctx context.Context, request *CreateCatalogRequest) (*CreateCatalogResponse, error) {
	c, err := s.catalogService.CreateCatalog(request.Catalog.Id, request.Catalog.Title, mapItems(request.Catalog.Items))
	if err != nil {
		return nil, err
	}

	return &CreateCatalogResponse{
		Id: c.Id,
	}, nil
}

func (s *CatalogsServiceServerImpl) AddCatalogItem(ctx context.Context, request *AddCatalogItemRequest) (*AddCatalogItemResponse, error) {
	err := s.catalogService.AddCatalogItem(request.CatalogId, request.Item.Id, request.Item.Value)
	if err != nil {
		return nil, err
	}

	return &AddCatalogItemResponse{}, nil
}

func (s *CatalogsServiceServerImpl) GetCatalogItems(ctx context.Context, request *GetCatalogItemsRequest) (*GetCatalogItemsResponse, error) {
	catalog, err := s.catalogService.GetCatalog(request.CatalogId)
	if err != nil {
		return nil, err
	}

	return &GetCatalogItemsResponse{
		Items: mapItemsResponse(catalog.Items),
	}, nil
}

func (s *CatalogsServiceServerImpl) mustEmbedUnimplementedCatalogsServiceServer() {
	panic("implement me")
}
