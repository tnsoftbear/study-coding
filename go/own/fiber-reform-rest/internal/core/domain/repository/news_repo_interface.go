package repository

import "fiber-reform-rest/internal/core/domain/model"

type NewsRepository interface {
	AssignCategory(newsID, catID int64)
	FindByID(id int64) *model.News
	LoadCategoryIDs(newsID int64) []int64
	LoadCollection(page, perPage int) []*model.News
	Save(*model.News) *model.News
	UnassignCategories(newsID int64, skipIDs []int64)
}
