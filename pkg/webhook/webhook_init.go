package webhook

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-logr/zapr"
	mfc "github.com/manifestival/client-go-client"
	mf "github.com/manifestival/manifestival"
	"github.com/tektoncd/operator/pkg/apis/operator/v1alpha1"
	clientset "github.com/tektoncd/operator/pkg/client/clientset/versioned"
	operatorclient "github.com/tektoncd/operator/pkg/client/injection/client"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/tektoncd/operator/pkg/reconciler/common"
	"go.uber.org/zap"
	"knative.dev/pkg/injection"
	"knative.dev/pkg/logging"
)

const WEBHOOK_INSTALLERSET_LABEL = "validating-defaulting-webhooks.operator.tekton.dev"

func CreateWebhookResources(ctx context.Context) {


	logger := logging.FromContext(ctx)
	mfclient, err := mfc.NewClient(injection.GetConfig(ctx))
	if err != nil {
		logger.Fatalw("error creating client from injected config", zap.Error(err))
	}
	mflogger := zapr.NewLogger(logger.Named("manifestival").Desugar())
	manifest, err := mf.ManifestFrom(mf.Slice{}, mf.UseClient(mfclient), mf.UseLogger(mflogger))
	if err != nil {
		logger.Fatalw("error creating initial manifest", zap.Error(err))
	}

	// Read manifests
	koDataDir := os.Getenv(common.KoEnvKey)
	validating_defaulting_webhooks := filepath.Join(koDataDir, "validating-defaulting-webhook")
	if err := common.AppendManifest(&manifest, validating_defaulting_webhooks); err != nil {
		logger.Fatalw("error creating initial manifest", zap.Error(err))
	}

	client := operatorclient.Get(ctx)

	err = checkAndDeleteInstallerSet(ctx, client)
	if err != nil {
		logger.Fatalw("error creating client from injected config", zap.Error(err))
	}

	if err := createInstallerSet(ctx, client, manifest); err != nil {
		logger.Fatalw("error creating client from injected config", zap.Error(err))
	}

	//// If installer set doesn't exist then create a new one
	//if !exist {
	//
	//	// make sure that openshift-pipelines namespace exists
	//	namespaceLocation := filepath.Join(koDataDir, "tekton-namespace")
	//	if err := common.AppendManifest(&oe.manifest, namespaceLocation); err != nil {
	//		return err
	//	}
	//
	//	// add inject CA bundles manifests
	//	cabundlesLocation := filepath.Join(koDataDir, "cabundles")
	//	if err := common.AppendManifest(&oe.manifest, cabundlesLocation); err != nil {
	//		return err
	//	}
	//
	//	// add pipelines-scc
	//	pipelinesSCCLocation := filepath.Join(koDataDir, "tekton-pipeline", "00-prereconcile")
	//	if err := common.AppendManifest(&oe.manifest, pipelinesSCCLocation); err != nil {
	//		return err
	//	}
	//
	//	if err := common.Transform(ctx, &oe.manifest, tp); err != nil {
	//		return err
	//	}
	//
	//	if err := createInstallerSet(ctx, oe.operatorClientSet, tp, oe.manifest, oe.version,
	//		prePipelineInstallerSet, "pre-pipeline"); err != nil {
	//		return err
	//	}
	//}

}

func checkAndDeleteInstallerSet(ctx context.Context, oc clientset.Interface) error {

	//// Check if installer set is already created
	//compInstallerSet, ok := tp.Status.ExtentionInstallerSets[component]
	//if !ok {
	//	return false, nil
	//}
	//
	//if compInstallerSet != "" {
		// if already created then check which version it is
	ctIs, err := oc.OperatorV1alpha1().TektonInstallerSets().
		List(ctx, metav1.ListOptions{
		LabelSelector: WEBHOOK_INSTALLERSET_LABEL,
	})
	if err != nil {
		if errors.IsNotFound(err) {
			return nil
		}
		return err
	}
	if len(ctIs.Items) >= 0 {
		for _, item := range ctIs.Items {
			err = oc.OperatorV1alpha1().TektonInstallerSets().
				Delete(ctx, item.Name, metav1.DeleteOptions{})
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func createInstallerSet(ctx context.Context, oc clientset.Interface, manifest mf.Manifest) error {

	is := makeInstallerSet(manifest)

	_, err := oc.OperatorV1alpha1().TektonInstallerSets().
		Create(ctx, is, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	return nil
}

func makeInstallerSet(manifest mf.Manifest) *v1alpha1.TektonInstallerSet {
	//ownerRef := *metav1.NewControllerRef(tp, tp.GetGroupVersionKind())
	return &v1alpha1.TektonInstallerSet{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: fmt.Sprintf("%s-", "validating-mutating-webhoook"),
			Labels: map[string]string{
				WEBHOOK_INSTALLERSET_LABEL: "",
			},
			Annotations: map[string]string{
				"releaseVersionKey":  "v1.6.0",
			},
			//OwnerReferences: []metav1.OwnerReference{ownerRef},
		},
		Spec: v1alpha1.TektonInstallerSetSpec{
			Manifests: manifest.Resources(),
		},
	}
}