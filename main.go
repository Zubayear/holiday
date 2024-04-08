package main

import (
	"github.com/Zubayear/holiday/api/routes"
	"github.com/Zubayear/holiday/pkg/external"
	jobs2 "github.com/Zubayear/holiday/pkg/jobs"
	"log"
	"net/http"
)

func main() {
	pool, cancel, err := external.ConnectToDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer cancel()
	jobsRepository := jobs2.NewRepository(pool)
	jobsService := jobs2.NewService(jobsRepository)

	r := routes.Router(jobsService)

	log.Fatal(http.ListenAndServe(":42069", r))
}
