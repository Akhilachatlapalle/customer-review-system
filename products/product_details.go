package products

import (
	"customer-review-system/helpers"
	"customer-review-system/models"
	"customer-review-system/store"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-redis/redis"
)

var productList = []*Product{&Product{id: "ProductA"}, &Product{id: "ProductB"}, &Product{id: "ProductC"}, &Product{id: "ProductD"}}

type Product struct {
	id      string
	Reviews []map[string]int
}
type cookieReq struct {
	Cookie string
}

func ProductList(w http.ResponseWriter, r *http.Request) {
	resp := models.Response{}
	req := cookieReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp.Error = "could not read json data"
		resp.Status = http.StatusBadRequest
		helpers.SendResp(w, resp, resp.Status)
		return
	}
	cookieCheck := helpers.IsEmpty(req.Cookie)
	_, err = store.Client.Get(req.Cookie).Result()
	if err == redis.Nil && err != nil {
		resp.Error = fmt.Sprint("Provided wrong cookie")
		resp.Status = http.StatusBadRequest
		helpers.SendResp(w, resp, resp.Status)
		return
	}
	if cookieCheck {
		resp.Error = fmt.Sprint("There is empty data.")
		resp.Status = http.StatusBadRequest
		helpers.SendResp(w, resp, resp.Status)
		return
	}
	list := productList
	idList := []string{}

	for _, each := range list {
		idList = append(idList, each.id)
	}
	resp.Data = fmt.Sprint(idList)
	resp.Status = http.StatusOK
	helpers.SendResp(w, resp, resp.Status)
	return
}

func ProductsRatings(w http.ResponseWriter, r *http.Request) {
	resp := models.Response{}
	req := cookieReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp.Error = "could not read json data"
		resp.Status = http.StatusBadRequest
		helpers.SendResp(w, resp, resp.Status)
		return
	}
	cookieCheck := helpers.IsEmpty(req.Cookie)
	_, err = store.Client.Get(req.Cookie).Result()
	if err == redis.Nil && err != nil {
		resp.Error = fmt.Sprint("Provided wrong cookie")
		resp.Status = http.StatusBadRequest
		helpers.SendResp(w, resp, resp.Status)
		return
	}
	if cookieCheck {
		resp.Error = fmt.Sprint("There is empty data.")
		resp.Status = http.StatusBadRequest
		helpers.SendResp(w, resp, resp.Status)
		return
	}
	pList := productList
	nList := []Product{}

	for _, each := range pList {
		nList = append(nList, *each)
	}
	resp.Data = fmt.Sprint(nList)
	resp.Status = http.StatusOK
	helpers.SendResp(w, resp, resp.Status)
	return
}

type PostReviewReq struct {
	Cookie  string
	Product string
	Rating  int
}

func GiveRating(w http.ResponseWriter, r *http.Request) {
	resp := models.Response{}
	req := PostReviewReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp.Error = "could not read json data"
		resp.Status = http.StatusBadRequest
		helpers.SendResp(w, resp, resp.Status)
		return
	}
	cookieCheck := helpers.IsEmpty(req.Cookie)
	productCheck := helpers.IsEmpty(req.Product)
	ratingCheck := allowedRating(req.Rating)
	u, err := store.Client.Get(req.Cookie).Result()
	if err == redis.Nil && err != nil {
		resp.Error = fmt.Sprint("Provided wrong cookie")
		resp.Status = http.StatusBadRequest
		helpers.SendResp(w, resp, resp.Status)
		return
	}
	if cookieCheck || productCheck {
		resp.Error = fmt.Sprint("There is empty data.")
		resp.Status = http.StatusBadRequest
		helpers.SendResp(w, resp, resp.Status)
		return
	}
	if !ratingCheck {
		resp.Error = fmt.Sprint("Wrong rating. Choose on  the scale of 1 to 5")
		resp.Status = http.StatusBadRequest
		helpers.SendResp(w, resp, resp.Status)
		return
	}

	err = updateRating(req.Product, req.Rating, u)
	if err != nil {
		resp.Error = err.Error()
		resp.Status = http.StatusBadRequest
		helpers.SendResp(w, resp, resp.Status)
		return
	}
	resp.Data = fmt.Sprint("Thanks for your rating!")
	resp.Status = http.StatusOK
	helpers.SendResp(w, resp, resp.Status)
	return
}

func updateRating(product string, rating int, u string) error {
	isExists := false
	for _, each := range productList {
		if product == each.id {
			isExists = true
			if len(each.Reviews) == 0 {
				k := make(map[string]int)
				k[u] = rating
				each.Reviews = append(each.Reviews, k)
			} else {
				for _, k := range each.Reviews {
					k[u] = rating
				}
			}
		}
	}
	if isExists {
		return nil
	}
	return fmt.Errorf("No Product available")
}

func allowedRating(i int) bool {
	if i == 1 || i == 2 || i == 3 || i == 4 || i == 5 {
		return true
	}
	return false
}
