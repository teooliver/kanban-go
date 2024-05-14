package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	models "github.com/teooliver/kanban/internal/models"
)

// TODO move to .env
const (
	host = "localhost"
	port = 5432
	user = "db_user"
	// password = "12345"
	dbname = "kanban-go"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	err := godotenv.Load(".env.dev")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, postgresPassword, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	// insert
	// hardcoded
	// insertStmt := `insert into "task"("title") values('hello_world2')`
	// _, e := db.Exec(insertStmt)
	// CheckError(e)
	models.ListTasks(db)

	// dynamic
	// insertDynStmt := `insert into "Students"("Name", "Roll") values($1, $2)`
	// _, e = db.Exec(insertDynStmt, "Jane", 2)
	// CheckError(e)

	// CHI
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Route("/task", func(r chi.Router) {
		// r.With(paginate).Get("/", listArticles)                           // GET /articles
		// r.With(paginate).Get("/{month}-{day}-{year}", listArticlesByDate) // GET /articles/01-16-2017

		r.Post("/", createTask) // POST /task

	})

	http.ListenAndServe(":3000", r)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	// here we read from the request context and fetch out `"user"` key set in
	// the MyMiddleware example above.
	// task := r.Context().Value("user").(string)
	// r.Body.Read(p []byte)

	// respond to the client
	w.Write([]byte(fmt.Sprintf("Task %s", "task")))
}