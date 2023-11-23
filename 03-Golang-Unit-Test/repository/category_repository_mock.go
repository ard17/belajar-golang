package repository

import (
	"golang-unit-test/entity"

	"github.com/stretchr/testify/mock"
)

type CategiryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategiryRepositoryMock) FindById(id string) *entity.Category {
	arguments := repository.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil
	} else {
		category := arguments.Get(0).(entity.Category)
		return &category
	}
}
