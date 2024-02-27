package code

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/render"
	"io"
	"net/http"
)

type Handlers struct {
	service codeGetter
}

func NewOrderHandler(service codeGetter) *Handlers {
	return &Handlers{service: service}
}

type Request struct {
	OrderID string `json:"order_id"`
}

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
	Codes  []byte `json:"payload,omitempty"`
}

type codeGetter interface {
	GetByOrderID(orderID string) ([]string, error)
}

func (s *Handlers) GetAllCodes(w http.ResponseWriter, r *http.Request) {

	var req Request

	err := json.NewDecoder(r.Body).Decode(&req)
	if errors.Is(err, io.EOF) {
		response := Response{
			Status: "error",
			Error:  err.Error(),
		}
		render.JSON(w, r, response)

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err != nil {
		response := Response{
			Status: "error",
			Error:  err.Error(),
		}
		render.JSON(w, r, response)

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	codes, err := s.service.GetByOrderID(req.OrderID)
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		response := Response{
			Status: "error",
			Error:  err.Error(),
		}
		render.JSON(w, r, response)

		return
	}
	res, err := json.Marshal(codes)
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		response := Response{
			Status: "error",
			Error:  err.Error(),
		}
		render.JSON(w, r, response)
	}

	response := Response{
		Status: "success",
		Codes:  res,
	}
	render.JSON(w, r, response)

}
