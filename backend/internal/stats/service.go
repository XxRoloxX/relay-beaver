package stats

import (
	"backend/internal/request"
	"backend/pkg/models"
	"slices"
	"time"
)

type StatsService struct {
	Repo request.RequestRepository
}

type StatisticEntry struct {
	Timestamp int     `json:"timestamp"`
	Value     float64 `json:"value"`
}

type HostStats struct {
	Host           string           `json:"host"`
	TotalRequests  []StatisticEntry `json:"totalRequests"`
	AverageLatency []StatisticEntry `json:"averageLatency"`
	BadRequests    []StatisticEntry `json:"badRequests"`
	ServerErrors   []StatisticEntry `json:"serverErrors"`
}

func (service StatsService) GetTimestampInterval(timestamp int, interval int) int {
	return timestamp - (timestamp % interval)
}

func (service StatsService) GetDefaultFrom() int {
	return int(time.Now().AddDate(0, 0, -1).Unix())
}
func (service StatsService) GetDefaultTo() int {
	return int(time.Now().Unix())
}
func (service StatsService) GetDefaultInterval() int {
	return 3600
}

func (service StatsService) ConvertMapCountToStatsEntry(data map[int]int) []StatisticEntry {
	entries := []StatisticEntry{}
	for timestamp, value := range data {
		entries = append(entries, StatisticEntry{
			Timestamp: timestamp,
			Value:     float64(value),
		})
	}
	return entries
}

func (service StatsService) ConvertMapAverageToStatsEntry(data map[int][]int) []StatisticEntry {
	entries := []StatisticEntry{}
	for timestamp, values := range data {
		average := float64(0)
		for _, value := range values {
			average += float64(value)
		}
		average = average / float64(len(values))
		entries = append(entries, StatisticEntry{
			Timestamp: timestamp,
			Value:     average,
		})
	}
	return entries
}

func (service StatsService) GetHostStats(host string, from int, to int, interval int) HostStats {
	requests, _ := service.Repo.FindAllByHost(host)

	totalRequests,
		averageLatency,
		badRequests,
		serverErrors := service.buildStatMaps(requests, from, to, interval)

	totalRequestsEntries := service.ConvertMapCountToStatsEntry(totalRequests)
	averageLatencyEntries := service.ConvertMapAverageToStatsEntry(averageLatency)
	badRequestsEntries := service.ConvertMapCountToStatsEntry(badRequests)
	serverErrorsEntries := service.ConvertMapCountToStatsEntry(serverErrors)

	hostStats := HostStats{
		Host:           host,
		TotalRequests:  totalRequestsEntries,
		AverageLatency: averageLatencyEntries,
		BadRequests:    badRequestsEntries,
		ServerErrors:   serverErrorsEntries,
	}

	return hostStats
}
func (service StatsService) GetHosts() []string {
	proxiedRequests, _ := service.Repo.FindAll()
	hosts := make([]string, 0)

	for _, proxiedRequest := range proxiedRequests {
		host := proxiedRequest.Host()
		if !slices.Contains(hosts, host) {
			hosts = append(hosts, host)
		}
	}
	return hosts
}

/*
GetStats returns statistics for all requests in the given time interval
Returns a tuple of total requests, average latency, bad requests and server errors
*/

// TODO: Refactor this to use a struct instead of returning a tuple
func (service StatsService) buildStatMaps(requests []models.ProxiedRequest, from int, to int, interval int) (map[int]int, map[int][]int, map[int]int, map[int]int) {
	var totalRequests map[int]int = make(map[int]int)
	var averageLatency map[int][]int = make(map[int][]int)
	var badRequests map[int]int = make(map[int]int)
	var serverErrors map[int]int = make(map[int]int)

	for _, request := range requests {
		if request.Timestamp() <= from || request.Timestamp() >= to {
			continue
		}

		timestamp := service.GetTimestampInterval(request.Timestamp(), interval)
		if request.Response.IsBadRequest() {
			if _, ok := badRequests[timestamp]; !ok {
				badRequests[timestamp] = 0
			}
			badRequests[timestamp]++
		}

		if request.Response.IsServerError() {
			if _, ok := serverErrors[timestamp]; !ok {
				serverErrors[timestamp] = 0
			}
			serverErrors[timestamp]++
		}

		if _, ok := totalRequests[timestamp]; !ok {
			totalRequests[timestamp] = 0
		}
		totalRequests[timestamp]++

		if _, ok := averageLatency[timestamp]; !ok {
			averageLatency[timestamp] = []int{}
		}
		averageLatency[timestamp] = append(averageLatency[timestamp], int(request.Latency()))
	}

	return totalRequests, averageLatency, badRequests, serverErrors
}
