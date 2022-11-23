package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/core"
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
	blockLag   int64
}

func respond(w http.ResponseWriter, body []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	_, err := w.Write(body)
	if err != nil {
		log.Error("[Readiness] Failed to write", "body", body, "error", err)
	}
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

	queryParam := url.QueryEscape("max(chain_head_block) by (job)")
	query := fmt.Sprintf("http://%s/api/v1/query?query=%s&time=%d", h.prometheus, queryParam, time.Now().Unix())
	resp, err := http.Get(query)
	if err != nil {
		log.Error("[Readiness] Failed to query to prometheus", "query", query, "error", err)
		respond(w, errorJSON("Unreachable Prometheus"), http.StatusOK)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("[Readiness] Failed to read response body", "error", err)
		respond(w, errorJSON("Failed to read response body"), http.StatusOK)
		return
	}

	var response Response
	err = json.Unmarshal(body, &response)
	// "value":[1669104409.589,"11849564"]}]}}, value return 2 different types, should bypass error here.

	if response.Status != "success" {
		log.Error("[Readiness] Process query to prometheus", "status", response.Status, "query", query)
		respond(w, errorJSON("Error when connecting to prometheus"), http.StatusOK)
		return
	}

	// Convert from string to int64
	maxHead, err := strconv.ParseInt(response.Data.Result[0].Value[1], 10, 64)
	if err != nil {
		log.Error("[Readiness] Failed to convert string to int64", "error", err)
		respond(w, errorJSON("Failed to convert string to int64"), http.StatusOK)
		return
	}
	currentBlockLag := maxHead - core.HeadFastBlockGauge.Value()

	log.Info("[Readiness] Current block number statics", "Max head", maxHead, "Current Block", core.HeadFastBlockGauge.Value(), "current Lag", currentBlockLag)

	if currentBlockLag > h.blockLag {
		respond(w, errorJSON("Block lag is larger than threshold"), http.StatusInternalServerError)
		return
	}

	respond(w, successJSON("Block lag catchs up"), http.StatusOK)
}

// New constructs a new Readiness service instance.
func NewReadinessHandler(stack *node.Node, cors, vhosts []string, prometheus string, blockLag int64) error {
	// check if http server with given endpoint exists and enable Readiness on it
	return newHandler(stack, cors, vhosts, prometheus, blockLag)
}

// newHandler returns a new `http.Handler` that will answer Readiness requests.
func newHandler(stack *node.Node, cors, vhosts []string, prometheus string, blockLag int64) error {
	h := handler{
		prometheus: prometheus,
		blockLag:   blockLag,
	}
	handler := node.NewHTTPHandlerStack(h, cors, vhosts)
	stack.RegisterHandler("readiness", "/readiness", handler)
	return nil
}
