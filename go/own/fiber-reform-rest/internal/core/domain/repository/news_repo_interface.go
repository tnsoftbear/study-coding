package repository

import "fiber-reform-rest/internal/core/domain/model"

type NewsRepository interface {
	LoadCollection(page, perPage int) []*model.News
	Save(*model.News) *model.News
	FindByID(id int64) *model.News
	SaveCategory(newsID, catID int64)
	LoadCategoryIDs(newsID int64) []int64
}
