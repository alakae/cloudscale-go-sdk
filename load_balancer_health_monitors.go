package cloudscale

import (
	"time"
)

const loadBalancerHealthMonitorBasePath = "v1/load-balancers/health-monitors"

type LoadBalancerHealthMonitor struct {
	TaggedResource
	// Just use omitempty everywhere. This makes it easy to use restful. Errors
	// will be coming from the API if something is disabled.
	HREF           string               `json:"href,omitempty"`
	UUID           string               `json:"uuid,omitempty"`
	Pool           LoadBalancerPoolStub `json:"pool,omitempty"`
	Delay          int                  `json:"delay,omitempty"`
	Timeout        int                  `json:"timeout,omitempty"`
	MaxRetries     int                  `json:"max_retries,omitempty"`
	MaxRetriesDown int                  `json:"max_retries_down,omitempty"`
	CreatedAt      time.Time            `json:"created_at,omitempty"`
}

type LoadBalancerHealthMonitorRequest struct {
	TaggedResourceRequest
	Pool           string `json:"pool,omitempty"`
	Delay          int    `json:"delay,omitempty"`
	Timeout        int    `json:"timeout,omitempty"`
	MaxRetries     int    `json:"max_retries,omitempty"`
	MaxRetriesDown int    `json:"max_retries_down,omitempty"`
	Type           string `json:"type,omitempty"`
}

type LoadBalancerHealthMonitorService interface {
	GenericCreateService[LoadBalancerHealthMonitor, LoadBalancerHealthMonitorRequest, LoadBalancerHealthMonitorRequest]
	GenericGetService[LoadBalancerHealthMonitor, LoadBalancerHealthMonitorRequest, LoadBalancerHealthMonitorRequest]
	GenericListService[LoadBalancerHealthMonitor, LoadBalancerHealthMonitorRequest, LoadBalancerHealthMonitorRequest]
	GenericUpdateService[LoadBalancerHealthMonitor, LoadBalancerHealthMonitorRequest, LoadBalancerHealthMonitorRequest]
	GenericDeleteService[LoadBalancerHealthMonitor, LoadBalancerHealthMonitorRequest, LoadBalancerHealthMonitorRequest]
}
