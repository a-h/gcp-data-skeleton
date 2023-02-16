terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = ">= 4.34.0"
    }
    google-beta = {
      source  = "hashicorp/google-beta"
      version = ">= 4.34.0"
    }
    random = {
      source  = "hashicorp/random"
      version = ">= 3.4.3"
    }
  }
  backend "gcs" {
    # Bucket is passed in via cli arg. Eg, terraform init -reconfigure -backend-configuration=dev.tfbackend
  }
}

provider "google" {
  project = var.project_id
  region  = var.region
}

provider "google-beta" {
  project = var.project_id
  region  = var.region
}
