// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	foldercustomrole "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/authorization/foldercustomrole"
	folderroleassignment "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/authorization/folderroleassignment"
	organizationcustomrole "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/authorization/organizationcustomrole"
	organizationroleassignment "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/authorization/organizationroleassignment"
	projectcustomrole "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/authorization/projectcustomrole"
	projectroleassignment "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/authorization/projectroleassignment"
	serviceaccountroleassignment "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/authorization/serviceaccountroleassignment"
	providerconfig "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/providerconfig"
	folder "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/resourcemanager/folder"
	project "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/resourcemanager/project"
	federatedidentityprovider "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/serviceaccount/federatedidentityprovider"
	key "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/serviceaccount/key"
	serviceaccount "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/serviceaccount/serviceaccount"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		foldercustomrole.Setup,
		folderroleassignment.Setup,
		organizationcustomrole.Setup,
		organizationroleassignment.Setup,
		projectcustomrole.Setup,
		projectroleassignment.Setup,
		serviceaccountroleassignment.Setup,
		providerconfig.Setup,
		folder.Setup,
		project.Setup,
		federatedidentityprovider.Setup,
		key.Setup,
		serviceaccount.Setup,
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
		foldercustomrole.SetupGated,
		folderroleassignment.SetupGated,
		organizationcustomrole.SetupGated,
		organizationroleassignment.SetupGated,
		projectcustomrole.SetupGated,
		projectroleassignment.SetupGated,
		serviceaccountroleassignment.SetupGated,
		providerconfig.SetupGated,
		folder.SetupGated,
		project.SetupGated,
		federatedidentityprovider.SetupGated,
		key.SetupGated,
		serviceaccount.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
