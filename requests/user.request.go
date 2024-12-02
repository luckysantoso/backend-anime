package requests

import "time"

type UserRequest struct {
	Name     string `json:"name"  binding:"required"`
	Email    string `json:"email"  binding:"required,email"` // tanpa spasi binding-nya
	Address  string `json:"address"  binding:"required"`
	BornDate time.Time `json:"born_date"  binding:"required"`
}