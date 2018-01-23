package weather

import "testing"

func TestMakeKeyFromNoaaTimeStamp(t *testing.T) {
	const station = "MSP"
	const correctExample = "2018-01-23T00:24:00+00:00"
	const badDate = "foo-2018-01-23T00:24:00+00:"
	const badTime = "2018-01-23T99:00:24:00:99+00:00"

	_, err := MakeKeyFromNoaaTimeStamp(station, correctExample)
	if err != nil {
		t.Errorf("MakeKeyFromNoaaTimeStamp errored on valid data (%v):  %v", correctExample, err)
	}

	_, err = MakeKeyFromNoaaTimeStamp(station, badDate)
	if err == nil {
		t.Errorf("MakeKeyFromNoaaTimeStamp FAILED TO ERROR on bad DATE: %v", badDate)
	}

	_, err = MakeKeyFromNoaaTimeStamp(station, badTime)
	if err == nil {
		t.Errorf("MakeKeyFromNoaaTimeStamp FAILED TO ERROR on bad TIME: %v", badTime)
	}
}
