terraform {
  required_providers {
    helmproxy = {
      version = "~> 1.0.0"
      source  = "example-helm.com/provider/helmproxy"
    }
  }
}


provider "helmproxy" {
  instance = var.instance
  zone = var.zone
  interface = var.interface
  project = var.project
  remote_port = var.remote_port

  kubernetes {
    // TODO: this needs to be changed to work cross cluster
    host           = ""
    token          = data.google_service_account_access_token.kube_sa.access_token
    config_path    = var.config_path
    config_context = var.config_context
  }
}

resource "helmproxy_release" "omnirpc_example" {
  name       = "omnirpc-example"
  chart = "../../../charts/omnirpc/"
}
