package noaa

import "testing"

func TestMakeKeyFromTimeStamp(t *testing.T) {
	const station = "MSP"
	const correctExample = "2018-01-23T00:24:00+00:00"
	const badDate = "foo-2018-01-23T00:24:00+00:"
	const badTime = "2018-01-23T99:00:24:00:99+00:00"

	_, err := MakeKeyFromTimeStamp(station, correctExample)
	if err != nil {
		t.Errorf("MakeKeyFromTimeStamp errored on valid data (%v):  %v", correctExample, err)
	}

	_, err = MakeKeyFromTimeStamp(station, badDate)
	if err == nil {
		t.Errorf("MakeKeyFromTimeStamp FAILED TO ERROR on bad DATE: %v", badDate)
	}

	_, err = MakeKeyFromTimeStamp(station, badTime)
	if err == nil {
		t.Errorf("MakeKeyFromTimeStamp FAILED TO ERROR on bad TIME: %v", badTime)
	}
}
