/*
Copyright 2025.

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

package v1

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	webv1 "github.com/pbguyBookworm/my-website/api/v1"
)

// nolint:unused
// log is for logging in this package.
var webpagelog = logf.Log.WithName("webpage-resource")

// SetupWebpageWebhookWithManager registers the webhook for Webpage in the manager.
func SetupWebpageWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&webv1.Webpage{}).
		WithValidator(&WebpageCustomValidator{}).
		WithDefaulter(&WebpageCustomDefaulter{}).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-web-mywebsite-com-v1-webpage,mutating=true,failurePolicy=fail,sideEffects=None,groups=web.mywebsite.com,resources=webpages,verbs=create;update,versions=v1,name=mwebpage-v1.kb.io,admissionReviewVersions=v1

// WebpageCustomDefaulter struct is responsible for setting default values on the custom resource of the
// Kind Webpage when those are created or updated.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as it is used only for temporary operations and does not need to be deeply copied.
type WebpageCustomDefaulter struct {
	// TODO(user): Add more fields as needed for defaulting
}

var _ webhook.CustomDefaulter = &WebpageCustomDefaulter{}

// Default implements webhook.CustomDefaulter so a webhook will be registered for the Kind Webpage.
func (d *WebpageCustomDefaulter) Default(_ context.Context, obj runtime.Object) error {
	webpage, ok := obj.(*webv1.Webpage)

	if !ok {
		return fmt.Errorf("expected an Webpage object but got %T", obj)
	}
	webpagelog.Info("Defaulting for Webpage", "name", webpage.GetName())

	// TODO(user): fill in your defaulting logic.
	// Ensure Bar is always Foo + 'bar'
	if webpage.Spec.Foo != nil {
		val := *webpage.Spec.Foo + "bar"
		webpage.Spec.Bar = &val
	} else {
		webpage.Spec.Bar = nil
	}

	return nil
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// NOTE: The 'path' attribute must follow a specific pattern and should not be modified directly here.
// Modifying the path for an invalid path can cause API server errors; failing to locate the webhook.
// +kubebuilder:webhook:path=/validate-web-mywebsite-com-v1-webpage,mutating=false,failurePolicy=fail,sideEffects=None,groups=web.mywebsite.com,resources=webpages,verbs=create;update,versions=v1,name=vwebpage-v1.kb.io,admissionReviewVersions=v1

// WebpageCustomValidator struct is responsible for validating the Webpage resource
// when it is created, updated, or deleted.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as this struct is used only for temporary operations and does not need to be deeply copied.
type WebpageCustomValidator struct {
	// TODO(user): Add more fields as needed for validation
}

var _ webhook.CustomValidator = &WebpageCustomValidator{}

// ValidateCreate implements webhook.CustomValidator so a webhook will be registered for the type Webpage.
func (v *WebpageCustomValidator) ValidateCreate(_ context.Context, obj runtime.Object) (admission.Warnings, error) {
	webpage, ok := obj.(*webv1.Webpage)
	if !ok {
		return nil, fmt.Errorf("expected a Webpage object but got %T", obj)
	}
	webpagelog.Info("Validation for Webpage upon creation", "name", webpage.GetName())

	// TODO(user): fill in your validation logic upon object creation.

	return nil, nil
}

// ValidateUpdate implements webhook.CustomValidator so a webhook will be registered for the type Webpage.
func (v *WebpageCustomValidator) ValidateUpdate(_ context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	webpage, ok := newObj.(*webv1.Webpage)
	if !ok {
		return nil, fmt.Errorf("expected a Webpage object for the newObj but got %T", newObj)
	}
	webpagelog.Info("Validation for Webpage upon update", "name", webpage.GetName())

	// TODO(user): fill in your validation logic upon object update.

	return nil, nil
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type Webpage.
func (v *WebpageCustomValidator) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	webpage, ok := obj.(*webv1.Webpage)
	if !ok {
		return nil, fmt.Errorf("expected a Webpage object but got %T", obj)
	}
	webpagelog.Info("Validation for Webpage upon deletion", "name", webpage.GetName())

	// TODO(user): fill in your validation logic upon object deletion.

	return nil, nil
}
