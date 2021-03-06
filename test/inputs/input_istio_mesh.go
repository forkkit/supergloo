package inputs

import (
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	v1 "github.com/solo-io/supergloo/pkg/api/v1"
)

const (
	IstioTestInstallNs = "istio-was-installed-herr"
	IstioTestVersion   = "1.0.6"
)

func IstioMesh(namespace string, secretRef *core.ResourceRef) *v1.Mesh {
	return IstioMeshWithInstallNs(namespace, IstioTestInstallNs, secretRef)
}

func IstioMeshWithVersion(namespace, version string, secretRef *core.ResourceRef) *v1.Mesh {
	return istioMesh(namespace, IstioTestInstallNs, version, secretRef)
}

func IstioMeshWithInstallNs(namespace, installNs string, secretRef *core.ResourceRef) *v1.Mesh {
	return istioMesh(namespace, installNs, "", secretRef)
}

func istioMesh(namespace, installNs, version string, secretRef *core.ResourceRef) *v1.Mesh {
	return &v1.Mesh{
		Metadata: core.Metadata{
			Namespace: namespace,
			Name:      "fancy-istio",
		},
		MeshType: &v1.Mesh_Istio{
			Istio: &v1.IstioMesh{
				InstallationNamespace: installNs,
				Version:               version,
			},
		},
		MtlsConfig: &v1.MtlsConfig{
			MtlsEnabled:     true,
			RootCertificate: secretRef,
		},
	}
}

func IstioMeshWithInstallNsPrometheus(namespace, installNs string, secretRef *core.ResourceRef, promCfgRefs []core.ResourceRef) *v1.Mesh {
	var monit *v1.MonitoringConfig
	if len(promCfgRefs) > 0 {
		monit = &v1.MonitoringConfig{
			PrometheusConfigmaps: promCfgRefs,
		}
	}
	return &v1.Mesh{
		Metadata: core.Metadata{
			Namespace: namespace,
			Name:      "fancy-istio",
		},
		MeshType: &v1.Mesh_Istio{
			Istio: &v1.IstioMesh{
				InstallationNamespace: installNs,
			},
		},
		MtlsConfig: &v1.MtlsConfig{
			MtlsEnabled:     true,
			RootCertificate: secretRef,
		},
		MonitoringConfig: monit,
	}
}
