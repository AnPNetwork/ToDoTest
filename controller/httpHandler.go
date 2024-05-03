package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"test/domain"
)

func (s HttpServer) MainHandler(w http.ResponseWriter, r *http.Request) {
	//p := path.Dir("./template/index.html")
	// set header
	w.Header().Set("Content-type", "text/html")
	//http.ServeFile(w, r, p)
	http.ServeFile(w, r, "./template/index.html")
}

func (s HttpServer) GetHandler(w http.ResponseWriter, r *http.Request) {

	result, err := s.db.GetTODOList()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ошибка sql (" + err.Error() + ")"))
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	jsonData, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ошибка конвертации в json (" + err.Error() + ")"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (s HttpServer) PostHandler(w http.ResponseWriter, r *http.Request) {

	type Data struct {
		Description string `json:"description"`
	}

	var resp Data

	// читаем тело запроса
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ошибка конвертации в json (" + err.Error() + ")"))
		return
	}

	defer r.Body.Close()

	// если все нормально - пишем по указателю в структуру
	err = json.Unmarshal(body, &resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ошибка конвертации из json (" + err.Error() + ")"))
		return
	}

	resp.Description = strings.TrimSpace(resp.Description)

	if resp.Description == "" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Description не может быть пустым"))
		return
	}

	if err = s.db.AddTODO(domain.TODO{Description: resp.Description}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ошибка sql (" + err.Error() + ")"))
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s HttpServer) PutHandler(w http.ResponseWriter, r *http.Request) {

	var resp domain.TODO

	// читаем тело запроса
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ошибка конвертации в json (" + err.Error() + ")"))
		return
	}

	defer r.Body.Close()

	// если все нормально - пишем по указателю в структуру
	err = json.Unmarshal(body, &resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ошибка конвертации из json (" + err.Error() + ")"))
		return
	}

	if resp.Id <= 0 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Id не может быть меньше или равно 0"))
		return
	}

	resp.Description = strings.TrimSpace(resp.Description)
	if resp.Description == "" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Description не может быть пустым"))
		return
	}

	if _, err = s.db.UpdateTODO(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ошибка sql (" + err.Error() + ")"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s HttpServer) DeleteHandler(w http.ResponseWriter, r *http.Request) {

	type Data struct {
		Id uint64 `json:"id"`
	}

	var resp Data

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ошибка конвертации в json (" + err.Error() + ")"))
		return
	}

	defer r.Body.Close()

	err = json.Unmarshal(body, &resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ошибка конвертации из json (" + err.Error() + ")"))
		return
	}

	if err = s.db.DeleteTODO(resp.Id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ошибка sql (" + err.Error() + ")"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
