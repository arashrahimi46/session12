package http

import (
	"net/http"
	"strconv"
	"testMongo/internal/model"
)

type HttpHandler interface {
	GetHomePage(w http.ResponseWriter, r *http.Request)
	AddUser(w http.ResponseWriter, r *http.Request)
}

type httpHandler struct {
	service model.Service
}

func NewHttpHandler(service model.Service) HttpHandler {
	return &httpHandler{
		service: service,
	}
}

func (h *httpHandler) GetHomePage(w http.ResponseWriter, r *http.Request) {
	homepageData := h.service.GetHomePage()
	w.Write([]byte(homepageData))
}

func (h *httpHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	age := r.FormValue("age")
	mobileNo := r.FormValue("mobile_no")
	//cast age from string to int
	numericAge, err := strconv.Atoi(age)
	user := model.User{
		Name:     name,
		Age:      numericAge,
		MobileNo: mobileNo,
	}

	err = h.service.AddUser(user)
	if err != nil {
		w.Write([]byte("cant add user because of error: " + err.Error()))
		return
	}
	w.Write([]byte("successfully added user"))
}
