package main

import (
	gnt "github.com/Erikokijay/go-gintonic"
	"github.com/gin-gonic/gin"
)

func main() {
	// Устанавливаем режим работы GIN
	gin.SetMode("release")

	// Создаем экземпляр маршрутизатора GIN
	eng := gin.Default()

	// Конфигурируем GO-GINTONIC
	gnt.Config(&gnt.ConfigSchema{SwaggerUrl: "/docs"}, eng)

	// Создаем роутер для группы маршрутов /api
	router := gnt.NewRouter(eng.Group("/api"))

	// Исправляем маршрут POST с /get на /post
	router.Post("/post", ping, gnt.RouteInfo{
		Tags:        []string{"Test"},
		Title:       "Route Title",
		Description: "Route Description",
	})

	// Добавляем маршрут GET /test
	router.Get("/test", ping2,
		gnt.ResultInfo{
			Code:   200,
			Output: ExampleResponse{},
		},
		gnt.ResultInfo{
			Code:   500,
			Output: "error",
		},
		gnt.RouteInfo{
			Tags:        []string{"Test", "First"},
			Title:       "Route Title",
			Description: "Route Description", // Запятая здесь обязательна
		}, // Здесь запятая НЕ нужна, так как это последний элемент
	)

	// Генерируем документацию Swagger
	gnt.GenerateSwagger() // Исправлено: убран аргумент

	// Запускаем сервер на порту 8080
	eng.Run(":8080")
}

// Обработчик для POST не /get а /post
func ping(c *gin.Context, data Resp) *Resp {
	return &Resp{Code: data.Code + 1, Msg: data.Msg + " modified"}
}

// Обработчик для GET /test
func ping2(c *gin.Context, data ExampleRequest) (int, interface{}) {
	return 200, &ExampleResponse{Msg: data.Name + " modified"}
}

// Структуры данных

// Resp - структура для входных и выходных данных маршрута POST /get
type Resp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// ExampleRequest - структура для входных данных маршрута GET /test
type ExampleRequest struct {
	Name string `json:"name" binding:"required"`
}

// ExampleResponse - структура для выходных данных маршрута GET /test
type ExampleResponse struct {
	Msg string `json:"msg"`
}
