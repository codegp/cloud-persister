package cloudpersister

import (
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/codegp/game-object-types/types"

	"golang.org/x/net/context"
)

// GetTerrainType retrieves a TerrainType by its ID.
func (c *CloudPersister) GetTerrainType(id int64) (*types.TerrainType, error) {
	ctx := context.Background()
	k := datastore.NewKey(ctx, "TerrainType", "", id, nil)
	TerrainType := &types.TerrainType{}
	if err := c.DatastoreClient().Get(ctx, k, TerrainType); err != nil {
		return nil, fmt.Errorf("datastoredb: could not get TerrainType: %v", err)
	}
	TerrainType.ID = id
	return TerrainType, nil
}

// AddTerrainType saves a given TerrainType, assigning it a new ID.
func (c *CloudPersister) AddTerrainType(b *types.TerrainType) (*types.TerrainType, error) {
	ctx := context.Background()
	k := datastore.NewIncompleteKey(ctx, "TerrainType", nil)
	k, err := c.DatastoreClient().Put(ctx, k, b)
	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not put TerrainType: %v", err)
	}
	b.ID = k.ID()
	return b, nil
}

// DeleteTerrainType removes a given TerrainType by its ID.
func (c *CloudPersister) DeleteTerrainType(id int64) error {
	ctx := context.Background()
	k := datastore.NewKey(ctx, "TerrainType", "", id, nil)
	if err := c.DatastoreClient().Delete(ctx, k); err != nil {
		return fmt.Errorf("datastoredb: could not delete TerrainType: %v", err)
	}
	return nil
}

// UpdateTerrainType updates the entry for a given TerrainType.
func (c *CloudPersister) UpdateTerrainType(b *types.TerrainType) error {
	ctx := context.Background()
	k := datastore.NewKey(ctx, "TerrainType", "", b.ID, nil)
	if _, err := c.DatastoreClient().Put(ctx, k, b); err != nil {
		return fmt.Errorf("datastoredb: could not update TerrainType: %v", err)
	}
	return nil
}

// ListTerrainTypes returns a list of TerrainTypes
func (c *CloudPersister) ListTerrainTypes() ([]*types.TerrainType, error) {
	ctx := context.Background()
	TerrainTypes := make([]*types.TerrainType, 0)
	q := datastore.NewQuery("TerrainType")

	keys, err := c.DatastoreClient().GetAll(ctx, q, &TerrainTypes)

	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not list TerrainTypes: %v", err)
	}

	for i, k := range keys {
		TerrainTypes[i].ID = k.ID()
	}

	return TerrainTypes, nil
}

//  QueryTerrainTypesByProp
func (c *CloudPersister) QueryTerrainTypesByProp(propName, value string) (*types.TerrainType, error) {
	ctx := context.Background()
	TerrainTypes := make([]*types.TerrainType, 0)
	q := datastore.NewQuery("TerrainType").Filter(fmt.Sprintf("%s =", propName), value)

	keys, err := c.DatastoreClient().GetAll(ctx, q, &TerrainTypes)

	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not list TerrainTypes: %v", err)
	}

	if len(TerrainTypes) == 0 {
		return nil, nil
	}

	TerrainTypes[0].ID = keys[0].ID()
	return TerrainTypes[0], nil
}
