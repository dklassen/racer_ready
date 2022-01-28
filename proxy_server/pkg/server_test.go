package server

import (
	"reflect"
	"testing"
)

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
				Bib:       39,
				Name:      "ALVES, Kayla",
				CheckedAt: "642966445.56509",
				Club:      "CASCA",
				Class:     "U14",
				Run1:      "DQg18",
				Run2:      "1:00.91",
				TotalTime: "",
			},
			err:    nil,
			reason: "testing the golden path"},
		{input: []string{"b=53", "m=SHELLY, Eva", "ms=642964974.21723", "t=CAN", "c=FORTU", "s=U14", "un=105987", "up=99900", "fp=0", "r1=59.60=59600", "r2=59.61=59610", "tt=1:59.21"},
			want: &Racer{
				Bib:       53,
				Name:      "SHELLY, Eva",
				CheckedAt: "642964974.21723",
				Club:      "FORTU",
				Class:     "U14",
				Run1:      "59.60",
				Run2:      "59.61",
				TotalTime: "1:59.21",
			},
			err:    nil,
			reason: "Is a longer format of the data we see coming in"},
		{input: []string{},
			want:   nil,
			err:    BadRawInputError{rawRecord: []string{}},
			reason: "Is a longer format of the data we see coming in"},
	}

	for _, tc := range tests {
		got, with_err := buildRacerStruct(tc.input)

		if with_err != nil && with_err.Error() != tc.err.Error() {
			t.Fatalf("expected: %+v, got: %+v", tc.err, with_err)
		}

		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %+v, got: %+v", tc.want, got)
		}
	}
}
