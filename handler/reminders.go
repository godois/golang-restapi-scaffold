package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/godois/golang-restapi-scaffold/entities"
)

func GetAllReminders(w rest.ResponseWriter, r *rest.Request) {
	obj := entities.Reminder{
		Id:        1,
		Message:   "Testando",
		CreatedAt: time.Now(),
		DeletedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	obj1 := entities.Reminder{
		Id:        1,
		Message:   "Testando2",
		CreatedAt: time.Now(),
		DeletedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	list := []entities.Reminder{}
	list = append(list, obj, obj1)

	w.WriteHeader(http.StatusOK)
	w.WriteJson(&list)
}

func GetReminder(w rest.ResponseWriter, r *rest.Request) {
	obj := entities.Reminder{
		Id:        1,
		Message:   "Testando",
		CreatedAt: time.Now(),
		DeletedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	w.WriteHeader(http.StatusOK)
	w.WriteJson(&obj)
}

func PostReminder(w rest.ResponseWriter, r *rest.Request) {
	reminder := entities.Reminder{}
	if err := r.DecodeJsonPayload(&reminder); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.WriteJson(&reminder)

}

func PutReminder(w rest.ResponseWriter, r *rest.Request) {
	reminder := entities.Reminder{}
	if err := r.DecodeJsonPayload(&reminder); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.WriteJson(&reminder)
}

func DeleteReminder(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err == nil {
		reminder := entities.Reminder{
			Id:        i,
			Message:   "Testando",
			CreatedAt: time.Now(),
			DeletedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		w.WriteJson(&reminder)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

}
