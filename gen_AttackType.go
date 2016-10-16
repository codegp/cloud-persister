package cloudpersister

import (
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/codegp/game-object-types/types"

	"golang.org/x/net/context"
)

// GetAttackType retrieves a AttackType by its ID.
func (c *CloudPersister) GetAttackType(id int64) (*types.AttackType, error) {
	ctx := context.Background()
	k := datastore.NewKey(ctx, "AttackType", "", id, nil)
	AttackType := &types.AttackType{}
	if err := c.DatastoreClient().Get(ctx, k, AttackType); err != nil {
		return nil, fmt.Errorf("datastoredb: could not get AttackType: %v", err)
	}
	AttackType.ID = id
	return AttackType, nil
}

// AddAttackType saves a given AttackType, assigning it a new ID.
func (c *CloudPersister) AddAttackType(b *types.AttackType) (*types.AttackType, error) {
	ctx := context.Background()
	k := datastore.NewIncompleteKey(ctx, "AttackType", nil)
	k, err := c.DatastoreClient().Put(ctx, k, b)
	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not put AttackType: %v", err)
	}
	b.ID = k.ID()
	return b, nil
}

// DeleteAttackType removes a given AttackType by its ID.
func (c *CloudPersister) DeleteAttackType(id int64) error {
	ctx := context.Background()
	k := datastore.NewKey(ctx, "AttackType", "", id, nil)
	if err := c.DatastoreClient().Delete(ctx, k); err != nil {
		return fmt.Errorf("datastoredb: could not delete AttackType: %v", err)
	}
	return nil
}

// UpdateAttackType updates the entry for a given AttackType.
func (c *CloudPersister) UpdateAttackType(b *types.AttackType) error {
	ctx := context.Background()
	k := datastore.NewKey(ctx, "AttackType", "", b.ID, nil)
	if _, err := c.DatastoreClient().Put(ctx, k, b); err != nil {
		return fmt.Errorf("datastoredb: could not update AttackType: %v", err)
	}
	return nil
}

// ListAttackTypes returns a list of AttackTypes
func (c *CloudPersister) ListAttackTypes() ([]*types.AttackType, error) {
	ctx := context.Background()
	AttackTypes := make([]*types.AttackType, 0)
	q := datastore.NewQuery("AttackType")

	keys, err := c.DatastoreClient().GetAll(ctx, q, &AttackTypes)

	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not list AttackTypes: %v", err)
	}

	for i, k := range keys {
		AttackTypes[i].ID = k.ID()
	}

	return AttackTypes, nil
}

//  QueryAttackTypesByProp
func (c *CloudPersister) QueryAttackTypesByProp(propName, value string) (*types.AttackType, error) {
	ctx := context.Background()
	AttackTypes := make([]*types.AttackType, 0)
	q := datastore.NewQuery("AttackType").Filter(fmt.Sprintf("%s =", propName), value)

	keys, err := c.DatastoreClient().GetAll(ctx, q, &AttackTypes)

	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not list AttackTypes: %v", err)
	}

	if len(AttackTypes) == 0 {
		return nil, nil
	}

	AttackTypes[0].ID = keys[0].ID()
	return AttackTypes[0], nil
}
