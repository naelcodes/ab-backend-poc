package repository

import . "neema.co.za/rest/utils/database"

type Repository struct {
	*RedisStore
}

func NewRepository(redisStore *RedisStore) *Repository {
	return &Repository{redisStore}
}
