package models

type Status struct {
	Watching    string `json:"watching"`
	Completed   string `json:"completed"`
	OnHold      string `json:"on_hold"`
	Dropped     string `json:"dropped"`
	PlanToWatch string `json:"plan_to_watch"`
}
