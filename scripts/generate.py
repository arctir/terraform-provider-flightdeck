# Copyright (c) Arctir, Inc.
# SPDX-License-Identifier: Apache-2.0



import json
import sys
import humps


class DatasourceMapping:

    def __init__(self, name, resource, ignore_fields=[], attribute_map={}):
        self.name = name
        self.resource = resource
        self.attribute_map = attribute_map
        self.ignore_fields = ignore_fields

    def map(self):
        schema = apispec["components"]["schemas"][self.resource]
        description = schema["description"] if "description" in schema else None
        res = to_resource(
            self.name,
            apispec["components"]["schemas"][self.resource],
            apispec["components"]["schemas"],
            ignore_fields=self.ignore_fields,
        )

        for a in res.attributes:
            if a.name in self.attribute_map:
                if self.attribute_map[a.name] == "required":
                    a.required = True
                if self.attribute_map[a.name] == "computed":
                    a.required = True
                    a.computed = True

        return Datasource(self.name, res, description=description)


class ResourceMapping:

    def __init__(
        self,
        name,
        create_resource,
        read_resource,
        create_ignore_fields=[],
        read_ignore_fields=[],
        extra_attributes=[],
    ):
        self.name = name
        self.create_resource = create_resource
        self.read_resource = read_resource
        self.extra_attributes = extra_attributes
        self.create_ignore_fields = create_ignore_fields
        self.read_ignore_fields = read_ignore_fields

    def map(self, spec):
        pair = ResourcePair(
            extra=self.extra_attributes,
        )
        pair.create = to_resource(
            self.name,
            spec["components"]["schemas"][self.create_resource],
            spec["components"]["schemas"],
            ignore_fields=self.create_ignore_fields,
        )
        pair.read = to_resource(
            self.name,
            spec["components"]["schemas"][self.read_resource],
            spec["components"]["schemas"],
            ignore_fields=self.read_ignore_fields,
        )

        schema = apispec["components"]["schemas"][self.read_resource]
        pair.read.description = (
            schema["description"] if "description" in schema else None
        )

        return pair


def serialize(obj):
    if hasattr(obj, "as_dict"):
        return obj.as_dict()
    return obj.__dict__


def merge_allof(d):
    if isinstance(d, dict):
        description = None
        for k, v in d.copy().items():
            if k == "description":
                description = v
            if isinstance(v, dict):
                d[k] = merge_allof(v)
            if isinstance(v, list):
                if k == "allOf":
                    new_props = {}
                    new_requires = []
                    for o in v:
                        if "properties" in o:
                            new_props = {**new_props, **o["properties"]}
                        if "required" in o:
                            new_requires = new_requires + o["required"]
                    d.pop(k)
                    d = {
                        "type": "object",
                        "description": description,
                        "properties": new_props,
                        "required": new_requires,
                    }
                else:
                    d[k] = merge_allof(v)
    return d


class Spec:
    def __init__(self, name, description, data_sources=[], resources=[]):
        self.name = name
        self.data_sources = data_sources
        self.resources = resources
        self.description = description

    def as_dict(self):
        return {
            "version": "0.1",
            "provider": {
                "name": self.name,
                "schema": {"description": self.description},
            },
            "datasources": self.data_sources,
            "resources": self.resources,
        }


class Datasource:
    def __init__(self, name, resource, extra=[], description=None):
        self.name = name
        self.resource = resource
        self.description = description

    def as_dict(self):
        out = {
            "name": self.resource.name,
            "schema": {
                "attributes": self.resource.attributes,
            },
        }
        if self.description:
            out["schema"]["description"] = self.description
        return out


class ResourcePair:
    def __init__(self, create=None, read=None, extra=[]):
        self.create = create
        self.read = read
        self.extra = extra

    def as_dict(self):
        attrs = self.extra + self.create.attributes

        known_attrs = []
        for a in attrs:
            known_attrs.append(a.name)

        if self.read:
            for a in self.read.attributes:
                if a.name not in known_attrs:
                    a.computed = True
                    attrs.append(a)

        out = {
            "name": self.create.name,
            "schema": {
                "attributes": attrs,
            },
        }
        if self.read and self.read.description:
            out["schema"]["description"] = self.read.description
        return out


class Resource:
    def __init__(self, name, type_name, attributes, nested=False, description=None):
        self.name = name
        self.attributes = attributes
        self.type_name = type_name
        self.nested = nested
        self.description = description

    def as_dict(self):
        t = "object" if self.type_name == "object" else None

        for a in self.attributes:
            a.bare = True

        o = {
            "name": self.name,
        }
        o[t] = {}
        o[t]["attribute_types"] = self.attributes
        return o


