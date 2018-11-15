// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	machineconfigurationopenshiftiov1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMachineConfigs implements MachineConfigInterface
type FakeMachineConfigs struct {
	Fake *FakeMachineconfigurationV1
}

var machineconfigsResource = schema.GroupVersionResource{Group: "machineconfiguration.openshift.io", Version: "v1", Resource: "machineconfigs"}

var machineconfigsKind = schema.GroupVersionKind{Group: "machineconfiguration.openshift.io", Version: "v1", Kind: "MachineConfig"}

// Get takes name of the machineConfig, and returns the corresponding machineConfig object, and an error if there is any.
func (c *FakeMachineConfigs) Get(name string, options v1.GetOptions) (result *machineconfigurationopenshiftiov1.MachineConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(machineconfigsResource, name), &machineconfigurationopenshiftiov1.MachineConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*machineconfigurationopenshiftiov1.MachineConfig), err
}

// List takes label and field selectors, and returns the list of MachineConfigs that match those selectors.
func (c *FakeMachineConfigs) List(opts v1.ListOptions) (result *machineconfigurationopenshiftiov1.MachineConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(machineconfigsResource, machineconfigsKind, opts), &machineconfigurationopenshiftiov1.MachineConfigList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &machineconfigurationopenshiftiov1.MachineConfigList{ListMeta: obj.(*machineconfigurationopenshiftiov1.MachineConfigList).ListMeta}
	for _, item := range obj.(*machineconfigurationopenshiftiov1.MachineConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested machineConfigs.
func (c *FakeMachineConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(machineconfigsResource, opts))
}

// Create takes the representation of a machineConfig and creates it.  Returns the server's representation of the machineConfig, and an error, if there is any.
func (c *FakeMachineConfigs) Create(machineConfig *machineconfigurationopenshiftiov1.MachineConfig) (result *machineconfigurationopenshiftiov1.MachineConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(machineconfigsResource, machineConfig), &machineconfigurationopenshiftiov1.MachineConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*machineconfigurationopenshiftiov1.MachineConfig), err
}

// Update takes the representation of a machineConfig and updates it. Returns the server's representation of the machineConfig, and an error, if there is any.
func (c *FakeMachineConfigs) Update(machineConfig *machineconfigurationopenshiftiov1.MachineConfig) (result *machineconfigurationopenshiftiov1.MachineConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(machineconfigsResource, machineConfig), &machineconfigurationopenshiftiov1.MachineConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*machineconfigurationopenshiftiov1.MachineConfig), err
}

// Delete takes name of the machineConfig and deletes it. Returns an error if one occurs.
func (c *FakeMachineConfigs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(machineconfigsResource, name), &machineconfigurationopenshiftiov1.MachineConfig{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMachineConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(machineconfigsResource, listOptions)

	_, err := c.Fake.Invokes(action, &machineconfigurationopenshiftiov1.MachineConfigList{})
	return err
}

// Patch applies the patch and returns the patched machineConfig.
func (c *FakeMachineConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *machineconfigurationopenshiftiov1.MachineConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(machineconfigsResource, name, data, subresources...), &machineconfigurationopenshiftiov1.MachineConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*machineconfigurationopenshiftiov1.MachineConfig), err
}
