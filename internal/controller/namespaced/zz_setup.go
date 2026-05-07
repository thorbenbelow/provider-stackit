// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	projectroleassignment "github.com/thorbenbelow/provider-stackit/internal/controller/namespaced/iam/projectroleassignment"
	serviceaccount "github.com/thorbenbelow/provider-stackit/internal/controller/namespaced/iam/serviceaccount"
	providerconfig "github.com/thorbenbelow/provider-stackit/internal/controller/namespaced/providerconfig"
	project "github.com/thorbenbelow/provider-stackit/internal/controller/namespaced/resourcemanager/project"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		projectroleassignment.Setup,
		serviceaccount.Setup,
		providerconfig.Setup,
		project.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		projectroleassignment.SetupGated,
		serviceaccount.SetupGated,
		providerconfig.SetupGated,
		project.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
