package repositories

import (
	"gym-management/classes/models/dtos"
	"gym-management/classes/models/entities"
	"gym-management/classes/models/errors"
	"time"
)

var classes = []entities.Class{
	{
		ClassId: 1,
		Name: "Class1",
		StartDate: time.Date(2024, time.January, 22, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(2024, time.January, 28,  0, 0, 0, 0, time.UTC),
		Capacity: 10,
	},
	{
		ClassId: 2,
		Name: "Class2",
		StartDate: time.Date(2024, time.January, 29, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(2024, time.January, 31,  0, 0, 0, 0, time.UTC),
		Capacity: 20,
	},
}

type ClassesRepositoryImpl struct {
}

func NewClassesRepository() *ClassesRepositoryImpl {
	return &ClassesRepositoryImpl{}
}

func (repo *ClassesRepositoryImpl) GetClasses() *[]dtos.ClassCompleteDTO{
	classesDTO := []dtos.ClassCompleteDTO{}
	for _, class := range classes {
		classesDTO = append(classesDTO, *class.ToClassCompleteDTO())
	}
	return &classesDTO
}

func (repo *ClassesRepositoryImpl) InsertNewClass(classSchedule *entities.Class) *dtos.ClassCompleteDTO{
	classSchedule.ClassId = len(classes) + 1
	classes = append(classes, *classSchedule)
	return (classSchedule.ToClassCompleteDTO())
}

func (repo *ClassesRepositoryImpl) GetClass(id int) (*entities.Class, error){
	for index, class := range classes {
		if class.ClassId == id{
			return &classes[index], nil
		}
	}
	return nil, errors.NewClassNotFoundError()
}

func (repo *ClassesRepositoryImpl) UpdateClass(id int, updatedClass *entities.Class) (*dtos.ClassCompleteDTO){
	//This has already been done in the service, but the ideia is to simulate an insert
	currentClass, _ := repo.GetClass(id)

	currentClass.Capacity = updatedClass.Capacity
	currentClass.StartDate = updatedClass.StartDate
	currentClass.EndDate = updatedClass.EndDate
	currentClass.Name = updatedClass.Name
	
	return currentClass.ToClassCompleteDTO()
}

func (repo *ClassesRepositoryImpl) DeleteClass(id int) (*dtos.ClassCompleteDTO, error) {
	for index, class := range classes {
		if class.ClassId == id{
			deletedClass := class
			classes = append(classes[:index], classes[index+1:]...)
			return deletedClass.ToClassCompleteDTO(), nil
		}
	}
	return nil, errors.NewClassNotFoundError()
}