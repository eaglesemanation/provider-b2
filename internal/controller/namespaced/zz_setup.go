// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	applicationkey "github.com/eaglesemanation/provider-b2/internal/controller/namespaced/b2/applicationkey"
	bucket "github.com/eaglesemanation/provider-b2/internal/controller/namespaced/b2/bucket"
	bucketfileversion "github.com/eaglesemanation/provider-b2/internal/controller/namespaced/b2/bucketfileversion"
	bucketnotificationrules "github.com/eaglesemanation/provider-b2/internal/controller/namespaced/b2/bucketnotificationrules"
	providerconfig "github.com/eaglesemanation/provider-b2/internal/controller/namespaced/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		applicationkey.Setup,
		bucket.Setup,
		bucketfileversion.Setup,
		bucketnotificationrules.Setup,
		providerconfig.Setup,
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
		applicationkey.SetupGated,
		bucket.SetupGated,
		bucketfileversion.SetupGated,
		bucketnotificationrules.SetupGated,
		providerconfig.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
