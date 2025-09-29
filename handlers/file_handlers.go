package handlers

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"file-loader/config"
	"file-loader/utils"

	"github.com/gin-gonic/gin"
)

type FileHandler struct{}

func NewFileHandler() *FileHandler {
	return &FileHandler{}
}

// List files
func (h *FileHandler) List(c *gin.Context) {
	files, err := ioutil.ReadDir(config.FileBasePath)
	if err != nil {
		utils.Respond(c, http.StatusInternalServerError, "Failed to list files", nil, err)
		return
	}

	var fileNames []string
	for _, f := range files {
		fileNames = append(fileNames, f.Name())
	}

	utils.Respond(c, http.StatusOK, "Files listed successfully", fileNames, nil)
}

// Read file
func (h *FileHandler) Read(c *gin.Context) {
	name := c.Param("name")
	path := filepath.Join(config.FileBasePath, name)

	data, err := ioutil.ReadFile(path)
	if err != nil {
		utils.Respond(c, http.StatusNotFound, "File not found", nil, err)
		return
	}

	utils.Respond(c, http.StatusOK, "File read successfully", string(data), nil)
}

// Write file
func (h *FileHandler) Write(c *gin.Context) {
	var req struct {
		Filename string `json:"filename" binding:"required"`
		Content  string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Respond(c, http.StatusBadRequest, "Invalid request", nil, err)
		return
	}

	path := filepath.Join(config.FileBasePath, req.Filename)

	if err := ioutil.WriteFile(path, []byte(req.Content), os.ModePerm); err != nil {
		utils.Respond(c, http.StatusInternalServerError, "Failed to write file", nil, err)
		return
	}

	utils.Respond(c, http.StatusCreated, "File written successfully", req.Filename, nil)
}
