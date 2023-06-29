package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

/*
	Перемещаем часть кода обработки ошибок во вспомогательные методы.
	Это поможет разделить проблемы и избавиться от повторения кода по мере развития программы
*/

// Помощник serverError записывает сообщение об ошибке в errorLog и
// затем отправляет пользователю ответ 500 "Внутренняя ошибка сервера".
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	// app.errorLog.Println(trace) // выводит данную строку текущего файла
	app.errorLog.Output(2, trace) // выводит номер строки и файл, где произошла ошибка
	/*
		В помощнике serverError() мы используем функцию debug.Stack(),
		чтобы получить трассировку стека для текущей горутины и добавить ее в логгер.
		Возможность видеть полный путь к приложению через трассировку стека может быть полезна при отладке возникнувших ошибок
	*/

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// Помощник clientError отправляет определенный код состояния и соответствующее описание
// пользователю. Мы будем использовать это в следующий уроках, чтобы отправлять ответы вроде 400 "Bad
// Request", когда есть проблема с пользовательским запросом.
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
	/*
		В помощнике clientError() мы используем функцию http.StatusText() для генерации понятного человеку
		текстового представления для кода состояния HTTP. К примеру, http.StatusText(400) вернет строку "Bad Request"
	*/
}

// Мы также реализуем помощник notFound. Это просто
// удобная оболочка вокруг clientError, которая отправляет пользователю ответ "404 Страница не найдена".
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

/*
	Мы начали использовать специальные константы из пакета net/http для кодов состояния HTTP вместо целых чисел.
	В помощнике serverError() мы использовали константу http.StatusInternalServerError вместо 500,
	а в помощнике notFound() — константу http.StatusNotFound вместо записи 404.

	Полный список констант кодовЖ https://pkg.go.dev/net/http#pkg-constants
*/
