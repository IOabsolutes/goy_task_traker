package models

import (
	"fmt"
	"time"
)

// SubscriptionType represents subscription levels with constant values
type SubscriptionType string

const (
	SubscriptionFree    SubscriptionType = "free"
	SubscriptionPremium SubscriptionType = "premium"
)

// IsValid checks if the subscription type is valid
func (s SubscriptionType) IsValid() bool {
	switch s {
	case SubscriptionFree, SubscriptionPremium:
		return true
	}
	return false
}

// GetTaskLimit returns the task limit for each subscription type
func (s SubscriptionType) GetTaskLimit() int {
	switch s {
	case SubscriptionFree:
		return 25
	case SubscriptionPremium:
		return 2000
	default:
		return 0
	}
}

// GetPrice returns the monthly price for each subscription type
func (s SubscriptionType) GetPrice() float64 {
	switch s {
	case SubscriptionFree:
		return 0.0
	case SubscriptionPremium:
		return 19.99
	default:
		return 0.0
	}
}

// UserStatus represents user account status
type UserStatus string

const (
	StatusActive    UserStatus = "active"
	StatusInactive  UserStatus = "inactive"
	StatusSuspended UserStatus = "suspended"
	StatusDeleted   UserStatus = "deleted"
)

// User represents a user in the task tracker system
type User struct {
	ID           string           `json:"id" db:"id"`
	Name         string           `json:"name" db:"name"`
	LastName     string           `json:"last_name" db:"last_name"`
	Email        string           `json:"email" db:"email"`
	Password     string           `json:"-" db:"password"` // "-" excludes from JSON
	Status       UserStatus       `json:"status" db:"status"`
	Subscription SubscriptionType `json:"subscription" db:"subscription"` // This is your subscription field!
	CreatedAt    time.Time        `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time        `json:"updated_at" db:"updated_at"`
	LastLoginAt  *time.Time       `json:"last_login_at,omitempty" db:"last_login_at"`
}

// NewUser creates a new user with default values
func NewUser(name, lastName, email, password string) *User {
	return &User{
		Name:         name,
		LastName:     lastName,
		Email:        email,
		Password:     password,
		Status:       StatusActive,
		Subscription: SubscriptionFree, // Default subscription
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

// UpgradeSubscription changes user's subscription to a new type
func (u *User) UpgradeSubscription(newSubscription SubscriptionType) error {
	if !newSubscription.IsValid() {
		return fmt.Errorf("invalid subscription type: %s", newSubscription)
	}
	u.Subscription = newSubscription
	u.UpdatedAt = time.Now()
	return nil
}

// CanCreateTasks checks if the user can create more tasks based on subscription
func (u *User) CanCreateTasks(currentTaskCount int) bool {
	limit := u.Subscription.GetTaskLimit()
	return limit == -1 || currentTaskCount < limit
}

// GetSubscriptionPrice returns the monthly price for user's current subscription
func (u *User) GetSubscriptionPrice() float64 {
	return u.Subscription.GetPrice()
}

// IsActive checks if the user account is active
func (u *User) IsActive() bool {
	return u.Status == StatusActive
}

// FullName returns the user's full name
func (u *User) FullName() string {
	return u.Name + " " + u.LastName
}
