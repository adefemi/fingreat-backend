package api

import (
	"database/sql"
	"fmt"
	db "github/adefemi/fingreat_backend/db/sqlc"
	"github/adefemi/fingreat_backend/utils"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
)

type Server struct {
	queries *db.Store
	router  *gin.Engine
	config  *utils.Config
}

var tokenController *utils.JWTToken

func NewServer(envPath string) *Server {

	config, err := utils.LoadConfig(envPath)
	if err != nil {
		panic(fmt.Sprintf("Couldn't load config: %v", err))
	}

	conn, err := sql.Open(config.DBdriver, config.DB_source_live)
	if err != nil {
		panic(fmt.Sprintf("Could not connect to database: %v", err))
	}

	tokenController = utils.NewJWTToken(config)

	q := db.NewStore(conn)

	g := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", currencyValidator)
	}

	g.Use(cors.Default())

	return &Server{
		queries: q,
		router:  g,
		config:  config,
	}
}

func (s *Server) Start(port int) {
	s.router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to Fingreat!"})
	})

	User{}.router(s)
	Auth{}.router(s)
	Account{}.router(s)

	s.router.Run(fmt.Sprintf(":%v", port))
}
