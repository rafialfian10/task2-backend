package handlers

import (
	"encoding/json"
	"net/http"
	dto "project/dto/result"
	tripsdto "project/dto/trips"
	"project/models"
	"project/repositories"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

var path_file = "http://localhost:3000/uploads/"

type handlerTrip struct {
	TripRepository repositories.TripRepository
}

func HandlerTrip(TripRepository repositories.TripRepository) *handlerTrip {
	return &handlerTrip{TripRepository}
}

func (h *handlerTrip) FindTrips(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	trips, err := h.TripRepository.FindTrips()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	for i, p := range trips {
		trips[i].Image = path_file + p.Image
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: trips}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTrip) GetTrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	trip, err := h.TripRepository.GetTrip(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	trip.Image = path_file + trip.Image

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: trip}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTrip) CreateTrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContex := r.Context().Value("dataFile") // add this code
	filename := dataContex.(string)

	//parse data
	countryId, _ := strconv.Atoi(r.FormValue("country_id"))
	day, _ := strconv.Atoi(r.FormValue("day"))
	night, _ := strconv.Atoi(r.FormValue("night"))
	price, _ := strconv.Atoi(r.FormValue("price"))
	quota, _ := strconv.Atoi(r.FormValue("quota"))

	request := tripsdto.CreateTripRequest{
		Title:          r.FormValue("title"),
		CountryId:      countryId,
		Accomodation:   r.FormValue("accomodation"),
		Transportation: r.FormValue("transportation"),
		Eat:            r.FormValue("eat"),
		Day:            day,
		Night:          night,
		DateTrip:       r.FormValue("datetrip"),
		Price:          price,
		Quota:          quota,
		Description:    r.FormValue("description"),
		Image:          filename,
	}

	// if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// parse DateTrip menjadi string
	date, _ := time.Parse("2 January 2006", request.DateTrip)

	trip := models.Trip{
		Title:          request.Title,
		CountryId:      request.CountryId,
		Accomodation:   request.Accomodation,
		Transportation: request.Transportation,
		Eat:            request.Eat,
		Day:            request.Day,
		Night:          request.Night,
		DateTrip:       date,
		Price:          request.Price,
		Quota:          request.Quota,
		Description:    request.Description,
		Image:          request.Image,
	}

	data, err := h.TripRepository.CreateTrip(trip)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTrip(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTrip) UpdateTrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(tripsdto.UpdateTripRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	trip, err := h.TripRepository.GetTrip(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Title != "" {
		trip.Title = request.Title
	}

	if request.Accomodation != "" {
		trip.Accomodation = request.Accomodation
	}

	if request.Transportation != "" {
		trip.Transportation = request.Transportation
	}

	if request.Eat != "" {
		trip.Eat = request.Eat
	}

	if request.Day != 0 {
		trip.Day = request.Day
	}

	if request.Night != 0 {
		trip.Night = request.Night
	}

	if request.DateTrip != "" {
		trip.DateTrip, _ = time.Parse("2 January 2006", request.DateTrip)
	}

	if request.Price != 0 {
		trip.Price = request.Price
	}

	if request.Quota != 0 {
		trip.Quota = request.Quota
	}

	if request.Description != "" {
		trip.Description = request.Description
	}

	if request.Image != "" {
		trip.Image = request.Image
	}

	data, err := h.TripRepository.UpdateTrip(trip)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTrip(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTrip) DeleteTrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user, err := h.TripRepository.GetTrip(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.TripRepository.DeleteTrip(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTrip(data)}
	json.NewEncoder(w).Encode(response)
}

func convertResponseTrip(u models.Trip) tripsdto.TripResponse {
	return tripsdto.TripResponse{
		Id:             u.Id,
		Title:          u.Title,
		CountryId:      u.CountryId,
		Accomodation:   u.Accomodation,
		Transportation: u.Transportation,
		Eat:            u.Eat,
		Day:            u.Day,
		Night:          u.Night,
		DateTrip:       u.DateTrip,
		Price:          u.Price,
		Quota:          u.Quota,
		Description:    u.Description,
		Image:          u.Image,
	}
}
