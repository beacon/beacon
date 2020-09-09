package types

// User def
type User struct {
	ID   ID
	Name string
}

// Team def
type Team struct {
	ID      ID
	Name    string
	OwnerID ID
}

// TeamMember team/member relationship
type TeamMember struct {
	TeamID    ID
	UserID    ID
	CreatedAt int64
}
