package server

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

const RACE_URL = "https://www.live-timing.com/includes/aj_race.php"

type BadRawInputError struct {
	rawRecord []string
}

func (b BadRawInputError) Error() string {
	return fmt.Sprintf("error parsing record string to struct: %v", b.rawRecord)
}

type Racer struct {
	Bib       int
	CheckedAt string
	Class     string
	Club      string
	Name      string
	Rank      string
	Run1      string
	Run2      string
	TotalTime string
}

func ConfigPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	return port
}

func buildRacerStruct(rawData []string) (*Racer, error) {
	record := map[string]string{}

	for _, v := range rawData {
		entry := strings.Split(v, "=")
		if len(entry) > 0 {
			record[entry[0]] = entry[1]
		}
	}

	if len(record) == 0 {
		return nil, BadRawInputError{rawData}
	}

	// lol error handling
	convertedBib, _ := strconv.Atoi(record["b"])

	return &Racer{
		Bib:       convertedBib,
		Name:      record["m"],
		CheckedAt: record["ms"],
		Club:      record["c"],
		Class:     record["s"],
		Run1:      record["r1"],
		Run2:      record["r2"],
		TotalTime: record["tt"],
	}, nil

}

func convertToJSON(rawData string) (racers []*Racer, err error) {
	result := strings.Split(rawData, "|")

	var a []string

	for _, v := range result {
		if strings.HasPrefix(v, "b") {
			if len(a) > 0 && a != nil {
				racer, err := buildRacerStruct(a)
				if err != nil {
					logrus.Error(err)
					continue
				}

				racers = append(racers, racer)

				a = a[:0]
			}

			a = append(a, v)
		}

		if len(a) > 0 {
			a = append(a, v)
		}
	}

	return racers, err
}

func checkResponseBody(response *http.Request) string {
	body, _ := ioutil.ReadAll(response.Body)
	return string(body)
}

func RequestLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		end := time.Now()
		latency := end.Sub(start)
		logrus.WithFields(logrus.Fields{
			"datetime":           start,
			"url":                r.URL,
			"ip":                 r.RemoteAddr,
			"latency_nanesecond": latency.Nanoseconds(),
			"http_user_agent":    r.UserAgent(),
			"http_method":        r.Method,
			"response_header":    w.Header(),
			"response_body":      checkResponseBody(r),
		}).Info("Served request")
	})
}

func NewHTTPClient() *http.Client {
	return &http.Client{}
}

func RetrieveRaceResults(w http.ResponseWriter, r *http.Request) {
	live_timing_request, _ := http.NewRequest("GET", RACE_URL, nil)
	parameter_passthrough := live_timing_request.URL.Query()

	values := r.URL.Query()
	for k := range values {
		parameter_passthrough.Add(k, r.URL.Query().Get(k))
	}

	live_timing_request.URL.RawQuery = parameter_passthrough.Encode()

	client := NewHTTPClient()

	resp, err := client.Do(live_timing_request)

	if err != nil {
		logrus.Error(err)
		http.Error(w, "Issue retrieving data from Live Timing", http.StatusBadRequest)
		return
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	jsonPayload, err := convertToJSON(string(body))

	if err != nil {
		logrus.Error(err)
		http.Error(w, "Issue parsing data from live timing", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(jsonPayload)
}
