package app

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"
)

type App struct {
	mux      *http.ServeMux
	template *template.Template
	env      env
	db       db
	msgStore messageStore
}

type env struct {
	dbConnStr       string
	msgStoreConnStr string
	port            string
	appEnv          string
	appName         string
}

//go:embed templates
var templateFiles embed.FS

func New() *App {
	app := App{mux: http.NewServeMux(), env: newEnv()}

	app.routes()
	app.template = template.Must(template.ParseFS(templateFiles, "templates/*"))

	return &app
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}

func (a *App) Start() {
	// msg, _ := newMessage("SOME TYPE", "{}", "{}")
	// err := a.msgStore.write("SOME NAME", msg, -1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	log.Println("📡 Listening on port " + a.env.port)
	log.Fatal(http.ListenAndServe(":"+a.env.port, a))
}

func (a *App) routes() {
	a.mux.Handle("/static/", withMiddleware(a.handlerStatic(), primeRequestContext, recoverWrap))
	a.mux.Handle("/", withMiddleware(a.handlerHome(), primeRequestContext, recoverWrap))
}

func newEnv() env {
	var e env

	e.port = os.Getenv("PORT")
	if e.port == "" {
		panic("PORT missing from env")
	}

	e.dbConnStr = os.Getenv("DB")
	if e.dbConnStr == "" {
		panic("DB missing from env")
	}

	e.msgStoreConnStr = os.Getenv("MESSAGE_STORE_DB")
	if e.msgStoreConnStr == "" {
		panic("MESSAGE_STORE_DB missing from env")
	}

	e.appName = os.Getenv("APP_NAME")
	if e.appName == "" {
		panic("APP_NAME missing from env")
	}

	e.appEnv = os.Getenv("APP_ENV")
	if e.appEnv == "" {
		e.appEnv = "development"
	}

	return e
}
