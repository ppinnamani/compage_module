package services

import (
	"github.com/ppinnamani/compage_module/compage_module_bheem/pkg/rest/server/daos"
	"github.com/ppinnamani/compage_module/compage_module_bheem/pkg/rest/server/models"
)

type CompageService struct {
	compageDao *daos.CompageDao
}

func NewCompageService() (*CompageService, error) {
	compageDao, err := daos.NewCompageDao()
	if err != nil {
		return nil, err
	}
	return &CompageService{
		compageDao: compageDao,
	}, nil
}

func (compageService *CompageService) CreateCompage(compage *models.Compage) (*models.Compage, error) {
	return compageService.compageDao.CreateCompage(compage)
}

func (compageService *CompageService) UpdateCompage(id int64, compage *models.Compage) (*models.Compage, error) {
	return compageService.compageDao.UpdateCompage(id, compage)
}

func (compageService *CompageService) DeleteCompage(id int64) error {
	return compageService.compageDao.DeleteCompage(id)
}

func (compageService *CompageService) ListCompages() ([]*models.Compage, error) {
	return compageService.compageDao.ListCompages()
}

func (compageService *CompageService) GetCompage(id int64) (*models.Compage, error) {
	return compageService.compageDao.GetCompage(id)
}
