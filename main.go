package main

import (
	"go_sample_api/infrastructure"
)

func main() {
	db := infrastructure.NewDB()
	r := infrastructure.NewRouting(db)
	r.Run()
}
