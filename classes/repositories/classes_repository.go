package repositories

import (
	"errors"
	"gym-management/classes/models/dtos"
	"gym-management/classes/models/entities"
	"time"
)

var classes = []entities.ClassSchedule{
	{
		Id: 1,
		Name: "Class1",
		StartDate: time.Date(2024, time.January, 22, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(2024, time.January, 28,  0, 0, 0, 0, time.UTC),
		Capacity: 10,
	},
	{
		Id: 2,
		Name: "Class2",
		StartDate: time.Date(2024, time.January, 22, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(2024, time.January, 28,  0, 0, 0, 0, time.UTC),
		Capacity: 20,
	},
}

type ClassesRepository struct {
	//db connection
}

func (repo *ClassesRepository) GetClassesSchedules() *[]dtos.ClassScheduleCompleteDTO{
	classesDTO := []dtos.ClassScheduleCompleteDTO{}
	for _, class := range classes {
		classesDTO = append(classesDTO, *class.ToClassSheduleDTO())
	}
	return &classesDTO
}

func (repo *ClassesRepository) InsertNewClassSchedule(classSchedule *entities.ClassSchedule) (*dtos.ClassScheduleCompleteDTO, error){
	classes = append(classes, *classSchedule)
	return (classSchedule.ToClassSheduleDTO()), nil
}

func (repo *ClassesRepository) GetClassSchedule(id int) (*entities.ClassSchedule, error){
	for index, class := range classes {
		if class.Id == id{
			return &classes[index], nil
		}
	}
	return nil, errors.New("class not found")
}

func (repo *ClassesRepository) UpdateClassSchedule(id int, updatedClass *entities.ClassSchedule) (*dtos.ClassScheduleCompleteDTO){
	//This has already been done in the service, but the ideia is to simulate an insert
	currentClass, _ := repo.GetClassSchedule(id)

	currentClass.Capacity = updatedClass.Capacity
	currentClass.StartDate = updatedClass.StartDate
	currentClass.EndDate = updatedClass.EndDate
	currentClass.Name = updatedClass.Name
	
	return currentClass.ToClassSheduleDTO()
}

func (repo *ClassesRepository) DeleteClassSchedule(id int) (*dtos.ClassScheduleCompleteDTO, error) {
	for index, class := range classes {
		if class.Id == id{
			deletedClass := class
			classes = append(classes[:index], classes[index+1:]...)
			return deletedClass.ToClassSheduleDTO(), nil
		}
	}
	return nil, errors.New("no booking to delete")
}