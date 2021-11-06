package grpc_service

import (
	"github.com/mayerkv/go-catalogs/domain"
)

func mapItems(items []*CatalogItem) []domain.CatalogItem {
	var res []domain.CatalogItem

	for _, i := range items {
		res = append(res, domain.CatalogItem{
			Id:    i.Id,
			Value: i.Value,
		})
	}

	return res
}

func mapItemsResponse(items map[string]domain.CatalogItem) []*CatalogItem {
	var res []*CatalogItem

	for _, i := range items {
		res = append(res, &CatalogItem{
			Id:    i.Id,
			Value: i.Value,
		})
	}

	return res
}
