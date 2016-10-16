package cloudpersister

import (
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/codegp/cloud-persister/models"

	"golang.org/x/net/context"
)

// GetGame retrieves a Game by its ID.
func (c *CloudPersister) GetGame(id int64) (*models.Game, error) {
	ctx := context.Background()
	k := datastore.NewKey(ctx, "Game", "", id, nil)
	Game := &models.Game{}
	if err := c.DatastoreClient().Get(ctx, k, Game); err != nil {
		return nil, fmt.Errorf("datastoredb: could not get Game: %v", err)
	}
	Game.ID = id
	return Game, nil
}

// AddGame saves a given Game, assigning it a new ID.
func (c *CloudPersister) AddGame(b *models.Game) (*models.Game, error) {
	ctx := context.Background()
	k := datastore.NewIncompleteKey(ctx, "Game", nil)
	k, err := c.DatastoreClient().Put(ctx, k, b)
	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not put Game: %v", err)
	}
	b.ID = k.ID()
	return b, nil
}

// DeleteGame removes a given Game by its ID.
func (c *CloudPersister) DeleteGame(id int64) error {
	ctx := context.Background()
	k := datastore.NewKey(ctx, "Game", "", id, nil)
	if err := c.DatastoreClient().Delete(ctx, k); err != nil {
		return fmt.Errorf("datastoredb: could not delete Game: %v", err)
	}
	return nil
}

// UpdateGame updates the entry for a given Game.
func (c *CloudPersister) UpdateGame(b *models.Game) error {
	ctx := context.Background()
	k := datastore.NewKey(ctx, "Game", "", b.ID, nil)
	if _, err := c.DatastoreClient().Put(ctx, k, b); err != nil {
		return fmt.Errorf("datastoredb: could not update Game: %v", err)
	}
	return nil
}

// ListGames returns a list of Games
func (c *CloudPersister) ListGames() ([]*models.Game, error) {
	ctx := context.Background()
	Games := make([]*models.Game, 0)
	q := datastore.NewQuery("Game")

	keys, err := c.DatastoreClient().GetAll(ctx, q, &Games)

	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not list Games: %v", err)
	}

	for i, k := range keys {
		Games[i].ID = k.ID()
	}

	return Games, nil
}

//  QueryGamesByProp
func (c *CloudPersister) QueryGamesByProp(propName, value string) (*models.Game, error) {
	ctx := context.Background()
	Games := make([]*models.Game, 0)
	q := datastore.NewQuery("Game").Filter(fmt.Sprintf("%s =", propName), value)

	keys, err := c.DatastoreClient().GetAll(ctx, q, &Games)

	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not list Games: %v", err)
	}

	if len(Games) == 0 {
		return nil, nil
	}

	Games[0].ID = keys[0].ID()
	return Games[0], nil
}
