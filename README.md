# gcp-data-skeleton/api

## Tasks

### build

```
go build ./...
```

### package

Package needs to `go mod vendor` any private modules.

This comment excludes the terraform and dataflow directories.

```
zip -r source.zip * -x "terraform/*" -x "dataflow/*" -x ".git"
```

### deploy-dev

dir: ./terraform

```
terraform apply -var-file="dev.tfvars"
```

### create-dataflow-docker-container

https://cloud.google.com/dataflow/docs/guides/templates/using-flex-templates

dir: ./dataflow
inputs: REGION, PROJECT_ID, REPO

```
gcloud builds submit --tag ${REGION}-docker.pkg.dev/${PROJECT_ID}/${REPO}/dataflow/wordcount:latest .
```

### create-dataflow-job-spec

https://cloud.google.com/dataflow/docs/guides/templates/using-flex-templates#create_a_flex_template

dir: ./dataflow
inputs: BUCKET, REGION, PROJECT_ID, REPO

```
gcloud dataflow flex-template build gs://${BUCKET}/dataflow/wordcount.json \
     --image "${REGION}-docker.pkg.dev/${PROJECT_ID}/${REPO}/dataflow/wordcount:latest" \
     --sdk-language "PYTHON" \
     --metadata-file "metadata.json"
```
