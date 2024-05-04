package utils

import "net/http"

func StatusOk(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	return
}

func BadRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	return
}

func Notfound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	return
}

func InternalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	return
}

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func UnAuthorized(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
	return
}

func GeneralStatus(w http.ResponseWriter, r *http.Request, code int) {
	w.WriteHeader(code)
	return
}
