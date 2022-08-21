package services

import "git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/repositories"

type AdminService interface {
}

type adminService struct {
	adminRepository repositories.AdminRepository
}

type ADSConfig struct {
	AdminRepository repositories.AdminRepository
}

func NewAdminServices(c *ADSConfig) *adminService {
	return &adminService{
		adminRepository: c.AdminRepository,
	}
}
