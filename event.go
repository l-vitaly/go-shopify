package goshopify

import (
	"fmt"
	"time"
)

const eventBasePath = "admin/events"

// EventService is an interface for interacting with the events
// endpoints of the Shopify API.
// See https://help.shopify.com/en/api/reference/events/event
type EventService interface {
	List(interface{}) ([]Event, error)
}

// EventServiceOp handles communication with the event related methods of the
// Shopify API.
type EventServiceOp struct {
	client *Client
}

// Event represents a Shopify event.
type Event struct {
	ID          int        `json:"id"`
	Arguments   []string   `json:"arguments"`
	Body        string     `json:"body"`
	CreatedAt   *time.Time `json:"created_at"`
	Path        string     `json:"path"`
	Message     string     `json:"message"`
	SubjectID   int        `json:"subject_id"`
	SubjectType string     `json:"subject_type"`
	Verb        string     `json:"verb"`
}

// EventResource represents the result from the events/X.json endpoint
type EventResource struct {
	Event *Event `json:"event"`
}

// EventsResource represents the result from the events.json endpoint
type EventsResource struct {
	Events []Event `json:"events"`
}

// EventListOptions A struct for all available order list options.
// See: https://help.shopify.com/en/api/reference/events/event#index
type EventListOptions struct {
	Limit        int       `json:"limit,omitempty"`
	Page         int       `json:"page,omitempty"`
	SinceID      int       `json:"since_id,omitempty"`
	CreatedAtMin time.Time `json:"created_at_min,omitempty"`
	CreatedAtMax time.Time `json:"created_at_max,omitempty"`
	Filter       string    `json:"filter,omitempty"`
	Verb         string    `json:"verb,omitempty"`
	Fields       string    `json:"fields,omitempty"`
}

// List events
func (s *EventServiceOp) List(options interface{}) ([]Event, error) {
	path := fmt.Sprintf("%s.json", eventBasePath)
	resource := new(EventsResource)
	err := s.client.Get(path, resource, options)
	return resource.Events, err
}
