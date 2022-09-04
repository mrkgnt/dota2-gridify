package api

import (
	"log"
	"net/http"

	"github.com/mrkgnt/dota2-gridify/internal/heroes"
	"github.com/pkg/errors"
)

type HeroesHandler interface {
	Get(http.ResponseWriter, *http.Request)
}

func NewHandler(heroesService heroes.HeroesService) HeroesHandler {
	return &handler{
		heroesService: heroesService,
	}
}

type handler struct {
	heroesService heroes.HeroesService
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	data, err := h.heroesService.GetHeroes()
	if err != nil {
		if errors.Cause(err) == heroes.ErrHeroesNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	setupResponse(w, contentType, data, http.StatusMovedPermanently)
}

func setupResponse(w http.ResponseWriter, contentType string, body []byte, statusCode int) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	_, err := w.Write(body)
	if err != nil {
		log.Println(err)
	}
}
