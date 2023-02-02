# gcp-data-skeleton/api

## Tasks

### build

dir: ./function

```
go build ./...
```

### package

Package needs to `go mod vendor` any private modules.

dir: ./function

```
zip -r function-source.zip *
```

### deploy-dev

dir: ./terraform

```
terraform apply -var-file="dev.tfvars"
```
