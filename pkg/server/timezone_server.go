package server

import (
	"encoding/json"
	httpCommon "myutility.com/m/v2/pkg/http"
	"myutility.com/m/v2/pkg/model"
	"net/http"
)

type TimeZoneServer struct {
	*httpCommon.Api
	port int
	name string
}

func (t *TimeZoneServer) Start() {
	api := &httpCommon.Api{
		Port: t.port,
	}
	t.Api = api
	api.Initialize()
	t.addRoutes()
	api.Start()
}

func (t *TimeZoneServer) addRoutes() {
	t.Api.Router.Handler("GET", "/health", t.healthCheck())
}

func NewTimeZoneServer(port int, name string) *TimeZoneServer {
	return &TimeZoneServer{
		port: port,
		name: name,
	}
}

func (t *TimeZoneServer) healthCheck() http.HandlerFunc {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		message := model.Message{
			Status: "healthy",
		}

		jsonData, err := json.Marshal(message)
		if err != nil {
			http.Error(w, "Json Parsing Error", http.StatusInternalServerError)
		}
		w.Write(jsonData)
	})

	return handler
}
