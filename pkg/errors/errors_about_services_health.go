package errors

import "fmt"

type SERVICE_HEALTH_CHECK_FAILED struct {
	DOMAIN string
	REASON string
}

func (e SERVICE_HEALTH_CHECK_FAILED) Error() error {
	return fmt.Errorf("%s health check failed: %s", e.DOMAIN, e.REASON)
}

func NewServiceHealthCheckFailedError(domain string, reason string) error {
	return SERVICE_HEALTH_CHECK_FAILED{
		DOMAIN: domain,
		REASON: reason,
	}.Error()
}
