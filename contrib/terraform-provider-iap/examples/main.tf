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

data "iap_tunnel_proxy" "tunnel_proxy" {
  hostname        = "my-hostname"
  zone = "us-west1-a"
  instance = "my-bastion-proxy"
  interface = "nic0"
  project = "my-fully-qualified-proxy"
  remote_port = 8888
}

output "tunnel_proxy" {
  value = data.iap_tunnel_proxy.tunnel_proxy.proxy_url
}
