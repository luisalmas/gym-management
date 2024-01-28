package repositories

import (
	"fmt"
	"gym-management/models/entities"
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

func (repo *ClassesRepository) GetClassesSchedules() *[]entities.ClassSchedule{
	return &classes
}

func (repo *ClassesRepository) InsertNewClassSchedule(classSchedule *entities.ClassSchedule) (*entities.ClassSchedule, error){
	classes = append(classes, *classSchedule)
	return classSchedule, nil
}

func (repo *ClassesRepository) GetClassSchedule(id int) (*entities.ClassSchedule, error){
	for index, class := range classes {
		if class.Id == id{
			return &classes[index], nil
		}
	}
	return nil, fmt.Errorf("class not found")
}

func (repo *ClassesRepository) UpdateClassSchedule(id int, updatedClass *entities.ClassSchedule) (*entities.ClassSchedule){
	currentClass, _ := repo.GetClassSchedule(id)

	currentClass.Capacity = updatedClass.Capacity
	currentClass.StartDate = updatedClass.StartDate
	currentClass.EndDate = updatedClass.EndDate
	currentClass.Name = updatedClass.Name
	
	return currentClass
}