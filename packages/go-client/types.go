package client

import gen "github.com/sorolens/sorolens/packages/go-client/internal/client"

// The generated client lives in an internal package so that regenerating it
// from docs/openapi.yaml never changes this package's import path. The aliases
// below re-export exactly the types that ApiClient's methods take or return,
// so callers outside this module can name them.

// ClientWithResponses is the full generated client returned by
// ApiClient.Generated.
type ClientWithResponses = gen.ClientWithResponses

// Request parameters.
type (
	// Cursor is the opaque pagination cursor taken from a previous page's
	// next_cursor field.
	Cursor = gen.Cursor
	// Limit is a page size (default 50, clamped to 200 by the API).
	Limit = gen.Limit
	// AlertCursor is the pagination cursor for watchdog alerts.
	AlertCursor = gen.AlertCursor
	// TagParam filters contracts by tag.
	TagParam = gen.TagParam

	ListContractsParams           = gen.ListContractsParams
	GetContractForecastParams     = gen.GetContractForecastParams
	ListContractEventsParams      = gen.ListContractEventsParams
	ListContractInvocationsParams = gen.ListContractInvocationsParams
	ListContractStorageParams     = gen.ListContractStorageParams
	ListWatchdogAlertsParams      = gen.ListWatchdogAlertsParams
	GetContractSnapshotParams     = gen.GetContractSnapshotParams
)

// Responses. Each carries the raw body and HTTP response plus one typed field
// per documented status code (JSON200, JSON404, ...).
type (
	ListContractsResponse           = gen.ListContractsResponse
	GetContractResponse             = gen.GetContractResponse
	GetGlobalStatsResponse          = gen.GetGlobalStatsResponse
	GetContractForecastResponse     = gen.GetContractForecastResponse
	ListContractEventsResponse      = gen.ListContractEventsResponse
	ListContractInvocationsResponse = gen.ListContractInvocationsResponse
	ListContractStorageResponse     = gen.ListContractStorageResponse
	ListWatchdogAlertsResponse      = gen.ListWatchdogAlertsResponse
	GetContractSnapshotResponse     = gen.GetContractSnapshotResponse
)

// Models carried by the responses above.
type (
	Contract         = gen.Contract
	ContractAlert    = gen.ContractAlert
	ContractSnapshot = gen.ContractSnapshot
	Event            = gen.Event
	Invocation       = gen.Invocation
	StorageEntry     = gen.StorageEntry
	ForecastSeries   = gen.ForecastSeries
	ForecastPoint    = gen.ForecastPoint

	// Error is the API's error envelope; the error responses below share it.
	Error         = gen.Error
	Error_Error   = gen.Error_Error
	InvalidInput  = gen.InvalidInput
	NotFound      = gen.NotFound
	InternalError = gen.InternalError
	RateLimited   = gen.RateLimited
)

// Enum types. Values without a type prefix in the generated code (for example
// the ListContractsParamsSort and ListContractsParamsDir values) are not
// re-exported because the generator renames them whenever another enum in the
// spec collides; convert a string instead, e.g. ListContractsParamsSort("label").
type (
	ContractAlertSeverity                = gen.ContractAlertSeverity
	ContractStatus                       = gen.ContractStatus
	ForecastSeriesMetric                 = gen.ForecastSeriesMetric
	ListContractsParamsNetwork           = gen.ListContractsParamsNetwork
	ListContractsParamsSort              = gen.ListContractsParamsSort
	ListContractsParamsDir               = gen.ListContractsParamsDir
	ListContractEventsParamsNetwork      = gen.ListContractEventsParamsNetwork
	ListContractInvocationsParamsNetwork = gen.ListContractInvocationsParamsNetwork
	ListContractInvocationsParamsStatus  = gen.ListContractInvocationsParamsStatus
	ListContractStorageParamsDurability  = gen.ListContractStorageParamsDurability
	ListContractStorageParamsNetwork     = gen.ListContractStorageParamsNetwork
	ListContractStorageParamsStatus      = gen.ListContractStorageParamsStatus
	ListWatchdogAlertsParamsNetwork      = gen.ListWatchdogAlertsParamsNetwork
	ListWatchdogAlertsParamsSeverity     = gen.ListWatchdogAlertsParamsSeverity
)

