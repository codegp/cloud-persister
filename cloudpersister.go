package cloudpersister

import (
	"fmt"

	"github.com/codegp/env"

	"cloud.google.com/go/datastore"
)

// CloudPersister is an object that can read/update cloud storage and datastore models
type CloudPersister struct {
	ds *datastore.Client
	fp FilePersister
}

// NewCloudPersister returns an instance of CloudPersister
func NewCloudPersister() (*CloudPersister, error) {
	ds, err := configureDatastore()
	if err != nil {
		return nil, err
	}

	fp, err := NewFilePersister()
	if err != nil {
		return nil, err
	}

	return &CloudPersister{
		ds: ds,
		fp: fp,
	}, nil
}

// DatastoreClient returns the cloud persisters instance of "cloud.google.com/go/datastore".Client
func (c *CloudPersister) DatastoreClient() *datastore.Client {
	return c.ds
}

func (c *CloudPersister) WriteMap(id int64, content []byte) error {
	name := fmt.Sprintf("map-%d.json", id)
	return c.fp.Write(name, content, "application/json", false)
}

func (c *CloudPersister) ReadMap(id int64) ([]byte, error) {
	name := fmt.Sprintf("map-%d.json", id)
	return c.fp.Read(name)
}

func (c *CloudPersister) WriteHistory(id int64, content []byte) error {
	name := fmt.Sprintf("history-%d.json", id)
	return c.fp.Write(name, content, "application/json", false)
}

func (c *CloudPersister) ReadHistory(id int64) ([]byte, error) {
	name := fmt.Sprintf("history-%d.json", id)
	return c.fp.Read(name)
}

func (c *CloudPersister) WriteGameTypeCode(id int64, content []byte) error {
	name := fmt.Sprintf("gametype-%d.go", id)
	return c.fp.Write(name, content, "text/plain", false)
}

func (c *CloudPersister) ReadGameTypeCode(id int64) ([]byte, error) {
	name := fmt.Sprintf("gametype-%d.go", id)
	return c.fp.Read(name)
}

func (c *CloudPersister) WriteProjectFile(id int64, filename string, content []byte) error {
	name := fmt.Sprintf("project-%d-%s", id, filename)
	return c.fp.Write(name, content, "text/plain", false)
}

func (c *CloudPersister) ReadProjectFile(id int64, filename string) ([]byte, error) {
	name := fmt.Sprintf("project-%d-%s", id, filename)
	return c.fp.Read(name)
}

func (c *CloudPersister) WriteIcon(id int64, content []byte) error {
	name := fmt.Sprintf("icon-%d.png", id)
	return c.fp.Write(name, content, "image/png", true)
}

func (c *CloudPersister) ReadIcon(id int64) ([]byte, error) {
	name := fmt.Sprintf("icon-%d.png", id)
	return c.fp.Read(name)
}

func (c *CloudPersister) WriteDocs(fname string, content []byte) error {
	name := fmt.Sprintf("%s.html", fname)
	return c.fp.Write(name, content, "text/html", true)
}

func (c *CloudPersister) ReadDocs(fname string) ([]byte, error) {
	name := fmt.Sprintf("%s.html", fname)
	return c.fp.Read(name)
}

func (c *CloudPersister) WriteGenCode(gameTypeID int64, lang env.Lang, content []byte) error {
	name := fmt.Sprintf("%d-%s.zip", gameTypeID, lang)
	return c.fp.Write(name, content, "application/zip", true)
}

func (c *CloudPersister) ReadGenCode(gameTypeID int64, lang env.Lang) ([]byte, error) {
	name := fmt.Sprintf("%d-%s.zip", gameTypeID, lang)
	return c.fp.Read(name)
}
