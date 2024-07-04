package services

import (
	"context"
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
	GetBoardById(boardId string) (schemas.BoardSchema, error)
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
	CreatedBy   string
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

func (s *BoardService) GetBoardById(boardId string) (schemas.BoardSchema, error) {
	coll := s.db.Collection(schemas.BoardCollection)

	filter := bson.D{{Key: "_id", Value: boardId}}

	var result schemas.BoardSchema
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return schemas.BoardSchema{}, err
	}

	return result, nil
}
