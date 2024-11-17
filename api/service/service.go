package service

import (
	"context"
	"errors"
	"kompre/api/repository"
	"kompre/models"
)

type CrudService interface {
	Create(ctx context.Context, mahasiswa models.KinerjaCrud) (models.KinerjaCrud, error)
	Get(ctx context.Context) ([]models.KinerjaCrud, error)
	Update(ctx context.Context, id int64, kinerja models.KinerjaCrud) error
	Delete(ctx context.Context, id int64) error
}

type CrudServiceImpl struct {
	repo repository.CrudRepository
}

// Create implements CrudService.
func (c *CrudServiceImpl) Create(ctx context.Context, kinerja models.KinerjaCrud) (models.KinerjaCrud, error) {

	kinerja.IndikatorKinerjaUtama = int(0.10*float64(kinerja.Kehadiran) +
		0.45*float64(kinerja.HasilKerja) +
		0.25*float64(kinerja.Inisiatif) +
		0.20*float64(kinerja.TeamWork))

	if kinerja.IndikatorKinerjaUtama >= 5 {
		kinerja.Penilaian = "Sangat Baik"
	} else if kinerja.IndikatorKinerjaUtama >= 4 {
		kinerja.Penilaian = "Baik"
	} else if kinerja.IndikatorKinerjaUtama >= 3 {
		kinerja.Penilaian = "Cukup Baik"
	} else if kinerja.IndikatorKinerjaUtama >= 2 {
		kinerja.Penilaian = "Kurang Baik"
	} else {
		kinerja.Penilaian = "Sangat Buruk"
	}

	createdKinerja, err := c.repo.Create(ctx, kinerja)
	if err != nil {
		return models.KinerjaCrud{}, errors.New("failed to create KinerjaCrud: " + err.Error())
	}

	return createdKinerja, nil

}

// Get implements CrudService.
func (c *CrudServiceImpl) Get(ctx context.Context) ([]models.KinerjaCrud, error) {
	return c.repo.Get(ctx)
}

// Update implements CrudService.
func (c *CrudServiceImpl) Update(ctx context.Context, id int64, kinerja models.KinerjaCrud) error {

	kinerja.IndikatorKinerjaUtama = int(0.10*float64(kinerja.Kehadiran) +
		0.45*float64(kinerja.HasilKerja) +
		0.25*float64(kinerja.Inisiatif) +
		0.20*float64(kinerja.TeamWork))

	if kinerja.IndikatorKinerjaUtama >= 5 {
		kinerja.Penilaian = "Sangat Baik"
	} else if kinerja.IndikatorKinerjaUtama >= 4 {
		kinerja.Penilaian = "Baik"
	} else if kinerja.IndikatorKinerjaUtama >= 3 {
		kinerja.Penilaian = "Cukup Baik"
	} else if kinerja.IndikatorKinerjaUtama >= 2 {
		kinerja.Penilaian = "Kurang Baik"
	} else {
		kinerja.Penilaian = "Sangat Buruk"
	}

	// Lakukan update ke database
	err := c.repo.Update(ctx, id, kinerja)
	if err != nil {
		return errors.New("failed to update KinerjaCrud: " + err.Error())
	}

	return nil
}

// Delete implements CrudService.
func (c *CrudServiceImpl) Delete(ctx context.Context, id int64) error {
	return c.repo.Delete(ctx, id)
}

func NewCrudService(repo repository.CrudRepository) CrudService {
	return &CrudServiceImpl{repo: repo}
}
