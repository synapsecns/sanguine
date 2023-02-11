# google provider used for kube access
provider "google" {
}

# token
data "google_service_account_access_token" "kube_sa" {
  target_service_account = var.service_account
  lifetime               = "1000s"
  scopes = [
    "https://www.googleapis.com/auth/cloud-platform",
    "https://www.googleapis.com/auth/userinfo.email"
  ]
}
