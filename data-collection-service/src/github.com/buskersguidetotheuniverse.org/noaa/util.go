package noaa

import (
	"errors"
	"fmt"
	"github.com/buskersguidetotheuniverse.org/types"
	"math"
	"strconv"
	"strings"
	"time"
)

// date-in format: 2018-01-22T21:05:00+00:00.  All timestamps are UTC
//
// note:  I'm pretty sure this isn't idiomatic Golang for generating errors...
//        It's also really verbose.  With more exposure to the std libs, it would be shorter.
func MakeKeyFromTimeStamp(station string, timeStamp string) (string, error) {
	var err error = nil
	s1 := strings.Replace(timeStamp, "T", "-", 1)
	s1 = strings.TrimSuffix(s1, "+00:00")

	dateParts := strings.Split(s1, "-")
	if len(dateParts) != 4 {
		return "", errors.New("Incoming timestamp DATE not in expected format: " + timeStamp)
	}

	year, _ := strconv.Atoi(dateParts[0])
	mm, _ := strconv.Atoi(dateParts[1])
	month := time.Month(mm)
	day, _ := strconv.Atoi(dateParts[2])

	timeParts := strings.Split(dateParts[3], ":")
	if len(timeParts) != 3 {
		return "", errors.New("Incoming timestamp TIME not in expected format: " + timeStamp)
	}

	hour, _ := strconv.Atoi(timeParts[0])
	minute, _ := strconv.Atoi(timeParts[1])
	second, _ := strconv.Atoi(timeParts[2])

	t := time.Date(year, time.Month(month), day, hour, minute, second, 0, time.UTC)
	unixTime := t.Unix()
	reverseOrderTimeStamp := math.MaxInt64 - unixTime

	stamp := fmt.Sprintf("%v%v", station, reverseOrderTimeStamp)
	return stamp, err
}

// stationID always takes the form of: https://api.weather.gov/stations/XXXX
func StationShortForm(stationLongForm string) string {
	return strings.TrimPrefix(stationLongForm, "https://api.weather.gov/stations/")
}

func ExtractIdsFromStationsResponse(response *types.StationsResponse, limit int) []string {
	stations := response.Features

	var stationIds []string
	for i := 0; i < limit && i < len(stations); i++ {
		id := stations[i].Properties.StationIdentifier
		stationIds = append(stationIds, id)
	}

	return stationIds
}
