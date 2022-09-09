package celeritas

import (
	"database/sql"
	"fmt"
	"github.com/guillospy92/utilsgolang/go-version-laravel/celeritas/session"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/CloudyKit/jet/v6"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"

	"github.com/guillospy92/utilsgolang/go-version-laravel/celeritas/render"
)

const version = "1.0.0"

// Accelerator name init app
type Accelerator struct {
	// init params accelerator struct
	AppName  string
	RootPath string

	// config accelerator
	Config config

	// debug message error and info accelerator
	ErrorLog *log.Logger
	InfoLog  *log.Logger

	// routes and renderer accelerator
	Routes  *chi.Mux
	Render  *render.Render
	JetView *jet.Set

	// sessions
	Session *scs.SessionManager

	// database info
	DB Database
}

// config access config attributes
type config struct {
	port        string
	renderer    string
	version     string
	debug       bool
	cookie      cookieConfig
	sessionType string
	database    dataBaseConfig
}

// NewAccelerator instance Accelerator
func NewAccelerator(rootPath string) *Accelerator {
	a := &Accelerator{}
	a.RootPath = rootPath

	// init folder start if no exists folder specif when accelerator create folders in the root path
	a.startCreateFolder()

	// implements and load function necessary for success run accelerator
	// startLogger() init logger accelerator
	// checkDotEnv() init load .env
	// startConfig() init config necessary for running accelerator
	// startJetView() init render and load view with jet
	// startRenderer() init render and load view accelerator
	// startSession() init session all accelerator
	a.startLogger()
	a.checkDotEnv()
	a.StartConfig()
	a.startJetView()
	a.startRenderer()
	a.startSession()
	a.startConnectionDataBase()

	// implements routes and server listening
	a.Routes = a.routes().(*chi.Mux)

	return a
}

// Init run accelerator
func (a *Accelerator) startCreateFolder() {
	folders := initPath{
		rootPath:    a.RootPath,
		folderNames: []string{"handlers", "migrations", "views", "data", "public", "tmp", "logs", "middleware"},
	}

	for _, path := range folders.folderNames {
		err := a.CreateDirNotExists(folders.rootPath + "/" + path)
		if err != nil {
			panic(err)
		}
	}
}

// StartServer init server
func (a *Accelerator) StartServer() {

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		ErrorLog:     a.ErrorLog,
		Handler:      a.Routes,
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 600 * time.Second,
	}

	defer func(Pool *sql.DB) {
		err := Pool.Close()
		if err != nil {
			a.ErrorLog.Println("error close connection", err)
		}
	}(a.DB.Pool)

	a.InfoLog.Printf("listen on ports %s", os.Getenv("PORT"))

	err := srv.ListenAndServe()

	if err != nil {
		a.ErrorLog.Fatal(err)
	}
}

// StartConfig config necessary accelerator
func (a *Accelerator) StartConfig() {
	debug := true
	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))

	if err != nil {
		a.ErrorLog.Println("error reading debug in file .env")
		debug = false
	}

	a.Config = config{
		port:     os.Getenv("PORT"),
		renderer: os.Getenv("RENDERER"),
		version:  version,
		debug:    debug,
		cookie: cookieConfig{
			name:     os.Getenv("COOKIE_NAME"),
			domain:   os.Getenv("COOKIE_DOMAIN"),
			lifeTime: os.Getenv("COOKIE_LIFETIME"),
			persists: os.Getenv("COOKIE_PERSISTS"),
			secure:   os.Getenv("COOKIE_SECURE"),
		},

		database: dataBaseConfig{
			dbType:     os.Getenv("DATABASE_TYPE"),
			dbHost:     os.Getenv("DATABASE_HOST"),
			dbName:     os.Getenv("DATABASE_NAME"),
			dbPort:     os.Getenv("DATABASE_PORT"),
			dbUser:     os.Getenv("DATABASE_USER"),
			dbPassword: os.Getenv("DATABASE_PASS"),
			dbSslMode:  os.Getenv("DATABASE_SSL_MODE"),
		},

		sessionType: os.Getenv("SESSION_TYPE"),
	}
}

// startLogger init load logger
func (a *Accelerator) startLogger() {
	a.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	a.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)
}

// startRenderer init load rendered
func (a *Accelerator) startRenderer() {
	a.Render = &render.Render{
		Renderer: a.Config.renderer,
		RootPath: a.RootPath,
		Port:     a.Config.port,
		JetView:  a.JetView,
	}
}

// startJetView start support jed views set
func (a *Accelerator) startJetView() {
	a.JetView = jet.NewSet(
		jet.NewOSFileSystemLoader(fmt.Sprintf("%s/views", a.RootPath)),
		jet.InDevelopmentMode(),
	)
}

// startSession configure all sessions
func (a *Accelerator) startSession() {
	sess := session.Session{
		CookieLifeTime: a.Config.cookie.lifeTime,
		CookiePersist:  a.Config.cookie.persists,
		CookieName:     a.Config.cookie.name,
		CookieDomain:   a.Config.cookie.domain,
		SessionType:    a.Config.sessionType,
	}

	a.Session = sess.InitSession()
}

// startConnectionDataBase connect database info
func (a *Accelerator) startConnectionDataBase() {
	if a.Config.database.dbType == "" {
		return
	}

	// create build dns
	dns := a.BuildDns()
	ql, err := a.openDB(dns)

	if err != nil {
		a.ErrorLog.Println("error open connection db")
		os.Exit(1)
	}

	a.DB = Database{
		DataType: a.Config.database.dbType,
		Pool:     ql,
	}
}

// checkDotEnv creating and reading .env
func (a *Accelerator) checkDotEnv() {
	path := a.RootPath + "/.env"
	if err := a.CreateFileNotExists(path); err != nil {
		a.ErrorLog.Println("error logs path", a.RootPath)
		panic(fmt.Errorf("error create or read file .env %v", err))
	}

	if err := godotenv.Load(path); err != nil {
		panic(fmt.Errorf("error load .env %v", err))
	}
}
