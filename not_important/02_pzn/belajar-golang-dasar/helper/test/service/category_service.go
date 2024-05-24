package service

import (
	"belajar-golang-dasar/helper/test/entity"
	"belajar-golang-dasar/helper/test/entity/repository"
	"errors"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) GetCategoryById(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return category, errors.New("Category not found")
	} else {
		return category, nil
	}
}
