package cloudpersister

import (
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/codegp/cloud-persister/models"

	"golang.org/x/net/context"
)

// GetGameType retrieves a GameType by its ID.
func (c *CloudPersister) GetGameType(id int64) (*models.GameType, error) {
	ctx := context.Background()
	k := datastore.NewKey(ctx, "GameType", "", id, nil)
	GameType := &models.GameType{}
	if err := c.DatastoreClient().Get(ctx, k, GameType); err != nil {
		return nil, fmt.Errorf("datastoredb: could not get GameType: %v", err)
	}
	GameType.ID = id
	return GameType, nil
}

// AddGameType saves a given GameType, assigning it a new ID.
func (c *CloudPersister) AddGameType(b *models.GameType) (*models.GameType, error) {
	ctx := context.Background()
	k := datastore.NewIncompleteKey(ctx, "GameType", nil)
	k, err := c.DatastoreClient().Put(ctx, k, b)
	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not put GameType: %v", err)
	}
	b.ID = k.ID()
	return b, nil
}

// DeleteGameType removes a given GameType by its ID.
func (c *CloudPersister) DeleteGameType(id int64) error {
	ctx := context.Background()
	k := datastore.NewKey(ctx, "GameType", "", id, nil)
	if err := c.DatastoreClient().Delete(ctx, k); err != nil {
		return fmt.Errorf("datastoredb: could not delete GameType: %v", err)
	}
	return nil
}

// UpdateGameType updates the entry for a given GameType.
func (c *CloudPersister) UpdateGameType(b *models.GameType) error {
	ctx := context.Background()
	k := datastore.NewKey(ctx, "GameType", "", b.ID, nil)
	if _, err := c.DatastoreClient().Put(ctx, k, b); err != nil {
		return fmt.Errorf("datastoredb: could not update GameType: %v", err)
	}
	return nil
}

// ListGameTypes returns a list of GameTypes
func (c *CloudPersister) ListGameTypes() ([]*models.GameType, error) {
	ctx := context.Background()
	GameTypes := make([]*models.GameType, 0)
	q := datastore.NewQuery("GameType")

	keys, err := c.DatastoreClient().GetAll(ctx, q, &GameTypes)

	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not list GameTypes: %v", err)
	}

	for i, k := range keys {
		GameTypes[i].ID = k.ID()
	}

	return GameTypes, nil
}

//  QueryGameTypesByProp
func (c *CloudPersister) QueryGameTypesByProp(propName, value string) (*models.GameType, error) {
	ctx := context.Background()
	GameTypes := make([]*models.GameType, 0)
	q := datastore.NewQuery("GameType").Filter(fmt.Sprintf("%s =", propName), value)

	keys, err := c.DatastoreClient().GetAll(ctx, q, &GameTypes)

	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not list GameTypes: %v", err)
	}

	if len(GameTypes) == 0 {
		return nil, nil
	}

	GameTypes[0].ID = keys[0].ID()
	return GameTypes[0], nil
}
