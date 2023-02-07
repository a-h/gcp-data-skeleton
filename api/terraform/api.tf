resource "random_id" "bucket_suffix" {
  byte_length = 8
}

resource "google_storage_bucket" "bucket" {
  name                        = "gcp-data-skeleton-${random_id.bucket_suffix.hex}"
  location                    = var.region
  uniform_bucket_level_access = true
  public_access_prevention    = "enforced"
}

resource "google_storage_bucket_object" "object" {
  name   = "function/function-source.zip"
  bucket = google_storage_bucket.bucket.name
  source = "../function/function-source.zip"
}

resource "google_cloudfunctions2_function" "function" {
  name        = "function-v2"
  location    = var.zone
  description = "a new function"

  build_config {
    runtime     = "go119"
    entry_point = "http"
    source {
      storage_source {
        bucket = google_storage_bucket.bucket.name
        object = google_storage_bucket_object.object.name
      }
    }
  }

  service_config {
    min_instance_count = 0
    max_instance_count = 100
    available_memory   = "256M"
    timeout_seconds    = 60
    environment_variables = {
      PROJECT_ID = var.project_id
      TOPIC_ID   = google_pubsub_topic.samples.name
    }
  }

  lifecycle {
    replace_triggered_by = [
      google_storage_bucket_object.object
    ]
  }
}

# Make the Cloud Function publicly accessible.
resource "google_cloud_run_service_iam_member" "public_access" {
  location = google_cloudfunctions2_function.function.location
  service  = google_cloudfunctions2_function.function.name
  role     = "roles/run.invoker"
  member   = "allUsers"
}

resource "google_pubsub_topic" "samples" {
  name                       = "samples"
  message_retention_duration = "86600s"
}

output "function_uri" {
  value = google_cloudfunctions2_function.function.service_config[0].uri
}
