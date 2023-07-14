package user

import (
	userModel "github.com/Hideinbruh/test-health/internal/model/user"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Писать User не обязательно, так как по названию пакета понятно, что он к пользователю относится
type Repository interface {
	// Названия методов должны соответствовать методам сервиса (хотя бы примерно)
	Create(user *userModel.UserInfo) (int64, error)
}

// можно назвать пул (типа пул соединений), а можно pgClient
type repository struct {
	pgClient *pgxpool.Pool
}

// Его нам создадут и передадут в мейне
func NewRepository(pgClient *pgxpool.Pool) *repository {
	return &repository{
		pgClient: pgClient,
	}
}

//Сами методы можно разбивать по папкам, как в сервисе, но не обязательно

// Погнали к реализации
func (r *repository) Create(user *userModel.UserInfo) (int64, error) {

	//Сначала нужно создать саму пбазу и поднять её в контейнере
	// для этого воспользуемся файлом docker-compose.yml
	query := ""

	return 1, nil
}
