package services

import (
	"bookstore_items-api/domains/item"

	"github.com/Azzamjiul/bookstore_utils-go/error_utils"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(item.Item) (*item.Item, *error_utils.RestErr)
	Get(int64) (*item.Item, *error_utils.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(item item.Item) (*item.Item, *error_utils.RestErr) {
	return nil, nil
}

func (s *itemsService) Get(id int64) (*item.Item, *error_utils.RestErr) {
	return nil, nil
}
