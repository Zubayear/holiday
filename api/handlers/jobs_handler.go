package handlers

import (
	"context"
	"github.com/Zubayear/holiday/api/presenter"
	jobs2 "github.com/Zubayear/holiday/pkg/jobs"
	"github.com/a-h/templ"
	"net/http"
)

func GetJobs(service jobs2.Service) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		jobs, err := service.GetJobs(context.Background())
		if err != nil {
			return
		}
		handler := presenter.Jobs(jobs)
		templ.Handler(handler).ServeHTTP(writer, request)
	}
}
