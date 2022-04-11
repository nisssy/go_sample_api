package infrastructure

import "github.com/gin-gonic/gin"

type Routing struct {
	DB   *DB
	Gin  *gin.Engine
	Port string
}

func NewRouting(db *DB) *Routing {
	c := NewConfig()
	r := &Routing{
		DB:   db,
		Gin:  gin.Default(),
		Port: c.Routing.Port,
	}

	return r
}

func (r *Routing) SetRouting() {
	usersController := ""
	r.Gin.GET("", func(ctx *gin.Context) {})
}

func (r *Routing) Run() {
	r.Gin.Run(r.Port)
}
