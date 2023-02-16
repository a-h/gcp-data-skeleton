{ pkgs ? import <nixpkgs> {} }:

{
 beam = pkgs.python39.withPackages(ps: with ps; [ 
    # [gcp] optionals.
    cachetools
    google-apitools
    google-auth
    google-auth-httplib2
    google-cloud-datastore
    google-cloud-pubsub
    # google-cloud-pubsublite - Not found
    google-cloud-bigquery
    google-cloud-bigquery-storage
    google-cloud-core
    google-cloud-bigtable
    google-cloud-spanner
    google-cloud-dlp
    google-cloud-language
    google-cloud-videointelligence
    google-cloud-vision
    # google-cloud-recommendations-ai - Not found
    # End of [gcp] section.
    google-cloud-storage
    apache-beam
    grpcio
    # Apache Beam tries to use pip. I don't know why.
    pip
  ]);
}
