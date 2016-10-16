# cloud-persister
A library that houses a client to perform crud operations on datastore entities and perform r/w to and from cloud storage

### generation
This library uses (google-cloud-go-transaction-generator)[https://github.com/codegp/google-cloud-go-transaction-generator] to generate code
to make datastore transactions for all of the codegp models. To generate the files:

```
go install github.com/codegp/google-cloud-go-transaction-generator
google-cloud-go-transaction-generator generatorConfig.yaml
```

Files need to be regenerated anytime [game object types](https://github.com/codegp/game-object-types) is updated
or the content of the models directory is updated.
