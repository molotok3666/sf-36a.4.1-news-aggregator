package api

import (
	"encoding/json"
	"net/http"
	"news-aggregator/pkg/storage"
	"strconv"

	"github.com/gorilla/mux"
)

// Программный интерфейс сервера GoNews
type API struct {
	db     storage.Interface
	router *mux.Router
}

// Конструктор объекта API
func New(db storage.Interface) *API {
	api := API{
		db: db,
	}
	api.router = mux.NewRouter()
	api.endpoints()
	return &api
}

// Регистрация обработчиков API.
func (api *API) endpoints() {
	api.router.HandleFunc("/news/{count}", api.newsHandler).Methods(http.MethodGet, http.MethodOptions)
	api.router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./webapp")))).Methods(http.MethodGet, http.MethodOptions)
}

// Получение маршрутизатора запросов.
// Требуется для передачи маршрутизатора веб-серверу.
func (api *API) Router() *mux.Router {
	return api.router
}

// Получение новостей из БД.
func (api *API) newsHandler(w http.ResponseWriter, r *http.Request) {
	count, err := strconv.Atoi(mux.Vars(r)["count"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	news, err := api.db.News(count)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	bytes, err := json.Marshal(news)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
	w.Header().Set("Content-Type", "application/json")
}
