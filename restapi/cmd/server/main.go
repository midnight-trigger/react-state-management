package main

import (
	"fmt"
	"net/http"
	"os"

	"restapi/src/common/di"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load("../../.env"); err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("host=db port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}

	taskHandler := di.InitTask(db)
	tagHandler := di.InitTag(db)
	router := httprouter.New()
	router.GET("/api/tasks", taskHandler.GetTasks)
	router.POST("/api/tasks", taskHandler.CreateTask)
	router.PUT("/api/tasks/:taskId", taskHandler.UpdateTask)
	router.DELETE("/api/tasks/:taskId", taskHandler.DeleteTask)
	router.GET("/api/tags", tagHandler.GetTags)
	router.POST("/api/tags", tagHandler.CreateTag)
	router.PUT("/api/tags/:tagId", tagHandler.UpdateTag)
	router.DELETE("/api/tags/:tagId", tagHandler.DeleteTag)

	if err := http.ListenAndServe(":8000", &Server{router}); err != nil {
		panic(err)
	}
}

type Server struct {
	r *httprouter.Router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Add("Access-Control-Allow-Headers", "Origin")
	w.Header().Add("Access-Control-Allow-Headers", "X-Requested-With")
	w.Header().Add("Access-Control-Allow-Headers", "Accept")
	w.Header().Add("Access-Control-Allow-Headers", "Accept-Language")
	w.Header().Set("Content-Type", "application/json")
	s.r.ServeHTTP(w, r)
}
