package backingresource

import (
	"path/filepath"

	"github.com/uccps-samples/library-go/pkg/assets"
	"github.com/uccps-samples/library-go/pkg/operator/staticpod/controller/backingresource/bindata"
)

const (
	manifestDir = "pkg/operator/staticpod/controller/backingresource"
)

func StaticPodManifests(targetNamespace string) func(name string) ([]byte, error) {
	return func(name string) ([]byte, error) {
		config := struct {
			TargetNamespace string
		}{
			TargetNamespace: targetNamespace,
		}
		return assets.MustCreateAssetFromTemplate(name, bindata.MustAsset(filepath.Join(manifestDir, name)), config).Data, nil
	}
}
