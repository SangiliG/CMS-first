package repo

// import (
// 	"github.com/SangiliG/cms-first/pkg/config"
// 	"github.com/SangiliG/cms-first/pkg/models"
// )

// func GetAllMovies(movie *models.Movie, pagination *models.Pagination) (*[]models.Movie, int64, error) {
// 	var movies []models.Movie
// 	var totalRows int64 = 0
// 	offset := (pagination.Page - 1) * pagination.Limit
// 	queryBuider := config.GetDB().Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
// 	result := queryBuider.Model(&models.Movie{}).Where(movie).Find(&movies)
// 	if result.Error != nil {
// 		msg := result.Error
// 		return nil, totalRows, msg
// 	}
// 	errCount := config.GetDB().Model(&models.Movie{}).Count(&totalRows).Error
// 	if errCount != nil {
// 		return nil, totalRows, errCount
// 	}
// 	return &movies, totalRows, nil
// }
