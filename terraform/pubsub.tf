resource "google_pubsub_topic" "samples" {
  name                       = "samples"
  message_retention_duration = "86600s"
}

