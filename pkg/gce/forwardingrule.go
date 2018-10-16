package gce

import (
	compute "google.golang.org/api/compute/v1"
)

type ForwardingRule interface {
	ForwardingRuleCreate(i *compute.ForwardingRule) error
	ForwardingRuleGet(i *compute.ForwardingRule) (*compute.ForwardingRule, error)
	ForwardingRuleDelete(i *compute.ForwardingRule) error
	ForwardingRuleUpdate(i *compute.ForwardingRule) error
}

// ForwardingRuleCreate an forwardingRule.
func (gce *GCEClient) ForwardingRuleCreate(payload compute.ForwardingRule) error {
	op, err := gce.service.ForwardingRules.Insert(gce.projectID, payload.Region, &payload).Do()
	if err != nil {
		return err
	}
	if err = gce.waitForRegionOp(op, payload.Region); err != nil {
		return err
	}
	return nil
}

// ForwardingRuleGet an ForwardingRule
func (gce *GCEClient) ForwardingRuleGet(payload compute.ForwardingRule) (*compute.ForwardingRule, error) {
	forwardingRule, err := gce.service.ForwardingRules.Get(gce.projectID, payload.Region, payload.Name).Do()
	if err != nil {
		if isHTTPErrorCode(err, 404) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return forwardingRule, nil
}

// ForwardingRuleDelete an forwardingRule
func (gce *GCEClient) ForwardingRuleDelete(payload compute.ForwardingRule) error {
	op, err := gce.service.ForwardingRules.Delete(gce.projectID, payload.Region, payload.Name).Do()
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

// ForwardingRuleUpdate an forwardingRule
// currently do not support updating an forwardingRule
func (gce *GCEClient) ForwardingRuleUpdate(payload compute.ForwardingRule) error {
	return nil
}