class ResourceAttribute:
    def __init__(
        self,
        name,
        type_name,
        required=False,
        computed=False,
        bare=False,
        nested=False,
        default=None,
        sensitive=False,
        description=None,
    ):
        self.name = humps.decamelize(name)
        self.type_name = type_name
        self.required = required
        self.computed = computed
        self.bare = bare
        self.nested = nested
        self.default = default
        self.sensitive = sensitive
        self.description = description

    def _computed_optional_required(self):
        if self.computed:
            if not self.required and self.computed:
                return "computed_optional"
            return "computed"
        if self.required:
            return "required"
        return "computed_optional"

    def _type_options(self, requirements=True):
        o = {}
        if requirements:
            o["computed_optional_required"] = self._computed_optional_required()
        if self.default and not self.nested:
            o["default"] = {"static": self.default}
        if self.sensitive and not self.nested:
            o["sensitive"] = True
        if self.description:
            o["description"] = self.description
        return o

    def type_spec(self, requirements=True):
        t = {}
        t[self.type_name] = self._type_options(requirements)
        return t

    def as_dict(self):
        o = {}
        if not self.bare:
            o["name"] = self.name
        if self.nested:
            o |= self.type_spec(False)
        else:
            o |= self.type_spec(True)
        return o


class ResourceBoolAttribute(ResourceAttribute):
    def __init__(
        self, name, required=False, bare=False, nested=False, description=None
    ):
        ResourceAttribute.__init__(
            self, name, "bool", required, bare, nested, description=description
        )


class ResourceStringAttribute(ResourceAttribute):
    def __init__(
        self,
        name,
        required=False,
        bare=False,
        nested=False,
        default=None,
        sensitive=False,
        description=None,
    ):
        ResourceAttribute.__init__(
            self,
            name,
            "string",
            required,
            bare,
            nested,
            default=default,
            sensitive=sensitive,
            description=description,
        )


class ResourceIntegerAttribute(ResourceAttribute):
    def __init__(
        self, name, required=False, bare=False, nested=False, description=None
    ):
        ResourceAttribute.__init__(
            self, name, "int64", required, bare, nested, description=description
        )


class ResourceListAttribute(ResourceAttribute):
    def __init__(
        self, name, subobj, required=False, bare=False, nested=False, description=None
    ):
        ResourceAttribute.__init__(
            self, name, "list_nested", required, bare, nested, description=description
        )
        self.subobj = subobj
        self.subobj.bare = True
        self.subobj.nested = True
        self.nested = nested

    def type_spec(self, requirements=True):
        o = super().type_spec(requirements)
        o[self.type_name]["element_type"] = self.subobj
        return o

    def as_dict(self):
        o = {
            "name": self.name,
        }
        if isinstance(self.subobj, ResourceObjectAttribute):
            o["list_nested"] = {
                "nested_object": self.subobj,
                "computed_optional_required": self._computed_optional_required(),
            }
            if self.description:
                o["list_nested"]["description"] = self.description

        else:
            o["list"] = {
                "element_type": self.subobj,
                "computed_optional_required": self._computed_optional_required(),
            }
            if self.description:
                o["list"]["description"] = self.description
        return o


class ResourceObjectAttribute(ResourceAttribute):
    def __init__(
        self,
        name,
        attributes,
        required=False,
        bare=False,
        nested=False,
        description=None,
    ):
        ResourceAttribute.__init__(
            self, name, "single_nested", required, bare, nested, description=description
        )
        self.attributes = attributes
        self.nested = nested

    def as_dict(self):
        o = {}
        if not self.nested:
            o = {
                "name": self.name,
                "single_nested": {
                    "attributes": self.attributes,
                    "computed_optional_required": self._computed_optional_required(),
                },
            }
        else:
            o = {
                "attributes": self.attributes,
                "computed_optional_required": self._computed_optional_required(),
            }
        return o


