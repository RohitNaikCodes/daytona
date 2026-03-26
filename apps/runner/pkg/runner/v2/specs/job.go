/*
 * Copyright 2025 Daytona Platforms Inc.
 * SPDX-License-Identifier: AGPL-3.0
 */

package specs

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	apiclient "github.com/daytonaio/daytona/libs/api-client-go"
	specsgen "github.com/daytonaio/runner/pkg/runner/v2/specs/gen"
)

// JobTypeFromAPIClient maps an apiclient.JobType string to a specsgen.JobType enum.
// The proto field name (e.g. CREATE_SANDBOX) matches the apiclient string constant.
func JobTypeFromAPIClient(t apiclient.JobType) specsgen.JobType {
	enumDesc := specsgen.JobType(0).Descriptor()
	val := enumDesc.Values().ByName(protoreflect.Name(string(t)))
	if val == nil {
		return specsgen.JobType_JOB_TYPE_UNSPECIFIED
	}
	return specsgen.JobType(val.Number())
}

// JobStatusFromAPIClient maps an apiclient.JobStatus string to specsgen.JobStatus.
func JobStatusFromAPIClient(s apiclient.JobStatus) specsgen.JobStatus {
	enumDesc := specsgen.JobStatus(0).Descriptor()
	val := enumDesc.Values().ByName(protoreflect.Name(string(s)))
	if val == nil {
		return specsgen.JobStatus_JOB_STATUS_UNSPECIFIED
	}
	return specsgen.JobStatus(val.Number())
}

// ResourceTypeFromAPIClient maps an apiclient ResourceType string to specsgen.ResourceType.
func ResourceTypeFromAPIClient(s string) specsgen.ResourceType {
	enumDesc := specsgen.ResourceType(0).Descriptor()
	val := enumDesc.Values().ByName(protoreflect.Name(s))
	if val == nil {
		return specsgen.ResourceType_RESOURCE_TYPE_UNSPECIFIED
	}
	return specsgen.ResourceType(val.Number())
}

// JobStatusToAPIClient converts a specsgen.JobStatus back to apiclient.JobStatus string.
func JobStatusToAPIClient(s specsgen.JobStatus) apiclient.JobStatus {
	return apiclient.JobStatus(s.Descriptor().Values().ByNumber(protoreflect.EnumNumber(s)).Name())
}

// JobFromAPIClient converts a polled *apiclient.Job into a *specsgen.Job.
// The trace context values that are strings are preserved; non-string values are dropped.
func JobFromAPIClient(j *apiclient.Job) *specsgen.Job {
	tc := make(map[string]string, len(j.TraceContext))
	for k, v := range j.TraceContext {
		if s, ok := v.(string); ok {
			tc[k] = s
		}
	}
	job := &specsgen.Job{
		Id:           j.Id,
		Type:         JobTypeFromAPIClient(j.Type),
		Status:       JobStatusFromAPIClient(j.Status),
		ResourceType: ResourceTypeFromAPIClient(j.ResourceType),
		ResourceId:   j.ResourceId,
		TraceContext: tc,
		CreatedAt:    j.CreatedAt,
	}
	if j.Payload != nil {
		job.Payload = j.Payload
	}
	if j.ErrorMessage != nil {
		job.ErrorMessage = j.ErrorMessage
	}
	if j.UpdatedAt != nil {
		job.UpdatedAt = j.UpdatedAt
	}
	return job
}
