package legacygroupification

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	appsv1 "github.com/uccps-samples/api/apps/v1"
	authorizationv1 "github.com/uccps-samples/api/authorization/v1"
	buildv1 "github.com/uccps-samples/api/build/v1"
	imagev1 "github.com/uccps-samples/api/image/v1"
	networkv1 "github.com/uccps-samples/api/network/v1"
	oauthv1 "github.com/uccps-samples/api/oauth/v1"
	projectv1 "github.com/uccps-samples/api/project/v1"
	quotav1 "github.com/uccps-samples/api/quota/v1"
	routev1 "github.com/uccps-samples/api/route/v1"
	securityv1 "github.com/uccps-samples/api/security/v1"
	templatev1 "github.com/uccps-samples/api/template/v1"
	userv1 "github.com/uccps-samples/api/user/v1"
)

// deprecated
func IsOAPI(gvk schema.GroupVersionKind) bool {
	if len(gvk.Group) > 0 {
		return false
	}

	_, ok := oapiKindsToGroup[gvk.Kind]
	return ok
}

// deprecated
func OAPIToGroupifiedGVK(gvk *schema.GroupVersionKind) {
	if len(gvk.Group) > 0 {
		return
	}

	newGroup, ok := oapiKindsToGroup[gvk.Kind]
	if !ok {
		return
	}
	gvk.Group = newGroup
}

