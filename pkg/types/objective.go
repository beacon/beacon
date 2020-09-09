package types

// ID type
type ID uint64

// Objective root
type Objective struct {
	ID          ID
	Title       string
	Description string
	StartTime   int64
	EndTime     int64
}

// KeyResult key result
type KeyResult struct {
	ObjID       ID
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
	TimeField
}
