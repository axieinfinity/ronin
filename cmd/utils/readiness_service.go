package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
)

type Metric struct {
	Job string
}
type Result struct {
	Metric Metric
	Value  []string
}

type Data struct {
	ResultType string
	Result     []Result
}
type Response struct {
	Status string
	Data   Data
}

type handler struct {
	prometheus string
	block_lag  int
}

func respond(w http.ResponseWriter, body []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	_, _ = w.Write(body)
}

func errorJSON(msg string) []byte {
	buf := bytes.Buffer{}
	fmt.Fprintf(&buf, `{"error": "%s"}`, msg)
	return buf.Bytes()
}

func successJSON(msg string) []byte {
	buf := bytes.Buffer{}
	fmt.Fprintf(&buf, `{"message": "%s"}`, msg)
	return buf.Bytes()
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		respond(w, errorJSON("only GET requests are supported"), http.StatusMethodNotAllowed)
		return
	}
	var threshold_lag int64
	threshold_lag = int64(h.block_lag)

	now := time.Now()
	sec := now.Unix()

	query := "http://" + h.prometheus + "/api/v1/query?query=max%28chain_head_block%29+by%28job%29&time=" + strconv.FormatInt(sec, 10)
	resp, err := http.Get(query)
	if err != nil {
		log.Error("[Readiness] Failed to query to prometheus", "query", query, "error", err)
		respond(w, errorJSON("Unreachable Prometheus"), http.StatusOK)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("[Readiness] Failed to parse body", "error", err)
		respond(w, errorJSON("Failed to parse json body"), http.StatusOK)
		return
	}

	var response Response
	err = json.Unmarshal(body, &response)

	if response.Status != "success" {
		log.Error("[Readiness] Process query to prometheus", "status", response.Status)
		respond(w, errorJSON("Error when connecting to prometheus"), http.StatusOK)
		return
	}

	// Convert from string to int64
	max_head, err := strconv.ParseInt(response.Data.Result[0].Value[1], 10, 64)
	if err != nil {
		log.Error("[Readiness] Failed to convert string to int64", "error", err)
		respond(w, errorJSON("Failed to convert string to int64"), http.StatusOK)
		return
	}
	current_block_lag := max_head - core.HeadFastBlockGauge.Value()

	log.Info("[Readiness] Current block number statics", "Max head", max_head, "Current Block", core.HeadFastBlockGauge.Value(), "current Lag", current_block_lag)

	if current_block_lag > threshold_lag {
		respond(w, errorJSON("Block lag is larger than threshold"), http.StatusInternalServerError)
		return
	}

	respond(w, successJSON("Block lag catchs up"), http.StatusOK)
}

// New constructs a new GraphQL service instance.
func New(stack *node.Node, backend ethapi.Backend, cors, vhosts []string, prometheus string, block_lag int) error {
	if backend == nil {
		panic("missing backend")
	}
	// check if http server with given endpoint exists and enable graphQL on it
	return newHandler(stack, backend, cors, vhosts, prometheus, block_lag)
}

// newHandler returns a new `http.Handler` that will answer GraphQL queries.
// It additionally exports an interactive query browser on the / endpoint.
func newHandler(stack *node.Node, backend ethapi.Backend, cors, vhosts []string, prometheus string, block_lag int) error {

	h := handler{
		prometheus: prometheus,
		block_lag:  block_lag,
	}
	handler := node.NewHTTPHandlerStack(h, cors, vhosts)

	stack.RegisterHandler("readiness", "/readiness", handler)

	return nil
}
