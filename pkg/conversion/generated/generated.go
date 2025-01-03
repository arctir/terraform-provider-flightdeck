// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.
//go:build !goverter

package generated

import (
	v1 "github.com/arctir/go-flightdeck/pkg/api/v1"
	conversion "github.com/arctir/terraform-provider-flightdeck/pkg/conversion"
	datasourcecluster "github.com/arctir/terraform-provider-flightdeck/pkg/provider/datasource_cluster"
	datasourceorganization "github.com/arctir/terraform-provider-flightdeck/pkg/provider/datasource_organization"
	datasourceportalversion "github.com/arctir/terraform-provider-flightdeck/pkg/provider/datasource_portal_version"
	resourceauthprovider "github.com/arctir/terraform-provider-flightdeck/pkg/provider/resource_auth_provider"
	resourcecatalogprovider "github.com/arctir/terraform-provider-flightdeck/pkg/provider/resource_catalog_provider"
	resourceconnection "github.com/arctir/terraform-provider-flightdeck/pkg/provider/resource_connection"
	resourceentitypagelayout "github.com/arctir/terraform-provider-flightdeck/pkg/provider/resource_entity_page_layout"
	resourceidentityprovider "github.com/arctir/terraform-provider-flightdeck/pkg/provider/resource_identity_provider"
	resourceintegration "github.com/arctir/terraform-provider-flightdeck/pkg/provider/resource_integration"
	resourcepluginconfiguration "github.com/arctir/terraform-provider-flightdeck/pkg/provider/resource_plugin_configuration"
	resourceportal "github.com/arctir/terraform-provider-flightdeck/pkg/provider/resource_portal"
	resourceportalproxy "github.com/arctir/terraform-provider-flightdeck/pkg/provider/resource_portal_proxy"
	resourcetenant "github.com/arctir/terraform-provider-flightdeck/pkg/provider/resource_tenant"
	resourcetenantuser "github.com/arctir/terraform-provider-flightdeck/pkg/provider/resource_tenant_user"
	basetypes "github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type AuthProviderConverterImpl struct{}

func (c *AuthProviderConverterImpl) ConvertToApiAuthProvider(source resourceauthprovider.AuthProviderModel) (v1.AuthProviderInput, error) {
	var v1AuthProviderInput v1.AuthProviderInput
	v1AuthProviderInput.Name = conversion.StringValueToString(source.Name)
	v1AuthProviderInput_ProviderConfig, err := conversion.SetApiAuthProviderConfig(source)
	if err != nil {
		return v1AuthProviderInput, err
	}
	v1AuthProviderInput.ProviderConfig = v1AuthProviderInput_ProviderConfig
	return v1AuthProviderInput, nil
}
func (c *AuthProviderConverterImpl) ConvertToTfAuthProvider(source v1.AuthProvider) (resourceauthprovider.AuthProviderModel, error) {
	resource_auth_providerAuthProviderModel, err := conversion.SetTfAuthProviderConfig(c, source)
	if err != nil {
		return resource_auth_providerAuthProviderModel, err
	}
	resource_auth_providerAuthProviderModel.CreatedAt = conversion.TimeToStringValue(source.CreatedAt)
	resource_auth_providerAuthProviderModel.Id = conversion.UUIDToStringValue(source.Id)
	resource_auth_providerAuthProviderModel.Name = conversion.StringToStringValue(source.Name)
	resource_auth_providerAuthProviderModel.OrganizationId = conversion.UUIDToStringValue(source.OrganizationId)
	resource_auth_providerAuthProviderModel.PortalName = conversion.StringToStringValue(source.PortalName)
	return resource_auth_providerAuthProviderModel, nil
}

type CatalogProviderConverterImpl struct{}

func (c *CatalogProviderConverterImpl) ConvertToApiCatalogProvider(source resourcecatalogprovider.CatalogProviderModel) (*v1.CatalogProviderInput, error) {
	var v1CatalogProviderInput v1.CatalogProviderInput
	v1CatalogProviderInput.Name = conversion.StringValueToString(source.Name)
	v1CatalogProviderInput_ProviderConfig, err := conversion.SetCatalogProviderConfig(source)
	if err != nil {
		return nil, err
	}
	v1CatalogProviderInput.ProviderConfig = v1CatalogProviderInput_ProviderConfig
	return &v1CatalogProviderInput, nil
}
func (c *CatalogProviderConverterImpl) ConvertToTfCatalogProvider(source v1.CatalogProvider) (resourcecatalogprovider.CatalogProviderModel, error) {
	resource_catalog_providerCatalogProviderModel, err := conversion.SetTfCatalogProviderConfig(c, source)
	if err != nil {
		return resource_catalog_providerCatalogProviderModel, err
	}
	resource_catalog_providerCatalogProviderModel.CreatedAt = conversion.TimeToStringValue(source.CreatedAt)
	resource_catalog_providerCatalogProviderModel.Id = conversion.UUIDToStringValue(source.Id)
	resource_catalog_providerCatalogProviderModel.Name = conversion.StringToStringValue(source.Name)
	resource_catalog_providerCatalogProviderModel.OrganizationId = conversion.UUIDToStringValue(source.OrganizationId)
	resource_catalog_providerCatalogProviderModel.PortalName = conversion.StringToStringValue(source.PortalName)
	return resource_catalog_providerCatalogProviderModel, nil
}

type ClusterConverterImpl struct{}

func (c *ClusterConverterImpl) ConvertToTfCluster(source v1.Cluster) datasourcecluster.ClusterModel {
	var datasource_clusterClusterModel datasourcecluster.ClusterModel
	datasource_clusterClusterModel.CreatedAt = conversion.TimeToStringValue(source.CreatedAt)
	datasource_clusterClusterModel.DisplayName = conversion.StringToStringValue(source.DisplayName)
	datasource_clusterClusterModel.Id = conversion.UUIDToStringValue(source.Id)
	datasource_clusterClusterModel.Name = conversion.StringToStringValue(source.Name)
	datasource_clusterClusterModel.Region = conversion.StringToStringValue(source.Region)
	return datasource_clusterClusterModel
}

type ConnectionConverterImpl struct{}

func (c *ConnectionConverterImpl) ConvertToApiConnection(source resourceconnection.ConnectionModel) (*v1.ConnectionInput, error) {
	var v1ConnectionInput v1.ConnectionInput
	v1ConnectionInput_ConnectionConfig, err := conversion.SetApiConnectionConfig(c, source)
	if err != nil {
		return nil, err
	}
	v1ConnectionInput.ConnectionConfig = v1ConnectionInput_ConnectionConfig
	v1ConnectionInput.Name = conversion.StringValueToString(source.Name)
	return &v1ConnectionInput, nil
}
func (c *ConnectionConverterImpl) ConvertToTfConnection(source v1.Connection) (resourceconnection.ConnectionModel, error) {
	resource_connectionConnectionModel, err := conversion.SetTfConnectionConfig(c, source)
	if err != nil {
		return resource_connectionConnectionModel, err
	}
	resource_connectionConnectionModel.CreatedAt = conversion.TimeToStringValue(source.CreatedAt)
	resource_connectionConnectionModel.Id = conversion.UUIDToStringValue(source.Id)
	resource_connectionConnectionModel.Name = conversion.StringToStringValue(source.Name)
	resource_connectionConnectionModel.OrganizationId = conversion.UUIDToStringValue(source.OrganizationId)
	resource_connectionConnectionModel.PortalName = conversion.StringToStringValue(source.PortalName)
	return resource_connectionConnectionModel, nil
}

type EntityPageLayoutConverterImpl struct{}

func (c *EntityPageLayoutConverterImpl) ConvertToApiEntityPageLayout(source resourceentitypagelayout.EntityPageLayoutModel) (v1.EntityPageLayoutInput, error) {
	var v1EntityPageLayoutInput v1.EntityPageLayoutInput
	v1EntityPageLayoutInput.Active = conversion.BoolValueToPtrBool(source.Active)
	pV1EntityPageCardSpecList, err := c.basetypesListValueToPV1EntityPageCardSpecList(source.CardOrder)
	if err != nil {
		return v1EntityPageLayoutInput, err
	}
	v1EntityPageLayoutInput.CardOrder = pV1EntityPageCardSpecList
	pV1EntityPageContentSpecList, err := c.basetypesListValueToPV1EntityPageContentSpecList(source.ContentOrder)
	if err != nil {
		return v1EntityPageLayoutInput, err
	}
	v1EntityPageLayoutInput.ContentOrder = pV1EntityPageContentSpecList
	v1EntityPageLayoutInput.Name = conversion.StringValueToString(source.Name)
	return v1EntityPageLayoutInput, nil
}
func (c *EntityPageLayoutConverterImpl) ConvertToTfEntityPageLayout(source v1.EntityPageLayout) (resourceentitypagelayout.EntityPageLayoutModel, error) {
	var resource_entity_page_layoutEntityPageLayoutModel resourceentitypagelayout.EntityPageLayoutModel
	resource_entity_page_layoutEntityPageLayoutModel.Active = conversion.PtrBoolToBoolValue(source.Active)
	basetypesListValue, err := conversion.SliceEntityPageCardSpecToListValue(source.CardOrder)
	if err != nil {
		return resource_entity_page_layoutEntityPageLayoutModel, err
	}
	resource_entity_page_layoutEntityPageLayoutModel.CardOrder = basetypesListValue
	basetypesListValue2, err := conversion.SliceEntityPageContentSpecToListValue(source.ContentOrder)
	if err != nil {
		return resource_entity_page_layoutEntityPageLayoutModel, err
	}
	resource_entity_page_layoutEntityPageLayoutModel.ContentOrder = basetypesListValue2
	resource_entity_page_layoutEntityPageLayoutModel.CreatedAt = conversion.TimeToStringValue(source.CreatedAt)
	resource_entity_page_layoutEntityPageLayoutModel.Id = conversion.UUIDToStringValue(source.Id)
	resource_entity_page_layoutEntityPageLayoutModel.Name = conversion.StringToStringValue(source.Name)
	resource_entity_page_layoutEntityPageLayoutModel.OrganizationId = conversion.UUIDToStringValue(source.OrganizationId)
	resource_entity_page_layoutEntityPageLayoutModel.PortalName = conversion.StringToStringValue(source.PortalName)
	return resource_entity_page_layoutEntityPageLayoutModel, nil
}
func (c *EntityPageLayoutConverterImpl) basetypesListValueToPV1EntityPageCardSpecList(source basetypes.ListValue) (*[]v1.EntityPageCardSpec, error) {
	v1EntityPageCardSpecList, err := conversion.ListValueToSliceEntityPageCardSpec(source)
	if err != nil {
		return nil, err
	}
	return &v1EntityPageCardSpecList, nil
}
func (c *EntityPageLayoutConverterImpl) basetypesListValueToPV1EntityPageContentSpecList(source basetypes.ListValue) (*[]v1.EntityPageContentSpec, error) {
	v1EntityPageContentSpecList, err := conversion.ListValueToSliceEntityPageContentSpec(source)
	if err != nil {
		return nil, err
	}
	return &v1EntityPageContentSpecList, nil
}

type IdentityProviderConverterImpl struct{}

func (c *IdentityProviderConverterImpl) ConvertToApiIdentityProvider(source resourceidentityprovider.IdentityProviderModel) (v1.IdentityProviderInput, error) {
	var v1IdentityProviderInput v1.IdentityProviderInput
	v1IdentityProviderInput.Name = conversion.StringValueToString(source.Name)
	v1IdentityProviderInput_ProviderConfig, err := conversion.SetApiIdentityProviderConfig(source)
	if err != nil {
		return v1IdentityProviderInput, err
	}
	v1IdentityProviderInput.ProviderConfig = v1IdentityProviderInput_ProviderConfig
	return v1IdentityProviderInput, nil
}
func (c *IdentityProviderConverterImpl) ConvertToTfIdentityProvider(source v1.IdentityProvider) (resourceidentityprovider.IdentityProviderModel, error) {
	var resource_identity_providerIdentityProviderModel resourceidentityprovider.IdentityProviderModel
	resource_identity_providerIdentityProviderModel.CreatedAt = conversion.TimeToStringValue(source.CreatedAt)
	resource_identity_providerIdentityProviderModel.Id = conversion.UUIDToStringValue(source.Id)
	resource_identity_providerIdentityProviderModel.Name = conversion.StringToStringValue(source.Name)
	resource_identity_providerIdentityProviderModel.OrganizationId = conversion.UUIDToStringValue(source.OrganizationId)
	resource_identity_providerIdentityProviderModel.TenantName = conversion.StringToStringValue(source.TenantName)
	return resource_identity_providerIdentityProviderModel, nil
}

type IntegrationConverterImpl struct{}

func (c *IntegrationConverterImpl) ConvertToApiIntegration(source resourceintegration.IntegrationModel) (*v1.IntegrationInput, error) {
	var v1IntegrationInput v1.IntegrationInput
	v1IntegrationInput_IntegrationConfig, err := conversion.SetApiIntegrationConfig(c, source)
	if err != nil {
		return nil, err
	}
	v1IntegrationInput.IntegrationConfig = v1IntegrationInput_IntegrationConfig
	v1IntegrationInput.Name = conversion.StringValueToString(source.Name)
	return &v1IntegrationInput, nil
}
func (c *IntegrationConverterImpl) ConvertToTfIntegration(source v1.Integration) (resourceintegration.IntegrationModel, error) {
	resource_integrationIntegrationModel, err := conversion.SetTfIntegrationConfig(c, source)
	if err != nil {
		return resource_integrationIntegrationModel, err
	}
	resource_integrationIntegrationModel.CreatedAt = conversion.TimeToStringValue(source.CreatedAt)
	resource_integrationIntegrationModel.Id = conversion.UUIDToStringValue(source.Id)
	resource_integrationIntegrationModel.Name = conversion.StringToStringValue(source.Name)
	resource_integrationIntegrationModel.OrganizationId = conversion.UUIDToStringValue(source.OrganizationId)
	resource_integrationIntegrationModel.PortalName = conversion.StringToStringValue(source.PortalName)
	return resource_integrationIntegrationModel, nil
}

type OrganizationConverterImpl struct{}

func (c *OrganizationConverterImpl) ConvertToTfOrganization(source v1.Organization) datasourceorganization.OrganizationModel {
	var datasource_organizationOrganizationModel datasourceorganization.OrganizationModel
	datasource_organizationOrganizationModel.ClusterId = conversion.UUIDToStringValue(source.ClusterId)
	datasource_organizationOrganizationModel.CreatedAt = conversion.TimeToStringValue(source.CreatedAt)
	datasource_organizationOrganizationModel.Id = conversion.UUIDToStringValue(source.Id)
	datasource_organizationOrganizationModel.Name = conversion.StringToStringValue(source.Name)
	datasource_organizationOrganizationModel.Owner = conversion.UUIDToStringValue(source.Owner)
	datasource_organizationOrganizationModel.Subdomain = conversion.StringToStringValue(source.Subdomain)
	return datasource_organizationOrganizationModel
}

type PluginConfigurationConverterImpl struct{}

func (c *PluginConfigurationConverterImpl) ConvertToApiPluginConfiguration(source resourcepluginconfiguration.PluginConfigurationModel) (v1.PluginConfigurationInput, error) {
	var v1PluginConfigurationInput v1.PluginConfigurationInput
	v1PluginConfigurationInput.BackendConfig = conversion.StringValueToPtrJSON(source.BackendConfig)
	v1PluginConfigurationInput.Definition = c.resource_plugin_configurationDefinitionValueToV1PluginConfigurationDefinitionSpec(source.Definition)
	v1PluginConfigurationInput.Enabled = conversion.BoolValueToBool(source.Enabled)
	v1PluginConfigurationInput.FrontendConfig = conversion.StringValueToPtrJSON(source.FrontendConfig)
	return v1PluginConfigurationInput, nil
}
func (c *PluginConfigurationConverterImpl) ConvertToTfPluginConfiguration(source v1.PluginConfiguration) (resourcepluginconfiguration.PluginConfigurationModel, error) {
	var resource_plugin_configurationPluginConfigurationModel resourcepluginconfiguration.PluginConfigurationModel
	resource_plugin_configurationPluginConfigurationModel.BackendConfig = conversion.PtrJSONToStringValue(source.BackendConfig)
	resource_plugin_configurationPluginConfigurationModel.CreatedAt = conversion.TimeToStringValue(source.CreatedAt)
	resource_plugin_configurationDefinitionValue, err := conversion.PluginConfigurationDefinitionSpecToDefinitionValue(source.Definition)
	if err != nil {
		return resource_plugin_configurationPluginConfigurationModel, err
	}
	resource_plugin_configurationPluginConfigurationModel.Definition = resource_plugin_configurationDefinitionValue
	resource_plugin_configurationPluginConfigurationModel.Enabled = conversion.BoolToBoolValue(source.Enabled)
	resource_plugin_configurationPluginConfigurationModel.FrontendConfig = conversion.PtrJSONToStringValue(source.FrontendConfig)
	resource_plugin_configurationPluginConfigurationModel.Id = conversion.UUIDToStringValue(source.Id)
	resource_plugin_configurationPluginConfigurationModel.OrganizationId = conversion.UUIDToStringValue(source.OrganizationId)
	resource_plugin_configurationPluginConfigurationModel.PortalName = conversion.StringToStringValue(source.PortalName)
	return resource_plugin_configurationPluginConfigurationModel, nil
}
func (c *PluginConfigurationConverterImpl) resource_plugin_configurationDefinitionValueToV1PluginConfigurationDefinitionSpec(source resourcepluginconfiguration.DefinitionValue) v1.PluginConfigurationDefinitionSpec {
	var v1PluginConfigurationDefinitionSpec v1.PluginConfigurationDefinitionSpec
	v1PluginConfigurationDefinitionSpec.Name = conversion.StringValueToString(source.Name)
	v1PluginConfigurationDefinitionSpec.PortalVersionId = conversion.StringValueToString(source.PortalVersionId)
	return v1PluginConfigurationDefinitionSpec
}

type PortalConverterImpl struct{}

func (c *PortalConverterImpl) ConvertToApiPortal(source resourceportal.PortalModel) (v1.PortalInput, error) {
	var v1PortalInput v1.PortalInput
	stringList, err := conversion.ListValueStringToSliceString(source.AlternateDomains)
	if err != nil {
		return v1PortalInput, err
	}
	v1PortalInput.AlternateDomains = stringList
	v1PortalInput.Domain = conversion.StringValueToString(source.Domain)
	v1PortalInput.Name = conversion.StringValueToString(source.Name)
	v1PortalInput.OrganizationName = conversion.StringValueToString(source.OrganizationName)
	v1PortalInput.TenantName = conversion.StringValueToString(source.TenantName)
	v1PortalInput.Title = conversion.StringValueToString(source.Title)
	v1PortalInput.VersionId = conversion.StringValueToString(source.VersionId)
	return v1PortalInput, nil
}
func (c *PortalConverterImpl) ConvertToTfPortal(source v1.Portal) (resourceportal.PortalModel, error) {
	var resource_portalPortalModel resourceportal.PortalModel
	basetypesListValue, err := conversion.SliceStringToListValueString(source.AlternateDomains)
	if err != nil {
		return resource_portalPortalModel, err
	}
	resource_portalPortalModel.AlternateDomains = basetypesListValue
	resource_portalPortalModel.CreatedAt = conversion.TimeToStringValue(source.CreatedAt)
	resource_portalPortalModel.Domain = conversion.StringToStringValue(source.Domain)
	resource_portalPortalModel.Hostname = conversion.StringToStringValue(source.Hostname)
	resource_portalPortalModel.Id = conversion.UUIDToStringValue(source.Id)
	resource_portalPortalModel.Identifier = conversion.StringToStringValue(source.Identifier)
	resource_portalPortalModel.Name = conversion.StringToStringValue(source.Name)
	resource_portalPortalModel.OrganizationId = conversion.UUIDToStringValue(source.OrganizationId)
	resource_portalPortalModel.OrganizationName = conversion.StringToStringValue(source.OrganizationName)
	resource_portalPortalModel.TenantName = conversion.StringToStringValue(source.TenantName)
	resource_portalPortalModel.Title = conversion.StringToStringValue(source.Title)
	resource_portalPortalModel.Url = conversion.StringToStringValue(source.Url)
	resource_portalPortalModel.VersionId = conversion.StringToStringValue(source.VersionId)
	return resource_portalPortalModel, nil
}

type PortalProxyConverterImpl struct{}

func (c *PortalProxyConverterImpl) ConvertToApiPortalProxy(source resourceportalproxy.PortalProxyModel) (v1.PortalProxyInput, error) {
	var v1PortalProxyInput v1.PortalProxyInput
	pStringList, err := conversion.ListValueStringToPtrSliceString(source.AllowedHeaders)
	if err != nil {
		return v1PortalProxyInput, err
	}
	v1PortalProxyInput.AllowedHeaders = pStringList
	pStringList2, err := conversion.ListValueStringToPtrSliceString(source.AllowedMethods)
	if err != nil {
		return v1PortalProxyInput, err
	}
	v1PortalProxyInput.AllowedMethods = pStringList2
	v1PortalProxyInput.ChangeOrigin = conversion.BoolValueToPtrBool(source.ChangeOrigin)
	v1PortalProxyInput.Credentials = conversion.StringValueToPortalProxyInputCredentials(source.Credentials)
	v1PortalProxyInput.Endpoint = conversion.StringValueToString(source.Endpoint)
	pV1PortalProxyHeaderList, err := c.basetypesListValueToPV1PortalProxyHeaderList(source.HttpHeaders)
	if err != nil {
		return v1PortalProxyInput, err
	}
	v1PortalProxyInput.HttpHeaders = pV1PortalProxyHeaderList
	v1PortalProxyInput.Name = conversion.StringValueToString(source.Name)
	pV1PortalProxyPathRewriteList, err := c.basetypesListValueToPV1PortalProxyPathRewriteList(source.PathRewrite)
	if err != nil {
		return v1PortalProxyInput, err
	}
	v1PortalProxyInput.PathRewrite = pV1PortalProxyPathRewriteList
	v1PortalProxyInput.Target = conversion.StringValueToString(source.Target)
	return v1PortalProxyInput, nil
}
func (c *PortalProxyConverterImpl) ConvertToTfPortalProxy(source v1.PortalProxy) (resourceportalproxy.PortalProxyModel, error) {
	var resource_portal_proxyPortalProxyModel resourceportalproxy.PortalProxyModel
	basetypesListValue, err := conversion.PtrSliceStringToListValueString(source.AllowedHeaders)
	if err != nil {
		return resource_portal_proxyPortalProxyModel, err
	}
	resource_portal_proxyPortalProxyModel.AllowedHeaders = basetypesListValue
	basetypesListValue2, err := conversion.PtrSliceStringToListValueString(source.AllowedMethods)
	if err != nil {
		return resource_portal_proxyPortalProxyModel, err
	}
	resource_portal_proxyPortalProxyModel.AllowedMethods = basetypesListValue2
	resource_portal_proxyPortalProxyModel.ChangeOrigin = conversion.PtrBoolToBoolValue(source.ChangeOrigin)
	resource_portal_proxyPortalProxyModel.CreatedAt = conversion.TimeToStringValue(source.CreatedAt)
	resource_portal_proxyPortalProxyModel.Credentials = conversion.PortalProxyCredentialsToStringValue(source.Credentials)
	resource_portal_proxyPortalProxyModel.Endpoint = conversion.StringToStringValue(source.Endpoint)
	basetypesListValue3, err := conversion.SlicePortalProxyHeadersToListValue(source.HttpHeaders)
	if err != nil {
		return resource_portal_proxyPortalProxyModel, err
	}
	resource_portal_proxyPortalProxyModel.HttpHeaders = basetypesListValue3
	resource_portal_proxyPortalProxyModel.Id = conversion.UUIDToStringValue(source.Id)
	resource_portal_proxyPortalProxyModel.Name = conversion.StringToStringValue(source.Name)
	resource_portal_proxyPortalProxyModel.OrganizationId = conversion.UUIDToStringValue(source.OrganizationId)
	basetypesListValue4, err := conversion.SlicePortalProxyPathRewriteToListValue(source.PathRewrite)
	if err != nil {
		return resource_portal_proxyPortalProxyModel, err
	}
	resource_portal_proxyPortalProxyModel.PathRewrite = basetypesListValue4
	resource_portal_proxyPortalProxyModel.PortalName = conversion.StringToStringValue(source.PortalName)
	resource_portal_proxyPortalProxyModel.Target = conversion.StringToStringValue(source.Target)
	return resource_portal_proxyPortalProxyModel, nil
}
func (c *PortalProxyConverterImpl) basetypesListValueToPV1PortalProxyHeaderList(source basetypes.ListValue) (*[]v1.PortalProxyHeader, error) {
	v1PortalProxyHeaderList, err := conversion.ListValueToSlicePortalProxyHeader(source)
	if err != nil {
		return nil, err
	}
	return &v1PortalProxyHeaderList, nil
}
func (c *PortalProxyConverterImpl) basetypesListValueToPV1PortalProxyPathRewriteList(source basetypes.ListValue) (*[]v1.PortalProxyPathRewrite, error) {
	v1PortalProxyPathRewriteList, err := conversion.ListValueToSlicePortalProxyPathRewrite(source)
	if err != nil {
		return nil, err
	}
	return &v1PortalProxyPathRewriteList, nil
}

type PortalVersionConverterImpl struct{}

func (c *PortalVersionConverterImpl) ConvertToTfPortalVersion(source v1.PortalVersion) datasourceportalversion.PortalVersionModel {
	var datasource_portal_versionPortalVersionModel datasourceportalversion.PortalVersionModel
	datasource_portal_versionPortalVersionModel.CreatedAt = conversion.TimeToStringValue(source.CreatedAt)
	datasource_portal_versionPortalVersionModel.Id = conversion.UUIDToStringValue(source.Id)
	datasource_portal_versionPortalVersionModel.Major = conversion.IntToInt64Value(source.Major)
	datasource_portal_versionPortalVersionModel.Minor = conversion.IntToInt64Value(source.Minor)
	datasource_portal_versionPortalVersionModel.Patch = conversion.IntToInt64Value(source.Patch)
	datasource_portal_versionPortalVersionModel.Rev = conversion.IntToInt64Value(source.Rev)
	datasource_portal_versionPortalVersionModel.Version = conversion.StringToStringValue(source.Version)
	return datasource_portal_versionPortalVersionModel
}

type TenantConverterImpl struct{}

func (c *TenantConverterImpl) ConvertToApiTenant(source resourcetenant.TenantModel) v1.TenantInput {
	var v1TenantInput v1.TenantInput
	v1TenantInput.DisplayName = conversion.StringValueToString(source.DisplayName)
	v1TenantInput.Name = conversion.StringValueToString(source.Name)
	return v1TenantInput
}
func (c *TenantConverterImpl) ConvertToTfTenant(source v1.Tenant) resourcetenant.TenantModel {
	var resource_tenantTenantModel resourcetenant.TenantModel
	resource_tenantTenantModel.CreatedAt = conversion.TimeToStringValue(source.CreatedAt)
	resource_tenantTenantModel.DisplayName = conversion.StringToStringValue(source.DisplayName)
	resource_tenantTenantModel.Id = conversion.UUIDToStringValue(source.Id)
	resource_tenantTenantModel.Identifier = conversion.StringToStringValue(source.Identifier)
	resource_tenantTenantModel.IssuerUrl = conversion.StringToStringValue(source.IssuerUrl)
	resource_tenantTenantModel.Name = conversion.StringToStringValue(source.Name)
	resource_tenantTenantModel.OrganizationId = conversion.UUIDToStringValue(source.OrganizationId)
	return resource_tenantTenantModel
}

type TenantUserConverterImpl struct{}

func (c *TenantUserConverterImpl) ConvertToApiTenantUser(source resourcetenantuser.TenantUserModel) v1.TenantUserInput {
	var v1TenantUserInput v1.TenantUserInput
	v1TenantUserInput.Email = conversion.StringValueToEmail(source.Email)
	v1TenantUserInput.Username = conversion.StringValueToString(source.Username)
	return v1TenantUserInput
}
func (c *TenantUserConverterImpl) ConvertToTfTenantUser(source v1.TenantUser) resourcetenantuser.TenantUserModel {
	var resource_tenant_userTenantUserModel resourcetenantuser.TenantUserModel
	resource_tenant_userTenantUserModel.CreatedAt = conversion.TimeToStringValue(source.CreatedAt)
	resource_tenant_userTenantUserModel.Email = conversion.EmailToTfStringValue(source.Email)
	resource_tenant_userTenantUserModel.Id = conversion.UUIDToStringValue(source.Id)
	resource_tenant_userTenantUserModel.OrganizationId = conversion.UUIDToStringValue(source.OrganizationId)
	resource_tenant_userTenantUserModel.TenantName = conversion.StringToStringValue(source.TenantName)
	resource_tenant_userTenantUserModel.Username = conversion.StringToStringValue(source.Username)
	return resource_tenant_userTenantUserModel
}
