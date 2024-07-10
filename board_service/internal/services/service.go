package services

import (
	"context"
	"errors"
	"time"

	"github.com/sweetloveinyourheart/miro-whiteboard/board_service/internal/db"
	"github.com/sweetloveinyourheart/miro-whiteboard/board_service/internal/schemas"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BoardService struct {
	client *mongo.Client
	db     *mongo.Database
}

type IBoardService interface {
	CreateBoard(newBoard BoardInfo) (bool, error)
	GetBoardById(userId int32, boardId string) (schemas.BoardSchema, error)
	DeleteBoard(userId int32, boardId string) (bool, error)
}

func NewBoardService(client *mongo.Client) IBoardService {
	db := client.Database(db.BoardDatabase)

	return &BoardService{
		client,
		db,
	}
}

type BoardInfo struct {
	Title       string
	Description string
	CreatedBy   int32
}

func (s *BoardService) CreateBoard(board BoardInfo) (bool, error) {
	coll := s.db.Collection(schemas.BoardCollection)

	newBoard := schemas.BoardSchema{
		Title:       board.Title,
		Description: board.Description,
		CreatedBy:   board.CreatedBy,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}

	_, err := coll.InsertOne(context.TODO(), newBoard)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *BoardService) GetBoardById(userId int32, boardId string) (schemas.BoardSchema, error) {
	coll := s.db.Collection(schemas.BoardCollection)

	filter := bson.D{{Key: "_id", Value: boardId}, {Key: "created_by", Value: userId}}

	var result schemas.BoardSchema
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return schemas.BoardSchema{}, err
	}

	return result, nil
}

func (s *BoardService) DeleteBoard(userId int32, boardId string) (bool, error) {
	coll := s.db.Collection(schemas.BoardCollection)

	filter := bson.D{{Key: "_id", Value: boardId}}

	var board schemas.BoardSchema
	err := coll.FindOne(context.TODO(), filter).Decode(board)
	if err != nil {
		return false, errors.New("board not found")
	}

	if board.CreatedBy != userId {
		return false, errors.New("cannot delete a board which is not your own")
	}

	result, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		return false, err
	}

	if result.DeletedCount < 1 {
		return false, errors.New("failed to delete board")
	}

	return true, nil
}
