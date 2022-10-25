package resourceapply

import (
	"bytes"
	"context"
	"fmt"
	"sort"
	"strings"

	"k8s.io/klog/v2"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	coreclientv1 "k8s.io/client-go/kubernetes/typed/core/v1"

	"github.com/uccps-samples/library-go/pkg/operator/events"
	"github.com/uccps-samples/library-go/pkg/operator/resource/resourcemerge"
)

// ApplyNamespace merges objectmeta, does not worry about anything else
func ApplyNamespace(ctx context.Context, client coreclientv1.NamespacesGetter, recorder events.Recorder, required *corev1.Namespace) (*corev1.Namespace, bool, error) {
	existing, err := client.Namespaces().Get(ctx, required.Name, metav1.GetOptions{})
	if apierrors.IsNotFound(err) {
		requiredCopy := required.DeepCopy()
		actual, err := client.Namespaces().
			Create(ctx, resourcemerge.WithCleanLabelsAndAnnotations(requiredCopy).(*corev1.Namespace), metav1.CreateOptions{})
		reportCreateEvent(recorder, requiredCopy, err)
		return actual, true, err
	}
	if err != nil {
		return nil, false, err
	}

	modified := resourcemerge.BoolPtr(false)
	existingCopy := existing.DeepCopy()

	resourcemerge.EnsureObjectMeta(modified, &existingCopy.ObjectMeta, required.ObjectMeta)
	if !*modified {
		return existingCopy, false, nil
	}

	if klog.V(4).Enabled() {
		klog.Infof("Namespace %q changes: %v", required.Name, JSONPatchNoError(existing, existingCopy))
	}

	actual, err := client.Namespaces().Update(ctx, existingCopy, metav1.UpdateOptions{})
	reportUpdateEvent(recorder, required, err)
	return actual, true, err
}

// ApplyService merges objectmeta and requires
// TODO, since this cannot determine whether changes are due to legitimate actors (api server) or illegitimate ones (users), we cannot update
// TODO I've special cased the selector for now
func ApplyService(ctx context.Context, client coreclientv1.ServicesGetter, recorder events.Recorder, required *corev1.Service) (*corev1.Service, bool, error) {
	existing, err := client.Services(required.Namespace).Get(ctx, required.Name, metav1.GetOptions{})
	if apierrors.IsNotFound(err) {
		requiredCopy := required.DeepCopy()
		actual, err := client.Services(requiredCopy.Namespace).
			Create(ctx, resourcemerge.WithCleanLabelsAndAnnotations(requiredCopy).(*corev1.Service), metav1.CreateOptions{})
		reportCreateEvent(recorder, requiredCopy, err)
		return actual, true, err
	}
	if err != nil {
		return nil, false, err
	}

	modified := resourcemerge.BoolPtr(false)
	existingCopy := existing.DeepCopy()

	resourcemerge.EnsureObjectMeta(modified, &existingCopy.ObjectMeta, required.ObjectMeta)
	selectorSame := equality.Semantic.DeepEqual(existingCopy.Spec.Selector, required.Spec.Selector)

	typeSame := false
	requiredIsEmpty := len(required.Spec.Type) == 0
	existingCopyIsCluster := existingCopy.Spec.Type == corev1.ServiceTypeClusterIP
	if (requiredIsEmpty && existingCopyIsCluster) || equality.Semantic.DeepEqual(existingCopy.Spec.Type, required.Spec.Type) {
		typeSame = true
	}

	if selectorSame && typeSame && !*modified {
		return existingCopy, false, nil
	}

	existingCopy.Spec.Selector = required.Spec.Selector
	existingCopy.Spec.Type = required.Spec.Type // if this is different, the update will fail.  Status will indicate it.

	if klog.V(4).Enabled() {
		klog.Infof("Service %q changes: %v", required.Namespace+"/"+required.Name, JSONPatchNoError(existing, required))
	}

	actual, err := client.Services(required.Namespace).Update(ctx, existingCopy, metav1.UpdateOptions{})
	reportUpdateEvent(recorder, required, err)
	return actual, true, err
}

