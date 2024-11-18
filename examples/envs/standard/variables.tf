locals {
  timestamp = timestamp()
}

variable "api_endpoint" {
  type    = string
  default = "https://api.arctir.cloud/v1"
}

variable "organization_id" {
  type = string
}

variable "tenant_name" {
  type    = string
  default = "default"
}

variable "tenant_display_name" {
  type    = string
  default = "Default Tenant"
}

variable "portal_version" {
  type    = string
  default = "1.32.3-arctir.6"
}

variable "portal_name" {
  type    = string
  default = "demo"
}

variable "portal_display_name" {
  type    = string
  default = "Demo Portal"
}

variable "portal_organization_name" {
  type    = string
  default = "Arctir"
}

variable "portal_version_id" {
  type    = string
  default = "1.32.3-arctir.3"
}

variable "portal_domain" {
  type    = string
  default = "demo"
}

variable "portal_alternate_domains" {
  type    = list(string)
  default = []
}

variable "github_integrations" {
  sensitive = true
  type = map(object({
    host         = string
    raw_base_url = string
    api_base_url = string
    token        = string
    apps = list(object({
      app_id         = string
      client_id      = string
      client_secret  = string
      private_key    = string
      webhook_secret = string
    }))
  }))
  default = {}
}

variable "gitlab_integrations" {
  sensitive = true
  type = map(object({
    host         = string
    base_url     = string
    api_base_url = string
    token        = string
  }))
  default = {}
}

variable "github_catalog_providers" {
  type = map(object({
    host         = string
    organization = string
    catalog_path = string
    filters = object({
      allow_forks = bool
      branch      = string
      repository  = string
      visibility  = list(string)
      topic = object({
        include = list(string)
        exclude = list(string)
      })
    })
  }))
  default = {}
}

variable "location_catalog_providers" {
  type = map(object({
    target = string
    allow  = list(string)
  }))
  default = {}
}

variable "github_auth_provider" {
  sensitive = true
  type = object({
    client_id               = string
    client_secret           = string
    callback_url            = string
    enterprise_instance_url = string
    additional_scopes       = list(string)
  })
  default = null
}

variable "tailscale_connection_auth_token" {
  type    = string
  default = null
}

variable "proxies" {
  sensitive = true
  type = map(object({
    endpoint      = string
    target        = string
    change_origin = bool
    http_headers = list(object({
      name  = string
      value = string
    }))
    credentials = string
  }))
}

variable "entity_page_layouts" {
  type = map(object({
    active = bool
    card_order = list(object({
      path   = string
      config = string
    }))
    content_order = list(object({
      path   = string
      config = string
    }))
  }))
}

variable "tenant_users" {
  type = list(object({
    email    = string
    username = string
  }))
}
