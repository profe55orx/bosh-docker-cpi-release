package cpi

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/cppforlife/bosh-cpi-go/apiv1"

	bvm "github.com/cppforlife/bosh-docker-cpi/vm"
)

type HasVMMethod struct {
	vmFinder bvm.Finder
}

func NewHasVMMethod(vmFinder bvm.Finder) HasVMMethod {
	return HasVMMethod{vmFinder: vmFinder}
}

func (a HasVMMethod) HasVM(vmCID apiv1.VMCID) (bool, error) {
	vm, err := a.vmFinder.Find(vmCID)
	if err != nil {
		return false, bosherr.WrapErrorf(err, "Finding VM '%s'", vmCID)
	}

	return vm.Exists()
}
