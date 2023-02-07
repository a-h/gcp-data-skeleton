resource "google_cloudfunctions2_function" "on_sample_published" {
  name     = "on-sample-published"
  location = var.zone

  build_config {
    runtime     = "go119"
    entry_point = "on-sample-published"
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

  event_trigger {
    trigger_region = var.zone
    event_type     = "google.cloud.pubsub.topic.v1.messagePublished"
    pubsub_topic   = google_pubsub_topic.samples.id
    retry_policy   = "RETRY_POLICY_DO_NOT_RETRY" # "RETRY_POLICY_RETRY"
  }

  lifecycle {
    replace_triggered_by = [
      google_storage_bucket_object.code
    ]
  }
}
