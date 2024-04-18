// Copyright (c) 2019-2024 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package virtualmachineclass

import (
	ctrlmgr "sigs.k8s.io/controller-runtime/pkg/manager"

	pkgctx "github.com/vmware-tanzu/vm-operator/pkg/context"
	"github.com/vmware-tanzu/vm-operator/webhooks/virtualmachineclass/mutation"
	"github.com/vmware-tanzu/vm-operator/webhooks/virtualmachineclass/validation"
)

func AddToManager(ctx *pkgctx.ControllerManagerContext, mgr ctrlmgr.Manager) error {
	if err := mutation.AddToManager(ctx, mgr); err != nil {
		return err
	}
	return validation.AddToManager(ctx, mgr)
}
