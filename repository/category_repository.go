package repository

import "github.com/fanialfi/golang-testing/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
