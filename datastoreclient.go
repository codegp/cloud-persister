package cloudpersister

import (
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/codegp/env"
	"golang.org/x/net/context"
)

// DSClient is a struct that implements methods performing datastore ops
// on codegp specific models
type DSClient struct {
	client *datastore.Client
}

// NewDatastoreClient is a public function to return a new instance
// of a DSClient. returns an error if we fail to connecto to the cloud
// datastore for any reason
func NewDatastoreClient() (*DSClient, error) {
	client, err := configureDatastore()
	if err != nil {
		return nil, err
	}

	return &DSClient{
		client: client,
	}, nil
}

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
