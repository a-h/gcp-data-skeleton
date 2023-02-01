# gcp-data-skeleton/api

## Tasks

### build

dir: ./helloGet

```
npm i
```

### package

dir: ./helloGet

```
zip function-source.zip ./**/*
```

### deploy-dev

dir: ./terraform

```
terraform apply -var-file="dev.tfvars"
```
