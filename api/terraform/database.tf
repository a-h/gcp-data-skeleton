resource "google_app_engine_application" "app" {
  location_id   = var.zone
  database_type = "CLOUD_FIRESTORE"
}
