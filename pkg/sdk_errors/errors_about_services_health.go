package sdk_errors

import "fmt"

type SERVICE_HEALTH_CHECK_FAILED struct {
	DOMAIN string
	REASON string
}

func (e SERVICE_HEALTH_CHECK_FAILED) Error() string {
	return fmt.Sprintf("%s health check failed: %s", e.DOMAIN, e.REASON)
}

func NewServiceHealthCheckFailedError(domain string, reason string) error {
	return SERVICE_HEALTH_CHECK_FAILED{
		DOMAIN: domain,
		REASON: reason,
	}
}
