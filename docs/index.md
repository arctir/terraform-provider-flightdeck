---
layout: ""
page_title: "Provider: Unifi"
description: |
  The Flightdeck provider provides resources to interact with the Flightdeck Developer Platform.
---

# Flightdeck Provider

The Flightdeck provider provides resources to interact with the Flightdeck Developer Platform.

## Configuration

This provider interfaces with the [Flightdeck API](https://github.com/arctir/flightdeck-api) and requires that the user provide a valid, 
authenticated, and authorized API token. This configuraiton may be easily obtained with the [Flightdeck CLI](https://github.com/arctir/flightdeck-cli).
After downloading the binray that matches your operating system and platform from the [Releases Page](https://github.com/arctir/flightdeck-cli/releases), 
simply execute:

```shell
flightdeck-cli auth login
```

This will open a browser window and prompt you to login to your Flightdeck account. After logging in, the CLI will save your API token to your local 
Flightdeck configuration file. 

This Terraform provider will automatically read this configuration file and use the API token to authenticate with the Flightdeck API.

## Example Usage

```terraform
terraform {
  required_providers {
    flightdeck = {
      source  = "arctir/flightdeck"
      version = "1.0.0"
    }
  }
}

provider "flightdeck" {
  api_endpoint = var.api_endpoint
}

data "flightdeck_organization" "default" {
  id = var.organization_id
}

resource "flightdeck_tenant" "default" {
  organization_id = data.flightdeck_organization.default.id
  name            = var.tenant_name
  display_name    = var.tenant_display_name
}

data "flightdeck_portal_version" "default" {
  version = var.portal_version
}

resource "flightdeck_portal" "demo" {
  organization_id   = data.flightdeck_organization.default.id
  tenant_name       = flightdeck_tenant.default.name
  name              = var.portal_name
  title             = var.portal_display_name
  organization_name = var.portal_organization_name
  version_id        = data.flightdeck_portal_version.default.id
  domain            = var.portal_domain
  alternate_domains = var.portal_alternate_domains
}

resource "flightdeck_integration" "github" {
  organization_id = data.flightdeck_organization.default.id
  portal_name     = flightdeck_portal.demo.name
  for_each        = nonsensitive(toset(keys(var.github_integrations)))
  name            = each.key
  github          = var.github_integrations[each.key]
}

resource "flightdeck_integration" "gitlab" {
  organization_id = data.flightdeck_organization.default.id
  portal_name     = flightdeck_portal.demo.name
  for_each        = nonsensitive(toset(keys(var.gitlab_integrations)))
  name            = each.key
  gitlab          = var.gitlab_integrations[each.key]
}

resource "flightdeck_catalog_provider" "location" {
  organization_id = data.flightdeck_organization.default.id
  portal_name     = flightdeck_portal.demo.name
  for_each        = nonsensitive(toset(keys(var.location_catalog_providers)))
  name            = each.key
  location        = var.location_catalog_providers[each.key]
}

resource "flightdeck_catalog_provider" "github" {
  organization_id = data.flightdeck_organization.default.id
  portal_name     = flightdeck_portal.demo.name
  for_each        = nonsensitive(toset(keys(var.github_catalog_providers)))
  name            = each.key
  github          = var.github_catalog_providers[each.key]
}

resource "flightdeck_tenant_user" "user" {
  organization_id = data.flightdeck_organization.default.id
  tenant_name     = flightdeck_tenant.default.name
  for_each        = { for k, v in var.tenant_users : k => v }
  email           = each.value.email
  username        = each.value.username
}

resource "flightdeck_auth_provider" "github" {
  count           = var.github_auth_provider != null ? 1 : 0
  organization_id = data.flightdeck_organization.default.id
  portal_name     = flightdeck_portal.demo.name
  name            = "github"
  github          = var.github_auth_provider
}

resource "flightdeck_connection" "tailscale" {
  count           = var.tailscale_connection_auth_token != null ? 1 : 0
  organization_id = data.flightdeck_organization.default.id
  portal_name     = flightdeck_portal.demo.name
  name            = "tailscale"
  tailscale = {
    auth_token = var.tailscale_connection_auth_token
  }
}

resource "flightdeck_portal_proxy" "proxy" {
  organization_id = data.flightdeck_organization.default.id
  portal_name     = flightdeck_portal.demo.name
  for_each        = nonsensitive(toset(keys(var.proxies)))
  name            = each.key
  endpoint        = var.proxies[each.key].endpoint
  target          = var.proxies[each.key].target
  change_origin   = var.proxies[each.key].change_origin
  http_headers    = var.proxies[each.key].http_headers
  credentials     = var.proxies[each.key].credentials
}

resource "flightdeck_entity_page_layout" "default" {
  organization_id = data.flightdeck_organization.default.id
  portal_name     = flightdeck_portal.demo.name
  for_each        = toset(keys(var.entity_page_layouts))
  name            = each.key
  active          = var.entity_page_layouts[each.key].active
  card_order      = var.entity_page_layouts[each.key].card_order
  content_order   = var.entity_page_layouts[each.key].content_order
}

resource "flightdeck_plugin_configuration" "argocd" {
  organization_id = data.flightdeck_organization.default.id
  portal_name     = flightdeck_portal.demo.name
  definition = {
    name              = "argocd"
    portal_version_id = data.flightdeck_portal_version.default.id
  }
  frontend_config = jsonencode({
    argocd = {
      revisionsToLoad = 3
    }
  })
  enabled = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `api_endpoint` (String) The endpoint for the flightdeck API
- `config_path` (String) The path to the flightdeck config file
