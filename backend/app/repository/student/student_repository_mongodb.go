package student

import (
	"context"
	"fmt"
	"sample/app/model"
	"sample/app/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBStudentRepository struct {
	collection *mongo.Collection
}

func NewMongoDBStudentRepository() IStudentRepository {
	return &MongoDBStudentRepository{
		collection: utils.StudentCollection,
	}
}

// GetStudents implements IStudentRepository.
func (m *MongoDBStudentRepository) GetStudents() ([]model.Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := m.collection.Find(ctx, bson.M{})
	if err != nil {
		return []model.Student{}, err
	}
	defer cursor.Close(ctx)

	var students []model.Student
	if err = cursor.All(ctx, &students); err != nil {
		return []model.Student{}, err
	}

	return students, nil
}

// AddStudent implements IStudentRepository.
func (m *MongoDBStudentRepository) AddStudent(req model.CreateStudentRequest) (model.Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get the next ID
	students, err := m.GetStudents()
	if err != nil {
		return model.Student{}, err
	}

	nextID := utils.GenerateNextID(students, func(s model.Student) int { return s.ID })

	student := model.Student{
		ID:    nextID,
		Name:  req.Name,
		Email: req.Email,
		Dob:   req.Dob,
		Gpa:   req.Gpa,
	}

	_, err = m.collection.InsertOne(ctx, student)
	if err != nil {
		return model.Student{}, err
	}

	return student, nil
}

// UpdateStudent implements IStudentRepository.
func (m *MongoDBStudentRepository) UpdateStudent(student model.Student) (model.Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"id": student.ID}
	update := bson.M{"$set": student}

	result := m.collection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After))
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return model.Student{}, fmt.Errorf("student with ID %d not found", student.ID)
		}
		return model.Student{}, result.Err()
	}

	var updatedStudent model.Student
	if err := result.Decode(&updatedStudent); err != nil {
		return model.Student{}, err
	}

	return updatedStudent, nil
}

// DeleteStudent implements IStudentRepository.
func (m *MongoDBStudentRepository) DeleteStudent(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"id": id}
	result, err := m.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("student with ID %d not found", id)
	}

	return nil
}

// GetStudentByID implements IStudentRepository.
func (m *MongoDBStudentRepository) GetStudentByID(id int) (model.Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"id": id}
	var student model.Student

	err := m.collection.FindOne(ctx, filter).Decode(&student)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Student{}, fmt.Errorf("student with ID %d not found", id)
		}
		return model.Student{}, err
	}

	return student, nil
}
