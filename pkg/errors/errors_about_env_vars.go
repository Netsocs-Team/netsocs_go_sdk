package errors

import "fmt"

type MISSING_INITIAL_ENVIRONMENT_VARIABLES struct {
	MISSING_VARS []string
	DOMAIN       string
}

func (e MISSING_INITIAL_ENVIRONMENT_VARIABLES) Error() error {
	return fmt.Errorf("missing %s environment variable(s): %v", e.DOMAIN, e.MISSING_VARS)
}

func NewMissingInitialEnvironmentVariablesError(domain string, missingVars []string) error {
	return MISSING_INITIAL_ENVIRONMENT_VARIABLES{
		MISSING_VARS: missingVars,
		DOMAIN:       domain,
	}.Error()
}