// ApplyPod merges objectmeta, does not worry about anything else
func ApplyPod(ctx context.Context, client coreclientv1.PodsGetter, recorder events.Recorder, required *corev1.Pod) (*corev1.Pod, bool, error) {
	existing, err := client.Pods(required.Namespace).Get(ctx, required.Name, metav1.GetOptions{})
	if apierrors.IsNotFound(err) {
		requiredCopy := required.DeepCopy()
		actual, err := client.Pods(requiredCopy.Namespace).
			Create(ctx, resourcemerge.WithCleanLabelsAndAnnotations(requiredCopy).(*corev1.Pod), metav1.CreateOptions{})
		reportCreateEvent(recorder, requiredCopy, err)
		return actual, true, err
	}
	if err != nil {
		return nil, false, err
	}

	modified := resourcemerge.BoolPtr(false)
	existingCopy := existing.DeepCopy()

	resourcemerge.EnsureObjectMeta(modified, &existingCopy.ObjectMeta, required.ObjectMeta)
	if !*modified {
		return existingCopy, false, nil
	}

	if klog.V(4).Enabled() {
		klog.Infof("Pod %q changes: %v", required.Namespace+"/"+required.Name, JSONPatchNoError(existing, required))
	}

	actual, err := client.Pods(required.Namespace).Update(ctx, existingCopy, metav1.UpdateOptions{})
	reportUpdateEvent(recorder, required, err)
	return actual, true, err
}

// ApplyServiceAccount merges objectmeta, does not worry about anything else
func ApplyServiceAccount(ctx context.Context, client coreclientv1.ServiceAccountsGetter, recorder events.Recorder, required *corev1.ServiceAccount) (*corev1.ServiceAccount, bool, error) {
	existing, err := client.ServiceAccounts(required.Namespace).Get(ctx, required.Name, metav1.GetOptions{})
	if apierrors.IsNotFound(err) {
		requiredCopy := required.DeepCopy()
		actual, err := client.ServiceAccounts(requiredCopy.Namespace).
			Create(ctx, resourcemerge.WithCleanLabelsAndAnnotations(requiredCopy).(*corev1.ServiceAccount), metav1.CreateOptions{})
		reportCreateEvent(recorder, requiredCopy, err)
		return actual, true, err
	}
	if err != nil {
		return nil, false, err
	}

	modified := resourcemerge.BoolPtr(false)
	existingCopy := existing.DeepCopy()

	resourcemerge.EnsureObjectMeta(modified, &existingCopy.ObjectMeta, required.ObjectMeta)
	if !*modified {
		return existingCopy, false, nil
	}
	if klog.V(4).Enabled() {
		klog.Infof("ServiceAccount %q changes: %v", required.Namespace+"/"+required.Name, JSONPatchNoError(existing, required))
	}
	actual, err := client.ServiceAccounts(required.Namespace).Update(ctx, existingCopy, metav1.UpdateOptions{})
	reportUpdateEvent(recorder, required, err)
	return actual, true, err
}

