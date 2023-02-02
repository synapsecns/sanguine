terraform {
  required_providers {
    iap = {
      version = "~> 1.0.0"
      source  = "example-iap.com/provider/iap"
    }
  }
}

provider "iap" {
}


provider "kubernetes" {
  // TODO: this needs to be changed to work cross cluster
  host                   = ""
  token                  = data.google_service_account_access_token.kube_sa.access_token
  proxy_url=resource.iap_tunnel_proxy.tunnel_proxy.proxy_url
  config_path=var.config_path
  config_context=var.config_context
}


resource "iap_tunnel_proxy" "tunnel_proxy" {
  zone = var.zone
  instance = var.instance
  interface = var.interface
  project = var.project
  remote_port = var.remote_port
}

output "tunnel_proxy" {
  value = resource.iap_tunnel_proxy.tunnel_proxy.proxy_url
}


# this will likely error witha  permission/not found error. This means the proxy is working
data "kubernetes_storage_class" "example" {
  metadata {
    name = "terraform-example"
  }

  depends_on = [resource.iap_tunnel_proxy.tunnel_proxy]
}

data "iap_tunnel_keep_alive" "keep_alive" {
  # keep alive for at least 100 seconds
  timeout = 100
  proxy_url = resource.iap_tunnel_proxy.tunnel_proxy

  for_each = {
    timestamp = "${timestamp()}"
  }
}
