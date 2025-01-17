# Go API client for caliri

Candlepin is a subscription management server written in Java. It helps with management of software subscriptions.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 4.4.15
- Package version: 4.4.15
- Generator version: 7.8.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import caliri "github.com/content-services/caliri/release/v4"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `caliri.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), caliri.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `caliri.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), caliri.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `caliri.ContextOperationServerIndices` and `caliri.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), caliri.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), caliri.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to */candlepin*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ActivationKeyAPI* | [**AddActivationKeyContentOverrides**](docs/ActivationKeyAPI.md#addactivationkeycontentoverrides) | **Put** /activation_keys/{activation_key_id}/content_overrides | 
*ActivationKeyAPI* | [**AddPoolToKey**](docs/ActivationKeyAPI.md#addpooltokey) | **Post** /activation_keys/{activation_key_id}/pools/{pool_id} | 
*ActivationKeyAPI* | [**AddProductIdToKey**](docs/ActivationKeyAPI.md#addproductidtokey) | **Post** /activation_keys/{activation_key_id}/product/{product_id} | 
*ActivationKeyAPI* | [**DeleteActivationKey**](docs/ActivationKeyAPI.md#deleteactivationkey) | **Delete** /activation_keys/{activation_key_id} | 
*ActivationKeyAPI* | [**DeleteActivationKeyContentOverrides**](docs/ActivationKeyAPI.md#deleteactivationkeycontentoverrides) | **Delete** /activation_keys/{activation_key_id}/content_overrides | 
*ActivationKeyAPI* | [**FindActivationKey**](docs/ActivationKeyAPI.md#findactivationkey) | **Get** /activation_keys | 
*ActivationKeyAPI* | [**GetActivationKey**](docs/ActivationKeyAPI.md#getactivationkey) | **Get** /activation_keys/{activation_key_id} | 
*ActivationKeyAPI* | [**GetActivationKeyPools**](docs/ActivationKeyAPI.md#getactivationkeypools) | **Get** /activation_keys/{activation_key_id}/pools | 
*ActivationKeyAPI* | [**ListActivationKeyContentOverrides**](docs/ActivationKeyAPI.md#listactivationkeycontentoverrides) | **Get** /activation_keys/{activation_key_id}/content_overrides | 
*ActivationKeyAPI* | [**RemovePoolFromKey**](docs/ActivationKeyAPI.md#removepoolfromkey) | **Delete** /activation_keys/{activation_key_id}/pools/{pool_id} | 
*ActivationKeyAPI* | [**RemoveProductIdFromKey**](docs/ActivationKeyAPI.md#removeproductidfromkey) | **Delete** /activation_keys/{activation_key_id}/product/{product_id} | 
*ActivationKeyAPI* | [**UpdateActivationKey**](docs/ActivationKeyAPI.md#updateactivationkey) | **Put** /activation_keys/{activation_key_id} | 
*AdminAPI* | [**GetQueueStats**](docs/AdminAPI.md#getqueuestats) | **Get** /admin/queues | 
*AdminAPI* | [**Initialize**](docs/AdminAPI.md#initialize) | **Get** /admin/init | 
*CdnAPI* | [**CreateCdn**](docs/CdnAPI.md#createcdn) | **Post** /cdn | 
*CdnAPI* | [**DeleteCdn**](docs/CdnAPI.md#deletecdn) | **Delete** /cdn/{label} | 
*CdnAPI* | [**GetContentDeliveryNetworks**](docs/CdnAPI.md#getcontentdeliverynetworks) | **Get** /cdn | 
*CdnAPI* | [**UpdateCdn**](docs/CdnAPI.md#updatecdn) | **Put** /cdn/{label} | 
*CertificateRevocationListAPI* | [**GetCurrentCrl**](docs/CertificateRevocationListAPI.md#getcurrentcrl) | **Get** /crl | 
*CertificateSerialAPI* | [**GetCertificateSerial**](docs/CertificateSerialAPI.md#getcertificateserial) | **Get** /serials/{id} | 
*CloudRegistrationAPI* | [**CancelCloudAccountJobs**](docs/CloudRegistrationAPI.md#cancelcloudaccountjobs) | **Delete** /cloud/jobs/orgsetup/{cloud_account_id} | 
*CloudRegistrationAPI* | [**CloudAuthorize**](docs/CloudRegistrationAPI.md#cloudauthorize) | **Post** /cloud/authorize | 
*CloudRegistrationAPI* | [**DeleteAnonymousConsumersByAccountId**](docs/CloudRegistrationAPI.md#deleteanonymousconsumersbyaccountid) | **Delete** /cloud/consumers/anonymous | 
*ConsumerAPI* | [**AddConsumerContentOverrides**](docs/ConsumerAPI.md#addconsumercontentoverrides) | **Put** /consumers/{consumer_uuid}/content_overrides | 
*ConsumerAPI* | [**Bind**](docs/ConsumerAPI.md#bind) | **Post** /consumers/{consumer_uuid}/entitlements | 
*ConsumerAPI* | [**ConsumerExists**](docs/ConsumerAPI.md#consumerexists) | **Head** /consumers/{consumer_uuid}/exists | 
*ConsumerAPI* | [**ConsumerExistsBulk**](docs/ConsumerAPI.md#consumerexistsbulk) | **Post** /consumers/exists | 
*ConsumerAPI* | [**CreateConsumer**](docs/ConsumerAPI.md#createconsumer) | **Post** /consumers | 
*ConsumerAPI* | [**DeleteConsumer**](docs/ConsumerAPI.md#deleteconsumer) | **Delete** /consumers/{consumer_uuid} | 
*ConsumerAPI* | [**DeleteConsumerContentOverrides**](docs/ConsumerAPI.md#deleteconsumercontentoverrides) | **Delete** /consumers/{consumer_uuid}/content_overrides | 
*ConsumerAPI* | [**DownloadExistingExport**](docs/ConsumerAPI.md#downloadexistingexport) | **Get** /consumers/{consumer_uuid}/export/{export_id} | 
*ConsumerAPI* | [**DryBind**](docs/ConsumerAPI.md#drybind) | **Get** /consumers/{consumer_uuid}/entitlements/dry-run | 
*ConsumerAPI* | [**ExportCertificates**](docs/ConsumerAPI.md#exportcertificates) | **Get** /consumers/{consumer_uuid}/certificates | 
*ConsumerAPI* | [**ExportData**](docs/ConsumerAPI.md#exportdata) | **Get** /consumers/{consumer_uuid}/export | 
*ConsumerAPI* | [**ExportDataAsync**](docs/ConsumerAPI.md#exportdataasync) | **Get** /consumers/{consumer_uuid}/export/async | 
*ConsumerAPI* | [**GetComplianceStatus**](docs/ConsumerAPI.md#getcompliancestatus) | **Get** /consumers/{consumer_uuid}/compliance | 
*ConsumerAPI* | [**GetComplianceStatusList**](docs/ConsumerAPI.md#getcompliancestatuslist) | **Get** /consumers/compliance | 
*ConsumerAPI* | [**GetConsumer**](docs/ConsumerAPI.md#getconsumer) | **Get** /consumers/{consumer_uuid} | 
*ConsumerAPI* | [**GetContentAccessBody**](docs/ConsumerAPI.md#getcontentaccessbody) | **Get** /consumers/{consumer_uuid}/accessible_content | 
*ConsumerAPI* | [**GetContentAccessForConsumer**](docs/ConsumerAPI.md#getcontentaccessforconsumer) | **Get** /consumers/{consumer_uuid}/content_access | 
*ConsumerAPI* | [**GetEntitlementCertificateSerials**](docs/ConsumerAPI.md#getentitlementcertificateserials) | **Get** /consumers/{consumer_uuid}/certificates/serials | 
*ConsumerAPI* | [**GetGuests**](docs/ConsumerAPI.md#getguests) | **Get** /consumers/{consumer_uuid}/guests | 
*ConsumerAPI* | [**GetHost**](docs/ConsumerAPI.md#gethost) | **Get** /consumers/{consumer_uuid}/host | 
*ConsumerAPI* | [**GetOwnerByConsumerUuid**](docs/ConsumerAPI.md#getownerbyconsumeruuid) | **Get** /consumers/{consumer_uuid}/owner | 
*ConsumerAPI* | [**GetRelease**](docs/ConsumerAPI.md#getrelease) | **Get** /consumers/{consumer_uuid}/release | 
*ConsumerAPI* | [**GetSystemPurposeComplianceStatus**](docs/ConsumerAPI.md#getsystempurposecompliancestatus) | **Get** /consumers/{consumer_uuid}/purpose_compliance | 
*ConsumerAPI* | [**ListConsumerContentOverrides**](docs/ConsumerAPI.md#listconsumercontentoverrides) | **Get** /consumers/{consumer_uuid}/content_overrides | 
*ConsumerAPI* | [**ListEntitlements**](docs/ConsumerAPI.md#listentitlements) | **Get** /consumers/{consumer_uuid}/entitlements | 
*ConsumerAPI* | [**RegenerateEntitlementCertificates**](docs/ConsumerAPI.md#regenerateentitlementcertificates) | **Put** /consumers/{consumer_uuid}/certificates | 
*ConsumerAPI* | [**RegenerateIdentityCertificates**](docs/ConsumerAPI.md#regenerateidentitycertificates) | **Post** /consumers/{consumer_uuid} | 
*ConsumerAPI* | [**RemoveDeletionRecord**](docs/ConsumerAPI.md#removedeletionrecord) | **Delete** /consumers/{consumer_uuid}/deletionrecord | 
*ConsumerAPI* | [**SearchConsumers**](docs/ConsumerAPI.md#searchconsumers) | **Get** /consumers | 
*ConsumerAPI* | [**UnbindAll**](docs/ConsumerAPI.md#unbindall) | **Delete** /consumers/{consumer_uuid}/entitlements | 
*ConsumerAPI* | [**UnbindByEntitlementId**](docs/ConsumerAPI.md#unbindbyentitlementid) | **Delete** /consumers/{consumer_uuid}/entitlements/{dbid} | 
*ConsumerAPI* | [**UnbindByPool**](docs/ConsumerAPI.md#unbindbypool) | **Delete** /consumers/{consumer_uuid}/entitlements/pool/{pool_id} | 
*ConsumerAPI* | [**UnbindBySerial**](docs/ConsumerAPI.md#unbindbyserial) | **Delete** /consumers/{consumer_uuid}/certificates/{serial} | 
*ConsumerAPI* | [**UpdateConsumer**](docs/ConsumerAPI.md#updateconsumer) | **Put** /consumers/{consumer_uuid} | 
*ConsumerTypeAPI* | [**CreateConsumerType**](docs/ConsumerTypeAPI.md#createconsumertype) | **Post** /consumertypes | 
*ConsumerTypeAPI* | [**DeleteConsumerType**](docs/ConsumerTypeAPI.md#deleteconsumertype) | **Delete** /consumertypes/{id} | 
*ConsumerTypeAPI* | [**GetConsumerType**](docs/ConsumerTypeAPI.md#getconsumertype) | **Get** /consumertypes/{id} | 
*ConsumerTypeAPI* | [**GetConsumerTypes**](docs/ConsumerTypeAPI.md#getconsumertypes) | **Get** /consumertypes | 
*ConsumerTypeAPI* | [**UpdateConsumerType**](docs/ConsumerTypeAPI.md#updateconsumertype) | **Put** /consumertypes/{id} | 
*ContentAPI* | [**GetContentByUuid**](docs/ContentAPI.md#getcontentbyuuid) | **Get** /content/{content_uuid} | 
*ContentAPI* | [**GetContents**](docs/ContentAPI.md#getcontents) | **Get** /content | 
*DeletedConsumerAPI* | [**ListByDate**](docs/DeletedConsumerAPI.md#listbydate) | **Get** /deleted_consumers | 
*DistributorVersionsAPI* | [**Create**](docs/DistributorVersionsAPI.md#create) | **Post** /distributor_versions | 
*DistributorVersionsAPI* | [**Delete**](docs/DistributorVersionsAPI.md#delete) | **Delete** /distributor_versions/{id} | 
*DistributorVersionsAPI* | [**GetVersions**](docs/DistributorVersionsAPI.md#getversions) | **Get** /distributor_versions | 
*DistributorVersionsAPI* | [**Update**](docs/DistributorVersionsAPI.md#update) | **Put** /distributor_versions/{id} | 
*EntitlementsAPI* | [**GetEntitlement**](docs/EntitlementsAPI.md#getentitlement) | **Get** /entitlements/{entitlement_id} | 
*EntitlementsAPI* | [**GetUpstreamCert**](docs/EntitlementsAPI.md#getupstreamcert) | **Get** /entitlements/{dbid}/upstream_cert | 
*EntitlementsAPI* | [**HasEntitlement**](docs/EntitlementsAPI.md#hasentitlement) | **Get** /entitlements/consumer/{consumer_uuid}/product/{product_id} | 
*EntitlementsAPI* | [**ListAllForConsumer**](docs/EntitlementsAPI.md#listallforconsumer) | **Get** /entitlements | 
*EntitlementsAPI* | [**MigrateEntitlement**](docs/EntitlementsAPI.md#migrateentitlement) | **Put** /entitlements/{entitlement_id}/migrate | 
*EntitlementsAPI* | [**RegenerateEntitlementCertificatesForProduct**](docs/EntitlementsAPI.md#regenerateentitlementcertificatesforproduct) | **Put** /entitlements/product/{product_id} | 
*EntitlementsAPI* | [**Unbind**](docs/EntitlementsAPI.md#unbind) | **Delete** /entitlements/{dbid} | 
*EntitlementsAPI* | [**UpdateEntitlement**](docs/EntitlementsAPI.md#updateentitlement) | **Put** /entitlements/{entitlement_id} | 
*EnvironmentAPI* | [**CreateConsumerInEnvironment**](docs/EnvironmentAPI.md#createconsumerinenvironment) | **Post** /environments/{env_id}/consumers | 
*EnvironmentAPI* | [**DeleteEnvironment**](docs/EnvironmentAPI.md#deleteenvironment) | **Delete** /environments/{env_id} | 
*EnvironmentAPI* | [**DeleteEnvironmentContentOverrides**](docs/EnvironmentAPI.md#deleteenvironmentcontentoverrides) | **Delete** /environments/{environment_id}/content_overrides | 
*EnvironmentAPI* | [**DemoteContent**](docs/EnvironmentAPI.md#demotecontent) | **Delete** /environments/{env_id}/content | 
*EnvironmentAPI* | [**GetEnvironment**](docs/EnvironmentAPI.md#getenvironment) | **Get** /environments/{env_id} | 
*EnvironmentAPI* | [**GetEnvironmentContentOverrides**](docs/EnvironmentAPI.md#getenvironmentcontentoverrides) | **Get** /environments/{environment_id}/content_overrides | 
*EnvironmentAPI* | [**PromoteContent**](docs/EnvironmentAPI.md#promotecontent) | **Post** /environments/{env_id}/content | 
*EnvironmentAPI* | [**PutEnvironmentContentOverrides**](docs/EnvironmentAPI.md#putenvironmentcontentoverrides) | **Put** /environments/{environment_id}/content_overrides | 
*EnvironmentAPI* | [**UpdateEnvironment**](docs/EnvironmentAPI.md#updateenvironment) | **Put** /environments/{env_id} | 
*GuestIdsAPI* | [**DeleteGuest**](docs/GuestIdsAPI.md#deleteguest) | **Delete** /consumers/{consumer_uuid}/guestids/{guest_id} | 
*GuestIdsAPI* | [**GetGuestId**](docs/GuestIdsAPI.md#getguestid) | **Get** /consumers/{consumer_uuid}/guestids/{guest_id} | 
*GuestIdsAPI* | [**GetGuestIds**](docs/GuestIdsAPI.md#getguestids) | **Get** /consumers/{consumer_uuid}/guestids | 
*GuestIdsAPI* | [**UpdateGuest**](docs/GuestIdsAPI.md#updateguest) | **Put** /consumers/{consumer_uuid}/guestids/{guest_id} | 
*GuestIdsAPI* | [**UpdateGuests**](docs/GuestIdsAPI.md#updateguests) | **Put** /consumers/{consumer_uuid}/guestids | 
*HypervisorsAPI* | [**HypervisorHeartbeatUpdate**](docs/HypervisorsAPI.md#hypervisorheartbeatupdate) | **Put** /hypervisors/{owner}/heartbeat | 
*HypervisorsAPI* | [**HypervisorUpdateAsync**](docs/HypervisorsAPI.md#hypervisorupdateasync) | **Post** /hypervisors/{owner} | 
*JobsAPI* | [**CancelJob**](docs/JobsAPI.md#canceljob) | **Delete** /jobs/{id} | 
*JobsAPI* | [**CleanupTerminalJobs**](docs/JobsAPI.md#cleanupterminaljobs) | **Delete** /jobs | 
*JobsAPI* | [**GetJobStatus**](docs/JobsAPI.md#getjobstatus) | **Get** /jobs/{id} | 
*JobsAPI* | [**GetSchedulerStatus**](docs/JobsAPI.md#getschedulerstatus) | **Get** /jobs/scheduler | 
*JobsAPI* | [**ListJobStatuses**](docs/JobsAPI.md#listjobstatuses) | **Get** /jobs | 
*JobsAPI* | [**ScheduleJob**](docs/JobsAPI.md#schedulejob) | **Post** /jobs/schedule/{jobKey} | 
*JobsAPI* | [**SetSchedulerStatus**](docs/JobsAPI.md#setschedulerstatus) | **Post** /jobs/scheduler | 
*OwnerAPI* | [**Claim**](docs/OwnerAPI.md#claim) | **Put** /owners/{anonymous_owner_key}/claim | 
*OwnerAPI* | [**CountConsumers**](docs/OwnerAPI.md#countconsumers) | **Get** /owners/{owner_key}/consumers/count | 
*OwnerAPI* | [**CreateActivationKey**](docs/OwnerAPI.md#createactivationkey) | **Post** /owners/{owner_key}/activation_keys | 
*OwnerAPI* | [**CreateEnvironment**](docs/OwnerAPI.md#createenvironment) | **Post** /owners/{owner_key}/environments | 
*OwnerAPI* | [**CreateOwner**](docs/OwnerAPI.md#createowner) | **Post** /owners | 
*OwnerAPI* | [**CreatePool**](docs/OwnerAPI.md#createpool) | **Post** /owners/{owner_key}/pools | 
*OwnerAPI* | [**CreateUeberCertificate**](docs/OwnerAPI.md#createuebercertificate) | **Post** /owners/{owner_key}/uebercert | 
*OwnerAPI* | [**DeleteLogLevel**](docs/OwnerAPI.md#deleteloglevel) | **Delete** /owners/{owner_key}/log | 
*OwnerAPI* | [**DeleteOwner**](docs/OwnerAPI.md#deleteowner) | **Delete** /owners/{owner_key} | 
*OwnerAPI* | [**GetConsumersSyspurpose**](docs/OwnerAPI.md#getconsumerssyspurpose) | **Get** /owners/{owner_key}/consumers_system_purpose | 
*OwnerAPI* | [**GetHypervisors**](docs/OwnerAPI.md#gethypervisors) | **Get** /owners/{owner_key}/hypervisors | 
*OwnerAPI* | [**GetImports**](docs/OwnerAPI.md#getimports) | **Get** /owners/{owner_key}/imports | 
*OwnerAPI* | [**GetOwner**](docs/OwnerAPI.md#getowner) | **Get** /owners/{owner_key} | 
*OwnerAPI* | [**GetOwnerContentAccess**](docs/OwnerAPI.md#getownercontentaccess) | **Get** /owners/{owner_key}/content_access | 
*OwnerAPI* | [**GetOwnerInfo**](docs/OwnerAPI.md#getownerinfo) | **Get** /owners/{owner_key}/info | 
*OwnerAPI* | [**GetOwnerSubscriptions**](docs/OwnerAPI.md#getownersubscriptions) | **Get** /owners/{owner_key}/subscriptions | 
*OwnerAPI* | [**GetSyspurpose**](docs/OwnerAPI.md#getsyspurpose) | **Get** /owners/{owner_key}/system_purpose | 
*OwnerAPI* | [**GetUeberCertificate**](docs/OwnerAPI.md#getuebercertificate) | **Get** /owners/{owner_key}/uebercert | 
*OwnerAPI* | [**GetUpstreamConsumers**](docs/OwnerAPI.md#getupstreamconsumers) | **Get** /owners/{owner_key}/upstream_consumers | 
*OwnerAPI* | [**HealEntire**](docs/OwnerAPI.md#healentire) | **Post** /owners/{owner_key}/entitlements | 
*OwnerAPI* | [**ImportManifestAsync**](docs/OwnerAPI.md#importmanifestasync) | **Post** /owners/{owner_key}/imports/async | 
*OwnerAPI* | [**ListConsumers**](docs/OwnerAPI.md#listconsumers) | **Get** /owners/{owner_key}/consumers | 
*OwnerAPI* | [**ListEnvironments**](docs/OwnerAPI.md#listenvironments) | **Get** /owners/{owner_key}/environments | 
*OwnerAPI* | [**ListOwnerPools**](docs/OwnerAPI.md#listownerpools) | **Get** /owners/{owner_key}/pools | 
*OwnerAPI* | [**ListOwners**](docs/OwnerAPI.md#listowners) | **Get** /owners | 
*OwnerAPI* | [**OwnerActivationKeys**](docs/OwnerAPI.md#owneractivationkeys) | **Get** /owners/{owner_key}/activation_keys | 
*OwnerAPI* | [**OwnerEntitlements**](docs/OwnerAPI.md#ownerentitlements) | **Get** /owners/{owner_key}/entitlements | 
*OwnerAPI* | [**OwnerServiceLevels**](docs/OwnerAPI.md#ownerservicelevels) | **Get** /owners/{owner_key}/servicelevels | 
*OwnerAPI* | [**RefreshPools**](docs/OwnerAPI.md#refreshpools) | **Put** /owners/{owner_key}/subscriptions | 
*OwnerAPI* | [**SetLogLevel**](docs/OwnerAPI.md#setloglevel) | **Put** /owners/{owner_key}/log | 
*OwnerAPI* | [**UndoImports**](docs/OwnerAPI.md#undoimports) | **Delete** /owners/{owner_key}/imports | 
*OwnerAPI* | [**UpdateOwner**](docs/OwnerAPI.md#updateowner) | **Put** /owners/{owner_key} | 
*OwnerAPI* | [**UpdatePool**](docs/OwnerAPI.md#updatepool) | **Put** /owners/{owner_key}/pools | 
*OwnerContentAPI* | [**CreateContent**](docs/OwnerContentAPI.md#createcontent) | **Post** /owners/{owner_key}/content | 
*OwnerContentAPI* | [**CreateContentBatch**](docs/OwnerContentAPI.md#createcontentbatch) | **Post** /owners/{owner_key}/content/batch | 
*OwnerContentAPI* | [**GetContentById**](docs/OwnerContentAPI.md#getcontentbyid) | **Get** /owners/{owner_key}/content/{content_id} | 
*OwnerContentAPI* | [**GetContentsByOwner**](docs/OwnerContentAPI.md#getcontentsbyowner) | **Get** /owners/{owner_key}/content | 
*OwnerContentAPI* | [**RemoveContent**](docs/OwnerContentAPI.md#removecontent) | **Delete** /owners/{owner_key}/content/{content_id} | 
*OwnerContentAPI* | [**UpdateContent**](docs/OwnerContentAPI.md#updatecontent) | **Put** /owners/{owner_key}/content/{content_id} | 
*OwnerProductAPI* | [**AddContentToProduct**](docs/OwnerProductAPI.md#addcontenttoproduct) | **Post** /owners/{owner_key}/products/{product_id}/content/{content_id} | 
*OwnerProductAPI* | [**AddContentsToProduct**](docs/OwnerProductAPI.md#addcontentstoproduct) | **Post** /owners/{owner_key}/products/{product_id}/batch_content | 
*OwnerProductAPI* | [**CreateProduct**](docs/OwnerProductAPI.md#createproduct) | **Post** /owners/{owner_key}/products | 
*OwnerProductAPI* | [**GetProductById**](docs/OwnerProductAPI.md#getproductbyid) | **Get** /owners/{owner_key}/products/{product_id} | 
*OwnerProductAPI* | [**GetProductCertificateById**](docs/OwnerProductAPI.md#getproductcertificatebyid) | **Get** /owners/{owner_key}/products/{product_id}/certificate | 
*OwnerProductAPI* | [**GetProductsByOwner**](docs/OwnerProductAPI.md#getproductsbyowner) | **Get** /owners/{owner_key}/products | 
*OwnerProductAPI* | [**RefreshPoolsForProduct**](docs/OwnerProductAPI.md#refreshpoolsforproduct) | **Put** /owners/{owner_key}/products/{product_id}/subscriptions | 
*OwnerProductAPI* | [**RemoveContentFromProduct**](docs/OwnerProductAPI.md#removecontentfromproduct) | **Delete** /owners/{owner_key}/products/{product_id}/content/{content_id} | 
*OwnerProductAPI* | [**RemoveContentsFromProduct**](docs/OwnerProductAPI.md#removecontentsfromproduct) | **Delete** /owners/{owner_key}/products/{product_id}/batch_content | 
*OwnerProductAPI* | [**RemoveProduct**](docs/OwnerProductAPI.md#removeproduct) | **Delete** /owners/{owner_key}/products/{product_id} | 
*OwnerProductAPI* | [**UpdateProduct**](docs/OwnerProductAPI.md#updateproduct) | **Put** /owners/{owner_key}/products/{product_id} | 
*PoolsAPI* | [**DeletePool**](docs/PoolsAPI.md#deletepool) | **Delete** /pools/{pool_id} | 
*PoolsAPI* | [**GetPool**](docs/PoolsAPI.md#getpool) | **Get** /pools/{pool_id} | 
*PoolsAPI* | [**GetPoolCdn**](docs/PoolsAPI.md#getpoolcdn) | **Get** /pools/{pool_id}/cdn | 
*PoolsAPI* | [**GetPoolEntitlements**](docs/PoolsAPI.md#getpoolentitlements) | **Get** /pools/{pool_id}/entitlements | 
*PoolsAPI* | [**GetSubCert**](docs/PoolsAPI.md#getsubcert) | **Get** /pools/{pool_id}/cert | 
*PoolsAPI* | [**ListEntitledConsumerUuids**](docs/PoolsAPI.md#listentitledconsumeruuids) | **Get** /pools/{pool_id}/entitlements/consumer_uuids | 
*PoolsAPI* | [**ListPools**](docs/PoolsAPI.md#listpools) | **Get** /pools | 
*ProductsAPI* | [**GetProductByUuid**](docs/ProductsAPI.md#getproductbyuuid) | **Get** /products/{product_uuid} | 
*ProductsAPI* | [**GetProducts**](docs/ProductsAPI.md#getproducts) | **Get** /products | 
*ProductsAPI* | [**RefreshPoolsForProducts**](docs/ProductsAPI.md#refreshpoolsforproducts) | **Put** /products/subscriptions | 
*RolesAPI* | [**AddRolePermission**](docs/RolesAPI.md#addrolepermission) | **Post** /roles/{role_name}/permissions | 
*RolesAPI* | [**AddUserToRole**](docs/RolesAPI.md#addusertorole) | **Post** /roles/{role_name}/users/{username} | 
*RolesAPI* | [**CreateRole**](docs/RolesAPI.md#createrole) | **Post** /roles | 
*RolesAPI* | [**DeleteRoleByName**](docs/RolesAPI.md#deleterolebyname) | **Delete** /roles/{role_name} | 
*RolesAPI* | [**DeleteUserFromRole**](docs/RolesAPI.md#deleteuserfromrole) | **Delete** /roles/{role_name}/users/{username} | 
*RolesAPI* | [**GetRoleByName**](docs/RolesAPI.md#getrolebyname) | **Get** /roles/{role_name} | 
*RolesAPI* | [**GetRoles**](docs/RolesAPI.md#getroles) | **Get** /roles | 
*RolesAPI* | [**RemoveRolePermission**](docs/RolesAPI.md#removerolepermission) | **Delete** /roles/{role_name}/permissions/{perm_id} | 
*RolesAPI* | [**UpdateRole**](docs/RolesAPI.md#updaterole) | **Put** /roles/{role_name} | 
*RootAPI* | [**GetRootResources**](docs/RootAPI.md#getrootresources) | **Get** / | 
*RulesAPI* | [**GetRules**](docs/RulesAPI.md#getrules) | **Get** /rules | 
*StatusAPI* | [**Status**](docs/StatusAPI.md#status) | **Get** /status | 
*SubscriptionAPI* | [**ActivateSubscription**](docs/SubscriptionAPI.md#activatesubscription) | **Post** /subscriptions | 
*SubscriptionAPI* | [**DeleteSubscription**](docs/SubscriptionAPI.md#deletesubscription) | **Delete** /subscriptions/{id} | 
*SubscriptionAPI* | [**GetSubscriptions**](docs/SubscriptionAPI.md#getsubscriptions) | **Get** /subscriptions | 
*UsersAPI* | [**CreateUser**](docs/UsersAPI.md#createuser) | **Post** /users | 
*UsersAPI* | [**DeleteUser**](docs/UsersAPI.md#deleteuser) | **Delete** /users/{username} | 
*UsersAPI* | [**GetUserInfo**](docs/UsersAPI.md#getuserinfo) | **Get** /users/{username} | 
*UsersAPI* | [**GetUserRoles**](docs/UsersAPI.md#getuserroles) | **Get** /users/{username}/roles | 
*UsersAPI* | [**ListUserOwners**](docs/UsersAPI.md#listuserowners) | **Get** /users/{username}/owners | 
*UsersAPI* | [**ListUsers**](docs/UsersAPI.md#listusers) | **Get** /users | 
*UsersAPI* | [**UpdateUser**](docs/UsersAPI.md#updateuser) | **Put** /users/{username} | 


## Documentation For Models

 - [AbstractCertificateDTO](docs/AbstractCertificateDTO.md)
 - [ActivationKeyDTO](docs/ActivationKeyDTO.md)
 - [ActivationKeyPoolDTO](docs/ActivationKeyPoolDTO.md)
 - [ActivationKeyProductDTO](docs/ActivationKeyProductDTO.md)
 - [AsyncJobStatusDTO](docs/AsyncJobStatusDTO.md)
 - [AttributeDTO](docs/AttributeDTO.md)
 - [BrandingDTO](docs/BrandingDTO.md)
 - [CapabilityDTO](docs/CapabilityDTO.md)
 - [CdnDTO](docs/CdnDTO.md)
 - [CertificateDTO](docs/CertificateDTO.md)
 - [CertificateSerialDTO](docs/CertificateSerialDTO.md)
 - [ClaimantOwner](docs/ClaimantOwner.md)
 - [CloudAuthenticationResultDTO](docs/CloudAuthenticationResultDTO.md)
 - [CloudRegistrationDTO](docs/CloudRegistrationDTO.md)
 - [ComplianceReasonDTO](docs/ComplianceReasonDTO.md)
 - [ComplianceStatusDTO](docs/ComplianceStatusDTO.md)
 - [ConsumerActivationKeyDTO](docs/ConsumerActivationKeyDTO.md)
 - [ConsumerDTO](docs/ConsumerDTO.md)
 - [ConsumerDTOArrayElement](docs/ConsumerDTOArrayElement.md)
 - [ConsumerInstalledProductDTO](docs/ConsumerInstalledProductDTO.md)
 - [ConsumerTypeDTO](docs/ConsumerTypeDTO.md)
 - [ConsumptionTypeCountsDTO](docs/ConsumptionTypeCountsDTO.md)
 - [ContentAccessDTO](docs/ContentAccessDTO.md)
 - [ContentDTO](docs/ContentDTO.md)
 - [ContentOverrideDTO](docs/ContentOverrideDTO.md)
 - [ContentToPromoteDTO](docs/ContentToPromoteDTO.md)
 - [DateRange](docs/DateRange.md)
 - [DeleteResult](docs/DeleteResult.md)
 - [DeletedConsumerDTO](docs/DeletedConsumerDTO.md)
 - [DistributorVersionCapabilityDTO](docs/DistributorVersionCapabilityDTO.md)
 - [DistributorVersionDTO](docs/DistributorVersionDTO.md)
 - [EntitlementDTO](docs/EntitlementDTO.md)
 - [EnvironmentContentDTO](docs/EnvironmentContentDTO.md)
 - [EnvironmentDTO](docs/EnvironmentDTO.md)
 - [ExceptionMessage](docs/ExceptionMessage.md)
 - [ExportResultDTO](docs/ExportResultDTO.md)
 - [GuestIdDTO](docs/GuestIdDTO.md)
 - [GuestIdDTOArrayElement](docs/GuestIdDTOArrayElement.md)
 - [HypervisorConsumerDTO](docs/HypervisorConsumerDTO.md)
 - [HypervisorIdDTO](docs/HypervisorIdDTO.md)
 - [HypervisorUpdateResultDTO](docs/HypervisorUpdateResultDTO.md)
 - [ImportRecordDTO](docs/ImportRecordDTO.md)
 - [ImportUpstreamConsumerDTO](docs/ImportUpstreamConsumerDTO.md)
 - [Link](docs/Link.md)
 - [NestedConsumerDTO](docs/NestedConsumerDTO.md)
 - [NestedEntitlementDTO](docs/NestedEntitlementDTO.md)
 - [NestedOwnerDTO](docs/NestedOwnerDTO.md)
 - [OwnerDTO](docs/OwnerDTO.md)
 - [OwnerInfo](docs/OwnerInfo.md)
 - [PermissionBlueprintDTO](docs/PermissionBlueprintDTO.md)
 - [PoolDTO](docs/PoolDTO.md)
 - [PoolQuantityDTO](docs/PoolQuantityDTO.md)
 - [ProductCertificateDTO](docs/ProductCertificateDTO.md)
 - [ProductContentDTO](docs/ProductContentDTO.md)
 - [ProductDTO](docs/ProductDTO.md)
 - [ProvidedProductDTO](docs/ProvidedProductDTO.md)
 - [QueueStatus](docs/QueueStatus.md)
 - [ReleaseVerDTO](docs/ReleaseVerDTO.md)
 - [RoleDTO](docs/RoleDTO.md)
 - [SchedulerStatusDTO](docs/SchedulerStatusDTO.md)
 - [StatusDTO](docs/StatusDTO.md)
 - [SubscriptionDTO](docs/SubscriptionDTO.md)
 - [SystemPurposeAttributesDTO](docs/SystemPurposeAttributesDTO.md)
 - [SystemPurposeComplianceStatusDTO](docs/SystemPurposeComplianceStatusDTO.md)
 - [TimestampedEntity](docs/TimestampedEntity.md)
 - [UeberCertificateDTO](docs/UeberCertificateDTO.md)
 - [UpstreamConsumerDTO](docs/UpstreamConsumerDTO.md)
 - [UpstreamConsumerDTOArrayElement](docs/UpstreamConsumerDTOArrayElement.md)
 - [UserDTO](docs/UserDTO.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### basicAuth

- **Type**: HTTP basic authentication

Example

```go
auth := context.WithValue(context.Background(), caliri.ContextBasicAuth, caliri.BasicAuth{
	UserName: "username",
	Password: "password",
})
r, err := client.Service.Operation(auth, args)
```

### ActivationKey

- **Type**: API key
- **API key parameter name**: activation-key
- **Location**: URL query string

Note, each API key must be added to a map of `map[string]APIKey` where the key is: activation-key and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		caliri.ContextAPIKeys,
		map[string]caliri.APIKey{
			"activation-key": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### ActivationKeyOwner

- **Type**: API key
- **API key parameter name**: owner
- **Location**: URL query string

Note, each API key must be added to a map of `map[string]APIKey` where the key is: owner and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		caliri.ContextAPIKeys,
		map[string]caliri.APIKey{
			"owner": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



