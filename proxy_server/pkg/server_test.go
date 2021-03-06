package server

import (
	"strings"
	"testing"

	"github.com/go-test/deep"
)

func createInt64Pointer(value int64) *int64 {
	pointer_value := new(int64)
	*pointer_value = value
	return pointer_value
}

func TestConfigPortWithPORTEnvVarSetPullsFromEnvironment(t *testing.T) {
	expectedPort := "1234"
	expected := ":" + expectedPort

	t.Setenv("PORT", expected)

	output := ConfigPort()
	if output != expected {
		t.Errorf("ConfigPort was incorrect, got: %s, want: %s.", output, expected)
	}

}

func TestConfigPortAddsExpectedPrefix(t *testing.T) {
	expectedPort := "1234"
	expected := ":" + expectedPort

	if strings.HasPrefix(expectedPort, ":") {
		t.Fatalf("Expected port test case should not start with :")
	}

	t.Setenv("PORT", expectedPort)

	output := ConfigPort()
	if output != expected {
		t.Errorf("ConfigPort was incorrect, got: %s, want: %s.", output, expected)
	}
}

func TestConfigPortWhenNotSetDefaultsTo8080(t *testing.T) {
	expected := ":8080"

	t.Setenv("PORT", "")

	output := ConfigPort()
	if output != expected {
		t.Errorf("ConfigPort was incorrect, got: %s, want: %s.", output, expected)
	}
}

func TestParseRawResponse(t *testing.T) {
	type test struct {
		input string
		want  *Race
		err   error
	}

	tests := []test{
		{
			input: "1=1=643758233.88783=N=30|hN=U=CALABOGIE U14 GS|hT=Giant Slalom=Women|hC=CAN=ON|hR=CALABOGIE PEAKS|hST=1/23/2022 10:00 AM|hM=|hLL=CalRacing.jpg|hLR=CalPeaks.jpg|hPr=99|hNreg=N|hID=227639|hP=pdf=Race Results|hP=pdf=Penalty Calculation|hP=pdf=Run 2 Condensed St|hP=pdf=Run 1 Condensed St|hE|b=39|m=ALVES, Kayla|ms=642966445.56509|c=CASCA|s=U14|up=99900|fp=0|r1=DQg18=2147483627|r2=1:00.91=60910",
			want: &Race{
				Name:      "CALABOGIE U14 GS",
				Technique: "Giant Slalom",
				Resort:    "CALABOGIE PEAKS",
				Gender:    "Women",
				Country:   "CAN",
				Province:  "ON",
				StartTime: "1/23/2022 10:00 AM",
				Racers: []*Racer{&Racer{
					Bib:         39,
					Name:        "ALVES, Kayla",
					CheckedAt:   "642966445.56509",
					Club:        "CASCA",
					Class:       "U14",
					Run1:        "DQg18",
					Run1ms:      nil,
					Run2:        "1:00.91",
					Run2ms:      createInt64Pointer(60910),
					TotalTime:   "",
					TotalTimems: nil,
				},
				},
			},
		},
	}

	for _, tc := range tests {
		got, with_err := parseRawResponse(tc.input)

		if with_err != nil && with_err.Error() != tc.err.Error() {
			t.Fatalf("expected: %+v, got: %+v", tc.err, with_err)
		}

		if diff := deep.Equal(tc.want, got); diff != nil {
			t.Error(diff)
		}
	}
}

func TestBuildRacerStruct(t *testing.T) {
	type test struct {
		input  []string
		want   *Racer
		err    error
		reason string
	}

	tests := []test{
		{input: []string{"b=39", "m=ALVES, Kayla", "ms=642966445.56509", "c=CASCA", "s=U14", "up=99900", "fp=0", "r1=DQg18=2147483627", "r2=1:00.91=60910"},
			want: &Racer{
				Bib:         39,
				Name:        "ALVES, Kayla",
				CheckedAt:   "642966445.56509",
				Club:        "CASCA",
				Class:       "U14",
				Run1:        "DQg18",
				Run1ms:      nil,
				Run2:        "1:00.91",
				Run2ms:      createInt64Pointer(60910),
				TotalTime:   "",
				TotalTimems: nil,
			},
			err:    nil,
			reason: "testing the golden path"},
		{input: []string{"b=53", "m=SHELLY, Eva", "ms=642964974.21723", "t=CAN", "c=FORTU", "s=U14", "un=105987", "up=99900", "fp=0", "r1=59.60=59600", "r2=59.61=59610", "tt=1:59.21"},
			want: &Racer{
				Bib:         53,
				Name:        "SHELLY, Eva",
				CheckedAt:   "642964974.21723",
				Club:        "FORTU",
				Class:       "U14",
				Run1:        "59.60",
				Run1ms:      createInt64Pointer(59600),
				Run2:        "59.61",
				Run2ms:      createInt64Pointer(59610),
				TotalTime:   "1:59.21",
				TotalTimems: createInt64Pointer(119210),
			},
			err:    nil,
			reason: "Is a longer format of the data we see coming in"},
		{input: []string{},
			want:   nil,
			err:    BadRawInputError{rawRecord: []string{}},
			reason: "Is a longer format of the data we see coming in"},
		{input: []string{"b=53", "m=SHELLY, Eva", "ms=642964974.21723", "t=CAN", "c=FORTU", "s=U14", "un=105987", "up=99900", "fp=0", "r1=On Course", "r2=59.61=59610", "tt=1:59.21"},
			want: &Racer{
				Bib:         53,
				Name:        "SHELLY, Eva",
				CheckedAt:   "642964974.21723",
				Club:        "FORTU",
				Class:       "U14",
				OnCourse:    true,
				Run1:        "On Course",
				Run1ms:      nil,
				Run2:        "59.61",
				Run2ms:      createInt64Pointer(59610),
				TotalTime:   "1:59.21",
				TotalTimems: createInt64Pointer(119210),
			},
			err:    nil,
			reason: ""},
		{input: []string{"b=53", "m=SHELLY, Eva", "ms=642964974.21723", "t=CAN", "c=FORTU", "s=U14", "un=105987", "up=99900", "fp=0", "r1=59.60=59600", "r2=On Course", "tt=1:59.21"},
			want: &Racer{
				Bib:         53,
				Name:        "SHELLY, Eva",
				CheckedAt:   "642964974.21723",
				Club:        "FORTU",
				Class:       "U14",
				OnCourse:    true,
				Run1:        "59.60",
				Run1ms:      createInt64Pointer(59600),
				Run2:        "On Course",
				Run2ms:      nil,
				TotalTime:   "1:59.21",
				TotalTimems: createInt64Pointer(119210),
			},
			err:    nil,
			reason: ""},
	}

	for _, tc := range tests {
		got, with_err := buildRacerStruct(tc.input)

		if with_err != nil && with_err.Error() != tc.err.Error() {
			t.Fatalf("expected: %+v, got: %+v", tc.err, with_err)
		}

		if diff := deep.Equal(tc.want, got); diff != nil {
			t.Error(diff)
		}
	}
}
