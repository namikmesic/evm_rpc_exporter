package collector

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

type EthBlockNumber struct {
	rpc      *rpc.Client
	desc     *prometheus.Desc
	provider string
}

func NewEthBlockNumber(rpc *rpc.Client, provider string) *EthBlockNumber {
	return &EthBlockNumber{
		rpc:      rpc,
		provider: provider,
		desc: prometheus.NewDesc(
			"eth_block_number",
			"number of the most recent block",
			[]string{"provider"},
			nil,
		),
	}
}

func (collector *EthBlockNumber) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.desc
}

func (collector *EthBlockNumber) Collect(ch chan<- prometheus.Metric) {
	var result hexutil.Uint64
	if err := collector.rpc.Call(&result, "eth_blockNumber"); err != nil {
		ch <- prometheus.NewInvalidMetric(collector.desc, err)
		return
	}

	value := float64(result)
	ch <- prometheus.MustNewConstMetric(collector.desc, prometheus.GaugeValue, value, collector.provider)
}