// deprecated
func OAPIToGroupified(uncast runtime.Object, gvk *schema.GroupVersionKind) {
	if len(gvk.Group) > 0 {
		return
	}

	switch obj := uncast.(type) {
	case *unstructured.Unstructured:
		newGroup := fixOAPIGroupKindInTopLevelUnstructured(obj.Object)
		if len(newGroup) > 0 {
			gvk.Group = newGroup
			uncast.GetObjectKind().SetGroupVersionKind(*gvk)
		}
	case *unstructured.UnstructuredList:
		newGroup := fixOAPIGroupKindInTopLevelUnstructured(obj.Object)
		if len(newGroup) > 0 {
			gvk.Group = newGroup
			uncast.GetObjectKind().SetGroupVersionKind(*gvk)
		}

	case *appsv1.DeploymentConfig, *appsv1.DeploymentConfigList,
		*appsv1.DeploymentConfigRollback,
		*appsv1.DeploymentLog,
		*appsv1.DeploymentRequest:
		gvk.Group = appsv1.GroupName
		uncast.GetObjectKind().SetGroupVersionKind(*gvk)

	case *authorizationv1.ClusterRoleBinding, *authorizationv1.ClusterRoleBindingList,
		*authorizationv1.ClusterRole, *authorizationv1.ClusterRoleList,
		*authorizationv1.Role, *authorizationv1.RoleList,
		*authorizationv1.RoleBinding, *authorizationv1.RoleBindingList,
		*authorizationv1.RoleBindingRestriction, *authorizationv1.RoleBindingRestrictionList,
		*authorizationv1.SubjectRulesReview, *authorizationv1.SelfSubjectRulesReview,
		*authorizationv1.ResourceAccessReview, *authorizationv1.LocalResourceAccessReview,
		*authorizationv1.SubjectAccessReview, *authorizationv1.LocalSubjectAccessReview:
		gvk.Group = authorizationv1.GroupName
		uncast.GetObjectKind().SetGroupVersionKind(*gvk)

	case *buildv1.BuildConfig, *buildv1.BuildConfigList,
		*buildv1.Build, *buildv1.BuildList,
		*buildv1.BuildLog,
		*buildv1.BuildRequest,
		*buildv1.BinaryBuildRequestOptions:
		gvk.Group = buildv1.GroupName
		uncast.GetObjectKind().SetGroupVersionKind(*gvk)

	case *imagev1.Image, *imagev1.ImageList,
		*imagev1.ImageSignature,
		*imagev1.ImageStreamImage,
		*imagev1.ImageStreamImport,
		*imagev1.ImageStreamMapping,
		*imagev1.ImageStream, *imagev1.ImageStreamList,
		*imagev1.ImageStreamTag:
		gvk.Group = imagev1.GroupName
		uncast.GetObjectKind().SetGroupVersionKind(*gvk)

	case *networkv1.ClusterNetwork, *networkv1.ClusterNetworkList,
		*networkv1.NetNamespace, *networkv1.NetNamespaceList,
		*networkv1.HostSubnet, *networkv1.HostSubnetList,
		*networkv1.EgressNetworkPolicy, *networkv1.EgressNetworkPolicyList:
		gvk.Group = networkv1.GroupName
		uncast.GetObjectKind().SetGroupVersionKind(*gvk)

	case *projectv1.Project, *projectv1.ProjectList,
		*projectv1.ProjectRequest:
		gvk.Group = projectv1.GroupName
		uncast.GetObjectKind().SetGroupVersionKind(*gvk)

	case *quotav1.ClusterResourceQuota, *quotav1.ClusterResourceQuotaList,
		*quotav1.AppliedClusterResourceQuota, *quotav1.AppliedClusterResourceQuotaList:
		gvk.Group = quotav1.GroupName
		uncast.GetObjectKind().SetGroupVersionKind(*gvk)

	case *oauthv1.OAuthAuthorizeToken, *oauthv1.OAuthAuthorizeTokenList,
		*oauthv1.OAuthClientAuthorization, *oauthv1.OAuthClientAuthorizationList,
		*oauthv1.OAuthClient, *oauthv1.OAuthClientList,
		*oauthv1.OAuthAccessToken, *oauthv1.OAuthAccessTokenList:
		gvk.Group = oauthv1.GroupName
		uncast.GetObjectKind().SetGroupVersionKind(*gvk)

	case *routev1.Route, *routev1.RouteList:
		gvk.Group = routev1.GroupName
		uncast.GetObjectKind().SetGroupVersionKind(*gvk)

	case *securityv1.SecurityContextConstraints, *securityv1.SecurityContextConstraintsList,
		*securityv1.PodSecurityPolicySubjectReview,
		*securityv1.PodSecurityPolicySelfSubjectReview,
		*securityv1.PodSecurityPolicyReview:
		gvk.Group = securityv1.GroupName
		uncast.GetObjectKind().SetGroupVersionKind(*gvk)

	case *templatev1.Template, *templatev1.TemplateList:
		gvk.Group = templatev1.GroupName
		uncast.GetObjectKind().SetGroupVersionKind(*gvk)

	case *userv1.Group, *userv1.GroupList,
		*userv1.Identity, *userv1.IdentityList,
		*userv1.UserIdentityMapping,
		*userv1.User, *userv1.UserList:
		gvk.Group = userv1.GroupName
		uncast.GetObjectKind().SetGroupVersionKind(*gvk)

	}
}

