package user

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

type MongoDBUserRepository struct {
	collection *mongo.Collection
}

func NewMongoDBUserRepository() IUserRepository {
	return &MongoDBUserRepository{
		collection: utils.UserCollection,
	}
}

// Register implements IUserRepository.
func (m *MongoDBUserRepository) Register(user model.User) (model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if username already exists
	_, err := m.GetUserByUsername(user.Username)
	if err == nil {
		return model.User{}, fmt.Errorf("username %s already exists", user.Username)
	}

	// Get the next ID
	users, err := m.getAllUsers()
	if err != nil {
		return model.User{}, err
	}

	nextID := utils.GenerateNextID(users, func(u model.User) int { return u.ID })
	user.ID = nextID

	_, err = m.collection.InsertOne(ctx, user)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

// Login implements IUserRepository.
func (m *MongoDBUserRepository) Login(username string, password string) (model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"username": username,
		"password": password,
	}

	var user model.User
	err := m.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.User{}, fmt.Errorf("invalid username or password")
		}
		return model.User{}, err
	}

	return user, nil
}

// GetUserByID implements IUserRepository.
func (m *MongoDBUserRepository) GetUserByID(id int) (model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"id": id}
	var user model.User

	err := m.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.User{}, fmt.Errorf("user with ID %d not found", id)
		}
		return model.User{}, err
	}

	return user, nil
}

// GetUserByUsername implements IUserRepository.
func (m *MongoDBUserRepository) GetUserByUsername(username string) (model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"username": username}
	var user model.User

	err := m.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.User{}, fmt.Errorf("user with username %s not found", username)
		}
		return model.User{}, err
	}

	return user, nil
}

// UpdateUser implements IUserRepository.
func (m *MongoDBUserRepository) UpdateUser(user model.User) (model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"id": user.ID}
	update := bson.M{"$set": user}

	result := m.collection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After))
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return model.User{}, fmt.Errorf("user with ID %d not found", user.ID)
		}
		return model.User{}, result.Err()
	}

	var updatedUser model.User
	if err := result.Decode(&updatedUser); err != nil {
		return model.User{}, err
	}

	return updatedUser, nil
}

// DeleteUser implements IUserRepository.
func (m *MongoDBUserRepository) DeleteUser(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"id": id}
	result, err := m.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("user with ID %d not found", id)
	}

	return nil
}

// getAllUsers is a helper method to get all users for ID generation
func (m *MongoDBUserRepository) getAllUsers() ([]model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := m.collection.Find(ctx, bson.M{})
	if err != nil {
		return []model.User{}, err
	}
	defer cursor.Close(ctx)

	var users []model.User
	if err = cursor.All(ctx, &users); err != nil {
		return []model.User{}, err
	}

	return users, nil
}
