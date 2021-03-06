# Go API client for swagger

Rainbond open api

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *http://0.0.0.0:7070*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*OpenapiApi* | [**OpenapiV1TeamsRegionsServicesDelete**](docs/OpenapiApi.md#openapiv1teamsregionsservicesdelete) | **Delete** /openapi/v1/teams/{team_id}/regions/{region_name}/services | 
*OpenapiAnnouncementApi* | [**OpenapiV1AnnouncementsCreate**](docs/OpenapiAnnouncementApi.md#openapiv1announcementscreate) | **Post** /openapi/v1/announcements | 
*OpenapiAnnouncementApi* | [**OpenapiV1AnnouncementsDelete**](docs/OpenapiAnnouncementApi.md#openapiv1announcementsdelete) | **Delete** /openapi/v1/announcements/{aid} | 
*OpenapiAnnouncementApi* | [**OpenapiV1AnnouncementsList**](docs/OpenapiAnnouncementApi.md#openapiv1announcementslist) | **Get** /openapi/v1/announcements | 
*OpenapiAnnouncementApi* | [**OpenapiV1AnnouncementsUpdate**](docs/OpenapiAnnouncementApi.md#openapiv1announcementsupdate) | **Put** /openapi/v1/announcements/{aid} | 
*OpenapiAppsApi* | [**OpenapiV1TeamsRegionsAppsCopyCreate**](docs/OpenapiAppsApi.md#openapiv1teamsregionsappscopycreate) | **Post** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/copy | 
*OpenapiAppsApi* | [**OpenapiV1TeamsRegionsAppsCopyList**](docs/OpenapiAppsApi.md#openapiv1teamsregionsappscopylist) | **Get** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/copy | 
*OpenapiAppsApi* | [**OpenapiV1TeamsRegionsAppsCreate**](docs/OpenapiAppsApi.md#openapiv1teamsregionsappscreate) | **Post** /openapi/v1/teams/{team_id}/regions/{region_name}/apps | 
*OpenapiAppsApi* | [**OpenapiV1TeamsRegionsAppsDelete**](docs/OpenapiAppsApi.md#openapiv1teamsregionsappsdelete) | **Delete** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id} | 
*OpenapiAppsApi* | [**OpenapiV1TeamsRegionsAppsHttpdomainsCreate**](docs/OpenapiAppsApi.md#openapiv1teamsregionsappshttpdomainscreate) | **Post** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/httpdomains | 
*OpenapiAppsApi* | [**OpenapiV1TeamsRegionsAppsInstallCreate**](docs/OpenapiAppsApi.md#openapiv1teamsregionsappsinstallcreate) | **Post** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/install | 
*OpenapiAppsApi* | [**OpenapiV1TeamsRegionsAppsList**](docs/OpenapiAppsApi.md#openapiv1teamsregionsappslist) | **Get** /openapi/v1/teams/{team_id}/regions/{region_name}/apps | 
*OpenapiAppsApi* | [**OpenapiV1TeamsRegionsAppsOperationsCreate**](docs/OpenapiAppsApi.md#openapiv1teamsregionsappsoperationscreate) | **Post** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/operations | 
*OpenapiAppsApi* | [**OpenapiV1TeamsRegionsAppsRead**](docs/OpenapiAppsApi.md#openapiv1teamsregionsappsread) | **Get** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id} | 
*OpenapiAppsApi* | [**OpenapiV1TeamsRegionsAppsServicesDelete**](docs/OpenapiAppsApi.md#openapiv1teamsregionsappsservicesdelete) | **Delete** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/services/{service_id} | 
*OpenapiAppsApi* | [**OpenapiV1TeamsRegionsAppsServicesEventsList**](docs/OpenapiAppsApi.md#openapiv1teamsregionsappsserviceseventslist) | **Get** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/services/{service_id}/events | 
*OpenapiAppsApi* | [**OpenapiV1TeamsRegionsAppsServicesList**](docs/OpenapiAppsApi.md#openapiv1teamsregionsappsserviceslist) | **Get** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/services | 
*OpenapiAppsApi* | [**OpenapiV1TeamsRegionsAppsServicesRead**](docs/OpenapiAppsApi.md#openapiv1teamsregionsappsservicesread) | **Get** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/services/{service_id} | 
*OpenapiAppstoreApi* | [**OpenapiV1AppstoresList**](docs/OpenapiAppstoreApi.md#openapiv1appstoreslist) | **Get** /openapi/v1/appstores | 
*OpenapiAppstoreApi* | [**OpenapiV1AppstoresUpdate**](docs/OpenapiAppstoreApi.md#openapiv1appstoresupdate) | **Put** /openapi/v1/appstores/{eid} | 
*OpenapiAuthApi* | [**OpenapiV1AuthTokenCreate**](docs/OpenapiAuthApi.md#openapiv1authtokencreate) | **Post** /openapi/v1/auth-token | 
*OpenapiEntrepriseApi* | [**OpenapiV1ConfigsList**](docs/OpenapiEntrepriseApi.md#openapiv1configslist) | **Get** /openapi/v1/configs | 
*OpenapiEntrepriseApi* | [**OpenapiV1EnterprisesList**](docs/OpenapiEntrepriseApi.md#openapiv1enterpriseslist) | **Get** /openapi/v1/enterprises | 
*OpenapiEntrepriseApi* | [**OpenapiV1EnterprisesRead**](docs/OpenapiEntrepriseApi.md#openapiv1enterprisesread) | **Get** /openapi/v1/enterprises/{eid} | 
*OpenapiEntrepriseApi* | [**OpenapiV1EnterprisesResourceList**](docs/OpenapiEntrepriseApi.md#openapiv1enterprisesresourcelist) | **Get** /openapi/v1/enterprises/{eid}/resource | 
*OpenapiEntrepriseApi* | [**OpenapiV1EnterprisesUpdate**](docs/OpenapiEntrepriseApi.md#openapiv1enterprisesupdate) | **Put** /openapi/v1/enterprises/{eid} | 
*OpenapiGatewayApi* | [**OpenapiV1HttpdomainsList**](docs/OpenapiGatewayApi.md#openapiv1httpdomainslist) | **Get** /openapi/v1/httpdomains | 
*OpenapiGatewayApi* | [**OpenapiV1TeamsRegionsAppsHttpdomainsDelete**](docs/OpenapiGatewayApi.md#openapiv1teamsregionsappshttpdomainsdelete) | **Delete** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/httpdomains/{rule_id} | 
*OpenapiGatewayApi* | [**OpenapiV1TeamsRegionsAppsHttpdomainsList**](docs/OpenapiGatewayApi.md#openapiv1teamsregionsappshttpdomainslist) | **Get** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/httpdomains | 
*OpenapiGatewayApi* | [**OpenapiV1TeamsRegionsAppsHttpdomainsUpdate**](docs/OpenapiGatewayApi.md#openapiv1teamsregionsappshttpdomainsupdate) | **Put** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/httpdomains/{rule_id} | 
*OpenapiOauthApi* | [**OpenapiV1OauthTypeList**](docs/OpenapiOauthApi.md#openapiv1oauthtypelist) | **Get** /openapi/v1/oauth/type | 
*OpenapiRegionApi* | [**OpenapiV1RegionsCreate**](docs/OpenapiRegionApi.md#openapiv1regionscreate) | **Post** /openapi/v1/regions | 
*OpenapiRegionApi* | [**OpenapiV1RegionsDelete**](docs/OpenapiRegionApi.md#openapiv1regionsdelete) | **Delete** /openapi/v1/regions/{region_id} | 
*OpenapiRegionApi* | [**OpenapiV1RegionsList**](docs/OpenapiRegionApi.md#openapiv1regionslist) | **Get** /openapi/v1/regions | 
*OpenapiRegionApi* | [**OpenapiV1RegionsRead**](docs/OpenapiRegionApi.md#openapiv1regionsread) | **Get** /openapi/v1/regions/{region_id} | 
*OpenapiRegionApi* | [**OpenapiV1RegionsStatusUpdate**](docs/OpenapiRegionApi.md#openapiv1regionsstatusupdate) | **Put** /openapi/v1/regions/{region_id}/status | 
*OpenapiRegionApi* | [**OpenapiV1RegionsUpdate**](docs/OpenapiRegionApi.md#openapiv1regionsupdate) | **Put** /openapi/v1/regions/{region_id} | 
*OpenapiTeamApi* | [**OpenapiV1TeamsCertificatesCreate**](docs/OpenapiTeamApi.md#openapiv1teamscertificatescreate) | **Post** /openapi/v1/teams/{team_id}/certificates | 
*OpenapiTeamApi* | [**OpenapiV1TeamsCertificatesDelete**](docs/OpenapiTeamApi.md#openapiv1teamscertificatesdelete) | **Delete** /openapi/v1/teams/{team_id}/certificates/{certificate_id} | 
*OpenapiTeamApi* | [**OpenapiV1TeamsCertificatesList**](docs/OpenapiTeamApi.md#openapiv1teamscertificateslist) | **Get** /openapi/v1/teams/{team_id}/certificates | 
*OpenapiTeamApi* | [**OpenapiV1TeamsCertificatesRead**](docs/OpenapiTeamApi.md#openapiv1teamscertificatesread) | **Get** /openapi/v1/teams/{team_id}/certificates/{certificate_id} | 
*OpenapiTeamApi* | [**OpenapiV1TeamsCertificatesUpdate**](docs/OpenapiTeamApi.md#openapiv1teamscertificatesupdate) | **Put** /openapi/v1/teams/{team_id}/certificates/{certificate_id} | 
*OpenapiTeamApi* | [**OpenapiV1TeamsCreate**](docs/OpenapiTeamApi.md#openapiv1teamscreate) | **Post** /openapi/v1/teams | 
*OpenapiTeamApi* | [**OpenapiV1TeamsDelete**](docs/OpenapiTeamApi.md#openapiv1teamsdelete) | **Delete** /openapi/v1/teams/{team_id} | 
*OpenapiTeamApi* | [**OpenapiV1TeamsList**](docs/OpenapiTeamApi.md#openapiv1teamslist) | **Get** /openapi/v1/teams | 
*OpenapiTeamApi* | [**OpenapiV1TeamsRead**](docs/OpenapiTeamApi.md#openapiv1teamsread) | **Get** /openapi/v1/teams/{team_id} | 
*OpenapiTeamApi* | [**OpenapiV1TeamsRegionsServicesList**](docs/OpenapiTeamApi.md#openapiv1teamsregionsserviceslist) | **Get** /openapi/v1/teams/{team_id}/regions/{region_name}/services | 
*OpenapiTeamApi* | [**OpenapiV1TeamsUpdate**](docs/OpenapiTeamApi.md#openapiv1teamsupdate) | **Put** /openapi/v1/teams/{team_id} | 
*OpenapiTeamApi* | [**OpenapiV1TeamsUsersCreate**](docs/OpenapiTeamApi.md#openapiv1teamsuserscreate) | **Post** /openapi/v1/teams/{team_id}/users/{user_id} | 
*OpenapiTeamApi* | [**OpenapiV1TeamsUsersDelete**](docs/OpenapiTeamApi.md#openapiv1teamsusersdelete) | **Delete** /openapi/v1/teams/{team_id}/users/{user_id} | 
*OpenapiTeamApi* | [**OpenapiV1TeamsUsersList**](docs/OpenapiTeamApi.md#openapiv1teamsuserslist) | **Get** /openapi/v1/teams/{team_id}/users | 
*OpenapiTeamApi* | [**OpenapiV1TeamsUsersUpdate**](docs/OpenapiTeamApi.md#openapiv1teamsusersupdate) | **Put** /openapi/v1/teams/{team_id}/users/{user_id} | 
*OpenapiTeamRegionApi* | [**OpenapiV1TeamsRegionsCreate**](docs/OpenapiTeamRegionApi.md#openapiv1teamsregionscreate) | **Post** /openapi/v1/teams/{team_id}/regions | 
*OpenapiTeamRegionApi* | [**OpenapiV1TeamsRegionsDelete**](docs/OpenapiTeamRegionApi.md#openapiv1teamsregionsdelete) | **Delete** /openapi/v1/teams/{team_id}/regions/{region_name} | 
*OpenapiTeamRegionApi* | [**OpenapiV1TeamsRegionsList**](docs/OpenapiTeamRegionApi.md#openapiv1teamsregionslist) | **Get** /openapi/v1/teams/{team_id}/regions | 
*OpenapiUploadApi* | [**OpenapiV1UploadFileCreate**](docs/OpenapiUploadApi.md#openapiv1uploadfilecreate) | **Post** /openapi/v1/upload-file | 
*OpenapiUserApi* | [**OpenapiV1AdministratorsCreate**](docs/OpenapiUserApi.md#openapiv1administratorscreate) | **Post** /openapi/v1/administrators | 
*OpenapiUserApi* | [**OpenapiV1AdministratorsList**](docs/OpenapiUserApi.md#openapiv1administratorslist) | **Get** /openapi/v1/administrators | 
*OpenapiUserApi* | [**OpenapiV1UserChangepwdUpdate**](docs/OpenapiUserApi.md#openapiv1userchangepwdupdate) | **Put** /openapi/v1/user/changepwd | 
*OpenapiUserApi* | [**OpenapiV1UsersAdministratorDelete**](docs/OpenapiUserApi.md#openapiv1usersadministratordelete) | **Delete** /openapi/v1/users/{user_id}/administrator | 
*OpenapiUserApi* | [**OpenapiV1UsersCreate**](docs/OpenapiUserApi.md#openapiv1userscreate) | **Post** /openapi/v1/users | 
*OpenapiUserApi* | [**OpenapiV1UsersDelete**](docs/OpenapiUserApi.md#openapiv1usersdelete) | **Delete** /openapi/v1/users/{user_id} | 
*OpenapiUserApi* | [**OpenapiV1UsersList**](docs/OpenapiUserApi.md#openapiv1userslist) | **Get** /openapi/v1/users | 
*OpenapiUserApi* | [**OpenapiV1UsersRead**](docs/OpenapiUserApi.md#openapiv1usersread) | **Get** /openapi/v1/users/{user_id} | 
*OpenapiUserApi* | [**OpenapiV1UsersTeamsList**](docs/OpenapiUserApi.md#openapiv1usersteamslist) | **Get** /openapi/v1/users/{user_id}/teams | 
*OpenapiUserApi* | [**OpenapiV1UsersUpdate**](docs/OpenapiUserApi.md#openapiv1usersupdate) | **Put** /openapi/v1/users/{user_id} | 


## Documentation For Models

 - [AddRegionRequest](docs/AddRegionRequest.md)
 - [AnnouncementRespSerilizer](docs/AnnouncementRespSerilizer.md)
 - [AppBaseInfo](docs/AppBaseInfo.md)
 - [AppCopyC](docs/AppCopyC.md)
 - [AppCopyCRes](docs/AppCopyCRes.md)
 - [AppCopyL](docs/AppCopyL.md)
 - [AppCopyModify](docs/AppCopyModify.md)
 - [AppInfo](docs/AppInfo.md)
 - [AppPostInfo](docs/AppPostInfo.md)
 - [AppServiceEvents](docs/AppServiceEvents.md)
 - [AppStoreImageHubBaseResp](docs/AppStoreImageHubBaseResp.md)
 - [AppStoreImageHubResp](docs/AppStoreImageHubResp.md)
 - [AppStoreInfo](docs/AppStoreInfo.md)
 - [AuthRequest](docs/AuthRequest.md)
 - [AutoSsl](docs/AutoSsl.md)
 - [CertificatesR](docs/CertificatesR.md)
 - [ChangePassWdUser](docs/ChangePassWdUser.md)
 - [CloudMarketBaseResp](docs/CloudMarketBaseResp.md)
 - [CreateAdminUserReq](docs/CreateAdminUserReq.md)
 - [CreateAncmReqSerilizer](docs/CreateAncmReqSerilizer.md)
 - [CreateTeamReq](docs/CreateTeamReq.md)
 - [CreateTeamUserReq](docs/CreateTeamUserReq.md)
 - [CreateUser](docs/CreateUser.md)
 - [CustomJwt](docs/CustomJwt.md)
 - [EnterpriseConfigSeralizer](docs/EnterpriseConfigSeralizer.md)
 - [EnterpriseHttpGatewayRule](docs/EnterpriseHttpGatewayRule.md)
 - [EnterpriseInfo](docs/EnterpriseInfo.md)
 - [EnterpriseSource](docs/EnterpriseSource.md)
 - [ExportAppBaseResp](docs/ExportAppBaseResp.md)
 - [Fail](docs/Fail.md)
 - [HttpGatewayRule](docs/HttpGatewayRule.md)
 - [Install](docs/Install.md)
 - [ListAnnouncementResp](docs/ListAnnouncementResp.md)
 - [ListAppStoreInfosResp](docs/ListAppStoreInfosResp.md)
 - [ListEntsResp](docs/ListEntsResp.md)
 - [ListRegionTeamServices](docs/ListRegionTeamServices.md)
 - [ListTeamRegionsResp](docs/ListTeamRegionsResp.md)
 - [ListTeamResp](docs/ListTeamResp.md)
 - [ListTeamUsersResp](docs/ListTeamUsersResp.md)
 - [ListUsersRespView](docs/ListUsersRespView.md)
 - [NewBieGuideBaseResp](docs/NewBieGuideBaseResp.md)
 - [OAuthType](docs/OAuthType.md)
 - [OauthServicesBaseResp](docs/OauthServicesBaseResp.md)
 - [OauthServicesResp](docs/OauthServicesResp.md)
 - [ObjectStorageBaseResp](docs/ObjectStorageBaseResp.md)
 - [ObjectStorageResp](docs/ObjectStorageResp.md)
 - [PostHttpGatewayRule](docs/PostHttpGatewayRule.md)
 - [RegionInfo](docs/RegionInfo.md)
 - [RegionInfoResp](docs/RegionInfoResp.md)
 - [RoleInfo](docs/RoleInfo.md)
 - [ServiceBaseInfo](docs/ServiceBaseInfo.md)
 - [ServiceGroupOperations](docs/ServiceGroupOperations.md)
 - [Success](docs/Success.md)
 - [TeamBaseInfo](docs/TeamBaseInfo.md)
 - [TeamCertificatesC](docs/TeamCertificatesC.md)
 - [TeamCertificatesL](docs/TeamCertificatesL.md)
 - [TeamCertificatesR](docs/TeamCertificatesR.md)
 - [TeamInfo](docs/TeamInfo.md)
 - [TeamRegionReq](docs/TeamRegionReq.md)
 - [TeamRegionsResp](docs/TeamRegionsResp.md)
 - [TeamServicesResp](docs/TeamServicesResp.md)
 - [TeamUser](docs/TeamUser.md)
 - [Token](docs/Token.md)
 - [UpdAppStoreInfoReq](docs/UpdAppStoreInfoReq.md)
 - [UpdateAncmReqSerilizer](docs/UpdateAncmReqSerilizer.md)
 - [UpdatePostHttpGatewayRule](docs/UpdatePostHttpGatewayRule.md)
 - [UpdateRegionReq](docs/UpdateRegionReq.md)
 - [UpdateRegionStatusReq](docs/UpdateRegionStatusReq.md)
 - [UpdateTeamInfoReq](docs/UpdateTeamInfoReq.md)
 - [UpdateUser](docs/UpdateUser.md)
 - [UploadFileResp](docs/UploadFileResp.md)
 - [UserInfo](docs/UserInfo.md)


## Documentation For Authorization

## Bearer
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```

## Author

barnett@goodrain.com

