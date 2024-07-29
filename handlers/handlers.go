package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/stxreocoma/travels-api/database"
	"github.com/stxreocoma/travels-api/models"

	"github.com/go-chi/chi"
)

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := new(models.User)
	resultDB := database.DB.Db.First(&user, id)
	if resultDB.Error != nil {
		http.Error(w, resultDB.Error.Error(), http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func GetLocationByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	location := new(models.Location)
	resultDB := database.DB.Db.First(&location, id)
	if resultDB.Error != nil {
		http.Error(w, resultDB.Error.Error(), http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(location)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func GetVisitByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	visit := new(models.Visit)
	resultDB := database.DB.Db.First(&visit, id)
	if resultDB.Error != nil {
		http.Error(w, resultDB.Error.Error(), http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(visit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

//https://github.com/MailRuChamps/hlcupdocs/blob/master/2017/TECHNICAL_TASK.md

func GetUserVisits(w http.ResponseWriter, r *http.Request) {
	visits := []models.Visit{}

	id := chi.URLParam(r, "id")

	params, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultDB := database.DB.Db.Where("user=", id).Where("visited_at BETWEEN", params["fromDate"], "AND", params["toDate"]).Where("location.country=", params["country"]).Where("location.distance<", params["distance"]).Find(&visits)
	if resultDB.Error != nil {
		http.Error(w, resultDB.Error.Error(), http.StatusNotFound)
		return
	}

	resp, err := json.Marshal(visits)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func GetLocationAverageMark(w http.ResponseWriter, r *http.Request) {
	var (
		visits []models.Visit
		mark   float64
	)

	id := chi.URLParam(r, "id")

	//	params, err := url.ParseQuery(r.URL.RawQuery)
	//	if err != nil {
	//		log.Println(err.Error())
	//		http.Error(w, err.Error(), http.StatusBadRequest)
	//		return
	//	}

	resultDB := database.DB.Db.Where("location=", id).First(&visits)
	if resultDB.Error != nil {
		http.Error(w, resultDB.Error.Error(), http.StatusNotFound)
		return
	}

	for _, visit := range visits {
		mark += float64(visit.Mark)
	}

	if mark != 0 {
		mark /= float64(len(visits))
	}

	resp, err := json.Marshal(mark)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func UpdateUserInfo(w http.ResponseWriter, r *http.Request) {

}

func UpdateLocationInfo(w http.ResponseWriter, r *http.Request) {

}

func UpdateVisitInfo(w http.ResponseWriter, r *http.Request) {

}

func CreateUser(w http.ResponseWriter, r *http.Request) {

}

func CreateLocation(w http.ResponseWriter, r *http.Request) {

}

func CreateVisit(w http.ResponseWriter, r *http.Request) {

}
