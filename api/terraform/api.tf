resource "random_id" "bucket_suffix" {
  byte_length = 8
}

resource "google_storage_bucket" "bucket" {
  name                        = "gcp-data-skeleton-${random_id.bucket_suffix.hex}"
  location                    = var.region
  uniform_bucket_level_access = true
  public_access_prevention = "enforced"
}

resource "google_storage_bucket_object" "object" {
  name   = "function/function-source.zip"
  bucket = google_storage_bucket.bucket.name
  source   = "../function/function-source.zip"
}

resource "google_cloudfunctions2_function" "function" {
  name        = "function-v2"
  location    = var.zone
  description = "a new function"

  build_config {
    runtime     = "go119"
    entry_point = "HelloGet"
    source {
      storage_source {
        bucket = google_storage_bucket.bucket.name
        object = google_storage_bucket_object.object.name
      }
    }
  }

  service_config {
    max_instance_count = 1
    available_memory   = "256M"
    timeout_seconds    = 60
  }
}

output "function_uri" {
  value = google_cloudfunctions2_function.function.service_config[0].uri
}
