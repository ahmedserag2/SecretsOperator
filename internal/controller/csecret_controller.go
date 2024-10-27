/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	csecretv1alpha1 "github.com/SecretsOperator/api/v1alpha1"


	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	smlistener "github.com/SecretsOperator/internal/gcpSecrets"
)

// CsecretReconciler reconciles a Csecret object
type CsecretReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=csecret.ssecrets,resources=csecrets,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=csecret.ssecrets,resources=csecrets/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=csecret.ssecrets,resources=csecrets/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Csecret object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.17.3/pkg/reconcile
func (r *CsecretReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {

	
	_ = log.FromContext(ctx)

	var SecretClient secretmanager.GCPSecretManagerService
    // Fetch the CSecret instance
	var CSecret csecretv1alpha1.Csecret
    if err := r.Get(ctx, req.NamespacedName, &CSecret); err != nil {
        log.Error(err, "unable to fetch CSecret")
        return ctrl.Result{}, client.IgnoreNotFound(err)
    }

    // Retrieve secretName and projectID from the CSecret spec
    secretName := CSecret.Spec.SecretName
    projectID := CSecret.Spec.ProjectID

	secret_payload = SecretClient.GetSecret(ctx, projectID, SecretName)    
	if err != nil {
        logger.Error(err, "failed to get secret from Google Secret Manager")
        return ctrl.Result{}, err
    }

	// Log the secret payload for debugging
    logger.Info("Retrieved secret payload", "payload", secretPayload)
	
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CsecretReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&csecretv1alpha1.Csecret{}).
		Complete(r)
}
