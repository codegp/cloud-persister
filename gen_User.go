package cloudpersister

import (
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/codegp/cloud-persister/models"

	"golang.org/x/net/context"
)

// GetUser retrieves a User by its ID.
func (c *CloudPersister) GetUser(id int64) (*models.User, error) {
	ctx := context.Background()
	k := datastore.NewKey(ctx, "User", "", id, nil)
	User := &models.User{}
	if err := c.DatastoreClient().Get(ctx, k, User); err != nil {
		return nil, fmt.Errorf("datastoredb: could not get User: %v", err)
	}
	User.ID = id
	return User, nil
}

// AddUser saves a given User, assigning it a new ID.
func (c *CloudPersister) AddUser(b *models.User) (*models.User, error) {
	ctx := context.Background()
	k := datastore.NewIncompleteKey(ctx, "User", nil)
	k, err := c.DatastoreClient().Put(ctx, k, b)
	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not put User: %v", err)
	}
	b.ID = k.ID()
	return b, nil
}

// DeleteUser removes a given User by its ID.
func (c *CloudPersister) DeleteUser(id int64) error {
	ctx := context.Background()
	k := datastore.NewKey(ctx, "User", "", id, nil)
	if err := c.DatastoreClient().Delete(ctx, k); err != nil {
		return fmt.Errorf("datastoredb: could not delete User: %v", err)
	}
	return nil
}

// UpdateUser updates the entry for a given User.
func (c *CloudPersister) UpdateUser(b *models.User) error {
	ctx := context.Background()
	k := datastore.NewKey(ctx, "User", "", b.ID, nil)
	if _, err := c.DatastoreClient().Put(ctx, k, b); err != nil {
		return fmt.Errorf("datastoredb: could not update User: %v", err)
	}
	return nil
}

// ListUsers returns a list of Users
func (c *CloudPersister) ListUsers() ([]*models.User, error) {
	ctx := context.Background()
	Users := make([]*models.User, 0)
	q := datastore.NewQuery("User")

	keys, err := c.DatastoreClient().GetAll(ctx, q, &Users)

	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not list Users: %v", err)
	}

	for i, k := range keys {
		Users[i].ID = k.ID()
	}

	return Users, nil
}

//  QueryUsersByProp
func (c *CloudPersister) QueryUsersByProp(propName, value string) (*models.User, error) {
	ctx := context.Background()
	Users := make([]*models.User, 0)
	q := datastore.NewQuery("User").Filter(fmt.Sprintf("%s =", propName), value)

	keys, err := c.DatastoreClient().GetAll(ctx, q, &Users)

	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not list Users: %v", err)
	}

	if len(Users) == 0 {
		return nil, nil
	}

	Users[0].ID = keys[0].ID()
	return Users[0], nil
}
