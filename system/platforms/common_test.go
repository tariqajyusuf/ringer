package platforms

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPackageNotFound_Error(t *testing.T) {
	e := &PackageNotFound{message: "not found"}
	assert.Equal(t, "not found", e.Error())
}

func TestInstallError_Error(t *testing.T) {
	e := &InstallError{message: "install failed"}
	assert.Equal(t, "install failed", e.Error())
}

func TestAuthorizationError_Error(t *testing.T) {
	e := &AuthorizationError{message: "permission denied"}
	assert.Equal(t, "permission denied", e.Error())
}

func TestErrorsSatisfyErrorInterface(t *testing.T) {
	var _ error = &PackageNotFound{message: "x"}
	var _ error = &InstallError{message: "x"}
	var _ error = &AuthorizationError{message: "x"}
}

func TestErrorsAs_PackageNotFound(t *testing.T) {
	err := error(&PackageNotFound{message: "pkg missing"})
	var target *PackageNotFound
	assert.True(t, errors.As(err, &target))
	assert.Equal(t, "pkg missing", target.message)
}

func TestErrorsAs_InstallError(t *testing.T) {
	err := error(&InstallError{message: "install failed"})
	var target *InstallError
	assert.True(t, errors.As(err, &target))
}

func TestErrorsAs_AuthorizationError(t *testing.T) {
	err := error(&AuthorizationError{message: "denied"})
	var target *AuthorizationError
	assert.True(t, errors.As(err, &target))
}
