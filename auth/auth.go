package auth

import (
	"customer-review-system/helpers"
	"customer-review-system/models"
	"customer-review-system/store"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-redis/redis"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

var uName, pwd, pwdConfirm string

type user struct {
	Username string
	Password string
}

type cookie struct {
	Username string
	UUID     uuid.UUID
}

type SignupReq struct {
	Username        string
	Password        string
	ConfirmPassword string
}

func Signup(w http.ResponseWriter, r *http.Request) {
	resp := models.Response{}
	req := SignupReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp.Error = "could not read json data"
		resp.Status = http.StatusBadRequest
		helpers.SendResp(w, resp, resp.Status)
		return
	}
	// Empty data checking
	uNameCheck := helpers.IsEmpty(req.Username)
	pwdCheck := helpers.IsEmpty(req.Password)
	pwdConfirmCheck := helpers.IsEmpty(req.ConfirmPassword)

	if uNameCheck || pwdCheck || pwdConfirmCheck {
		resp.Error = fmt.Sprintf("There is empty data.")
		resp.Status = http.StatusBadRequest
		helpers.SendResp(w, resp, resp.Status)
		return
	}
	userInfo := user{
		Username: req.Username,
		Password: req.Password,
	}

	json := unMarshal(userInfo)
	if pwd == pwdConfirm {
		// Save to database (username and password)
		if !isExists(req.Username) {
			err := store.Client.Set(req.Username, json, 0).Err()
			if err != nil {
				resp.Status = http.StatusBadRequest
				resp.Error = fmt.Sprint("Couldn't store the details. Please try again!")
				helpers.SendResp(w, resp, resp.Status)
				return
			}
			resp.Status = http.StatusOK
			resp.Error = fmt.Sprint("Registration successful.")
			helpers.SendResp(w, resp, resp.Status)
			return
		}
		resp.Error = fmt.Sprint("UserName already exists. Provide new one.")
		resp.Status = http.StatusBadRequest
		helpers.SendResp(w, resp, resp.Status)
		return
	}

	resp.Status = http.StatusBadRequest
	resp.Error = fmt.Sprint("Password information must be the same.")
	helpers.SendResp(w, resp, resp.Status)
	return

}

func Login(w http.ResponseWriter, r *http.Request) {
	resp := models.Response{}
	req := user{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp.Error = "could not read json data"
		resp.Status = http.StatusBadRequest
		helpers.SendResp(w, resp, resp.Status)
		return
	}
	// Empty data checking
	uNameCheck := helpers.IsEmpty(req.Username)
	pwdCheck := helpers.IsEmpty(req.Password)

	if uNameCheck || pwdCheck {
		resp.Error = fmt.Sprint("There is empty data.")
		resp.Status = http.StatusBadRequest
		helpers.SendResp(w, resp, resp.Status)
		return
	}
	var usr = user{}
	val, err := store.Client.Get(req.Username).Result()
	if err == redis.Nil {
		resp.Error = fmt.Sprint("Login failed, UserName doesnot exist in the system.")
		resp.Status = http.StatusBadRequest
		helpers.SendResp(w, resp, resp.Status)
		return
	} else if err != nil {
		resp.Error = fmt.Sprint("Some error occured! please try again: ", err)
		resp.Status = http.StatusBadRequest
		helpers.SendResp(w, resp, resp.Status)
		return
	}

	err = json.Unmarshal([]byte(val), &usr)
	if err != nil {
		resp.Error = fmt.Sprint("Unmarshal error!")
		resp.Status = http.StatusBadRequest
		helpers.SendResp(w, resp, resp.Status)
		return
	}
	ok := strings.EqualFold(usr.Password, req.Password)
	if !ok {
		resp.Error = fmt.Sprint("Provided Wrong password, ", req.Password)
		resp.Status = http.StatusBadRequest
		helpers.SendResp(w, resp, resp.Status)
		return
	}
	uuid, _ := uuid.NewV4()
	err = Store(uuid.String(), req.Username)
	if err != nil {
		resp.Error = fmt.Sprint("Unmarshal error!")
		resp.Status = http.StatusBadRequest
		helpers.SendResp(w, resp, resp.Status)
		return
	}

	resp.Data = fmt.Sprint("Cookie to use is: ", uuid)
	resp.Status = http.StatusOK
	helpers.SendResp(w, resp, resp.Status)
	return
}

func isExists(uName string) bool {
	_, err := store.Client.Get(uName).Result()
	if err == redis.Nil {
		return false
	}
	logrus.Error("Username already Exists")
	return true
}

func unMarshal(userInfo user) []byte {
	json, err := json.Marshal(userInfo)
	if err != nil {
		panic(err)
	}
	return json
}

func Store(c, u string) error {
	err := store.Client.Set(c, u, 5*time.Minute).Err()
	if err != nil {
		return err
	}
	return nil
}
