terraform {
  required_providers {
    boilerplate = {
      version = "~> 1.0.0"
      source  = "boilerplate-example.com/boilerplateprovider/boilerplate"
    }
  }
}

resource "boilerplate_uuid" "id" {
  uuid_count = 1
}

resource "boilerplate_person" "student" {
  pid   = boilerplate_uuid.id.uuid_count
  name = "precious"
  city = "lagos"
  age  = 20
}