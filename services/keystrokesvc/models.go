package keystrokesvc

type RequestBody struct {
	Data TypingMetrics `json:"data"`
}

type TypingMetrics struct {
	UserPID    string     `json:"userPID"`
	SampleText string     `json:"sampleText"`
	InputText  string     `json:"inputText"`
	Metrics    Metrics    `json:"metrics"`
	MetaData   MetaData   `json:"metaData"`
	DeviceInfo DeviceInfo `json:"deviceInfo"`
}

type Metrics struct {
	RawMetrics  RawMetrics `json:"rawMetrics"`
	WPM         int        `json:"wpm"`
	TotalEvents int        `json:"totalEvents"`
	UniqueKeys  int        `json:"uniqueKeys"`
}

type RawMetrics struct {
	DwellTimes      []float64 `json:"dwellTimes"`
	FlightTimes     []float64 `json:"flightTimes"`
	UpToUpTimes     []float64 `json:"upToUpTimes"`
	DownToDownTimes []float64 `json:"downToDownTimes"`
}

type MetaData struct {
	RecordingStartedAt int64 `json:"recordingStartedAt"`
	RecordingEndedAt   int64 `json:"recordingEndedAt"`
	TotalKeystrokes    int   `json:"totalKeystrokes"`
	TextLength         int   `json:"textLength"`
	TotalTimeInMs      int   `json:"totalTimeInMs"`
}

type DeviceInfo struct {
	Browser          string `json:"browser"`
	Version          string `json:"version"`
	OS               string `json:"os"`
	OSVersion        string `json:"osVersion"`
	DeviceType       string `json:"device_type"`
	ScreenResolution string `json:"screen_resolution"`
}
