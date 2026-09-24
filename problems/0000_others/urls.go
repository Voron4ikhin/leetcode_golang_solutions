package _000_others

import (
	"context"
	"fmt"
	"net/http"
	"sync"
)

// 1. Поочередно выполнит http запросы по предложенному списку ссылок
// • в случае получения http-кода ответа на запрос "200 OK" печатаем на экране "адрес url -
// ok"
// • в случае получения http-кода ответа на запрос отличного от "200 OK" либо в случае
// ошибки печатаем на экране "адрес url - not ok"
// 2. Модифицируйте программу таким образом, чтобы использовались каналы для
// коммуникации основного потока с горутинами. Пример:
// • Запросы по списку выполняются в горутинах.
// • Печать результатов на экран происходит в основном потоке
// 3. Модифицируйте программу таким образом, чтобы нигде не использовалась длина
// слайса урлов. Считайте, что урлы приходят из внешнего источника. Сколько их будет
// заранее - неизвестно. Предложите идиоматичный вариант, как ваша программа будет
// узнавать об окончании списка и передавать сигнал об окончании действий далее.
func SendUrls() {
	var urls = []string{
		"http://ozon.ru",
		"https://ozon.ru",
		"http://google.com",
		"http://somesite.com",
		"http://non-existent.domain.tld",
		"https://ya.ru",
		"http://ya.ru",
		"http://ёёёё",
	}

	urlsCh := make(chan string)
	go func() {
		for _, url := range urls {
			urlsCh <- url
		}
		close(urlsCh)
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resCh := make(chan string)
	var wg sync.WaitGroup
	workers := 5
	wg.Add(workers)

	for numWorker := range workers {
		go func() {
			defer wg.Done()
			worker(ctx, urlsCh, resCh, numWorker)
		}()
	}

	go func() {
		wg.Wait()
		close(resCh)
	}()

	for resp := range resCh {
		fmt.Println(resp)
	}
}

func worker(ctx context.Context, urlsCh <-chan string, resCh chan<- string, worker int) {
	for {
		select {
		case <-ctx.Done():
			return
		case url, ok := <-urlsCh:
			if !ok {
				return
			}
			res, err := http.Get(url)
			if err != nil {
				resCh <- fmt.Sprintf("worker %d: адрес url - %s not ok\n", worker, url)
				continue
			}
			defer res.Body.Close()
			if res.StatusCode != 200 {
				resCh <- fmt.Sprintf("worker %d: адрес url - %s not ok\n", worker, url)
				continue
			}
			resCh <- fmt.Sprintf("worker %d: адрес url - %s ok\n", worker, url)
		}
	}
}
