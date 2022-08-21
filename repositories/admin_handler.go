package repositories

import "gorm.io/gorm"

type AdminRepository interface {
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
