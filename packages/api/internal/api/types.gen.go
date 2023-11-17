// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package api

import (
	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	AccessTokenAuthScopes = "AccessTokenAuth.Scopes"
	ApiKeyAuthScopes      = "ApiKeyAuth.Scopes"
)

// Defines values for EnvironmentBuildStatus.
const (
	EnvironmentBuildStatusBuilding EnvironmentBuildStatus = "building"
	EnvironmentBuildStatusError    EnvironmentBuildStatus = "error"
	EnvironmentBuildStatusReady    EnvironmentBuildStatus = "ready"
)

// Environment defines model for Environment.
type Environment struct {
	// Aliases Aliases of the environment
	Aliases *[]string `json:"aliases,omitempty"`

	// BuildID Identifier of the last successful build for given environment
	BuildID string `json:"buildID"`

	// EnvID Identifier of the environment
	EnvID string `json:"envID"`

	// Public Whether the environment is public or only accessible by the team
	Public bool `json:"public"`
}

// EnvironmentBuild defines model for EnvironmentBuild.
type EnvironmentBuild struct {
	// BuildID Identifier of the build
	BuildID string `json:"buildID"`

	// EnvID Identifier of the environment
	EnvID string `json:"envID"`

	// Logs Build logs
	Logs []string `json:"logs"`

	// Status Status of the environment
	Status *EnvironmentBuildStatus `json:"status,omitempty"`
}

// EnvironmentBuildStatus Status of the environment
type EnvironmentBuildStatus string

// Error defines model for Error.
type Error struct {
	// Code Error code
	Code int32 `json:"code"`

	// Message Error
	Message string `json:"message"`
}

// Instance defines model for Instance.
type Instance struct {
	// ClientID Identifier of the client
	ClientID string `json:"clientID"`

	// EnvID Identifier of the environment from which is the instance created
	EnvID string `json:"envID"`

	// InstanceID Identifier of the instance
	InstanceID string `json:"instanceID"`
}

// NewInstance defines model for NewInstance.
type NewInstance struct {
	// EnvID Identifier of the required environment
	EnvID string `json:"envID"`
}

// BuildID defines model for buildID.
type BuildID = string

// EnvID defines model for envID.
type EnvID = string

// InstanceID defines model for instanceID.
type InstanceID = string

// N400 defines model for 400.
type N400 = Error

// N401 defines model for 401.
type N401 = Error

// N404 defines model for 404.
type N404 = Error

// N500 defines model for 500.
type N500 = Error

// PostEnvsMultipartBody defines parameters for PostEnvs.
type PostEnvsMultipartBody struct {
	// Alias Alias of the environment
	Alias *string `json:"alias,omitempty"`

	// BuildContext Docker build context
	BuildContext openapi_types.File `json:"buildContext"`

	// Dockerfile Dockerfile content
	Dockerfile string `json:"dockerfile"`

	// StartCmd Start command to execute in the template after the build
	StartCmd *string `json:"startCmd,omitempty"`
}

// PostEnvsEnvIDMultipartBody defines parameters for PostEnvsEnvID.
type PostEnvsEnvIDMultipartBody struct {
	// Alias Alias of the environment
	Alias *string `json:"alias,omitempty"`

	// BuildContext Docker build context
	BuildContext openapi_types.File `json:"buildContext"`

	// Dockerfile Dockerfile content
	Dockerfile string `json:"dockerfile"`

	// StartCmd Start command to execute in the template after the build
	StartCmd *string `json:"startCmd,omitempty"`
}

// GetEnvsEnvIDBuildsBuildIDParams defines parameters for GetEnvsEnvIDBuildsBuildID.
type GetEnvsEnvIDBuildsBuildIDParams struct {
	// LogsOffset Index of the starting build log that should be returned with the environment
	LogsOffset *int `form:"logsOffset,omitempty" json:"logsOffset,omitempty"`
}

// PostEnvsEnvIDBuildsBuildIDLogsJSONBody defines parameters for PostEnvsEnvIDBuildsBuildIDLogs.
type PostEnvsEnvIDBuildsBuildIDLogsJSONBody struct {
	// ApiSecret API secret
	ApiSecret string   `json:"apiSecret"`
	Logs      []string `json:"logs"`
}

// PostInstancesInstanceIDRefreshesJSONBody defines parameters for PostInstancesInstanceIDRefreshes.
type PostInstancesInstanceIDRefreshesJSONBody struct {
	// Duration Duration for which the instance should be kept alive in seconds
	Duration *int `json:"duration,omitempty"`
}

// PostEnvsMultipartRequestBody defines body for PostEnvs for multipart/form-data ContentType.
type PostEnvsMultipartRequestBody PostEnvsMultipartBody

// PostEnvsEnvIDMultipartRequestBody defines body for PostEnvsEnvID for multipart/form-data ContentType.
type PostEnvsEnvIDMultipartRequestBody PostEnvsEnvIDMultipartBody

// PostEnvsEnvIDBuildsBuildIDLogsJSONRequestBody defines body for PostEnvsEnvIDBuildsBuildIDLogs for application/json ContentType.
type PostEnvsEnvIDBuildsBuildIDLogsJSONRequestBody PostEnvsEnvIDBuildsBuildIDLogsJSONBody

// PostInstancesJSONRequestBody defines body for PostInstances for application/json ContentType.
type PostInstancesJSONRequestBody = NewInstance

// PostInstancesInstanceIDRefreshesJSONRequestBody defines body for PostInstancesInstanceIDRefreshes for application/json ContentType.
type PostInstancesInstanceIDRefreshesJSONRequestBody PostInstancesInstanceIDRefreshesJSONBody
