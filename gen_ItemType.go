package cloudpersister

import (
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/codegp/game-object-types/types"

	"golang.org/x/net/context"
)

// GetItemType retrieves a ItemType by its ID.
func (c *CloudPersister) GetItemType(id int64) (*types.ItemType, error) {
	ctx := context.Background()
	k := datastore.NewKey(ctx, "ItemType", "", id, nil)
	ItemType := &types.ItemType{}
	if err := c.DatastoreClient().Get(ctx, k, ItemType); err != nil {
		return nil, fmt.Errorf("datastoredb: could not get ItemType: %v", err)
	}
	ItemType.ID = id
	return ItemType, nil
}

// GetMultiItemType retrieves a list of ItemTypes by their ID.
func (c *CloudPersister) GetMultiItemType(ids []int64) ([]*types.ItemType, error) {
	if len(ids) == 0 {
		return []*types.ItemType{}, nil
	}
	ctx := context.Background()
	ks := make([]*datastore.Key, len(ids))
	ItemTypes := make([]*types.ItemType, len(ids))
	for i, id := range ids {
		ks[i] = datastore.NewKey(ctx, "ItemType", "", id, nil)
		ItemTypes[i] = &types.ItemType{}
	}
	if err := c.DatastoreClient().GetMulti(ctx, ks, ItemTypes); err != nil {
		return nil, fmt.Errorf("datastoredb: could not get ItemTypes: %v", err)
	}
	for i, id := range ids {
		ItemTypes[i].ID = id
	}
	return ItemTypes, nil
}

// AddItemType saves a given ItemType, assigning it a new ID.
func (c *CloudPersister) AddItemType(b *types.ItemType) (*types.ItemType, error) {
	ctx := context.Background()
	k := datastore.NewIncompleteKey(ctx, "ItemType", nil)
	k, err := c.DatastoreClient().Put(ctx, k, b)
	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not put ItemType: %v", err)
	}
	b.ID = k.ID()
	return b, nil
}

// DeleteItemType removes a given ItemType by its ID.
func (c *CloudPersister) DeleteItemType(id int64) error {
	ctx := context.Background()
	k := datastore.NewKey(ctx, "ItemType", "", id, nil)
	if err := c.DatastoreClient().Delete(ctx, k); err != nil {
		return fmt.Errorf("datastoredb: could not delete ItemType: %v", err)
	}
	return nil
}

// UpdateItemType updates the entry for a given ItemType.
func (c *CloudPersister) UpdateItemType(b *types.ItemType) error {
	ctx := context.Background()
	k := datastore.NewKey(ctx, "ItemType", "", b.ID, nil)
	if _, err := c.DatastoreClient().Put(ctx, k, b); err != nil {
		return fmt.Errorf("datastoredb: could not update ItemType: %v", err)
	}
	return nil
}

// ListItemTypes returns a list of ItemTypes
func (c *CloudPersister) ListItemTypes() ([]*types.ItemType, error) {
	ctx := context.Background()
	ItemTypes := make([]*types.ItemType, 0)
	q := datastore.NewQuery("ItemType")

	keys, err := c.DatastoreClient().GetAll(ctx, q, &ItemTypes)

	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not list ItemTypes: %v", err)
	}

	for i, k := range keys {
		ItemTypes[i].ID = k.ID()
	}

	return ItemTypes, nil
}

//  QueryItemTypesByProp
func (c *CloudPersister) QueryItemTypesByProp(propName, value string) (*types.ItemType, error) {
	ctx := context.Background()
	ItemTypes := make([]*types.ItemType, 0)
	q := datastore.NewQuery("ItemType").Filter(fmt.Sprintf("%s =", propName), value)

	keys, err := c.DatastoreClient().GetAll(ctx, q, &ItemTypes)

	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not list ItemTypes: %v", err)
	}

	if len(ItemTypes) == 0 {
		return nil, nil
	}

	ItemTypes[0].ID = keys[0].ID()
	return ItemTypes[0], nil
}
