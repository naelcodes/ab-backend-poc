package sessions

import (
	"crypto/rand"
	"fmt"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis/v3"
	"neema.co.za/rest/utils/logger"
)

type AppSessionStore struct {
	*session.Store
}

func (s *AppSessionStore) GenerateSessionID() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	return fmt.Sprintf("%x", b), err
}

func NewAppSessionStore() *AppSessionStore {

	logger.Info(fmt.Sprintf("REDIS_HOST : %s", os.Getenv("REDIS_HOST")))
	logger.Info(fmt.Sprintf("REDIS_PORT : %s", os.Getenv("REDIS_PORT")))

	port, _ := strconv.Atoi(os.Getenv("REDIS_PORT"))
	store := session.New(session.Config{
		Storage: redis.New(redis.Config{
			Host: os.Getenv("REDIS_HOST"),
			Port: port,
		}),
	})

	return &AppSessionStore{store}
}
