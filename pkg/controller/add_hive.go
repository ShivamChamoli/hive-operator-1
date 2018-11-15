package controller

import (
	"github.com/openshift/hive-operator/pkg/controller/hive"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, hive.Add)
}
