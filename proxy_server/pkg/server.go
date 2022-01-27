package server

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// Live timing query parameters
// r = RACE ID
// m = LAST UPDATED AT
// u = UPDATE INTERVAL

const RACE_URL = "https://www.live-timing.com/includes/aj_race.php"

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

func buildRacerStruct(rawData []string) *Racer {
	record := map[string]string{}

	for _, v := range rawData {
		entry := strings.Split(v, "=")
		if len(entry) > 0 {
			record[entry[0]] = entry[1] //: len(entry)-1]
		}
	}

	if len(record) == 0 {
		return nil
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
	}

}

// Records seem to come in two forms
// [b=39 m=ALVES, Kayla ms=642966445.56509 c=CASCA s=U14 up=99900 fp=0 r1=DQg18=2147483627 r2=1:00.91=60910]
// [b=53|m=SHELLY, Eva|ms=642964974.21723|t=CAN|c=FORTU|s=U14|un=105987|up=99900|fp=0|r1=59.60=59600|r2=59.61=59610|tt=1:59.21
// first is

func convertToJSON(rawData string) (racer []*Racer) {

	result := strings.Split(rawData, "|")

	var a []string
	var racers []*Racer

	for _, v := range result {
		if strings.HasPrefix(v, "b") {
			if len(a) > 0 && a != nil {
				racers = append(racers, buildRacerStruct(a))
				a = a[:0]
			}

			a = append(a, v)
		}

		if len(a) > 0 {
			a = append(a, v)
		}
	}

	return racers
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

func RetrieveRaceResults(w http.ResponseWriter, r *http.Request) {
	live_timing_request, _ := http.NewRequest("GET", RACE_URL, nil)
	parameter_passthrough := live_timing_request.URL.Query()

	values := r.URL.Query()
	for k := range values {
		parameter_passthrough.Add(k, r.URL.Query().Get(k))
	}

	live_timing_request.URL.RawQuery = parameter_passthrough.Encode()

	client := &http.Client{}

	resp, err := client.Do(live_timing_request)

	if err != nil {
		logrus.Error(err)
		http.Error(w, "Issue retrieving data from Live Timing", http.StatusBadRequest)
		return
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	jsonPayload := convertToJSON(string(body))

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(jsonPayload)
}
