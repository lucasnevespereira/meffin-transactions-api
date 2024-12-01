package services

import (
	"context"
	"fmt"
	"meffin-transactions-api/internal/models"
	"meffin-transactions-api/internal/repository"
	"strconv"
)

type ICategoryService interface {
	Create(ctx context.Context, request models.CreateCategoryRequest) (*models.Category, error)
	GetUserCategories(ctx context.Context, userId string) ([]*models.Category, error)
	DeleteCategory(ctx context.Context, categoryID string) error
	UpdateCategory(ctx context.Context, category *models.Category) (*models.Category, error)
}

type CategoryService struct {
	repository repository.IRepository
}

func NewCategoryService(transactionRepository repository.IRepository) ICategoryService {
	return &CategoryService{
		repository: transactionRepository,
	}
}

func (s *CategoryService) Create(ctx context.Context, request models.CreateCategoryRequest) (*models.Category, error) {
	createdCategory, err := s.repository.CreateCategory(ctx, &repository.RowCategory{
		UserID: request.UserID,
		Name:   request.Name,
		Type:   request.Type,
		Color:  request.Color,
	})
	if err != nil {
		return nil, err
	}

	return &models.Category{
		ID:     strconv.FormatInt(createdCategory.ID, 10),
		UserID: createdCategory.UserID,
		Name:   createdCategory.Name,
		Type:   createdCategory.Type,
		Color:  createdCategory.Color,
	}, nil
}

func (s *CategoryService) GetUserCategories(ctx context.Context, userId string) ([]*models.Category, error) {
	rowCategories, err := s.repository.GetCategoriesByUserID(ctx, userId)
	if err != nil {
		return nil, err
	}

	return toCategories(rowCategories), nil
}

func (s *CategoryService) DeleteCategory(ctx context.Context, categoryID string) error {
	id, err := strconv.ParseInt(categoryID, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid category ID: %v", err)
	}
	return s.repository.DeleteCategory(ctx, id)
}

func (s *CategoryService) UpdateCategory(ctx context.Context, category *models.Category) (*models.Category, error) {
	categoryID, err := strconv.ParseInt(category.ID, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid category ID: %v", err)
	}

	updatedCategory, err := s.repository.UpdateCategory(ctx, &repository.RowCategory{
		ID:     categoryID,
		UserID: category.UserID,
		Name:   category.Name,
		Type:   category.Type,
		Color:  category.Color,
	})
	if err != nil {
		return nil, err
	}

	return &models.Category{
		ID:     strconv.FormatInt(updatedCategory.ID, 10),
		UserID: updatedCategory.UserID,
		Name:   updatedCategory.Name,
		Type:   updatedCategory.Type,
		Color:  updatedCategory.Color,
	}, nil
}

func toCategories(rowCategories []*repository.RowCategory) []*models.Category {
	var categories []*models.Category

	for _, rowCategory := range rowCategories {
		categories = append(categories, &models.Category{
			ID:     strconv.FormatInt(rowCategory.ID, 10),
			UserID: rowCategory.UserID,
			Name:   rowCategory.Name,
			Type:   rowCategory.Type,
			Color:  rowCategory.Color,
		})
	}

	return categories
}
