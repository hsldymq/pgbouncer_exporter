# PgBouncer exporter

Prometheus exporter for PgBouncer.

## As an exporter
### Building and running
```shell
go build -o pgbouncer_exporter ./cmd/pgbouncer_exporter.go

./pgbouncer_exporter \
    -web.listen-address=:9127 \
    -web.telemetry-path=/metrics \
    -pgBouncer.connectionString=postgres://user:password@host:port/dbname?sslmode=disable
```

It exports metrics at `0.0.0.0:9127/metrics`

## As a library



## Metrics

Metric | tags | Description
---------|----|-------------
stats.avg_query_count | database | Average queries per second in last stat period
stats.avg_query | database | The average query duration, shown as microsecond
stats.avg_query_time | database | Average query duration in microseconds
stats.avg_recv | database | Average received (from clients) bytes per second
stats.avg_req | database | The average number of requests per second in last stat period, shown as request/second
stats.avg_sent | database | Average sent (to clients) bytes per second
stats.avg_wait_time | database | Time spent by clients waiting for a server in microseconds (average per second)
stats.avg_xact_count | database | Average transactions per second in last stat period
stats.avg_xact_time | database | Average transaction duration in microseconds
stats.bytes_received_per_second | database | The total network traffic received, shown as byte/second
stats.bytes_sent_per_second | database | The total network traffic sent, shown as byte/second
stats.total_query_count | database | Total number of SQL queries pooled
stats.total_query_time | database | Total number of microseconds spent by pgbouncer when actively connected to PostgreSQL, executing queries
stats.total_received | database | Total volume in bytes of network traffic received by pgbouncer, shown as bytes
stats.total_requests | database | Total number of SQL requests pooled by pgbouncer, shown as requests
stats.total_sent | database | Total volume in bytes of network traffic sent by pgbouncer, shown as bytes
stats.total_wait_time | database | Time spent by clients waiting for a server in microseconds
stats.total_xact_count | database | Total number of SQL transactions pooled
stats.total_xact_time | database | Total number of microseconds spent by pgbouncer when connected to PostgreSQL in a transaction, either idle in transaction or executing queries
pools.cl_active | database | Client connections linked to server connection and able to process queries, shown as connection
pools.cl_waiting | database | Client connections waiting on a server connection, shown as connection
pools.sv_active | database | Server connections linked to a client connection, shown as connection
pools.sv_idle | database | Server connections idle and ready for a client query, shown as connection
pools.sv_used | database | Server connections idle more than server_check_delay, needing server_check_query, shown as connection
pools.sv_tested | database | Server connections currently running either server_reset_query or server_check_query, shown as connection
pools.sv_login | database | Server connections currently in the process of logging in, shown as connection
pools.maxwait | database | Age of oldest unserved client connection, shown as second
