/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package framework

import "time"

var defaultTimeouts = TimeoutContext{
	PodStart:                  5 * time.Minute,
	PodStartShort:             2 * time.Minute,
	PodStartSlow:              15 * time.Minute,
	PodDelete:                 5 * time.Minute,
	ClaimProvision:            5 * time.Minute,
	ClaimProvisionShort:       1 * time.Minute,
	DataSourceProvision:       5 * time.Minute,
	ClaimBound:                3 * time.Minute,
	PVReclaim:                 3 * time.Minute,
	PVBound:                   3 * time.Minute,
	PVCreate:                  3 * time.Minute,
	PVDelete:                  5 * time.Minute,
	PVDeleteSlow:              20 * time.Minute,
	SnapshotCreate:            5 * time.Minute,
	SnapshotDelete:            5 * time.Minute,
	SnapshotControllerMetrics: 5 * time.Minute,
}

// TimeoutContext contains timeout settings for several actions.
type TimeoutContext struct {
	// PodStart is how long to wait for the pod to be started.
	PodStart time.Duration

	// PodStartShort is same as `PodStart`, but shorter.
	// Use it in a case-by-case basis, mostly when you are sure pod start will not be delayed.
	PodStartShort time.Duration

	// PodStartSlow is same as `PodStart`, but longer.
	// Use it in a case-by-case basis, mostly when you are sure pod start will take longer than usual.
	PodStartSlow time.Duration

	// PodDelete is how long to wait for the pod to be deleted.
	PodDelete time.Duration

	// ClaimProvision is how long claims have to become dynamically provisioned.
	ClaimProvision time.Duration

	// DataSourceProvision is how long claims have to become dynamically provisioned from source claim.
	DataSourceProvision time.Duration

	// ClaimProvisionShort is the same as `ClaimProvision`, but shorter.
	ClaimProvisionShort time.Duration

	// ClaimBound is how long claims have to become bound.
	ClaimBound time.Duration

	// PVReclaim is how long PVs have to become reclaimed.
	PVReclaim time.Duration

	// PVBound is how long PVs have to become bound.
	PVBound time.Duration

	// PVCreate is how long PVs have to be created.
	PVCreate time.Duration

	// PVDelete is how long PVs have to become deleted.
	PVDelete time.Duration

	// PVDeleteSlow is the same as PVDelete, but slower.
	PVDeleteSlow time.Duration

	// SnapshotCreate is how long for snapshot to create snapshotContent.
	SnapshotCreate time.Duration

	// SnapshotDelete is how long for snapshot to delete snapshotContent.
	SnapshotDelete time.Duration

	// SnapshotControllerMetrics is how long to wait for snapshot controller metrics.
	SnapshotControllerMetrics time.Duration
}

// NewTimeoutContextWithDefaults returns a TimeoutContext with default values.
func NewTimeoutContextWithDefaults() *TimeoutContext {
	// Make a copy, otherwise the caller would have the ability to
	// modify the defaults
	copy := defaultTimeouts
	return &copy
}
