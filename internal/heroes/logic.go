package heroes

import (
	"errors"
)

var (
	ErrHeroesNotFound = errors.New("heroes could not be retrieved")
)

type heroesService struct {
	heroesRepository HeroesRepository
}

func NewHeroesService(heroesRepo HeroesRepository) HeroesService {
	return &heroesService{
		heroesRepository: heroesRepo,
	}
}

func (h *heroesService) GetHeroes() ([]byte, error) {
	return h.heroesRepository.GetHeroes()
}
