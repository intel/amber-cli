/*
 * Copyright (C) 2022 Intel Corporation
 * SPDX-License-Identifier: BSD-3-Clause
 */

package constants

const (
	ExecLink              = "/usr/bin/tenantctl"
	HomeDir               = "/opt/tac/"
	ConfigDir             = "/etc/tac/"
	LogDir                = "/var/log/tac/"
	DefaultConfigFilePath = ConfigDir + "config.yaml"
	ConfigFileName        = "config"
	ConfigFileExtension   = "yaml"
	LogFilePath           = LogDir + "tac.log"
	DefaultFilePermission = 0640
	ExplicitCLIName       = "Tenant CLI"
)

//Command and parameter names
const (
	ApiKeyParamName          = "api-key"
	TenantIdParamName        = "tenant-id"
	UserIdParamName          = "user-id"
	ServiceIdParamName       = "service-id"
	ServiceOfferIdParamName  = "service-offer-id"
	ProductIdParamName       = "product-id"
	ApiClientIdParamName     = "api-client-id"
	ActivationStatus         = "status"
	PolicyIdsParamName       = "policy-ids"
	ApiClientNameParamName   = "api-client-name"
	ServiceNameParamName     = "service-name"
	EmailIdParamName         = "email-id"
	UserRoleParamName        = "user-role"
	PolicyFileParamName      = "policy-file"
	PolicyIdParamName        = "policy-id"
	TagNameParamName         = "tag-name"
	TagKeyAndValuesParamName = "tag-key-value"
	SetExpiryDateParamName   = "set-expiry"
	EnvFileParamName         = "env-file"

	RootCmd        = "tenantctl"
	CreateCmd      = "create"
	ListCmd        = "list"
	DeleteCmd      = "delete"
	UpdateCmd      = "update"
	UninstallCmd   = "uninstall"
	VersionCmd     = "version"
	SetupConfigCmd = "config"
)

//Resource names
const (
	PolicyCmd       = "policy"
	UserCmd         = "user"
	ProductCmd      = "product"
	ServiceOfferCmd = "serviceOffer"
	ServiceCmd      = "service"
	ApiClientCmd    = "apiClient"
	TagCmd          = "tag"
	RoleCmd         = "role"
)

const (
	AmberBaseUrl      = "amber-base-url"
	TenantId          = "tenant-id"
	HttpClientTimeout = "http-client-timeout"
	Loglevel          = "log-level"

	DefaultLogLevel          = "info"
	DefaultHttpClientTimeout = 10
	ApiClientStatusActive    = "Active"
	ApiClientStatusInactive  = "Inactive"
	ApiClientStatusCancelled = "Cancelled"
	TenantAdminRole          = "Tenant Admin"
	UserRole                 = "User"
	ExpiryDateInputFormat    = "2006-01-02"
)

//HTTP constants
const (
	HTTPMediaTypeJson        = "application/json"
	HTTPHeaderKeyContentType = "Content-Type"
	HTTPHeaderKeyAccept      = "Accept"
	HTTPHeaderKeyApiKey      = "X-API-KEY"
	HTTPHeaderKeyCreatedBy   = "Created-By"
	HTTPHeaderKeyUpdatedBy   = "Updated-By"
)

//API endpoint
const (
	TmsBaseUrl                = "/management/v1"
	PmsBaseUrl                = "/management/v1"
	PolicyApiEndpoint         = "/policies"
	ServiceApiEndpoint        = "/services"
	ServiceOfferApiEndpoint   = "/service-offers"
	ApiClientResourceEndpoint = "/api-clients"
	UserApiEndpoint           = "/users"
	ProductApiEndpoint        = "/products"
	TagApiEndpoint            = "/tags"
)
