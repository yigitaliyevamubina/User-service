package storage

import (
	"database/sql"
	"user-service/storage/postgres"
	"user-service/storage/repo"
)

// IStorage
type IStorage interface {
	User() repo.UserStorageI
}
type storagePg struct {
	db       *sql.DB
	userRepo repo.UserStorageI
}

func NewStoragePg(db *sql.DB) *storagePg {
	return &storagePg{
		db:       db,
		userRepo: postgres.NewUserRepo(db),
	}
}

func (s *storagePg) User() repo.UserStorageI {
	return s.userRepo
}
