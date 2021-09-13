package config

var (
	GitToken = "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"
	ConnStr  = "user=postgres password=Qweasdzxc1 dbname=users sslmode=disable"
)

// 1
// Сделать структуру репозитория с двумя методами
// Create (insert) func .. Create(..) error
// List (select списка) func ... List(..)
// Открыть соединение с базой один раз
// Пробросить сущность базы аргументом в конструкторе репозитория
