package repository

import "hello-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
