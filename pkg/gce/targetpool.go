package gce

import (
	compute "google.golang.org/api/compute/v1"
)

type TargetPool interface {
	TargetPoolCreate(i *compute.TargetPool) error
	TargetPoolGet(i *compute.TargetPool) (*compute.TargetPool, error)
	TargetPoolDelete(i *compute.TargetPool) error
	TargetPoolUpdate(i *compute.TargetPool) error
}

// TargetPoolCreate an targetPool.
func (gce *GCEClient) TargetPoolCreate(payload compute.TargetPool) error {
	op, err := gce.compute.TargetPools.Insert(gce.projectID, payload.Region, &payload).Do()
	if err != nil {
		return err
	}
	if err = gce.waitForRegionOp(op, payload.Region); err != nil {
		return err
	}
	return nil
}

// TargetPoolGet an TargetPool
func (gce *GCEClient) TargetPoolGet(payload compute.TargetPool) (*compute.TargetPool, error) {
	targetPool, err := gce.compute.TargetPools.Get(gce.projectID, payload.Region, payload.Name).Do()
	if err != nil {
		if isHTTPErrorCode(err, 404) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return targetPool, nil
}

// TargetPoolDelete an targetPool
func (gce *GCEClient) TargetPoolDelete(payload compute.TargetPool) error {
	op, err := gce.compute.TargetPools.Delete(gce.projectID, payload.Region, payload.Name).Do()
	if err != nil {
		if isHTTPErrorCode(err, 404) {
			return nil
		}
		return err
	}
	if err = gce.waitForRegionOp(op, payload.Region); err != nil {
		return err
	}
	return nil
}

// TargetPoolUpdate an targetPool
// currently do not support updating an targetPool
func (gce *GCEClient) TargetPoolUpdate(payload compute.TargetPool) error {
	return nil
}
