package models

import (
	"errors"
	"time"
)

// Данный файл используется для того, чтобы определить типы данных верхнего уровня,
// которые модель базы данных будет использовать и возвращать.

var ErrNoRecord = errors.New("models: подходящей записи не найдено")

// названия полей структуры Snippet соответствуют полям в MySQL таблице snippets
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
