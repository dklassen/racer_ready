package server

import (
	"encoding/json"
	"errors"
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

type Race struct {
	Name      string
	Technique string
	Gender    string
	Resort    string
	StartTime string
	Country   string
	Province  string
	Racers    []*Racer
}

type Racer struct {
	Bib         int
	CheckedAt   string
	Class       string
	Club        string
	Name        string
	OnCourse    bool
	Rank        string
	Run1        string
	Run1ms      *int64
	Run2        string
	Run2ms      *int64
	TotalTime   string
	TotalTimems *int64
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

func convertTimeStringtoMS(rawTime string) (*int64, error) {
	if strings.HasPrefix(rawTime, "DQ") {
		return nil, fmt.Errorf("tried parsing invalid string to duration %s", rawTime)
	}
	if strings.HasPrefix(rawTime, "DNS") {
		return nil, errors.New("tried parsing an invalid string to duration")
	}
	if strings.HasPrefix(rawTime, "DNF") {
		return nil, errors.New("tried parsing an invalid string to duration")
	}

	convertedDuration := strings.Replace(rawTime, ":", "m", 1)
	convertedDuration += "s"
	runTime, err := time.ParseDuration(convertedDuration)
	if err != nil {
		return nil, err
	}

	milliseconds := runTime.Milliseconds()
	return &milliseconds, nil
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

	racer := &Racer{
		Bib:       convertedBib,
		Name:      record["m"],
		CheckedAt: record["ms"],
		Club:      record["c"],
		Class:     record["s"],
		OnCourse:  false,
		Run1:      record["r1"],
		Run2:      record["r2"],
		TotalTime: record["tt"],
	}

	if run1ms, err := convertTimeStringtoMS(racer.Run1); err == nil {
		racer.Run1ms = run1ms
	}
	if run2ms, err := convertTimeStringtoMS(racer.Run2); err == nil {
		racer.Run2ms = run2ms
	}

	if racer.Run1 == "On Course" || racer.Run2 == "On Course" {
		racer.OnCourse = true
	}

	if totalTimems, err := convertTimeStringtoMS(racer.TotalTime); err == nil {
		racer.TotalTimems = totalTimems
	}

	return racer, nil
}

func parseNameField(field string) string {
	fieldParts := strings.Split(field, "=")
	return strings.TrimSpace(fieldParts[1])
}

func parseTechniqueField(field string) string {
	fieldParts := strings.Split(field, "=")
	return strings.TrimSpace(fieldParts[0])
}

func parseGenderField(field string) string {
	fieldParts := strings.Split(field, "=")
	return strings.TrimSpace(fieldParts[1])
}

func parseCountryField(field string) string {
	fieldParts := strings.Split(field, "=")
	return fieldParts[0]
}

func parseProvinceField(field string) string {
	fieldParts := strings.Split(field, "=")
	return fieldParts[1]
}

func parseResortField(field string) string {
	return strings.TrimSpace(field)
}

func parseStartTime(field string) string {
	return field
}

func parseRawResponse(rawData string) (*Race, error) {
	result := strings.Split(rawData, "|")

	race := &Race{}
	var tempRacer []string

	for _, v := range result {

		if !strings.ContainsAny(v, "=") {
			continue
		}

		field := strings.SplitN(v, "=", 2)

		switch field[0] {
		case "hN":
			race.Name = parseNameField(field[1])
		case "hT":
			race.Technique = parseTechniqueField(field[1])
			race.Gender = parseGenderField(field[1])
		case "hC":
			race.Country = parseCountryField(field[1])
			race.Province = parseProvinceField(field[1])
		case "hR":
			race.Resort = parseResortField(field[1])
		case "hST":
			race.StartTime = parseStartTime(field[1])
		case "hM":
			//
		case "b":
			if len(tempRacer) > 0 && strings.HasPrefix(tempRacer[0], "b=") {
				racer, err := buildRacerStruct(tempRacer)
				if err != nil {
					logrus.Error(err)
					continue
				}
				race.Racers = append(race.Racers, racer)
				tempRacer = tempRacer[:0]
			} else {
				tempRacer = tempRacer[:0]
			}
			tempRacer = append(tempRacer, v)
		default:
			tempRacer = append(tempRacer, v)
		}
	}

	if len(tempRacer) > 0 && strings.HasPrefix(tempRacer[0], "b=") {
		racer, err := buildRacerStruct(tempRacer)
		if err != nil {
			logrus.Error(err)
		}
		race.Racers = append(race.Racers, racer)
	}

	return race, nil

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

	jsonPayload, err := parseRawResponse(string(body))

	if err != nil {
		logrus.Error(err)
		http.Error(w, "Issue parsing data from live timing", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(jsonPayload)
}
