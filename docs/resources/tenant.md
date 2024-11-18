---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "flightdeck_tenant Resource - terraform-provider-flightdeck"
subcategory: ""
description: |-
  Represents a Tenant resource.
---

# flightdeck_tenant (Resource)

Represents a Tenant resource.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `display_name` (String) The display name of the Tenant.
- `name` (String) The name of the Tenant.
- `organization_id` (String)

### Read-Only

- `created_at` (String) The date and time of the resources creation.
- `id` (String) The ID of the Flightdeck resource.
- `identifier` (String) The internal identifier of the Tenant.
- `issuer_url` (String) The URL of the Tenant's OIDC Issuer.