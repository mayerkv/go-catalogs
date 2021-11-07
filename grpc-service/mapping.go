package grpc_service

import (
	"fmt"
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
	res := make([]*CatalogItem, 0)

	for _, i := range items {
		res = append(res, &CatalogItem{
			Id:    i.Id,
			Value: i.Value,
		})
	}

	fmt.Println(res)

	return res
}
