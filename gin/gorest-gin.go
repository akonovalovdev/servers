// REST-сервер реализован с помощью Gin.

package main

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/akonovalovdev/servers/gin/internal/taskstore"
	"github.com/gin-gonic/gin"
)

type taskServer struct {
	store *taskstore.TaskStore
}

func NewTaskServer() *taskServer {
	store := taskstore.New()
	return &taskServer{store: store}
}

/*
У обработчиков, используемых в Gin, нет стандартных сигнатур HTTP-обработчиков Go. Они просто принимают объект
gin.Context, который может быть использован для анализа запроса и для формирования ответа. Но в Gin
есть механизмы для взаимодействия со стандартными обработчиками — вспомогательные функции gin.WrapF и gin.WrapH

В отличие от ранней версии нашего сервера, тут нет нужды вручную писать в журнал сведения о запросах,
так как стандартный механизм логирования Gin, представленный ПО промежуточного
уровня, сам решает эту задачу (и делается это с использованием всяческих полезных мелочей,
вроде оформления вывода разными цветами и включения в журнал сведений о времени обработки запросов).

Нам, кроме того, больше не нужно самостоятельно реализовывать вспомогательную функцию renderJSON,
так как в Gin есть собственный механизм Context.JSON, который позволяет формировать JSON-ответы.
*/
func (ts *taskServer) getAllTasksHandler(c *gin.Context) {
	allTasks := ts.store.GetAllTasks()
	c.JSON(http.StatusOK, allTasks)
}

func (ts *taskServer) deleteAllTasksHandler(c *gin.Context) {
	ts.store.DeleteAllTasks()
}

// createTaskHandler обрабатывает запросы, которые включают в себя особые данные
// Тут под «привязкой» понимается обработка содержимого запросов (которое может быть
// представлено данными в различных форматах, например — JSON и YAML),
// проверка полученных данных и запись соответствующих значений в структуры Go
func (ts *taskServer) createTaskHandler(c *gin.Context) {
	// RequestTask, привязка где проверка данных не используется
	type RequestTask struct {
		Text string    `json:"text"`
		Tags []string  `json:"tags"`
		Due  time.Time `json:"due"`
	}

	var rt RequestTask
	// за разбор JSON-данных запроса отвечает ShouldBindJSON
	if err := c.ShouldBindJSON(&rt); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	id := ts.store.CreateTask(rt.Text, rt.Tags, rt.Due)
	/*
	 В Gin нам не нужно пользоваться «одноразовой» структурой для ID ответа.
	 Вместо этого мы используем gin.H — псевдоним для map[string]interface{}; это
	 позволяет конструировать ответы, используя небольшие объёмы кода
	*/
	c.JSON(http.StatusOK, gin.H{"Id": id})
}

func (ts *taskServer) getTaskHandler(c *gin.Context) {
	// Gin позволяет обращаться к параметрам маршрута (к тому, что начинается с двоеточия, вроде :id) через Context.Params
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	// Gin не поддерживает регулярные выражения в маршрутах
	// поэтому нам нужно самим позаботиться о разборе данных(таких как целые числа - индификаторы задач)
	/*
		Речь идёт об отсутствии поддержки регулярных выражений в системе маршрутизации Gin. А это значит,
		что обработка любого необычного маршрута, его разбор и проверка, потребуют писать больше кода
	*/
	task, err := ts.store.GetTask(id)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, task)
}

func (ts *taskServer) deleteTaskHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if err = ts.store.DeleteTask(id); err != nil {
		c.String(http.StatusNotFound, err.Error())
	}
}

func (ts *taskServer) tagHandler(c *gin.Context) {
	tag := c.Params.ByName("tag")
	tasks := ts.store.GetTasksByTag(tag)
	c.JSON(http.StatusOK, tasks)
}

func (ts *taskServer) dueHandler(c *gin.Context) {
	badRequestError := func() {
		c.String(http.StatusBadRequest, "expect /due/<year>/<month>/<day>, got %v", c.FullPath())
	}

	year, err := strconv.Atoi(c.Params.ByName("year"))
	if err != nil {
		badRequestError()
		return
	}

	month, err := strconv.Atoi(c.Params.ByName("month"))
	if err != nil || month < int(time.January) || month > int(time.December) {
		badRequestError()
		return
	}

	day, err := strconv.Atoi(c.Params.ByName("day"))
	if err != nil {
		badRequestError()
		return
	}

	tasks := ts.store.GetTasksByDueDate(year, time.Month(month), day)
	c.JSON(http.StatusOK, tasks)
}

func main() {
	router := gin.Default()
	server := NewTaskServer()

	router.POST("/task/", server.createTaskHandler)
	router.GET("/task/", server.getAllTasksHandler)
	router.DELETE("/task/", server.deleteAllTasksHandler)
	router.GET("/task/:id", server.getTaskHandler)
	router.DELETE("/task/:id", server.deleteTaskHandler)
	router.GET("/tag/:tag", server.tagHandler)
	router.GET("/due/:year/:month/:day", server.dueHandler)

	router.Run("localhost:" + os.Getenv("SERVERPORT"))
}
