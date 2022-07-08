package models

type RandomizerForm struct {
	Genres          []string
	MeanScoreFrom   string
	MeanScoreTo     string
	NumEpisodesFrom string
	NumEpisodesTo   string
	YearFrom        string
	YearTo          string
	DurationFrom    string
	DurationTo      string
}
