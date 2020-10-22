package fixtures

type Basic struct {
	WeekendInfo WeekendInfo `yaml:"WeekendInfo"`
}

type WeekendInfo struct {
	TrackName string `yaml:"TrackName"`
	TrackID   int    `yaml:"TrackID"`
}
