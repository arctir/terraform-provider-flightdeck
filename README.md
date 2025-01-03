# terraform-provider-flightdeck

![flightdeck logo](https://github.com/arctir/terraform-provider-flightdeck/blob/main/img/flightdeck.png?raw=true)

A [Terraform](https://terraform.io) provider for the Flightdeck Developer Platform.

### About

This Terraform provider is largely auto-generated with the [Flightdeck OpenAPI Specification](https://github.com/arctir/flightdeck-api). This ensures that as the API changes, new resources and/or capabilities may be integrated into this Terraform provider.

### Configuration

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

### Examples

Example usage may be found under the `/examples` directory of this repository
