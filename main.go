package main

import (
	"customer-review-system/auth"
	"customer-review-system/products"
	"customer-review-system/store"
	"fmt"
	"net/http"
)

func init() {
	store.Client = store.InitDB()
}
func main() {
	ApplicationPort := "8080"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hey Welcome! Please Signup! using: /signup")
	})
	http.HandleFunc("/signup", auth.Signup)
	http.HandleFunc("/login", auth.Login)
	http.HandleFunc("/products", products.ProductList)
	http.HandleFunc("/product/rating", products.GiveRating)
	http.HandleFunc("/products/ratings", products.ProductsRatings)
	http.ListenAndServe(":"+ApplicationPort, nil)
}
