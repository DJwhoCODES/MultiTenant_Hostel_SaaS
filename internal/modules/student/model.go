package student

import "time"

type Student struct {
	ID        string    `bson:"_id,omitempty" json:"id"`
	TenantID  string    `bson:"tenantId" json:"tenantId"`
	Name      string    `bson:"name" json:"name"`
	Phone     string    `bson:"phone" json:"phone"`
	RoomID    string    `bson:"roomId,omitempty" json:"roomId"`
	BedNumber int       `bson:"bedNumber,omitempty" json:"bedNumber"`
	CheckIn   time.Time `bson:"checkIn" json:"checkIn"`
	Status    string    `bson:"status" json:"status"` // active, left
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}
