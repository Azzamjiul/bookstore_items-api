package controllers

import (
	"bookstore_items-api/domains/item"
	"bookstore_items-api/services"
	"bookstore_items-api/utils/http_utils"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Azzamjiul/bookstore_utils-go/error_utils"
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
		http_utils.ResponseErrorAsJSON(w, err)
		return
	}

	var itemRequest item.Item

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		restErr := error_utils.NewBadRequestError("invalid request body")
		http_utils.ResponseErrorAsJSON(w, restErr)
		return
	}
	defer r.Body.Close()

	// Unmarshal request into the item struct
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		restErr := error_utils.NewBadRequestError("invalid request body")
		http_utils.ResponseErrorAsJSON(w, restErr)
		return
	}

	itemRequest.Seller = oauth.GetCallerId(r)

	result, createErr := services.ItemsService.Create(itemRequest)
	if createErr != nil {
		http_utils.ResponseErrorAsJSON(w, createErr)
		return
	}

	http_utils.ResponseAsJSON(w, http.StatusCreated, result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
