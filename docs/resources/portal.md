---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "flightdeck_portal Resource - terraform-provider-flightdeck"
subcategory: ""
description: |-
  Represents a Portal resource.
---

# flightdeck_portal (Resource)

Represents a Portal resource.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `alternate_domains` (List of String) A list of alternate domains for the Portal.
- `domain` (String) The primary domain of the Portal.
- `name` (String) The name of the Portal.
- `organization_id` (String)
- `organization_name` (String) The name of the Organization operating this Portal.
- `tenant_name` (String) The name of the Tenant providing user idenity to this Portal.
- `title` (String) The HTML title of the Portal.
- `version_id` (String) The ID of the Portal Version.

### Optional

- `status` (Attributes) (see [below for nested schema](#nestedatt--status))

### Read-Only

- `created_at` (String) The date and time of the resources creation.
- `hostname` (String) The hostname of the Portal.
- `id` (String) The ID of the Flightdeck resource.
- `identifier` (String) The identifier of the Portal.
- `url` (String) The primary URL of the Portal.

<a id="nestedatt--status"></a>
### Nested Schema for `status`

Required:

- `detail` (String) A detailed message about the current status of the Portal.
- `status` (String) The current status of the Portal.