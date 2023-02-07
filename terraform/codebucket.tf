resource "random_id" "bucket_suffix" {
  byte_length = 8
}

resource "google_storage_bucket" "bucket" {
  name                        = "gcp-data-skeleton-${random_id.bucket_suffix.hex}"
  location                    = var.region
  uniform_bucket_level_access = true
  public_access_prevention    = "enforced"
}

resource "google_storage_bucket_object" "code" {
  name   = "source.zip"
  bucket = google_storage_bucket.bucket.name
  source = "../source.zip"
}

