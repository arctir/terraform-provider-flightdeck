---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "flightdeck_auth_provider Resource - terraform-provider-flightdeck"
subcategory: ""
description: |-
  Represents a Flightdeck Auth Provider resource.
---

# flightdeck_auth_provider (Resource)

Represents a Flightdeck Auth Provider resource.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the Auth Provider resource.
- `organization_id` (String) The ID of the Flightdeck Organization resource.
- `portal_name` (String) The name of the Flightdeck Portal resource.

### Optional

- `github` (Attributes) (see [below for nested schema](#nestedatt--github))
- `gitlab` (Attributes) (see [below for nested schema](#nestedatt--gitlab))
- `google` (Attributes) (see [below for nested schema](#nestedatt--google))

### Read-Only

- `created_at` (String) The date and time of the resources creation.
- `id` (String) The ID of the Flightdeck resource.

<a id="nestedatt--github"></a>
### Nested Schema for `github`

Required:

- `client_id` (String) The client ID for the Github OAuth2 provider.
- `client_secret` (String, Sensitive) The client secret for the Github OAuth2 provider.

Optional:

- `additional_scopes` (List of String) Any additional scopes necessary for the Github OAuth2 provider.
- `callback_url` (String) The callback URL for the Github OAuth2 provider.
- `enterprise_instance_url` (String) The instance URL for the Github Enterprise instance.


<a id="nestedatt--gitlab"></a>
### Nested Schema for `gitlab`

Required:

- `client_id` (String) The client ID for the Gitlab OAuth2 provider.
- `client_secret` (String, Sensitive) The client secret for the Gitlab OAuth2 provider.

Optional:

- `additional_scopes` (List of String) Any additional scopes necessary for the Gitlab OAuth2 provider.
- `audience` (String) The audience for the Gitlab OAuth2 provider.
- `callback_url` (String) The callback URL for the Gitlab OAuth2 provider.


<a id="nestedatt--google"></a>
### Nested Schema for `google`

Required:

- `client_id` (String) The client ID for the Google OAuth2 provider.
- `client_secret` (String, Sensitive) The client secret for the Google OAuth2 provider.

Optional:

- `additional_scopes` (List of String) Any additional scopes necessary for the Google OAuth2 provider.
- `callback_url` (String) The callback URL for the Google OAuth2 provider.
