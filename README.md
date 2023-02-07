# gcp-data-skeleton/api

## Tasks

### build

```
go build ./...
```

### package

Package needs to `go mod vendor` any private modules.

```
zip -r source.zip * -x "terraform/*" -x ".git"
```

### deploy-dev

dir: ./terraform

```
terraform apply -var-file="dev.tfvars"
```
