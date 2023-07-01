package main

import "github.com/akonovalovdev/servers/snippetbox/pkg/models"

// Создаем тип templateData, который будет действовать как хранилище для
// любых динамических данных, которые нужно передать HTML-шаблонам.
// На данный момент он содержит только одно поле, но мы добавим в него другие
// по мере развития нашего приложения.
// Добавляем поле Snippets в структуру templateData
type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
