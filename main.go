package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	rkboot "github.com/rookie-ninja/rk-boot"
	rkbootmux "github.com/rookie-ninja/rk-boot/mux"
	rkmuxinter "github.com/rookie-ninja/rk-mux/interceptor"
)

type ResultResponse struct {
	Message string
}

// Response.
type ServiceResponse struct {
	Result string `json:"result"`
}

// Response.
type GreeterResponse struct {
	Message string
}

var (
	requestsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "myapp_requests_total",
			Help: "Total number of requests received",
		},
	)
	requestsFail = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "myapp_requests_fails",
			Help: "Total number of requests fails",
		},
	)
	requestsCreateService = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "myapp_requests_add",
			Help: "Total number of requests to add a service",
		},
	)
	requestsCancelService = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "myapp_requests_cancel",
			Help: "Total number of requests to cancel a service",
		},
	)
	requestsHealth = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "myapp_requests_health",
			Help: "Total number of requests to verify the health of the app",
		},
	)
	requestsMain = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "myapp_requests_main",
			Help: "Total number of requests to the main page",
		},
	)
)

func init() {
	prometheus.MustRegister(requestsTotal)
	prometheus.MustRegister(requestsFail)
	prometheus.MustRegister(requestsCreateService)
	prometheus.MustRegister(requestsCancelService)
	prometheus.MustRegister(requestsHealth)
	prometheus.MustRegister(requestsMain)
}

// @Summary Greeter service
// @Id 1
// @version 1.0
// @produce application/json
// @Param name query string true "Input name"
// @Success 200 {object} GreeterResponse
// @Router /greeter [get]
func Greeter(writer http.ResponseWriter, request *http.Request) {
	requestsTotal.Inc()
	rkmuxinter.WriteJson(writer, http.StatusOK, &GreeterResponse{
		Message: fmt.Sprintf("Hello %s!", request.URL.Query().Get("name")),
	})
}

// @Summary Create service
// @Id 2
// @version 1.0
// @produce application/json
// @Param name formData string true "name"
// @Param address formData string true "address"
// @Success 200 {object} ServiceResponse
// @Router /service/create [post]
func CreateService(writer http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	if name == "" {
		http.Error(writer, "Invalid value for name", http.StatusBadRequest)
		return
	}

	address := request.FormValue("address")
	if address == "" {
		http.Error(writer, "Invalid value for address", http.StatusBadRequest)
		return
	}

	response := "Hi " + name + " your service was succesfully created to the address " + address + " see you soon!"
	requestsCreateService.Inc()
	requestsTotal.Inc()
	rkmuxinter.WriteJson(writer, http.StatusOK, &ServiceResponse{
		Result: response,
	})
}

// @Summary Cancel service
// @Id 4
// @version 1.0
// @produce application/json
// @Param serviceID formData int true "serviceID"
// @Success 200 {object} ServiceResponse
// @Router /service/cancel [delete]
func CancelService(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(request.FormValue("serviceID"))
	if err != nil {
		requestsFail.Inc()
		http.Error(writer, "Invalid value for serviceID", http.StatusBadRequest)
		return
	}

	response := "Service " + strconv.Itoa(id) + " was successfully cancelled"

	requestsTotal.Inc()
	requestsCancelService.Inc()
	rkmuxinter.WriteJson(writer, http.StatusOK, &ServiceResponse{
		Result: response,
	})

}

// @Summary Health service
// @Id 7
// @version 1.0
// @produce application/json
// @Success 200 {object} GreeterResponse
// @Router /health [get]
func Health(writer http.ResponseWriter, request *http.Request) {
	rkmuxinter.WriteJson(writer, http.StatusOK, &GreeterResponse{
		Message: fmt.Sprintf("Health ok!"),
	})
	requestsTotal.Inc()
	requestsHealth.Inc()
}

func Test(writer http.ResponseWriter, request *http.Request) {
	rkmuxinter.WriteJson(writer, http.StatusOK, &GreeterResponse{
		Message: fmt.Sprintf("testing v3.000!"),
	})
	requestsTotal.Inc()
}

func MainPage(writer http.ResponseWriter, request *http.Request) {
	rkmuxinter.WriteJson(writer, http.StatusOK, &GreeterResponse{
		Message: fmt.Sprintf("Hello, please go to /docs to see the documentation!!!"),
	})
	requestsTotal.Inc()
	requestsMain.Inc()
}

func main() {
	// Create a new boot instance.
	boot := rkboot.NewBoot()

	// Register handler
	entry := rkbootmux.GetMuxEntry("greeter")
	entry.Router.NewRoute().Methods(http.MethodGet).Path("/greeter").HandlerFunc(Greeter)
	entry.Router.NewRoute().Methods(http.MethodGet).Path("/health").HandlerFunc(Health)
	entry.Router.NewRoute().Methods(http.MethodGet).Path("/test").HandlerFunc(Test)
	entry.Router.NewRoute().Methods(http.MethodGet).Path("/").HandlerFunc(MainPage)
	entry.Router.NewRoute().Methods(http.MethodPost).Path("/service/create").HandlerFunc(CreateService)
	entry.Router.NewRoute().Methods(http.MethodDelete).Path("/service/cancel").HandlerFunc(CancelService)
	entry.Router.NewRoute().Path("/metrics").Handler(promhttp.Handler())

	// Bootstrap
	boot.Bootstrap(context.TODO())

	boot.WaitForShutdownSig(context.TODO())
}
