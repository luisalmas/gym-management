package repositories

import (
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

func (repo *ClassesRepository) InsertNewClassSchedule(classScheduleDTO *dtos.ClassScheduleDTO) (*entities.ClassSchedule, error){
	entity := &entities.ClassSchedule{}
	classScheduleEntity, error := entity.New(classScheduleDTO)

	if error != nil{
		return nil, error
	}

	classes = append(classes, *classScheduleEntity)
	return classScheduleEntity, nil
}