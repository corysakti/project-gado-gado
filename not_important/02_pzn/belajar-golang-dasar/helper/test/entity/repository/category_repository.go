package repository

import "belajar-golang-dasar/helper/test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
