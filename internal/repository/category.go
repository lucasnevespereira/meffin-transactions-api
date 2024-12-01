package repository

import (
	"context"
	"fmt"
)

func (r *Repository) CreateCategory(ctx context.Context, rowCategory *RowCategory) (*RowCategory, error) {
	result := r.db.WithContext(ctx).Create(rowCategory)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to create category: %v", result.Error)
	}

	return rowCategory, nil
}

func (r *Repository) GetCategoriesByUserID(ctx context.Context, userId string) ([]*RowCategory, error) {
	var rowCategories []*RowCategory
	result := r.db.WithContext(ctx).Where("user_id = ?", userId).Find(&rowCategories)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get categories: %v", result.Error)
	}

	return rowCategories, nil
}

func (r *Repository) DeleteCategory(ctx context.Context, categoryID int64) error {
	err := r.db.Delete(&RowCategory{}, categoryID).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateCategory(ctx context.Context, updatedCategory *RowCategory) (*RowCategory, error) {
	err := r.db.Model(updatedCategory).Updates(updatedCategory).Error
	if err != nil {
		return nil, err
	}

	return updatedCategory, nil
}
