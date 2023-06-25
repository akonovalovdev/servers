package main

import (
	"fmt"
	"net/http"
	"sync"
)

type otvet struct {
	res string
	ok  bool
}

func main() {
	var urls = []string{
		"http://ozon.ru",
		"https://ozon.ru",
		"http://google.com",
		"http://somesite.com",
		"http://non-existent.domain.tld",
		"https://ya.ru",
		"http://ya.ru",
		"http://ёёёё",
		"http://ozon.ru",
		"https://ozon.ru",
		"http://google.com",
		"http://somesite.com",
		"http://non-existent.domain.tld",
		"https://ya.ru",
		"http://ya.ru",
		"http://ёёёё",
		"http://ozon.ru",
		"https://ozon.ru",
		"http://google.com",
		"http://somesite.com",
		"http://non-existent.domain.tld",
		"https://ya.ru",
		"http://ya.ru",
		"http://ёёёё",
		"http://ozon.ru",
		"https://ozon.ru",
		"http://google.com",
		"http://somesite.com",
		"http://non-existent.domain.tld",
		"https://ya.ru",
		"http://ya.ru",
		"http://ёёёё",
		"http://ozon.ru",
		"https://ozon.ru",
		"http://google.com",
		"http://somesite.com",
		"http://non-existent.domain.tld",
		"https://ya.ru",
		"http://ya.ru",
		"http://ёёёё",
	}

	ch := make(chan otvet)

	client := &http.Client{}

	var wg sync.WaitGroup

	wg.Add(len(urls)) // увеличиваем счётчик

	for _, url := range urls {
		go func(url string) {
			defer wg.Done() // вычитаем счётчик на 1

			// создаём обект запроса, если содержит не существующий метод или не верный url, возвращаем ошибку
			request, err := http.NewRequest(http.MethodGet, url, nil)
			if err != nil {
				ch <- otvet{
					res: fmt.Sprintf("%s - not ok", url),
					ok:  false,
				}
				return
			}

			// отправляем запрос от клиента и получаем ответ - бывают разные ошибки. Например:
			// отсутствие ответа от сервера или ошибка при отправке и тп
			resp, err := client.Do(request)
			if err != nil {
				ch <- otvet{
					res: fmt.Sprintf("%s - not ok", url),
					ok:  false,
				}
				return
			}
			// закрыли тело после прочтения
			defer resp.Body.Close()

			if resp.StatusCode != 200 {
				ch <- otvet{
					res: fmt.Sprintf("%s - not ok", url),
					ok:  false,
				}
				return
			}

			ch <- otvet{
				res: fmt.Sprintf("%s - ok", url),
				ok:  true,
			}
		}(url)
	}

	go func() {
		// когда счётчик обнулится основная горутина продолжит работу
		wg.Wait()
		// закрыли канал
		close(ch)
	}()

	oks := 0
	for val := range ch { // range заканчивается когда прочитал всё из канала и канал закрыт
		fmt.Println(val)
		if val.ok {
			oks++
		}
		if oks == 2 {
			return
		}

	}

}
