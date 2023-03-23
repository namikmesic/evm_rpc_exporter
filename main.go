package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/namikmesic/evm_rpc_exporter/internal/collector"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	endpoint := os.Getenv("RPC_ENDPOINT")
	provider := os.Getenv("RPC_PROVIDER")

	rpc, err := rpc.Dial(endpoint)
	if err != nil {
		log.Fatal(err)
	}

	registry := prometheus.NewPedanticRegistry()

	registry.MustRegister(
		collector.NewEthBlockNumber(rpc, provider),
	)

	handler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{
		ErrorLog:      log.New(os.Stderr, log.Prefix(), log.Flags()),
		ErrorHandling: promhttp.ContinueOnError,
	})

	http.Handle("/metrics", handler)
	log.Fatal(http.ListenAndServe(":9368", nil))
}
