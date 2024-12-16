package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/fixme_my_friend/hw02_fix_app/types"
)

// ReadJSON читает JSON файл и возвращает массив сотрудников
func ReadJSON(filePath string, limit int) ([]types.Employee, error) {
	// Открываем файл
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файл: %w", err)
	}
	defer file.Close() // Закрытие файла после завершения функции

	// Читаем содержимое файла
	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать файл: %w", err)
	}

	// Парсим JSON в структуру данных
	var employees []types.Employee
	err = json.Unmarshal(bytes, &employees)
	if err != nil {
		return nil, fmt.Errorf("не удалось распарсить JSON: %w", err)
	}

	// Применяем лимит, если он задан
	if limit > 0 && len(employees) > limit {
		employees = employees[:limit]
	}

	return employees, nil
}
