package cloudpersister

import (
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/codegp/cloud-persister/models"

	"golang.org/x/net/context"
)

// GetProject retrieves a Project by its ID.
func (c *CloudPersister) GetProject(id int64) (*models.Project, error) {
	ctx := context.Background()
	k := datastore.NewKey(ctx, "Project", "", id, nil)
	Project := &models.Project{}
	if err := c.DatastoreClient().Get(ctx, k, Project); err != nil {
		return nil, fmt.Errorf("datastoredb: could not get Project: %v", err)
	}
	Project.ID = id
	return Project, nil
}

// GetMultiProject retrieves a list of Projects by their ID.
func (c *CloudPersister) GetMultiProject(ids []int64) ([]*models.Project, error) {
	if len(ids) == 0 {
		return []*models.Project{}, nil
	}
	ctx := context.Background()
	ks := make([]*datastore.Key, len(ids))
	Projects := make([]*models.Project, len(ids))
	for i, id := range ids {
		ks[i] = datastore.NewKey(ctx, "Project", "", id, nil)
		Projects[i] = &models.Project{}
	}
	if err := c.DatastoreClient().GetMulti(ctx, ks, Projects); err != nil {
		return nil, fmt.Errorf("datastoredb: could not get Projects: %v", err)
	}
	for i, id := range ids {
		Projects[i].ID = id
	}
	return Projects, nil
}

// AddProject saves a given Project, assigning it a new ID.
func (c *CloudPersister) AddProject(b *models.Project) (*models.Project, error) {
	ctx := context.Background()
	k := datastore.NewIncompleteKey(ctx, "Project", nil)
	k, err := c.DatastoreClient().Put(ctx, k, b)
	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not put Project: %v", err)
	}
	b.ID = k.ID()
	return b, nil
}

// DeleteProject removes a given Project by its ID.
func (c *CloudPersister) DeleteProject(id int64) error {
	ctx := context.Background()
	k := datastore.NewKey(ctx, "Project", "", id, nil)
	if err := c.DatastoreClient().Delete(ctx, k); err != nil {
		return fmt.Errorf("datastoredb: could not delete Project: %v", err)
	}
	return nil
}

// UpdateProject updates the entry for a given Project.
func (c *CloudPersister) UpdateProject(b *models.Project) error {
	ctx := context.Background()
	k := datastore.NewKey(ctx, "Project", "", b.ID, nil)
	if _, err := c.DatastoreClient().Put(ctx, k, b); err != nil {
		return fmt.Errorf("datastoredb: could not update Project: %v", err)
	}
	return nil
}

// ListProjects returns a list of Projects
func (c *CloudPersister) ListProjects() ([]*models.Project, error) {
	ctx := context.Background()
	Projects := make([]*models.Project, 0)
	q := datastore.NewQuery("Project")

	keys, err := c.DatastoreClient().GetAll(ctx, q, &Projects)

	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not list Projects: %v", err)
	}

	for i, k := range keys {
		Projects[i].ID = k.ID()
	}

	return Projects, nil
}

//  QueryProjectsByProp
func (c *CloudPersister) QueryProjectsByProp(propName, value string) (*models.Project, error) {
	ctx := context.Background()
	Projects := make([]*models.Project, 0)
	q := datastore.NewQuery("Project").Filter(fmt.Sprintf("%s =", propName), value)

	keys, err := c.DatastoreClient().GetAll(ctx, q, &Projects)

	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not list Projects: %v", err)
	}

	if len(Projects) == 0 {
		return nil, nil
	}

	Projects[0].ID = keys[0].ID()
	return Projects[0], nil
}
