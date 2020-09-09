package dao

import "context"

type TypeUID uint32
type TypeID uint32

// Goal goals
type Goal struct {
	GoalInfo
	GoalDetail
}

// GoalInfo brief info
type GoalInfo struct {
	UID   uint32
	ID    uint32
	Title string
	TimeField
}

// GoalDetail detail
type GoalDetail struct {
	Detail   string
	Labels   []string
	Depends  []uint32
	Assignee string
}

// GoalDB objectives
type GoalDB interface {
	Create(ctx context.Context, g *Goal) error
	Update(ctx context.Context, g *Goal) error
	Get(ctx context.Context, uid TypeUID, id TypeID)
	Delete(ctx context.Context, uid TypeUID, id TypeID)
}
