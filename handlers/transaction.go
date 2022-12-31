package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	dto "project/dto/result"
	transactionsdto "project/dto/transactions.go"
	"project/models"
	"project/repositories"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

var path_file_trans = "http://localhost:5000/uploads/"

type handlerTransaction struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *handlerTransaction {
	return &handlerTransaction{TransactionRepository}
}

func (h *handlerTransaction) FindTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	transaction, err := h.TransactionRepository.FindTransactions()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	for i, p := range transaction {
		transaction[i].Image = path_file_trans + p.Image
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: transaction}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) GetTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	trans, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	trans.Image = path_file_trans + trans.Image

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: trans}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get data user token
	// userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	// userId := int(userInfo["id"].(float64))

	// // middleware upload file
	// dataContex := r.Context().Value("dataFile")
	// filename := dataContex.(string)

	//parse data
	counterQty, _ := strconv.Atoi(r.FormValue("qty"))
	total, _ := strconv.Atoi(r.FormValue("total"))
	tripId, _ := strconv.Atoi(r.FormValue("trip_id"))
	// UserId, _ := strconv.Atoi(r.FormValue("user_id"))

	request := transactionsdto.CreateTransactionRequest{
		CounterQty: counterQty,
		Total:      total,
		// Status:     r.FormValue("status"),
		// Image:      filename,
		TripId: tripId,
		UserId: 88,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Create Unique Transaction Id here ...
	var transIdIsMatch = false
	var transactionId int

	for !transIdIsMatch {
		transactionId = int(time.Now().Unix())
		transaction, _ := h.TransactionRepository.GetTransaction(transactionId)
		if transaction.Id == 0 {
			transIdIsMatch = true
		}
	}

	transaction := models.Transaction{
		Id:         transactionId,
		CounterQty: request.CounterQty,
		Total:      request.Total,
		Status:     "pending",
		// Image:      request.Image,
		TripId: request.TripId,
		UserId: request.UserId,
	}

	// panggil Transaction repository dan masukkan transaction ke dalam kedalam function CreateTransaction
	data, err := h.TransactionRepository.CreateTransaction(transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	// panggil function getTrip agar setelah data di create data id akan keluar response
	transactionResponse, err := h.TransactionRepository.GetTransaction(data.Id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// snap
	var s = snap.Client{}
	s.New("SB-Mid-server-CBYg0a0CWSxQrUrIYbcaHJvM", midtrans.Sandbox)

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transactionResponse.Id),
			GrossAmt: int64(transactionResponse.Total),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: transactionResponse.User.Name,
			Email: transactionResponse.User.Email,
		},
	}

	snapResp, _ := s.CreateTransaction(req)
	fmt.Println(snapResp)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: snapResp}
	json.NewEncoder(w).Encode(response)
}

// function notification
func (h *handlerTransaction) Notification(w http.ResponseWriter, r *http.Request) {
	var notificationPayload map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&notificationPayload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	transactionStatus := notificationPayload["transaction_status"].(string)
	fraudStatus := notificationPayload["fraud_status"].(string)
	orderId := notificationPayload["order_id"].(string)

	orderIdConvert, _ := strconv.Atoi(orderId)

	transaction, _ := h.TransactionRepository.GetTransaction(orderIdConvert)
	fmt.Println(transactionStatus, fraudStatus, orderId, transaction)

	if transactionStatus == "capture" {
		if fraudStatus == "challenge" {

			h.TransactionRepository.UpdateTransaction("pending", transaction.Id)
		} else if fraudStatus == "accept" {

			h.TransactionRepository.UpdateTransaction("success", transaction.Id)
		}
	} else if transactionStatus == "settlement" {

		h.TransactionRepository.UpdateTransaction("success", transaction.Id)
	} else if transactionStatus == "deny" {

		h.TransactionRepository.UpdateTransaction("failed", transaction.Id)
	} else if transactionStatus == "cancel" || transactionStatus == "expire" {

		h.TransactionRepository.UpdateTransaction("failed", transaction.Id)
	} else if transactionStatus == "pending" {

		h.TransactionRepository.UpdateTransaction("pending", transaction.Id)
	}

	w.WriteHeader(http.StatusOK)
}

// func (h *handlerTransaction) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	id, _ := strconv.Atoi(mux.Vars(r)["id"])
// 	transaction, err := h.TransactionRepository.GetTransaction(int(id))
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	// middleware
// 	dataContex := r.Context().Value("dataFile")
// 	filename := dataContex.(string)

// 	// request image agar nantinya image dapat diupdate
// 	request := transactionsdto.UpdateTransactionRequest{
// 		Image: filename,
// 	}

// 	// parse counter qty
// 	counterQty, _ := strconv.Atoi(r.FormValue("qty"))
// 	if counterQty != 0 {
// 		transaction.CounterQty = counterQty
// 	}

// 	// parse counter total
// 	total, _ := strconv.Atoi(r.FormValue("total"))
// 	if total != 0 {
// 		transaction.Total = total
// 	}

// 	// status
// 	if r.FormValue("status") != "" {
// 		transaction.Status = r.FormValue("status")
// 	}

// 	// image
// 	if request.Image != "" {
// 		transaction.Image = request.Image
// 	}

// 	// parse trip id
// 	tripId, _ := strconv.Atoi(r.FormValue("trip_id"))
// 	if tripId != 0 {
// 		transaction.TripId = tripId
// 	}

// 	// panggil Transaction repository dan masukkan transaction ke dalam kedalam function UpdateTransaction
// 	data, err := h.TransactionRepository.UpdateTransaction(transaction)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	// panggil function getTrip agar setelah data di create data id akan keluar response
// 	newtransactionResponse, err := h.TransactionRepository.GetTransaction(data.Id)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTransaction(newtransactionResponse)}
// 	json.NewEncoder(w).Encode(response)
// }

func (h *handlerTransaction) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.TransactionRepository.DeleteTransaction(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTransaction(data)}
	json.NewEncoder(w).Encode(response)
}

func convertResponseTransaction(u models.Transaction) transactionsdto.TransactionResponse {
	return transactionsdto.TransactionResponse{
		Id:         u.Id,
		CounterQty: u.CounterQty,
		Total:      u.Total,
		Status:     u.Status,
		// Image:      u.Image,
		TripId: u.TripId,
	}
}
