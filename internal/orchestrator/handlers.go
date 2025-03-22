package orchestrator

import (
	// api types

	seterav1 "github/setera/pkg/api/setera.com/v1"
)

type Event string

const (
	AddEvent     Event = "Add"
	UpdateEvent  Event = "Update"
	DeleteEvent  Event = "Delete"
	UnknownEvent Event = "Unknown"
)

// add tenant key to the workqueue - add event
func (o *Orchestrator) addTenantHandler(obj any) {

	tenant, ok := obj.(*seterav1.Tenant)
	if !ok {
		o.logger.Info("Failed to cast object to tenant in add handler")

	}

	o.logger.Info("Adding tenant", tenant.Name)

	o.enqueue(tenant, AddEvent)

}

// add tenant key to the workqueue - update event
func (o *Orchestrator) updateTenantHandler(oldObj, newObj any) {

	oldTenant, ok := oldObj.(*seterav1.Tenant)
	if !ok {
		o.logger.Info("Failed to cast old object to tenant in update handler")
		return
	}

	newTenant, ok := newObj.(*seterav1.Tenant)
	if !ok {
		o.logger.Info("Failed to cast new object to tenant in update handler")
		return
	}

	if oldTenant.ResourceVersion == newTenant.ResourceVersion {
		o.logger.Info("No change in tenant", newTenant.Name)
		return
	}

	// Only add update event for tenants where the number of nodes has changed
	if len(oldTenant.Spec.Nodes) != len(newTenant.Spec.Nodes) {

		o.logger.Info("Number of nodes changed in tenant", newTenant.Name)
		o.logger.Info("Add tenant to queue - Update event", newTenant.Name)
		o.enqueue(newTenant, UpdateEvent)
		return

	}

}

// add tenant key to the workqueue - deletion event
func (o *Orchestrator) deleteTenantHandler(obj any) {

	tenant, ok := obj.(*seterav1.Tenant)
	if !ok {
		o.logger.Info("Failed to cast object to tenant in delete handler")
		return
	}

	o.logger.Info("Deleting tenant", tenant.Name)
	o.enqueue(tenant, DeleteEvent)

}
