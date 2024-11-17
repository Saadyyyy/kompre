package repository

import (
	"context"
	"kompre/models"

	"gorm.io/gorm"
)

type CrudRepository interface {
	Create(ctx context.Context, mahasiswa models.KinerjaCrud) (models.KinerjaCrud, error)
	Get(ctx context.Context) ([]models.KinerjaCrud, error)
	Update(ctx context.Context, id int64, data models.KinerjaCrud) error
	Delete(ctx context.Context, id int64) error
}

type CrudRepositoryImpl struct {
	db *gorm.DB
}

func NewCrudRepository(db *gorm.DB) CrudRepository {
	return &CrudRepositoryImpl{db: db}
}

// Create implements CrudRepository.
func (c *CrudRepositoryImpl) Create(ctx context.Context, mahasiswa models.KinerjaCrud) (models.KinerjaCrud, error) {
	if err := c.db.WithContext(ctx).Create(&mahasiswa).Error; err != nil {
		return models.KinerjaCrud{}, err
	}
	return mahasiswa, nil
}

// Get implements CrudRepository.
func (c *CrudRepositoryImpl) Get(ctx context.Context) ([]models.KinerjaCrud, error) {
	var mahasiswas []models.KinerjaCrud
	if err := c.db.WithContext(ctx).Find(&mahasiswas).Error; err != nil {
		return nil, err
	}
	return mahasiswas, nil
}

// Update implements CrudRepository.
func (c *CrudRepositoryImpl) Update(ctx context.Context, id int64, data models.KinerjaCrud) error {
	if err := c.db.WithContext(ctx).Model(&models.KinerjaCrud{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

// Delete implements CrudRepository.
func (c *CrudRepositoryImpl) Delete(ctx context.Context, id int64) error {
	if err := c.db.WithContext(ctx).Delete(&models.KinerjaCrud{}, id).Error; err != nil {
		return err
	}
	return nil
}
