package main

import (
	"log"
	"net/http"
	"path/filepath"
)

// Тип для настраиваемой файловой системы, который включает в себя http.FileSystem
type neuteredFileSystem struct {
	fs http.FileSystem
}

func main() {
	// Регистрируем два новых обработчика и соответствующие URL-шаблоны в
	// маршрутизаторе servemux
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
	/*
			Ограничение просмотра файлов из директории
		1-способ
		Добавьте пустой файл index.html в ту директорию, где требуется отключить
		вывод списка файлов. Веб-сервер всегда ищет файл index.html,
		и пользователь увидит пустую страницу с кодом состояния 200 OK.
		2-способ
		Настраиваемая имплементация файловой системы http.FileSystem,
		с помощью которой будет возвращаться ошибка os.ErrNotExist
		для любого HTTP запроса напрямую к папке
	*/
	// Инициализируем FileServer, он будет обрабатывать
	// HTTP-запросы к статическим файлам из папки "./ui/static".
	// Использование настраиваемой файловой системы, добавив новую структуру neuteredFileSystem и метод Open для неё
	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./snippetbox/ui/static/")})
	/*
		Иногда может потребоваться обслужить только один статический файл. Для этой задачи,
		функция http.ServeFile() используется следующим образом:

		func downloadHandler(w http.ResponseWriter, r *http.Request) {
		    http.ServeFile(w, r, "./ui/static/file.zip")
		}

		Внимание: Обработчик http.ServeFile() автоматически не очищает путь файла. Если вы указываете
		в данную функцию путь к файлу полученный от пользователя напрямую, во избежание атак обхода директории,
		перед использование этих данных, обязательно очистите их с помощью функции filepath.Clean().
		Иначе, пользователь сможет скачать различные файлы с сервера включая файл с настройками к базе данных.
	*/

	// Используем функцию mux.Handle() для регистрации обработчика для
	// всех запросов, которые начинаются с "/static/". Мы убираем
	// префикс "/static" перед тем как запрос достигнет http.FileServer
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

// Создаем метод Open(), который вызывается каждый раз, когда http.FileServer получает запрос
func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	// Открываем вызываемый путь
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	// Используя метод IsDir() мы проверим если вызываемый путь является папкой или нет
	if s.IsDir() {
		// Если это папка, то с помощью метода Stat("index.html") мы проверим если файл index.html существует внутри данной папки
		index := filepath.Join(path, "index.html")
		// Если файл index.html не существует, то метод вернет ошибку os.ErrNotExist
		// которая, в свою очередь, будет преобразована через http.FileServer в ответ 404 страница не найдена
		if _, err := nfs.fs.Open(index); err != nil {
			// Мы также вызываем метод Close() для закрытия только, что открытого index.html файла,
			// чтобы избежать утечки файлового дескриптора
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}
	// Во всех остальных случаях мы просто возвращаем файл и даем http.FileServer сделать то, что он должен.
	return f, nil
}
