package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"restapi/src/common/di"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	dns := fmt.Sprintf("host=db port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	db, err := sql.Open("postgres", dns)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	taskHandler := di.InitTask(db)
	router := httprouter.New()
	router.GET("/api/tasks", taskHandler.GetTasks)

	if err := http.ListenAndServe(":8000", &Server{router}); err != nil {
		panic(err)
	}
}

type Server struct {
	r *httprouter.Router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Add("Access-Control-Allow-Headers", "Origin")
	w.Header().Add("Access-Control-Allow-Headers", "X-Requested-With")
	w.Header().Add("Access-Control-Allow-Headers", "Accept")
	w.Header().Add("Access-Control-Allow-Headers", "Accept-Language")
	w.Header().Set("Content-Type", "application/json")
	s.r.ServeHTTP(w, r)
}
