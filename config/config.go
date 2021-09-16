package config

var (
	GitToken = "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"
	ConnStr  = "user=postgres password=Qweasdzxc1 dbname=users sslmode=disable"
)

type Config struct {
	User     string
	Password string
	Dbname   string
	Sslmode  string
}

// 1
// Сделать структуру репозитория с двумя методами
// Create (insert) func .. Create(..) error
// List (select списка) func ... List(..)
// Открыть соединение с базой один раз
// Пробросить сущность базы аргументом в конструкторе репозитория
