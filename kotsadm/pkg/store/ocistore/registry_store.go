package ocistore

import (
	"fmt"

	"github.com/pkg/errors"
	registrytypes "github.com/replicatedhq/kots/kotsadm/pkg/registry/types"
)

const (
	AppRegistryDetailsSecretNamePrefix = "kotsadm-app-registry"
)

// GetRegistryDetailsForApp will return the locally configured image registry for the
// given appID. The password is included in the response, encrypted, and the caller
// can decrypt it with the key
// Registry details are stored in a secret
func (s OCIStore) GetRegistryDetailsForApp(appID string) (*registrytypes.RegistrySettings, error) {
	secret, err := s.getSecret(AppRegistryDetailsSecretNamePrefix)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get app registry details secret")
	}

	registryHostname, ok := secret.Data[fmt.Sprintf("%s.hostname", appID)]
	if !ok {
		return nil, nil
	}

	registrySettings := registrytypes.RegistrySettings{
		Hostname: string(registryHostname),
	}

	registryUsername, ok := secret.Data[fmt.Sprintf("%s.username", appID)]
	if ok {
		registrySettings.Username = string(registryUsername)
	}

	registryPasswordEnc, ok := secret.Data[fmt.Sprintf("%s.password", appID)]
	if ok {
		registrySettings.PasswordEnc = string(registryPasswordEnc)
	}

	registryNamespace, ok := secret.Data[fmt.Sprintf("%s.namespace", appID)]
	if ok {
		registrySettings.Namespace = string(registryNamespace)
	}

	return &registrySettings, nil
}

func (s OCIStore) UpdateRegistry(appID string, hostname string, username string, password string, namespace string) error {
	return ErrNotImplemented
}
