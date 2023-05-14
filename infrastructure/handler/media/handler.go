package media

import (
	"fmt"
	"github.com/arthur8davis/housematch-api/domain/model"
	"github.com/arthur8davis/housematch-api/domain/services/media"
	"github.com/arthur8davis/housematch-api/domain/services/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type handler struct {
	useCase media.UseCaseMedia
}

func newHandler(useCase media.UseCaseMedia) handler {
	return handler{useCase}
}

func (h handler) upload(c *gin.Context) {
	file, err := c.FormFile("fileUpload")
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Error to get file: %s", err.Error()))
		return
	}

	generateFilename, err := uuid.NewUUID()
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Error al convertir la cadena en UUID: %s\n", err))
		return
	}

	filename := generateFilename.String() + filepath.Ext(file.Filename)
	destinationFile := filepath.Join(os.Getenv("FILES"), filename)

	err = c.SaveUploadedFile(file, destinationFile)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Error to download file: %s", err.Error()))
		return
	}

	fileModel, err := createFileModel(file, filename, destinationFile)
	if err != nil {
		fmt.Printf("Error al convertir la cadena en UUID: %s\n", err)
		return
	}

	m, err := h.useCase.Create(fileModel)
	if err != nil {
		c.JSON(response.Wrong(model.ResponseError{Error: err.Error()}))
	}

	c.JSON(response.Created(m))
}

func (h handler) delete(c *gin.Context) {
	id := c.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		fmt.Printf("Error al convertir la cadena en UUID: %s\n", err)
		return
	}

	deleted, err := h.useCase.Delete(uid)
	if err != nil {
		c.JSON(response.Wrong(model.ResponseError{Error: err.Error()}))
		return
	}

	c.JSON(response.Deleted(deleted))
}

func (h handler) files(c *gin.Context) {
	c.File("/home/arthur/files" + c.Param("filepath"))
}

func createFileModel(file *multipart.FileHeader, filename, dest string) (model.Media, error) {
	newId, err := uuid.NewUUID()
	if err != nil {
		return model.Media{}, fmt.Errorf("Error generate UUID: %s\n", err)
	}

	return model.Media{
		ID:   newId,
		Name: filename,
		URL:  dest,
		Type: file.Header.Get("Content-Type"),
		Size: file.Size,
	}, nil
}