// ApplyConfigMap merges objectmeta, requires data
func ApplyConfigMap(ctx context.Context, client coreclientv1.ConfigMapsGetter, recorder events.Recorder, required *corev1.ConfigMap) (*corev1.ConfigMap, bool, error) {
	existing, err := client.ConfigMaps(required.Namespace).Get(ctx, required.Name, metav1.GetOptions{})
	if apierrors.IsNotFound(err) {
		requiredCopy := required.DeepCopy()
		actual, err := client.ConfigMaps(requiredCopy.Namespace).
			Create(ctx, resourcemerge.WithCleanLabelsAndAnnotations(requiredCopy).(*corev1.ConfigMap), metav1.CreateOptions{})
		reportCreateEvent(recorder, requiredCopy, err)
		return actual, true, err
	}
	if err != nil {
		return nil, false, err
	}

	modified := resourcemerge.BoolPtr(false)
	existingCopy := existing.DeepCopy()

	resourcemerge.EnsureObjectMeta(modified, &existingCopy.ObjectMeta, required.ObjectMeta)

	caBundleInjected := required.Labels["config.uccp.io/inject-trusted-cabundle"] == "true"
	_, newCABundleRequired := required.Data["ca-bundle.crt"]

	var modifiedKeys []string
	for existingCopyKey, existingCopyValue := range existingCopy.Data {
		// if we're injecting a ca-bundle and the required isn't forcing the value, then don't use the value of existing
		// to drive a diff detection. If required has set the value then we need to force the value in order to have apply
		// behave predictably.
		if caBundleInjected && !newCABundleRequired && existingCopyKey == "ca-bundle.crt" {
			continue
		}
		if requiredValue, ok := required.Data[existingCopyKey]; !ok || (existingCopyValue != requiredValue) {
			modifiedKeys = append(modifiedKeys, "data."+existingCopyKey)
		}
	}
	for existingCopyKey, existingCopyBinValue := range existingCopy.BinaryData {
		if requiredBinValue, ok := required.BinaryData[existingCopyKey]; !ok || !bytes.Equal(existingCopyBinValue, requiredBinValue) {
			modifiedKeys = append(modifiedKeys, "binaryData."+existingCopyKey)
		}
	}
	for requiredKey := range required.Data {
		if _, ok := existingCopy.Data[requiredKey]; !ok {
			modifiedKeys = append(modifiedKeys, "data."+requiredKey)
		}
	}
	for requiredBinKey := range required.BinaryData {
		if _, ok := existingCopy.BinaryData[requiredBinKey]; !ok {
			modifiedKeys = append(modifiedKeys, "binaryData."+requiredBinKey)
		}
	}

	dataSame := len(modifiedKeys) == 0
	if dataSame && !*modified {
		return existingCopy, false, nil
	}
	existingCopy.Data = required.Data
	existingCopy.BinaryData = required.BinaryData
	// if we're injecting a cabundle, and we had a previous value, and the required object isn't setting the value, then set back to the previous
	if existingCABundle, existedBefore := existing.Data["ca-bundle.crt"]; caBundleInjected && existedBefore && !newCABundleRequired {
		if existingCopy.Data == nil {
			existingCopy.Data = map[string]string{}
		}
		existingCopy.Data["ca-bundle.crt"] = existingCABundle
	}

	actual, err := client.ConfigMaps(required.Namespace).Update(ctx, existingCopy, metav1.UpdateOptions{})

	var details string
	if !dataSame {
		sort.Sort(sort.StringSlice(modifiedKeys))
		details = fmt.Sprintf("cause by changes in %v", strings.Join(modifiedKeys, ","))
	}
	if klog.V(4).Enabled() {
		klog.Infof("ConfigMap %q changes: %v", required.Namespace+"/"+required.Name, JSONPatchNoError(existing, required))
	}
	reportUpdateEvent(recorder, required, err, details)
	return actual, true, err
}

