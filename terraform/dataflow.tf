resource "google_artifact_registry_repository" "dataflow_repo" {
  location      = var.zone
  repository_id = "dataflow"
  format        = "DOCKER"
}

resource "google_storage_bucket_object" "wordcount_template" {
  name   = "/dataflow/wordcount.json"
  bucket = google_storage_bucket.bucket.name
  source = "../dataflow/wordcount.json"
}

# https://cloud.google.com/dataflow/docs/guides/templates/using-flex-templates#metadata
resource "google_dataflow_flex_template_job" "wordcount_job" {
  provider                = google-beta
  name                    = "wordcount"
  container_spec_gcs_path = "gs://${google_storage_bucket.bucket.name}/dataflow/wordcount.json"
  region                  = "europe-west1" # Belgium
}

output "dataflow_repo" {
  value = google_artifact_registry_repository.dataflow_repo.name
}

