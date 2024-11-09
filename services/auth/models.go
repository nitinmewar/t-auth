package authsvc

/* --------------------------------- SIGNUP --------------------------------- */
type SingupObject struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginObject struct {
	Email     string       `json:"email"`
	TypingDNA KeystrokeDNA `json:"typing_dna"`
	Password  string       `json:"password"`
}

type EmailCheckRequest struct {
	Email string `json:"email"`
}

type EMailCheckResponse struct {
	Exist bool `json:"exist"`
}

type UserObject struct {
	PID             string `json:"id"`
	Email           string `json:"email"`
	FirstName       string `json:"first_name,omitempty"`
	LastName        string `json:"last_name,omitempty"`
	IsKeystrokeDone bool   `json:"is_keystroke_done"`
	CreatedAt       int    `json:"created_at"`
}

/* ------------------------------- KEY STROKE ------------------------------- */
type KeystrokeDNA struct {
	UserPID    string     `json:"user_id"`
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
