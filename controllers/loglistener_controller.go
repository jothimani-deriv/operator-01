package controllers

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	ctrllog "sigs.k8s.io/controller-runtime/pkg/log"

	cachev1alpha1 "github.com/jothimani-deriv/operator-01/api/v1alpha1"
)

// LogListenerReconciler reconciles a LogListener object
type LogListenerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

func (r *LogListenerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := ctrllog.FromContext(ctx)
	// Fetch the Memcached instance
	loglistenerValue := &cachev1alpha1.LogListener{}
	err := r.Get(ctx, req.NamespacedName, loglistenerValue)
	if err != nil {
		if errors.IsNotFound(err) {
			log.Info("Memcached resource not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request.
		log.Error(err, "Failed to get Memcached")
		return ctrl.Result{}, err
	}

	logMessage := loglistenerValue.Spec.Trigger
	fmt.Println("Debug: Trigger value is:", logMessage)
	log.Info("Applied size: ", "trigger", logMessage)
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *LogListenerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&cachev1alpha1.LogListener{}).
		Complete(r)
}
