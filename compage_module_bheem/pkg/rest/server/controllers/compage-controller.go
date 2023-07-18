package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ppinnamani/compage_module/compage_module_bheem/pkg/rest/server/models"
	"github.com/ppinnamani/compage_module/compage_module_bheem/pkg/rest/server/services"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type CompageController struct {
	compageService *services.CompageService
}

func NewCompageController() (*CompageController, error) {
	compageService, err := services.NewCompageService()
	if err != nil {
		return nil, err
	}
	return &CompageController{
		compageService: compageService,
	}, nil
}

func (compageController *CompageController) CreateCompage(context *gin.Context) {
	// validate input
	var input models.Compage
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger compage creation
	if _, err := compageController.compageService.CreateCompage(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Compage created successfully"})
}

func (compageController *CompageController) UpdateCompage(context *gin.Context) {
	// validate input
	var input models.Compage
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger compage update
	if _, err := compageController.compageService.UpdateCompage(id, &input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Compage updated successfully"})
}

func (compageController *CompageController) FetchCompage(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger compage fetching
	compage, err := compageController.compageService.GetCompage(id)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, compage)
}

func (compageController *CompageController) DeleteCompage(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger compage deletion
	if err := compageController.compageService.DeleteCompage(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Compage deleted successfully",
	})
}

func (compageController *CompageController) ListCompages(context *gin.Context) {
	// trigger all compages fetching
	compages, err := compageController.compageService.ListCompages()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, compages)
}

func (*CompageController) PatchCompage(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PATCH",
	})
}

func (*CompageController) OptionsCompage(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "OPTIONS",
	})
}

func (*CompageController) HeadCompage(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "HEAD",
	})
}