// ApplySecret merges objectmeta, requires data
func ApplySecret(ctx context.Context, client coreclientv1.SecretsGetter, recorder events.Recorder, requiredInput *corev1.Secret) (*corev1.Secret, bool, error) {
	// copy the stringData to data.  Error on a data content conflict inside required.  This is usually a bug.
	required := requiredInput.DeepCopy()
	if required.Data == nil {
		required.Data = map[string][]byte{}
	}
	for k, v := range required.StringData {
		if dataV, ok := required.Data[k]; ok {
			if string(dataV) != v {
				return nil, false, fmt.Errorf("Secret.stringData[%q] conflicts with Secret.data[%q]", k, k)
			}
		}
		required.Data[k] = []byte(v)
	}
	required.StringData = nil

	existing, err := client.Secrets(required.Namespace).Get(ctx, required.Name, metav1.GetOptions{})
	if apierrors.IsNotFound(err) {
		requiredCopy := required.DeepCopy()
		actual, err := client.Secrets(requiredCopy.Namespace).
			Create(ctx, resourcemerge.WithCleanLabelsAndAnnotations(requiredCopy).(*corev1.Secret), metav1.CreateOptions{})
		reportCreateEvent(recorder, requiredCopy, err)
		return actual, true, err
	}
	if err != nil {
		return nil, false, err
	}

	existingCopy := existing.DeepCopy()

	resourcemerge.EnsureObjectMeta(resourcemerge.BoolPtr(false), &existingCopy.ObjectMeta, required.ObjectMeta)

	switch required.Type {
	case corev1.SecretTypeServiceAccountToken:
		// Secrets for ServiceAccountTokens will have data injected by kube controller manager.
		// We will apply only the explicitly set keys.
		if existingCopy.Data == nil {
			existingCopy.Data = map[string][]byte{}
		}

		for k, v := range required.Data {
			existingCopy.Data[k] = v
		}

	default:
		existingCopy.Data = required.Data
	}

	existingCopy.Type = required.Type

	// Server defaults some values and we need to do it as well or it will never equal.
	if existingCopy.Type == "" {
		existingCopy.Type = corev1.SecretTypeOpaque
	}

	if equality.Semantic.DeepEqual(existingCopy, existing) {
		return existing, false, nil
	}

	if klog.V(4).Enabled() {
		klog.Infof("Secret %s/%s changes: %v", required.Namespace, required.Name, JSONPatchSecretNoError(existing, existingCopy))
	}

	var actual *corev1.Secret
	/*
	 * Kubernetes validation silently hides failures to update secret type.
	 * https://github.com/kubernetes/kubernetes/blob/98e65951dccfd40d3b4f31949c2ab8df5912d93e/pkg/apis/core/validation/validation.go#L5048
	 * We need to explicitly opt for delete+create in that case.
	 */
	if existingCopy.Type == existing.Type {
		actual, err = client.Secrets(required.Namespace).Update(ctx, existingCopy, metav1.UpdateOptions{})
		reportUpdateEvent(recorder, existingCopy, err)

		if err == nil {
			return actual, true, err
		}
		if !strings.Contains(err.Error(), "field is immutable") {
			return actual, true, err
		}
	}

	// if the field was immutable on a secret, we're going to be stuck until we delete it.  Try to delete and then create
	deleteErr := client.Secrets(required.Namespace).Delete(ctx, existingCopy.Name, metav1.DeleteOptions{})
	reportDeleteEvent(recorder, existingCopy, deleteErr)

	// clear the RV and track the original actual and error for the return like our create value.
	existingCopy.ResourceVersion = ""
	actual, err = client.Secrets(required.Namespace).Create(ctx, existingCopy, metav1.CreateOptions{})
	reportCreateEvent(recorder, existingCopy, err)

	return actual, true, err
}

// SyncConfigMap applies a ConfigMap from a location `sourceNamespace/sourceName` to `targetNamespace/targetName`
func SyncConfigMap(ctx context.Context, client coreclientv1.ConfigMapsGetter, recorder events.Recorder, sourceNamespace, sourceName, targetNamespace, targetName string, ownerRefs []metav1.OwnerReference) (*corev1.ConfigMap, bool, error) {
	return SyncPartialConfigMap(ctx, client, recorder, sourceNamespace, sourceName, targetNamespace, targetName, nil, ownerRefs)
}

// SyncPartialConfigMap does what SyncConfigMap does but it only synchronizes a subset of keys given by `syncedKeys`.
// SyncPartialConfigMap will delete the target if `syncedKeys` are set but the source does not contain any of these keys.
func SyncPartialConfigMap(ctx context.Context, client coreclientv1.ConfigMapsGetter, recorder events.Recorder, sourceNamespace, sourceName, targetNamespace, targetName string, syncedKeys sets.String, ownerRefs []metav1.OwnerReference) (*corev1.ConfigMap, bool, error) {
	source, err := client.ConfigMaps(sourceNamespace).Get(ctx, sourceName, metav1.GetOptions{})
	switch {
	case apierrors.IsNotFound(err):
		modified, err := deleteConfigMapSyncTarget(ctx, client, recorder, targetNamespace, targetName)
		return nil, modified, err
	case err != nil:
		return nil, false, err
	default:
		if len(syncedKeys) > 0 {
			for sourceKey := range source.Data {
				if !syncedKeys.Has(sourceKey) {
					delete(source.Data, sourceKey)
				}
			}
			for sourceKey := range source.BinaryData {
				if !syncedKeys.Has(sourceKey) {
					delete(source.BinaryData, sourceKey)
				}
			}

			// remove the synced CM if the requested fields are not present in source
			if len(source.Data)+len(source.BinaryData) == 0 {
				modified, err := deleteConfigMapSyncTarget(ctx, client, recorder, targetNamespace, targetName)
				return nil, modified, err
			}
		}

		source.Namespace = targetNamespace
		source.Name = targetName
		source.ResourceVersion = ""
		source.OwnerReferences = ownerRefs
		return ApplyConfigMap(ctx, client, recorder, source)
	}
}

