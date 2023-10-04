package web

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var wg sync.WaitGroup

func StartWebServer() {
	// Инициализация веб-сайта
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Добро пожаловать на веб-сайт!")
	})

	r.POST("/send_message", func(c *gin.Context) {
		// Обработка POST-запросов для отправки сообщений
		// ...

		c.String(http.StatusOK, "Сообщение отправлено асинхронно.")
	})

	// Запуск веб-сайта в горутине
	go func() {
		if err := r.Run(":8080"); err != nil {
			log.Fatal("Ошибка запуска веб-сайта:", err)
		}
	}()

	log.Println("Веб-сайт запущен на :8080")

	// Ожидание завершения всех горутин перед выходом
	wg.Wait()
}
