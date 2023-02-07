resource "google_cloudfunctions2_function" "http" {
  name     = "http"
  location = var.zone

  build_config {
    runtime     = "go119"
    entry_point = "http"
    source {
      storage_source {
        bucket = google_storage_bucket.bucket.name
        object = google_storage_bucket_object.code.name
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
      google_storage_bucket_object.code
    ]
  }
}

# Make the Cloud Function publicly accessible.
resource "google_cloud_run_service_iam_member" "public_access" {
  location = google_cloudfunctions2_function.http.location
  service  = google_cloudfunctions2_function.http.name
  role     = "roles/run.invoker"
  member   = "allUsers"

  lifecycle {
    replace_triggered_by = [
      google_cloudfunctions2_function.http
    ]
  }
}

output "function_uri" {
  value = google_cloudfunctions2_function.http.service_config[0].uri
}

