package handler

import "github.com/ant0ine/go-json-rest/rest"

func GetAllReminders(w rest.ResponseWriter, r *rest.Request) {
	reminders := []entities.Reminder{}
	// i.DB.Find(&reminders)
	w.WriteJson(&reminders)
}
