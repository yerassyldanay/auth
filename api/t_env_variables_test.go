package api

import (
	"auth/utils/config"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEnvVariables(t *testing.T) {
	configuration, err := config.LoadConfig("../environment/")
	require.NoError(t, err)

	require.NotZero(t, configuration.BackendHost)
	require.NotZero(t, configuration.BackendPort)
	require.NotZero(t, configuration.BackendScheme)

	require.NotZero(t, configuration.JwtSignUpEmailAudience)
	require.NotZero(t, configuration.JwtSignUpEmailSecretKey)

	require.NotZero(t, configuration.SmtpServerHost)
	require.NotZero(t, configuration.SmtpServerPort)
	require.NotZero(t, configuration.SmtpServerUsername)
	require.NotZero(t, configuration.SmtpServerPassword)

	require.NotZero(t, configuration.DbName)
	require.NotZero(t, configuration.DbSource)
}
