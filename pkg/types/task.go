package types

// Task tasks
type Task struct {
	Parent      ID
	ID          ID
	Title       string
	Description string
	Comment     string
	StartTime   int64
	EndTime     int64
	Assignee    []string

	Cost         uint16
	EstimateTime uint16
}
