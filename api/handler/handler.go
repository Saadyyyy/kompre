package handler

import (
	"context"
	"kompre/api/service"
	"kompre/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CrudHandler struct {
	service service.CrudService
}

func NewCrudHandler(service service.CrudService) *CrudHandler {
	return &CrudHandler{service: service}
}

func (h *CrudHandler) Create(c echo.Context) error {
	var mahasiswa models.KinerjaCrud

	if err := c.Bind(&mahasiswa); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
	}

	createdMahasiswa, err := h.service.Create(context.Background(), mahasiswa)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create mahasiswa"})
	}

	return c.JSON(http.StatusOK, createdMahasiswa)
}

func (h *CrudHandler) Get(c echo.Context) error {

	mahasiswas, err := h.service.Get(context.Background())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to fetch mahasiswas"})
	}

	return c.JSON(http.StatusOK, mahasiswas)
}

func (h *CrudHandler) Update(c echo.Context) error {
	var mahasiswa models.KinerjaCrud

	if err := c.Bind(&mahasiswa); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
	}

	if err := h.service.Update(context.Background(), id, mahasiswa); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to update mahasiswa"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Mahasiswa updated successfully"})
}

func (h *CrudHandler) Delete(c echo.Context) error {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
	}

	if err := h.service.Delete(context.Background(), id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to delete mahasiswa"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Mahasiswa deleted successfully"})
}
