package cloudpersister

import (
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/codegp/game-object-types/types"

	"golang.org/x/net/context"
)

// GetBotType retrieves a BotType by its ID.
func (c *CloudPersister) GetBotType(id int64) (*types.BotType, error) {
	ctx := context.Background()
	k := datastore.NewKey(ctx, "BotType", "", id, nil)
	BotType := &types.BotType{}
	if err := c.DatastoreClient().Get(ctx, k, BotType); err != nil {
		return nil, fmt.Errorf("datastoredb: could not get BotType: %v", err)
	}
	BotType.ID = id
	return BotType, nil
}

// AddBotType saves a given BotType, assigning it a new ID.
func (c *CloudPersister) AddBotType(b *types.BotType) (*types.BotType, error) {
	ctx := context.Background()
	k := datastore.NewIncompleteKey(ctx, "BotType", nil)
	k, err := c.DatastoreClient().Put(ctx, k, b)
	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not put BotType: %v", err)
	}
	b.ID = k.ID()
	return b, nil
}

// DeleteBotType removes a given BotType by its ID.
func (c *CloudPersister) DeleteBotType(id int64) error {
	ctx := context.Background()
	k := datastore.NewKey(ctx, "BotType", "", id, nil)
	if err := c.DatastoreClient().Delete(ctx, k); err != nil {
		return fmt.Errorf("datastoredb: could not delete BotType: %v", err)
	}
	return nil
}

// UpdateBotType updates the entry for a given BotType.
func (c *CloudPersister) UpdateBotType(b *types.BotType) error {
	ctx := context.Background()
	k := datastore.NewKey(ctx, "BotType", "", b.ID, nil)
	if _, err := c.DatastoreClient().Put(ctx, k, b); err != nil {
		return fmt.Errorf("datastoredb: could not update BotType: %v", err)
	}
	return nil
}

// ListBotTypes returns a list of BotTypes
func (c *CloudPersister) ListBotTypes() ([]*types.BotType, error) {
	ctx := context.Background()
	BotTypes := make([]*types.BotType, 0)
	q := datastore.NewQuery("BotType")

	keys, err := c.DatastoreClient().GetAll(ctx, q, &BotTypes)

	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not list BotTypes: %v", err)
	}

	for i, k := range keys {
		BotTypes[i].ID = k.ID()
	}

	return BotTypes, nil
}

//  QueryBotTypesByProp
func (c *CloudPersister) QueryBotTypesByProp(propName, value string) (*types.BotType, error) {
	ctx := context.Background()
	BotTypes := make([]*types.BotType, 0)
	q := datastore.NewQuery("BotType").Filter(fmt.Sprintf("%s =", propName), value)

	keys, err := c.DatastoreClient().GetAll(ctx, q, &BotTypes)

	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not list BotTypes: %v", err)
	}

	if len(BotTypes) == 0 {
		return nil, nil
	}

	BotTypes[0].ID = keys[0].ID()
	return BotTypes[0], nil
}