def to_resource(name, component, all_components, ignore_fields=[], basepath=""):
    attributes = []

    path = ".".join([basepath, name])
    if path in ignore_fields:
        return None

    def decode_type(field, spec, required, ignore_fields=[]):
        r = []

        if path in ignore_fields:
            return r

        default = False
        sensitive = spec["format"] == "password" if "format" in spec else False
        description = spec["description"] if "description" in spec else None

        if "oneOf" in spec and "discriminator" in spec:
            for dk, dv in spec["discriminator"]["mapping"].items():
                foo = dv.split("/")[-1]
                subtype = all_components[foo]
                ignore = ignore_fields + [
                    ".".join([path, field, spec["discriminator"]["propertyName"]])
                ]
                res = to_resource(
                    dk, subtype, all_components, ignore_fields=ignore, basepath=path
                )
                if res:
                    r.append(res)

        elif spec["type"] == "string":
            r.append(
                ResourceStringAttribute(
                    field,
                    required,
                    default=default,
                    sensitive=sensitive,
                    description=description,
                )
            )
        elif spec["type"] == "boolean":
            r.append(ResourceBoolAttribute(field, required, description=description))
        elif spec["type"] == "integer":
            r.append(ResourceIntegerAttribute(field, required, description=description))
        elif spec["type"] == "array":
            listobj = decode_type(
                field, spec["items"], True, ignore_fields=ignore_fields
            )
            if len(listobj) != 1:
                raise Exception("List should be a single element")
            listobj[0].nested = True

            nested = False
            r.append(
                ResourceListAttribute(
                    k, listobj[0], required, False, nested, description=description
                )
            )
        elif spec["type"] == "object":
            res = to_resource(
                field, spec, all_components, basepath=path, ignore_fields=ignore_fields
            )
            if res:
                r.append(res)
        else:
            print(spec)
        return r

    if component["type"] == "object":
        description = component["description"] if "description" in component else None
        if not "properties" in component:
            # treat this as a json blob

            return ResourceStringAttribute(name, description=description)
        else:
            for k, v in component["properties"].items():
                prop_path = ".".join([path, k])
                if prop_path in ignore_fields:
                    continue
                required = "required" in component and k in component["required"]
                a = decode_type(k, v, required, ignore_fields)
                attributes += a
            return ResourceObjectAttribute(
                name,
                attributes,
                description=description,
            )
    else:
        raise ("Could not decode type")


# load the spec file and read the yaml
apispec = None
with open(sys.argv[1]) as f:
    apispec = json.loads(f.read())

apispec = merge_allof(apispec)

data_sources_mappings = [
    DatasourceMapping(
        "cluster",
        "Cluster",
        attribute_map={
            "id": "computed",
            "created_at": "computed",
            "name": "required",
            "display_name": "computed",
            "region": "computed",
        },
    ),
    DatasourceMapping(
        "organization",
        "Organization",
        attribute_map={
            "name": "computed",
            "id": "required",
            "created_at": "computed",
            "owner": "computed",
            "subdomain": "computed",
            "cluster_id": "computed",
        },
    ),
    DatasourceMapping(
        "portal_version",
        "PortalVersion",
        attribute_map={
            "version": "required",
            "id": "computed",
            "created_at": "computed",
            "major": "computed",
            "minor": "computed",
            "patch": "computed",
            "rev": "computed",
        },
    ),
]

