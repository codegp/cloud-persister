package cloudpersister

import (
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/codegp/game-object-types/types"

	"golang.org/x/net/context"
)

// GetMoveType retrieves a MoveType by its ID.
func (c *CloudPersister) GetMoveType(id int64) (*types.MoveType, error) {
	ctx := context.Background()
	k := datastore.NewKey(ctx, "MoveType", "", id, nil)
	MoveType := &types.MoveType{}
	if err := c.DatastoreClient().Get(ctx, k, MoveType); err != nil {
		return nil, fmt.Errorf("datastoredb: could not get MoveType: %v", err)
	}
	MoveType.ID = id
	return MoveType, nil
}

// GetMultiMoveType retrieves a list of MoveTypes by their ID.
func (c *CloudPersister) GetMultiMoveType(ids []int64) ([]*types.MoveType, error) {
	if len(ids) == 0 {
		return []*types.MoveType{}, nil
	}
	ctx := context.Background()
	ks := make([]*datastore.Key, len(ids))
	MoveTypes := make([]*types.MoveType, len(ids))
	for i, id := range ids {
		ks[i] = datastore.NewKey(ctx, "MoveType", "", id, nil)
		MoveTypes[i] = &types.MoveType{}
	}
	if err := c.DatastoreClient().GetMulti(ctx, ks, MoveTypes); err != nil {
		return nil, fmt.Errorf("datastoredb: could not get MoveTypes: %v", err)
	}
	for i, id := range ids {
		MoveTypes[i].ID = id
	}
	return MoveTypes, nil
}

// AddMoveType saves a given MoveType, assigning it a new ID.
func (c *CloudPersister) AddMoveType(b *types.MoveType) (*types.MoveType, error) {
	ctx := context.Background()
	k := datastore.NewIncompleteKey(ctx, "MoveType", nil)
	k, err := c.DatastoreClient().Put(ctx, k, b)
	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not put MoveType: %v", err)
	}
	b.ID = k.ID()
	return b, nil
}

// DeleteMoveType removes a given MoveType by its ID.
func (c *CloudPersister) DeleteMoveType(id int64) error {
	ctx := context.Background()
	k := datastore.NewKey(ctx, "MoveType", "", id, nil)
	if err := c.DatastoreClient().Delete(ctx, k); err != nil {
		return fmt.Errorf("datastoredb: could not delete MoveType: %v", err)
	}
	return nil
}

// UpdateMoveType updates the entry for a given MoveType.
func (c *CloudPersister) UpdateMoveType(b *types.MoveType) error {
	ctx := context.Background()
	k := datastore.NewKey(ctx, "MoveType", "", b.ID, nil)
	if _, err := c.DatastoreClient().Put(ctx, k, b); err != nil {
		return fmt.Errorf("datastoredb: could not update MoveType: %v", err)
	}
	return nil
}

// ListMoveTypes returns a list of MoveTypes
func (c *CloudPersister) ListMoveTypes() ([]*types.MoveType, error) {
	ctx := context.Background()
	MoveTypes := make([]*types.MoveType, 0)
	q := datastore.NewQuery("MoveType")

	keys, err := c.DatastoreClient().GetAll(ctx, q, &MoveTypes)

	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not list MoveTypes: %v", err)
	}

	for i, k := range keys {
		MoveTypes[i].ID = k.ID()
	}

	return MoveTypes, nil
}

//  QueryMoveTypesByProp
func (c *CloudPersister) QueryMoveTypesByProp(propName, value string) (*types.MoveType, error) {
	ctx := context.Background()
	MoveTypes := make([]*types.MoveType, 0)
	q := datastore.NewQuery("MoveType").Filter(fmt.Sprintf("%s =", propName), value)

	keys, err := c.DatastoreClient().GetAll(ctx, q, &MoveTypes)

	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not list MoveTypes: %v", err)
	}

	if len(MoveTypes) == 0 {
		return nil, nil
	}

	MoveTypes[0].ID = keys[0].ID()
	return MoveTypes[0], nil
}