func deleteConfigMapSyncTarget(ctx context.Context, client coreclientv1.ConfigMapsGetter, recorder events.Recorder, targetNamespace, targetName string) (bool, error) {
	err := client.ConfigMaps(targetNamespace).Delete(ctx, targetName, metav1.DeleteOptions{})
	if apierrors.IsNotFound(err) {
		return false, nil
	}
	if err == nil {
		recorder.Eventf("TargetConfigDeleted", "Deleted target configmap %s/%s because source config does not exist", targetNamespace, targetName)
		return true, nil
	}
	return false, err
}

// SyncSecret applies a Secret from a location `sourceNamespace/sourceName` to `targetNamespace/targetName`
func SyncSecret(ctx context.Context, client coreclientv1.SecretsGetter, recorder events.Recorder, sourceNamespace, sourceName, targetNamespace, targetName string, ownerRefs []metav1.OwnerReference) (*corev1.Secret, bool, error) {
	return SyncPartialSecret(ctx, client, recorder, sourceNamespace, sourceName, targetNamespace, targetName, nil, ownerRefs)
}

// SyncPartialSecret does what SyncSecret does but it only synchronizes a subset of keys given by `syncedKeys`.
// SyncPartialSecret will delete the target if `syncedKeys` are set but the source does not contain any of these keys.
func SyncPartialSecret(ctx context.Context, client coreclientv1.SecretsGetter, recorder events.Recorder, sourceNamespace, sourceName, targetNamespace, targetName string, syncedKeys sets.String, ownerRefs []metav1.OwnerReference) (*corev1.Secret, bool, error) {
	source, err := client.Secrets(sourceNamespace).Get(ctx, sourceName, metav1.GetOptions{})
	switch {
	case apierrors.IsNotFound(err):
		modified, err := deleteSecretSyncTarget(ctx, client, recorder, targetNamespace, targetName)
		return nil, modified, err
	case err != nil:
		return nil, false, err
	default:
		if source.Type == corev1.SecretTypeServiceAccountToken {

			// Make sure the token is already present, otherwise we have to wait before creating the target
			if len(source.Data[corev1.ServiceAccountTokenKey]) == 0 {
				return nil, false, fmt.Errorf("secret %s/%s doesn't have a token yet", source.Namespace, source.Name)
			}

			if source.Annotations != nil {
				// When syncing a service account token we have to remove the SA annotation to disable injection into copies
				delete(source.Annotations, corev1.ServiceAccountNameKey)
				// To make it clean, remove the dormant annotations as well
				delete(source.Annotations, corev1.ServiceAccountUIDKey)
			}

			// SecretTypeServiceAccountToken implies required fields and injection which we do not want in copies
			source.Type = corev1.SecretTypeOpaque
		}

		if len(syncedKeys) > 0 {
			for sourceKey := range source.Data {
				if !syncedKeys.Has(sourceKey) {
					delete(source.Data, sourceKey)
				}
			}
			for sourceKey := range source.StringData {
				if !syncedKeys.Has(sourceKey) {
					delete(source.StringData, sourceKey)
				}
			}

			// remove the synced secret if the requested fields are not present in source
			if len(source.Data)+len(source.StringData) == 0 {
				modified, err := deleteSecretSyncTarget(ctx, client, recorder, targetNamespace, targetName)
				return nil, modified, err
			}
		}

		source.Namespace = targetNamespace
		source.Name = targetName
		source.ResourceVersion = ""
		source.OwnerReferences = ownerRefs
		return ApplySecret(ctx, client, recorder, source)
	}
}

func deleteSecretSyncTarget(ctx context.Context, client coreclientv1.SecretsGetter, recorder events.Recorder, targetNamespace, targetName string) (bool, error) {
	err := client.Secrets(targetNamespace).Delete(ctx, targetName, metav1.DeleteOptions{})
	if apierrors.IsNotFound(err) {
		return false, nil
	}
	if err == nil {
		recorder.Eventf("TargetSecretDeleted", "Deleted target secret %s/%s because source config does not exist", targetNamespace, targetName)
		return true, nil
	}
	return false, err
}
