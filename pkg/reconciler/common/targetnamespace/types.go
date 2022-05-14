package targetnamespace

import (
	"context"
	"fmt"
	"strings"

	corev1 "k8s.io/api/core/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	mf "github.com/manifestival/manifestival"
	"github.com/tektoncd/operator/pkg/reconciler/common"
	"k8s.io/apimachinery/pkg/labels"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

const (
	targetNamespaceLabelKey = "operator.tekton.dev/namespace-in-use-by-%s"
)

type targetNamespaceManager struct {
	client    v1.NamespaceInterface
	component string
	nsName    string
}

func NewTargetNamespaceManager(nsClient v1.NamespaceInterface, component, nsName string) *targetNamespaceManager {
	return &targetNamespaceManager{
		client:    nsClient,
		component: component,
		nsName:    nsName,
	}
}

func inUseLabelKey(component string) string {
	return fmt.Sprintf(targetNamespaceLabelKey, component)
}

func InjectNamespaceInUseLabel(ns, component string) mf.Transformer {
	lSet := labels.Set{
		inUseLabelKey(component): "true",
	}

	conditions := []mf.Predicate{
		mf.ByKind("Namespace"),
		mf.ByName(ns),
	}
	return common.InjectLabelOverwriteExisting(lSet, conditions...)
}

func (tnm *targetNamespaceManager) CheckNamespaceInUseLabel(ctx context.Context) (bool, error) {
	ns, err := tnm.GetNamespace(ctx)
	if err != nil {
		return false, nil
	}
	labels := ns.GetLabels()
	checkLabelKey := inUseLabelKey(tnm.component)
	for k, v := range labels {
		if k == checkLabelKey && v == "true" {
			return true, nil
		}
	}
	return false, nil
}

func (tnm *targetNamespaceManager) GetNamespace(ctx context.Context) (*corev1.Namespace, error) {
	return tnm.client.Get(ctx, tnm.nsName, metav1.GetOptions{})
}

func (tnm *targetNamespaceManager) RemoveNamespaceInUseLabel(ctx context.Context) error {
	ns, err := tnm.GetNamespace(ctx)
	if err != nil {
		return err
	}
	labels := ns.GetLabels()
	inUsekey := inUseLabelKey(tnm.component)
	delete(labels, inUsekey)
	ns.SetLabels(labels)
	_, err = tnm.client.Update(ctx, ns, metav1.UpdateOptions{})
	return err
}

func (tnm *targetNamespaceManager) DeleteNamespaceIfNoOtherUsers(ctx context.Context) error {
	err := tnm.RemoveNamespaceInUseLabel(ctx)
	if err != nil {
		return err
	}
	ok, err := tnm.CheckAnyOtherComponentPresent(ctx)
	if err != nil {
		return err
	}
	if ok {
		return nil
	}
	return tnm.client.Delete(ctx, tnm.nsName, metav1.DeleteOptions{})
}

func (tnm *targetNamespaceManager) CheckAnyOtherComponentPresent(ctx context.Context) (bool, error) {
	ns, err := tnm.GetNamespace(ctx)
	if err != nil {
		return false, err
	}
	labels := ns.GetLabels()
	for key := range labels {
		if isInUseLabel(key) {
			return true, nil
		}
	}
	return false, nil
}

func isInUseLabel(key string) bool {
	inUseLabelKeyPrefix := fmt.Sprintf(targetNamespaceLabelKey, "")
	return strings.HasPrefix(key, inUseLabelKeyPrefix)
}
