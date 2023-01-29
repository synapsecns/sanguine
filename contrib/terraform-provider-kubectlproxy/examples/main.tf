terraform {
  required_providers {
    kubectlproxy = {
      version = "~> 1.0.0"
      source  = "example-kube.com/provider/kubectlproxy"
    }
  }
}


provider "kubectlproxy" {
  instance = var.instance
  zone = var.zone
  interface = var.interface
  project = var.project
  remote_port = var.remote_port

  // TODO: this needs to be changed to work cross cluster
  host           = ""
  token          = data.google_service_account_access_token.kube_sa.access_token
  config_path    = var.config_path
  config_context = var.config_context
}

resource "kubectlproxy_secret" "example" {
  metadata {
    name = "basic-auth"
  }

  data = {
    username = "admin"
    password = "P4ssw0rd"
  }

  type = "kubernetes.io/basic-auth"
}
