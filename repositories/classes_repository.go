package repositories

import (
	"fmt"
	"gym-management/models/dtos"
	"gym-management/models/entities"
	"time"
)

var classes = []entities.ClassSchedule{
	{
		Id: 1,
		Name: "Class1",
		Start_date: time.Date(2024, time.January, 22, 0, 0, 0, 0, time.UTC),
		End_date: time.Date(2024, time.January, 28,  0, 0, 0, 0, time.UTC),
		Capacity: 10,
	},
	{
		Id: 2,
		Name: "Class2",
		Start_date: time.Date(2024, time.January, 22, 0, 0, 0, 0, time.UTC),
		End_date: time.Date(2024, time.January, 28,  0, 0, 0, 0, time.UTC),
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
	for _, class := range classes {
		if class.Id == id{
			return &class, nil
		}
	}
	return nil, fmt.Errorf("class not found")
}

func (repo *ClassesRepository) UpdateClassSchedule(id int, updatedClass *dtos.ClassScheduleDTO) (*entities.ClassSchedule, error){
	currentClass, _ := repo.GetClassSchedule(id)

	currentClass.Capacity = updatedClass.Capacity
	currentClass.Start_date = updatedClass.Start_date
	currentClass.End_date = updatedClass.End_date
	currentClass.Name = updatedClass.Name

	return currentClass, nil
}