package fixtures

type Basic struct {
	WeekendInfo WeekendInfo        `yaml:"WeekendInfo"`
	SessionInfo SessionInfoDetails `yaml:"SessionInfo"`
}

type WeekendInfo struct {
	TrackName string `yaml:"TrackName"`
	TrackID   int    `yaml:"TrackID"`
}

type SessionInfoDetails struct {
	Sessions []SessionInfoSession `yaml:"Sessions"`
}

type SessionInfoSession struct {
	SessionType      string                   `yaml:"SessionType"`
	ResultsPositions []SessionResultsPosition `yaml:"ResultsPositions"`
}

type SessionResultsPosition struct {
	Position      int     `yaml:"Position"`
	ClassPosition int     `yaml:"ClassPosition"`
	Time          float64 `yaml:"Time"`
	FastestTime   float32 `yaml:"FastestTime"`
}
