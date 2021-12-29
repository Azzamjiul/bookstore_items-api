package controllers

import (
	"bookstore_items-api/domains/item"
	"bookstore_items-api/services"
	"fmt"
	"net/http"

	"github.com/azzamjiul/bookstore_oauth-go/oauth"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		// TODO: return error to the caller
		return
	}

	item := item.Item{
		Seller: oauth.GetCallerId(r),
	}

	result, err := services.ItemsService.Create(item)
	if err != nil {
		// TODO: return error json to the user
		return
	}

	fmt.Println(result)
	// TODO: return created item as json with HTTP status 201 - Created
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