resource_mappings = [
    ResourceMapping(
        "organization",
        "OrganizationInput",
        "Organization",
    ),
    ResourceMapping(
        "tenant",
        "TenantInput",
        "Tenant",
        extra_attributes=[ResourceStringAttribute("organization_id", True)],
    ),
    ResourceMapping(
        "identity_provider",
        "IdentityProviderInput",
        "IdentityProvider",
        extra_attributes=[
            ResourceStringAttribute(
                "organization_id",
                True,
                description="The ID of the Flightdeck Organization resource.",
            ),
            ResourceStringAttribute(
                "tenant_name",
                True,
                description="The name of the Flightdeck Tenant resource.",
            ),
        ],
    ),
    ResourceMapping(
        "tenant_user",
        "TenantUserInput",
        "TenantUser",
        extra_attributes=[
            ResourceStringAttribute(
                "organization_id",
                True,
                description="The ID of the Flightdeck Organization resource.",
            ),
            ResourceStringAttribute(
                "tenant_name",
                True,
                description="The name of the Flightdeck Tenant resource.",
            ),
        ],
    ),
    ResourceMapping(
        "portal",
        "PortalInput",
        "Portal",
        extra_attributes=[ResourceStringAttribute("organization_id", True)],
        read_ignore_fields=[
            ".portal.version",
        ],
    ),
    ResourceMapping(
        "integration",
        "IntegrationInput",
        "Integration",
        extra_attributes=[
            ResourceStringAttribute(
                "organization_id",
                True,
                description="The ID of the Flightdeck Organization resource.",
            ),
            ResourceStringAttribute(
                "portal_name",
                True,
                description="The name of the Flightdeck Portal resource.",
            ),
        ],
        create_ignore_fields=[
            ".integration.github.configType",
            ".integration.gitlab.configType",
            ".integration.location.configType",
        ],
        read_ignore_fields=[
            ".integration.github.configType",
            ".integration.gitlab.configType",
            ".integration.location.configType",
        ],
    ),
    ResourceMapping(
        "catalog_provider",
        "CatalogProviderInput",
        "CatalogProvider",
        extra_attributes=[
            ResourceStringAttribute(
                "organization_id",
                True,
                description="The ID of the Flightdeck Organization resource.",
            ),
            ResourceStringAttribute(
                "portal_name",
                True,
                description="The name of the Flightdeck Portal resource.",
            ),
        ],
        create_ignore_fields=[
            ".catalog_provider.github.configType",
            ".catalog_provider.gitlab.configType",
            ".catalog_provider.location.configType",
            ".catalog_provider.github.schedule",
            ".catalog_provider.gitlab.schedule",
        ],
        read_ignore_fields=[
            ".catalog_provider.github.configType",
            ".catalog_provider.gitlab.configType",
            ".catalog_provider.location.configType",
            ".catalog_provider.github.schedule",
            ".catalog_provider.gitlab.schedule",
        ],
    ),
    ResourceMapping(
        "plugin_configuration",
        "PluginConfigurationInput",
        "PluginConfiguration",
        extra_attributes=[
            ResourceStringAttribute(
                "organization_id",
                True,
                description="The ID of the Flightdeck Organization resource.",
            ),
            ResourceStringAttribute(
                "portal_name",
                True,
                description="The name of the Flightdeck Portal resource.",
            ),
        ],
    ),
    ResourceMapping(
        "connection",
        "ConnectionInput",
        "Connection",
        extra_attributes=[
            ResourceStringAttribute(
                "organization_id",
                True,
                description="The ID of the Flightdeck Organization resource.",
            ),
            ResourceStringAttribute(
                "portal_name",
                True,
                description="The name of the Flightdeck Portal resource.",
            ),
        ],
        create_ignore_fields=[
            ".connection.tailscale.configType",
        ],
        read_ignore_fields=[
            ".connection.tailscale.configType",
        ],
    ),
    ResourceMapping(
        "auth_provider",
        "AuthProviderInput",
        "AuthProvider",
        extra_attributes=[
            ResourceStringAttribute(
                "organization_id",
                True,
                description="The ID of the Flightdeck Organization resource.",
            ),
            ResourceStringAttribute(
                "portal_name",
                True,
                description="The name of the Flightdeck Portal resource.",
            ),
        ],
        create_ignore_fields=[
            ".auth_provider.github.configType",
            ".auth_provider.gitlab.configType",
            ".auth_provider.google.configType",
        ],
        read_ignore_fields=[
            ".auth_provider.github.configType",
            ".auth_provider.gitlab.configType",
            ".auth_provider.google.configType",
        ],
    ),
    ResourceMapping(
        "portal_proxy",
        "PortalProxyInput",
        "PortalProxy",
        extra_attributes=[
            ResourceStringAttribute(
                "organization_id",
                True,
                description="The ID of the Flightdeck Organization resource.",
            ),
            ResourceStringAttribute(
                "portal_name",
                True,
                description="The name of the Flightdeck Portal resource.",
            ),
        ],
        create_ignore_fields=[],
        read_ignore_fields=[".created_at"],
    ),
    ResourceMapping(
        "entity_page_layout",
        "EntityPageLayoutInput",
        "EntityPageLayout",
        extra_attributes=[
            ResourceStringAttribute(
                "organization_id",
                True,
                description="The ID of the Flightdeck Organization resource.",
            ),
            ResourceStringAttribute(
                "portal_name",
                True,
                description="The name of the Flightdeck Portal resource.",
            ),
        ],
        create_ignore_fields=[".created_at"],
        read_ignore_fields=[".created_at"],
    ),
]

data_sources = []
for mapping in data_sources_mappings:
    source = mapping.map()
    data_sources.append(source)

resources = []
for resource_mapping in resource_mappings:
    resources.append(resource_mapping.map(apispec))

tfspec = Spec(
    "flightdeck", "A Terraform provider for Flightdeck", data_sources, resources
)
print(json.dumps(tfspec, default=serialize, indent=2))
