variable "service_account" {
  type        = string
  description = "The service account to impersonate"
}

variable "config_path" {
  type = string
  description = "The path to the kube config file"
  default = "~/.kube/config"
}

variable "config_context" {
  type = string
  description = "The context to use in the kube config file"
}

variable "zone" {
  type = string
  description = "The zone of the bastion proxy"
}


variable "instance" {
  type = string
  description = "The instance to use for the bastion proxy"
}

variable "interface" {
  type = string
  description = "The interface to use for the bastion proxy"
  default = "nic0"
}

variable "project" {
  type = string
  description = "The project of the bastion proxy"
}

variable "remote_port" {
  type = string
  description = "The remote_port of the bastion proxy"
  # tiny proxy default
  default = 8888
}


