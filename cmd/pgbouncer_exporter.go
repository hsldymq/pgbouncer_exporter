package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/hsldymq/pgbouncer_exporter/collector"
	"github.com/hsldymq/pgbouncer_exporter/logger"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/version"
)

const (
	indexHTML = `
	<html>
		<head>
			<title>PgBouncer Exporter</title>
		</head>
		<body>
			<h1>PgBouncer Exporter</h1>
			<p>
			<a href='%s'>Metrics</a>
			</p>
		</body>
	</html>`
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	var (
		listenAddress           = flag.String("web.listen-address", ":9127", "Address on which to expose metrics and web interface.")
		connectionStringPointer = flag.String("pgBouncer.connectionString", "postgres://postgres:postgres@localhost:6543/pgbouncer?sslmode=disable",
			"Connection string for accessing pgBouncer. Can also be set using environment variable DATA_SOURCE_NAME")
		metricsPath      = flag.String("web.telemetry-path", "/metrics", "Path under which to expose metrics.")
		metricsNamespace = flag.String("telemetry.namespace", "pgbouncer", "metrics namespace")
	)
	flag.Parse()

	connectionString := getEnv("DATA_SOURCE_NAME", *connectionStringPointer)
	e := collector.NewExporter(connectionString, *metricsNamespace)
	prometheus.MustRegister(e)

	logger.Entry().Infoln("Starting pgbouncer exporter version: ", version.Info())

	http.Handle(*metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(fmt.Sprintf(indexHTML, *metricsPath)))
	})

	if err := http.ListenAndServe(*listenAddress, nil); err != nil {
		logger.Entry().WithError(err).Fatal("exporter shutdown")
	}
}
