package opensearchrepository

import (
	"errors"

	"github.com/rancher/opni/pkg/opensearch/certs"
	osapi "github.com/rancher/opni/pkg/opensearch/opensearch/types"
	opensearch "github.com/rancher/opni/pkg/opensearch/reconciler"
	"github.com/rancher/opni/pkg/util/meta"
	"k8s.io/client-go/util/retry"
	opensearchv1 "opensearch.opster.io/api/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

func (r *Reconciler) reconcileOpensearchObjects(cluster *opensearchv1.OpenSearchCluster) error {
	certMgr := certs.NewCertMgrOpensearchCertManager(
		r.ctx,
		certs.WithNamespace(cluster.Namespace),
		certs.WithCluster(cluster.Name),
	)

	osReconciler, err := opensearch.NewReconciler(
		r.ctx,
		opensearch.ReconcilerConfig{
			CertReader:            certMgr,
			OpensearchServiceName: cluster.Spec.General.ServiceName,
		},
	)
	if err != nil {
		return err
	}

	settings := osapi.RepositoryRequest{}
	switch {
	case r.repository.Spec.Settings.S3 != nil:
		settings.Type = osapi.RepositoryTypeS3
		settings.Settings.S3Settings = &osapi.S3Settings{
			Bucket: r.repository.Spec.Settings.S3.Bucket,
			Path:   r.repository.Spec.Settings.S3.Folder,
		}
	case r.repository.Spec.Settings.FileSystem != nil:
		settings.Type = osapi.RepositoryTypeFileSystem
		settings.Settings.FileSystemSettings = &osapi.FileSystemSettings{
			Location: r.repository.Spec.Settings.FileSystem.Location,
		}
	default:
		return errors.New("invalid repository settings")
	}

	return osReconciler.MaybeUpdateRepository(r.repository.Name, settings)
}

func (r *Reconciler) deleteOpensearchObjects(cluster *opensearchv1.OpenSearchCluster) error {
	if cluster != nil {
		certMgr := certs.NewCertMgrOpensearchCertManager(
			r.ctx,
			certs.WithNamespace(cluster.Namespace),
			certs.WithCluster(cluster.Name),
		)

		osReconciler, err := opensearch.NewReconciler(
			r.ctx,
			opensearch.ReconcilerConfig{
				CertReader:            certMgr,
				OpensearchServiceName: cluster.Spec.General.ServiceName,
			},
		)
		if err != nil {
			return err
		}

		err = osReconciler.MaybeDeleteRepository(r.repository.Name)
		if err != nil {
			return err
		}
	}

	return retry.RetryOnConflict(retry.DefaultRetry, func() error {
		if err := r.client.Get(r.ctx, client.ObjectKeyFromObject(r.repository), r.repository); err != nil {
			return err
		}
		controllerutil.RemoveFinalizer(r.repository, meta.OpensearchFinalizer)
		return r.client.Update(r.ctx, r.repository)
	})
}