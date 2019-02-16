package goshopify

import (
	"fmt"
	"time"
)

const inventoryLevelBasePath = "admin/inventory_levels"

// InventoryLevelService is an interface for interacting with the inventory level
// endpoints of the Shopify API.
// See https://help.shopify.com/en/api/reference/inventory/inventorylevel
type InventoryLevelService interface {
	List(interface{}) ([]InventoryLevel, error)
	Set(InventoryLevelSet) (*InventoryLevel, error)
}

// InventoryLevelServiceOp handles communication with the inventory level related methods of the
// Shopify API.
type InventoryLevelServiceOp struct {
	client *Client
}

// InventoryLevel represents a Shopify inventory level.
type InventoryLevel struct {
	InventoryItemID   int        `json:"inventory_item_id"`
	LocationID        int        `json:"location_id"`
	UpdatedAt         *time.Time `json:"updated_at"`
	AdminGraphqlApIID string     `json:"admin_graphql_api_id"`
}

// InventoryLevelSet sets the inventory level data
type InventoryLevelSet struct {
	LocationID            int
	InventoryItemID       int
	Available             int
	DisconnectIfNecessary bool
}

// InventoryLevelResource represents the result from the inventory_levels/X.json endpoint
type InventoryLevelResource struct {
	InventoryLevel *InventoryLevel `json:"inventory_level"`
}

// InventoryLevelsResource represents the result from the inventory_levels.json endpoint
type InventoryLevelsResource struct {
	InventoryLevels []InventoryLevel `json:"inventory_levels"`
}

// InventoryLevelListOptions A struct for all available order list options.
// See: https://help.shopify.com/en/api/reference/inventory/inventorylevel#index
type InventoryLevelListOptions struct {
	InventoryItemIDs string `json:"inventory_item_ids,omitempty"`
	LocationIDs      string `json:"location_ids,omitempty"`
	Limit            int    `json:"limit,omitempty"`
	Page             int    `json:"page,omitempty"`
}

// List inventory levels
func (s *InventoryLevelServiceOp) List(options interface{}) ([]InventoryLevel, error) {
	path := fmt.Sprintf("%s.json", inventoryLevelBasePath)
	resource := new(InventoryLevelsResource)
	err := s.client.Get(path, resource, options)
	return resource.InventoryLevels, err
}

// Set inventory level for an inventory item at a location
func (s *InventoryLevelServiceOp) Set(inventoryLevelSet InventoryLevelSet) (*InventoryLevel, error) {
	path := fmt.Sprintf("%s/set.json", inventoryLevelBasePath)
	resource := new(InventoryLevelResource)
	err := s.client.Post(path, inventoryLevelSet, resource)
	return resource.InventoryLevel, err
}
