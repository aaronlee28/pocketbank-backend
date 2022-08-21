package repositories

import (
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/models"
	"gorm.io/gorm"
)

type AdminRepository interface {
	UsersList() (*[]models.User, error)
}

type adminRepository struct {
	db *gorm.DB
}

type ADRConfig struct {
	DB *gorm.DB
}

func NewAdminRepository(c *ADRConfig) adminRepository {
	return adminRepository{db: c.DB}
}

func (w *adminRepository) UsersList() (*[]models.User, error) {
	var u *[]models.User
	err := w.db.Find(&u).Error

	return u, err
}
