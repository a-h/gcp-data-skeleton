resource "google_pubsub_subscription" "to_bigquery" {
  name  = "to_bigquery"
  topic = google_pubsub_topic.samples.name

  bigquery_config {
    table = "${google_bigquery_table.samples.project}:${google_bigquery_table.samples.dataset_id}.${google_bigquery_table.samples.table_id}"
  }

  depends_on = [google_project_iam_member.viewer, google_project_iam_member.editor]
}

data "google_project" "project" {
}

resource "google_project_iam_member" "viewer" {
  project = data.google_project.project.project_id
  role   = "roles/bigquery.metadataViewer"
  member = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-pubsub.iam.gserviceaccount.com"
}

resource "google_project_iam_member" "editor" {
  project = data.google_project.project.project_id
  role   = "roles/bigquery.dataEditor"
  member = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-pubsub.iam.gserviceaccount.com"
}

resource "google_bigquery_dataset" "samples" {
  dataset_id = "samples"
}

resource "google_bigquery_table" "samples" {
  deletion_protection = false
  table_id   = "samples"
  dataset_id = google_bigquery_dataset.samples.dataset_id

  schema = <<EOF
[
  {
    "name": "data",
    "type": "JSON",
    "mode": "NULLABLE",
    "description": "The data"
  }
]
EOF
}

