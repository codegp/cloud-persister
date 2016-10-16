package cloudpersister

import (
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/codegp/cloud-persister/models"

	"golang.org/x/net/context"
)

// GetMap retrieves a Map by its ID.
func (c *CloudPersister) GetMap(id int64) (*models.Map, error) {
	ctx := context.Background()
	k := datastore.NewKey(ctx, "Map", "", id, nil)
	Map := &models.Map{}
	if err := c.DatastoreClient().Get(ctx, k, Map); err != nil {
		return nil, fmt.Errorf("datastoredb: could not get Map: %v", err)
	}
	Map.ID = id
	return Map, nil
}

// AddMap saves a given Map, assigning it a new ID.
func (c *CloudPersister) AddMap(b *models.Map) (*models.Map, error) {
	ctx := context.Background()
	k := datastore.NewIncompleteKey(ctx, "Map", nil)
	k, err := c.DatastoreClient().Put(ctx, k, b)
	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not put Map: %v", err)
	}
	b.ID = k.ID()
	return b, nil
}

// DeleteMap removes a given Map by its ID.
func (c *CloudPersister) DeleteMap(id int64) error {
	ctx := context.Background()
	k := datastore.NewKey(ctx, "Map", "", id, nil)
	if err := c.DatastoreClient().Delete(ctx, k); err != nil {
		return fmt.Errorf("datastoredb: could not delete Map: %v", err)
	}
	return nil
}

// UpdateMap updates the entry for a given Map.
func (c *CloudPersister) UpdateMap(b *models.Map) error {
	ctx := context.Background()
	k := datastore.NewKey(ctx, "Map", "", b.ID, nil)
	if _, err := c.DatastoreClient().Put(ctx, k, b); err != nil {
		return fmt.Errorf("datastoredb: could not update Map: %v", err)
	}
	return nil
}

// ListMaps returns a list of Maps
func (c *CloudPersister) ListMaps() ([]*models.Map, error) {
	ctx := context.Background()
	Maps := make([]*models.Map, 0)
	q := datastore.NewQuery("Map")

	keys, err := c.DatastoreClient().GetAll(ctx, q, &Maps)

	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not list Maps: %v", err)
	}

	for i, k := range keys {
		Maps[i].ID = k.ID()
	}

	return Maps, nil
}

//  QueryMapsByProp
func (c *CloudPersister) QueryMapsByProp(propName, value string) (*models.Map, error) {
	ctx := context.Background()
	Maps := make([]*models.Map, 0)
	q := datastore.NewQuery("Map").Filter(fmt.Sprintf("%s =", propName), value)

	keys, err := c.DatastoreClient().GetAll(ctx, q, &Maps)

	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not list Maps: %v", err)
	}

	if len(Maps) == 0 {
		return nil, nil
	}

	Maps[0].ID = keys[0].ID()
	return Maps[0], nil
}
