package repositories

import (
	"gym-management/classes/models/dtos"
	"gym-management/classes/models/entities"

	"github.com/stretchr/testify/mock"
)

type MockClassesRepository struct {
	mock.Mock
}

func (mockClassesRepo *MockClassesRepository) GetClassesSchedules() *[]dtos.ClassCompleteDTO {
	args := mockClassesRepo.Called()
	return args.Get(0).(*[]dtos.ClassCompleteDTO)
}
func (mockClassesRepo *MockClassesRepository) InsertNewClassSchedule(classSchedule *entities.Class) (*dtos.ClassCompleteDTO) {
	args := mockClassesRepo.Called(classSchedule)
	return args.Get(0).(*dtos.ClassCompleteDTO)
}
func (mockClassesRepo *MockClassesRepository) GetClassSchedule(id int) (*entities.Class, error) {
	args := mockClassesRepo.Called(id)
	if args.Get(0) != nil && args.Get(1) != nil {
		return args.Get(0).(*entities.Class), args.Get(1).(error)
	}

	if args.Get(1) == nil{
		return args.Get(0).(*entities.Class), nil
	}

	if args.Get(0) == nil{
		return nil, args.Get(1).(error)
	}

	return nil, nil
}
func (mockClassesRepo *MockClassesRepository) UpdateClassSchedule(id int, updatedClass *entities.Class) *dtos.ClassCompleteDTO {
	args := mockClassesRepo.Called(id, updatedClass)
	return args.Get(0).(*dtos.ClassCompleteDTO)
}
func (mockClassesRepo *MockClassesRepository) DeleteClassSchedule(id int) (*dtos.ClassCompleteDTO, error) {
	args := mockClassesRepo.Called(id)
	return args.Get(0).(*dtos.ClassCompleteDTO), args.Get(1).(error)
}