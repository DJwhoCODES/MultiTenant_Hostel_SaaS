package auth

import "time"

type User struct {
	ID        string    `bson:"_id,omitempty" json:"id"`
	TenantID  string    `bson:"tenantId" json:"tenantId"`
	Name      string    `bson:"name" json:"name"`
	Email     string    `bson:"email" json:"email"`
	Password  string    `bson:"password,omitempty"`
	Role      string    `bson:"role" json:"role"` // owner, manager
	CreatedAt time.Time `bson:"createdAt"`
}

type Tenant struct {
	ID        string    `bson:"_id,omitempty" json:"id"`
	Name      string    `bson:"name" json:"name"`
	Subdomain string    `bson:"subdomain" json:"subdomain"`
	CreatedAt time.Time `bson:"createdAt"`
}