const (
	ContractAlertSeverityCritical = gen.ContractAlertSeverityCritical
	ContractAlertSeverityInfo     = gen.ContractAlertSeverityInfo
	ContractAlertSeverityWarning  = gen.ContractAlertSeverityWarning

	ContractStatusActive      = gen.ContractStatusActive
	ContractStatusBackfilling = gen.ContractStatusBackfilling
	ContractStatusError       = gen.ContractStatusError
	ContractStatusPaused      = gen.ContractStatusPaused
	ContractStatusPending     = gen.ContractStatusPending

	ListContractsParamsNetworkFuturenet  = gen.ListContractsParamsNetworkFuturenet
	ListContractsParamsNetworkMainnet    = gen.ListContractsParamsNetworkMainnet
	ListContractsParamsNetworkStandalone = gen.ListContractsParamsNetworkStandalone
	ListContractsParamsNetworkTestnet    = gen.ListContractsParamsNetworkTestnet

	ListContractEventsParamsNetworkFuturenet  = gen.ListContractEventsParamsNetworkFuturenet
	ListContractEventsParamsNetworkMainnet    = gen.ListContractEventsParamsNetworkMainnet
	ListContractEventsParamsNetworkStandalone = gen.ListContractEventsParamsNetworkStandalone
	ListContractEventsParamsNetworkTestnet    = gen.ListContractEventsParamsNetworkTestnet

	ListContractInvocationsParamsNetworkFuturenet  = gen.ListContractInvocationsParamsNetworkFuturenet
	ListContractInvocationsParamsNetworkMainnet    = gen.ListContractInvocationsParamsNetworkMainnet
	ListContractInvocationsParamsNetworkStandalone = gen.ListContractInvocationsParamsNetworkStandalone
	ListContractInvocationsParamsNetworkTestnet    = gen.ListContractInvocationsParamsNetworkTestnet
	ListContractInvocationsParamsStatusFAILED      = gen.ListContractInvocationsParamsStatusFAILED
	ListContractInvocationsParamsStatusNOTFOUND    = gen.ListContractInvocationsParamsStatusNOTFOUND
	ListContractInvocationsParamsStatusSUCCESS     = gen.ListContractInvocationsParamsStatusSUCCESS

	ListContractStorageParamsDurabilityInstance   = gen.ListContractStorageParamsDurabilityInstance
	ListContractStorageParamsDurabilityPersistent = gen.ListContractStorageParamsDurabilityPersistent
	ListContractStorageParamsDurabilityTemporary  = gen.ListContractStorageParamsDurabilityTemporary
	ListContractStorageParamsNetworkFuturenet     = gen.ListContractStorageParamsNetworkFuturenet
	ListContractStorageParamsNetworkMainnet       = gen.ListContractStorageParamsNetworkMainnet
	ListContractStorageParamsNetworkStandalone    = gen.ListContractStorageParamsNetworkStandalone
	ListContractStorageParamsNetworkTestnet       = gen.ListContractStorageParamsNetworkTestnet
	ListContractStorageParamsStatusArchived       = gen.ListContractStorageParamsStatusArchived
	ListContractStorageParamsStatusDeleted        = gen.ListContractStorageParamsStatusDeleted
	ListContractStorageParamsStatusLive           = gen.ListContractStorageParamsStatusLive

	ListWatchdogAlertsParamsNetworkFuturenet  = gen.ListWatchdogAlertsParamsNetworkFuturenet
	ListWatchdogAlertsParamsNetworkMainnet    = gen.ListWatchdogAlertsParamsNetworkMainnet
	ListWatchdogAlertsParamsNetworkStandalone = gen.ListWatchdogAlertsParamsNetworkStandalone
	ListWatchdogAlertsParamsNetworkTestnet    = gen.ListWatchdogAlertsParamsNetworkTestnet
	ListWatchdogAlertsParamsSeverityCritical  = gen.ListWatchdogAlertsParamsSeverityCritical
	ListWatchdogAlertsParamsSeverityInfo      = gen.ListWatchdogAlertsParamsSeverityInfo
	ListWatchdogAlertsParamsSeverityWarning   = gen.ListWatchdogAlertsParamsSeverityWarning
)
