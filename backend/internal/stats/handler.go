package stats

import (
	"backend/internal/database"
	"backend/internal/logger"
	"backend/internal/request"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var HOST_QUERY_PARAM = "host"
var FROM_QUERY_PARAM = "from"
var TO_QUERY_PARAM = "to"
var INTERVAL_QUERY_PARAM = "interval"

type StatsHandler struct {
	Service StatsService
	Logger  logger.HttpLogger
}

func NewStatsHandler() StatsHandler {
	return StatsHandler{
		Service: StatsService{
			Repo: request.RequestMongoRepository{
				Db: *database.Db,
			},
		},
		Logger: logger.HttpLogger{},
	}
}

func (h *StatsHandler) GetStats(w http.ResponseWriter, r *http.Request) {
	logger := h.Logger.Request(r)
	host := r.URL.Query().Get(HOST_QUERY_PARAM)

	if host == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Host parameter is required"))
		logger.Error("Host parameter is required")
		return
	}

	from, err := strconv.Atoi(r.URL.Query().Get(FROM_QUERY_PARAM))
	if err != nil {
		from = h.Service.GetDefaultFrom()
	}
	to, err := strconv.Atoi(r.URL.Query().Get(TO_QUERY_PARAM))
	if err != nil {
		to = h.Service.GetDefaultTo()
	}
	interval, err := strconv.Atoi(r.URL.Query().Get(INTERVAL_QUERY_PARAM))
	if err != nil {
		interval = h.Service.GetDefaultInterval()
	}

	stats := h.Service.GetHostStats(host, from, to, interval)

	serialized, error := json.Marshal(stats)

	if error != nil {
		logger.Error(fmt.Sprintf("Error serializing stats: %s", error.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error serializing stats"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(serialized)
}
