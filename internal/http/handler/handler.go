package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/vandi37/Calculator/pkg/calc_service"
)

const (
	NotFound         = "page not found"
	InvalidBody      = "invalid body"
	MethodNotAllowed = "method not allowed"
)

type Handler struct {
	*http.ServeMux
	calc *calc_service.Calculator
}

func SendJson(w io.Writer, v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	return err
}

func NewHandler(path string, calc *calc_service.Calculator) *Handler {
	res := &Handler{http.NewServeMux(), calc}
	res.HandleFunc(path, ContentType(CheckPath(path, CheckMethod(http.MethodPost, res.CalcHandler))))
	res.HandleFunc("/coffee/", ContentType(SendCoffeeHandler))
	if path != "/" {
		res.HandleFunc("/", ContentType(NotFoundHandler))
	}
	return res
}
