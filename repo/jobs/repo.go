package repo

import (
	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	"gorm.io/gorm"
)

type JobsRepository interface {
	CheckIfJobExists(jobID string) error
	CheckIfApplied(applicantID, jobID string) (*models.JobApplication, error)
	CreateApplication(applicationID, applicantID, jobID string) error
	GetAllJobs() (*[]models.Job, error)
}

type JobRepo struct {
	Driver *gorm.DB
}

func NewJobRepo(driver *gorm.DB) *JobRepo {
	return &JobRepo{Driver: driver}
}
