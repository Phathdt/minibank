package postgresql

import "minibank/auth"

func (u User) MapToEntity() auth.User {
	return auth.User{
		ID:         u.ID,
		Email:      u.Email,
		Password:   u.Password,
		Role:       u.Role,
		InsertedAt: u.InsertedAt,
		UpdatedAt:  u.UpdatedAt,
	}
}
