package users

import (
	"github.com/forderation/user-task/internal/db/gormgen"
	"gorm.io/gorm"
)

type Repository struct {
	DB    *gorm.DB
	Query *gormgen.Query
}

func NewRepository(db *gorm.DB) *Repository {
	var query *gormgen.Query
	if db != nil {
		query = gormgen.Use(db)
	}
	return &Repository{
		DB:    db,
		Query: query,
	}
}