var oapiKindsToGroup = map[string]string{
	"DeploymentConfigRollback": "apps.uccp.io",
	"DeploymentConfig":         "apps.uccp.io", "DeploymentConfigList": "apps.uccp.io",
	"DeploymentLog":      "apps.uccp.io",
	"DeploymentRequest":  "apps.uccp.io",
	"ClusterRoleBinding": "authorization.uccp.io", "ClusterRoleBindingList": "authorization.uccp.io",
	"ClusterRole": "authorization.uccp.io", "ClusterRoleList": "authorization.uccp.io",
	"RoleBindingRestriction": "authorization.uccp.io", "RoleBindingRestrictionList": "authorization.uccp.io",
	"RoleBinding": "authorization.uccp.io", "RoleBindingList": "authorization.uccp.io",
	"Role": "authorization.uccp.io", "RoleList": "authorization.uccp.io",
	"SubjectRulesReview": "authorization.uccp.io", "SelfSubjectRulesReview": "authorization.uccp.io",
	"ResourceAccessReview": "authorization.uccp.io", "LocalResourceAccessReview": "authorization.uccp.io",
	"SubjectAccessReview": "authorization.uccp.io", "LocalSubjectAccessReview": "authorization.uccp.io",
	"BuildConfig": "build.uccp.io", "BuildConfigList": "build.uccp.io",
	"Build": "build.uccp.io", "BuildList": "build.uccp.io",
	"BinaryBuildRequestOptions": "build.uccp.io",
	"BuildLog":                  "build.uccp.io",
	"BuildRequest":              "build.uccp.io",
	"Image":                     "image.uccp.io", "ImageList": "image.uccp.io",
	"ImageSignature":     "image.uccp.io",
	"ImageStreamImage":   "image.uccp.io",
	"ImageStreamImport":  "image.uccp.io",
	"ImageStreamMapping": "image.uccp.io",
	"ImageStream":        "image.uccp.io", "ImageStreamList": "image.uccp.io",
	"ImageStreamTag": "image.uccp.io", "ImageStreamTagList": "image.uccp.io",
	"ClusterNetwork": "network.uccp.io", "ClusterNetworkList": "network.uccp.io",
	"EgressNetworkPolicy": "network.uccp.io", "EgressNetworkPolicyList": "network.uccp.io",
	"HostSubnet": "network.uccp.io", "HostSubnetList": "network.uccp.io",
	"NetNamespace": "network.uccp.io", "NetNamespaceList": "network.uccp.io",
	"OAuthAccessToken": "oauth.uccp.io", "OAuthAccessTokenList": "oauth.uccp.io",
	"OAuthAuthorizeToken": "oauth.uccp.io", "OAuthAuthorizeTokenList": "oauth.uccp.io",
	"OAuthClientAuthorization": "oauth.uccp.io", "OAuthClientAuthorizationList": "oauth.uccp.io",
	"OAuthClient": "oauth.uccp.io", "OAuthClientList": "oauth.uccp.io",
	"Project": "project.uccp.io", "ProjectList": "project.uccp.io",
	"ProjectRequest":       "project.uccp.io",
	"ClusterResourceQuota": "quota.uccp.io", "ClusterResourceQuotaList": "quota.uccp.io",
	"AppliedClusterResourceQuota": "quota.uccp.io", "AppliedClusterResourceQuotaList": "quota.uccp.io",
	"Route": "route.uccp.io", "RouteList": "route.uccp.io",
	"SecurityContextConstraints": "security.uccp.io", "SecurityContextConstraintsList": "security.uccp.io",
	"PodSecurityPolicySubjectReview":     "security.uccp.io",
	"PodSecurityPolicySelfSubjectReview": "security.uccp.io",
	"PodSecurityPolicyReview":            "security.uccp.io",
	"Template":                           "template.uccp.io", "TemplateList": "template.uccp.io",
	"Group": "user.uccp.io", "GroupList": "user.uccp.io",
	"Identity": "user.uccp.io", "IdentityList": "user.uccp.io",
	"UserIdentityMapping": "user.uccp.io",
	"User":                "user.uccp.io", "UserList": "user.uccp.io",
}

func fixOAPIGroupKindInTopLevelUnstructured(obj map[string]interface{}) string {
	kind, ok := obj["kind"]
	if !ok {
		return ""
	}
	kindStr, ok := kind.(string)
	if !ok {
		return ""
	}
	newGroup, ok := oapiKindsToGroup[kindStr]
	if !ok {
		return ""
	}

	apiVersion, ok := obj["apiVersion"]
	if !ok {
		return newGroup
	}
	apiVersionStr, ok := apiVersion.(string)
	if !ok {
		return newGroup
	}

	if apiVersionStr != "v1" {
		return newGroup
	}
	obj["apiVersion"] = newGroup + "/v1"

	return newGroup
}
