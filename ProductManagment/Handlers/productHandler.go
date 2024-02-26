package handlers

import (
	helpers "apiModules/ProductManagment/Helpers"
	modals "apiModules/ProductManagment/Modals"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// ürünlerin
var productStore = make(map[string]modals.Product)
var productId int = 0

// HTTP POST -/api/products ekle
func PostProductHandler(w http.ResponseWriter, r *http.Request) {
	var product modals.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	helpers.CheckError(err)

	product.CreatedOn = time.Now()
	productId++
	product.ID = productId
	key := strconv.Itoa(productId)
	productStore[key] = product

	jsonProduct, err := json.Marshal(product)
	helpers.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonProduct)
}

// HTTP GET -/api/products ürünleri getir
func GetProductsHandler(w http.ResponseWriter, r *http.Request) {

	var productsArr []modals.Product

	for _, product := range productStore {
		productsArr = append(productsArr, product)
	}

	jsonProductsArr, err := json.Marshal(productsArr)
	helpers.CheckError(err)

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	w.Write(jsonProductsArr)

}

// HTTP GET -/api/products/{id} tekil ürün getir
func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	var product modals.Product
	vars := mux.Vars(r)
	paramsProductId, _ := strconv.Atoi(vars["id"])
	for _, value := range productStore {
		if value.ID == paramsProductId {
			product = value
		}
	}
	w.Header().Set("Content-Type", "application/json")
	jsonProduct, err := json.Marshal(product)
	helpers.CheckError(err)
	w.Write(jsonProduct)
	w.WriteHeader(http.StatusOK)

}

// HTTP PUT -api/products/{id} Tekil ürün Güncelleme
func PutProductHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var updatedProduct modals.Product
	vars := mux.Vars(r)
	key := vars["id"]

	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
	helpers.CheckError(err)

	if _, ok := productStore[key]; ok {
		updatedProduct.ID, _ = strconv.Atoi(key)
		updatedProduct.ChangedOn = time.Now()
		delete(productStore, key)
		productStore[key] = updatedProduct
	} else {
		fmt.Printf("Ürün Bulunamadı key : %s", key)
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	data, err := json.Marshal(updatedProduct)
	helpers.CheckError(err)
	w.Write(data)

}

// HTTP Delete -api/products/{id} Tekil ürün silme
func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	key := vars["id"]
	var deletedProduct modals.Product = productStore[key]
	data, err := json.Marshal(deletedProduct)
	helpers.CheckError(err)
	if _, ok := productStore[key]; ok {
		delete(productStore, key)
	} else {
		fmt.Printf("Ürün Bulunamadı %s ", key)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
