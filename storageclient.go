package cloudpersister

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/codegp/env"
	"cloud.google.com/go/storage"
	"golang.org/x/net/context"
)

const buketname = "codegp"

type FilePersister interface {
	Read(string) ([]byte, error)
	Write(string, []byte, string, bool) error
}

func NewFilePersister() (FilePersister, error) {
	if env.IsLocal() {
		return NewFSPersister(), nil
	}

	return NewCloudStoragePersister()
}

type FSPersister struct{}

func NewFSPersister() *FSPersister {
	return &FSPersister{}
}

func (fs *FSPersister) Read(name string) ([]byte, error) {
	return ioutil.ReadFile(fmt.Sprintf("/localstore/%s", name))
}
func (fs *FSPersister) Write(name string, content []byte, _ string, _ bool) error {
	return ioutil.WriteFile(fmt.Sprintf("/localstore/%s", name), content, os.ModePerm)
}

type CloudStoragePersister struct {
	bucketHandle *storage.BucketHandle
}

func NewCloudStoragePersister() (*CloudStoragePersister, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	bucketHandle := client.Bucket(buketname)
	return &CloudStoragePersister{
		bucketHandle: bucketHandle,
	}, nil
}

func (p *CloudStoragePersister) Read(name string) ([]byte, error) {
	ctx := context.Background()
	reader, err := p.bucketHandle.Object(name).NewReader(ctx)
	if err != nil {
		return []byte{}, err
	}

	result := make([]byte, reader.Size())
	_, err = reader.Read(result)
	if err != nil {
		return []byte{}, err
	}
	return result, reader.Close()
}

func (p *CloudStoragePersister) Write(name string, content []byte, mimeType string, public bool) error {
	ctx := context.Background()
	writer := p.bucketHandle.Object(name).NewWriter(ctx)
	writer.ContentType = mimeType

	aclEnt := storage.AllAuthenticatedUsers
	if public {
		aclEnt = storage.AllUsers
	}

	writer.ACL = []storage.ACLRule{{aclEnt, storage.RoleReader}}
	_, err := writer.Write(content)
	if err != nil {
		return err
	}
	return writer.Close()
}
