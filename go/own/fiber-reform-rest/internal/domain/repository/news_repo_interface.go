package repository

import "fiber-reform-rest/internal/domain/model"

type NewsRepository interface {
	Load(page, perPage int) []*model.News
	Save(*model.News) (*model.News, error)
}
