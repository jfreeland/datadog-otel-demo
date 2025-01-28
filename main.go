package main

import (
	"log"
	"net/http"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	span := tracer.StartSpan("health_check",
		tracer.ResourceName("/health"),
		tracer.SpanType(ext.SpanTypeWeb),
	)
	defer span.Finish()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	span := tracer.StartSpan("get_user",
		tracer.ResourceName("/user"),
		tracer.SpanType(ext.SpanTypeWeb),
	)
	defer span.Finish()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"id": 1, "name": "Test User"}`))
}

func main() {
	tracer.Start(tracer.WithServiceName("apm-demo"))
	defer tracer.Stop()

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/user", userHandler)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
