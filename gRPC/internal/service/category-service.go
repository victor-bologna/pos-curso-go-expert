package service

import (
	"context"
	"io"

	"github.com/pos-curso-go-expert-grpc/internal/database"
	"github.com/pos-curso-go-expert-grpc/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{CategoryDB: categoryDB}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}
	return &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}

func (c *CategoryService) ListCategories(ctx context.Context, in *pb.Blank) (*pb.ListCategoryResponse, error) {
	categoriesDB, err := c.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}

	var categories []*pb.Category

	for _, categoryDB := range categoriesDB {
		categories = append(categories, &pb.Category{
			Id:          categoryDB.ID,
			Name:        categoryDB.Name,
			Description: categoryDB.Description,
		})
	}
	return &pb.ListCategoryResponse{Category: categories}, nil
}

func (c *CategoryService) GetCategory(ctx context.Context, in *pb.CategoryGetRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.FindByCategoryID(in.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}

func (c *CategoryService) StreamCategories(in pb.CategoryService_StreamCategoriesServer) error {
	categories := &pb.ListCategoryResponse{}

	for {
		category, err := in.Recv()
		if err == io.EOF {
			return in.SendAndClose(categories)
		}
		if err != nil {
			return err
		}

		categoryResp, err := c.CreateCategory(context.TODO(), category)
		if err != nil {
			return err
		}
		categories.Category = append(categories.Category, categoryResp)
	}
}

func (c *CategoryService) BiStreamCategories(in pb.CategoryService_BiStreamCategoriesServer) error {
	for {
		category, err := in.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		categoryResp, err := c.CreateCategory(context.TODO(), category)
		if err != nil {
			return err
		}
		err = in.Send(categoryResp)
		if err != nil {
			return err
		}
	}
}
