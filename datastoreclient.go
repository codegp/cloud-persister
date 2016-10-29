package cloudpersister

import (
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/codegp/env"
	"golang.org/x/net/context"
)

func getProjectID() (string, error) {
	projID := env.GCloudProjectID()
	if projID == "" {
		return projID, fmt.Errorf("No project id found in the environment")
	}

	return projID, nil
}

func configureDatastore() (*datastore.Client, error) {
	projID, err := getProjectID()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	client, err := datastore.NewClient(ctx, projID)
	if err != nil {
		return nil, err
	}

	// test the client, error if we cant start a new transaction
	t, err := client.NewTransaction(ctx)
	if err != nil {
		return nil, fmt.Errorf("Failed to start test transaction: %v", err)
	}
	if err := t.Rollback(); err != nil {
		return nil, fmt.Errorf("Failed to rollback test transaction: %v", err)
	}

	return client, nil
}
