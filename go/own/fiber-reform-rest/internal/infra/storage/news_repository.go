package storage

import (
	"fiber-reform-rest/internal/domain/model"

	"gopkg.in/reform.v1"
)

type NewsRepository struct {
	reformDB *reform.DB
}

func NewNewsRepository(reformDB *reform.DB) *NewsRepository {
	return &NewsRepository{
		reformDB: reformDB,
	}
}

func (r *NewsRepository) Load(page, perPage int) []*model.News {
	// offset := (page-1) * perPage
	// limit := perPage
	return []*model.News{
		{
			ID:      1,
			Title:   "aaa",
			Content: "content aaa",
		},
		{
			ID:      2,
			Title:   "bbb",
			Content: "content bbb",
		},
	}
}

func (r *NewsRepository) Save(*model.News) (*model.News, error) {
	return &model.News{
		ID:      3,
		Title:   "ccc",
		Content: "content ccc",
	}, nil
}
