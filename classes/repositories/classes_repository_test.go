package repositories

import (
	"gym-management/classes/models/dtos"
	"gym-management/classes/models/entities"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClassesRepository(t *testing.T) {

	var firstClass = &entities.Class{
			ClassId: 1,
			Name: "Class1",
			StartDate: time.Date(2024, time.January, 22, 0, 0, 0, 0, time.UTC),
			EndDate: time.Date(2024, time.January, 28,  0, 0, 0, 0, time.UTC),
			Capacity: 10,
		}

	var secondClass = &entities.Class{
			ClassId: 2,
			Name: "Class2",
			StartDate: time.Date(2024, time.January, 29, 0, 0, 0, 0, time.UTC),
			EndDate: time.Date(2024, time.January, 31,  0, 0, 0, 0, time.UTC),
			Capacity: 20,
		}

	var insertClass = &entities.Class{
		ClassId: 3,
		Name: "Class3",
		StartDate: time.Date(2024, time.February, 01, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(2024, time.January, 20,  0, 0, 0, 0, time.UTC),
		Capacity: 30,
	}

	var updateClass = &entities.Class{
		ClassId: 1,
		Name: "Class3",
		StartDate: time.Date(2024, time.February, 01, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(2024, time.January, 20,  0, 0, 0, 0, time.UTC),
		Capacity: 30,
	}

	//===================== GetClasses tests ==============================================

	t.Run("GetClasses", func(t *testing.T){
		classesRepository := NewClassesRepository()
		classes := classesRepository.GetClasses()

		assert.Equal(t, []dtos.ClassCompleteDTO{*firstClass.ToClassCompleteDTO(), *secondClass.ToClassCompleteDTO()}, *classes)
	})

	//===================== GetClass tests ==============================================

	t.Run("GetClass", func(t *testing.T){
		classesRepository := NewClassesRepository()
		class, err := classesRepository.GetClass(1)

		assert.Nil(t, err)
		assert.Equal(t, firstClass, class)
	})

	t.Run("GetClassNotFound", func(t *testing.T){
		classesRepository := NewClassesRepository()
		class, err := classesRepository.GetClass(0)

		assert.NotNil(t, err)
		assert.Nil(t, class)
	})

	//===================== InsertClass tests ==============================================

	t.Run("InsertNewClass", func(t *testing.T){
		classesRepository := NewClassesRepository()
		class := classesRepository.InsertNewClass(insertClass)

		assert.Equal(t, insertClass.ToClassCompleteDTO(), class)

		getClass, err := classesRepository.GetClass(insertClass.ClassId)

		assert.Nil(t, err)
		assert.Equal(t, insertClass, getClass)
	})

	//===================== UpdateClass tests ==============================================

	t.Run("UpdateClass", func(t *testing.T){
		classesRepository := NewClassesRepository()
		class := classesRepository.UpdateClass(updateClass.ClassId,updateClass)

		assert.Equal(t, updateClass .ToClassCompleteDTO(), class)

		getClass, err := classesRepository.GetClass(updateClass.ClassId)

		assert.Nil(t, err)
		assert.Equal(t, updateClass, getClass)
	})

	//===================== DeleteClass tests ==============================================

	t.Run("DeleteClass", func(t *testing.T){
		classesRepository := NewClassesRepository()
		deletedClass, err := classesRepository.DeleteClass(2)

		assert.Nil(t, err)
		assert.Equal(t, secondClass.ToClassCompleteDTO(), deletedClass)
	})

	t.Run("DeleteClassNotFound", func(t *testing.T){
		classesRepository := NewClassesRepository()
		deletedClass, err := classesRepository.DeleteClass(0)

		assert.NotNil(t, err)
		assert.Nil(t, deletedClass)
	})
}