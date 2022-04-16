package infrastructure

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"go_sample_api/interfaces/controllers"
)

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

	// TODOセッションの設定
	store := cookie.NewStore([]byte("secret"))
	r.Gin.Use(sessions.Sessions("mysession", store))

	// ルーターの設定
	// set User Router
	r.setRoutingUser()
	// set Place Router
	r.setRoutingPlace()
	// set Booking Router
	r.setRoutingBooking()
	// set Review Router
	r.setRoutingReview()
	return r
}

// TODO Auth関連のコードの追加
// UserRouting
func (r *Routing) setRoutingUser() {
	userController := controllers.NewUserController(r.DB)
	// Create
	r.Gin.POST("user", func(c *gin.Context) { userController.CreateUser(c) })
	// Read
	r.Gin.GET("/users", func(c *gin.Context) { userController.GetUserList(c) })
	r.Gin.GET("/user/:id", func(c *gin.Context) { userController.GetUserByID(c) })
	// Update
	r.Gin.PUT("/user/:id", func(c *gin.Context) { userController.UpdateUser(c) })
	// Delete
	r.Gin.DELETE("/user/:id", func(c *gin.Context) { userController.DeleteUser(c) })
}

// PlaceRouting
func (r *Routing) setRoutingPlace() {
	// userController := controllers.NewUserController(r.DB)
	// r.Gin.GET("/users/:id", func(c *gin.Context) { userController.Get(c) })
}

// BookingRouting
func (r *Routing) setRoutingBooking() {
	// userController := controllers.NewUserController(r.DB)
	// r.Gin.GET("/users/:id", func(c *gin.Context) { userController.Get(c) })
}

// ReviewRouting
func (r *Routing) setRoutingReview() {
	// userController := controllers.NewUserController(r.DB)
	// r.Gin.GET("/users/:id", func(c *gin.Context) { userController.Get(c) })
}

func (r *Routing) Run() {
	r.Gin.Run(r.Port)
}
