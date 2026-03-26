/*
 * Copyright 2025 Daytona Platforms Inc.
 * SPDX-License-Identifier: AGPL-3.0
 *
 * Converters from proto-generated specsgen types to the Docker-client dto
 * types used by the v2 executor's underlying Docker operations.
 *
 * These functions are the boundary between the proto-defined job payload
 * schema (apps/runner/specs/runner.proto) and the existing Docker client
 * layer (apps/runner/pkg/docker).  No business logic lives here — only
 * field-by-field mapping between equivalent structs.
 */

package specs

import (
	"github.com/daytonaio/runner/pkg/api/dto"
	specsgen "github.com/daytonaio/runner/pkg/runner/v2/specs/gen"
)

// ─── Shared helpers ───────────────────────────────────────────────────────────

// RegistryToDTO converts a proto RegistryInfo to a dto.RegistryDTO pointer.
// Returns nil if r is nil.
func RegistryToDTO(r *specsgen.RegistryInfo) *dto.RegistryDTO {
	if r == nil {
		return nil
	}
	return &dto.RegistryDTO{
		Url:      r.GetUrl(),
		Project:  r.Project,
		Username: r.Username,
		Password: r.Password,
	}
}

// RegistryToDTOValue converts a proto RegistryInfo to a dto.RegistryDTO value
// (non-pointer).  Missing fields are left as zero values.
func RegistryToDTOValue(r *specsgen.RegistryInfo) dto.RegistryDTO {
	if r == nil {
		return dto.RegistryDTO{}
	}
	return dto.RegistryDTO{
		Url:      r.GetUrl(),
		Project:  r.Project,
		Username: r.Username,
		Password: r.Password,
	}
}

// VolumesToDTO converts a slice of proto VolumeMount pointers to dto.VolumeDTO
// values, skipping nil entries.
func VolumesToDTO(vs []*specsgen.VolumeMount) []dto.VolumeDTO {
	if len(vs) == 0 {
		return nil
	}
	out := make([]dto.VolumeDTO, 0, len(vs))
	for _, v := range vs {
		if v == nil {
			continue
		}
		out = append(out, dto.VolumeDTO{
			VolumeId:  v.GetVolumeId(),
			MountPath: v.GetMountPath(),
			Subpath:   v.Subpath,
		})
	}
	return out
}

// ─── Sandbox payload converters ───────────────────────────────────────────────

// CreateSandboxPayloadToDTO converts a parsed CreateSandboxPayload to the dto
// type expected by docker.Create.
func CreateSandboxPayloadToDTO(p *specsgen.CreateSandboxPayload) dto.CreateSandboxDTO {
	return dto.CreateSandboxDTO{
		Id:               p.GetId(),
		FromVolumeId:     p.GetFromVolumeId(),
		UserId:           p.GetUserId(),
		Snapshot:         p.GetSnapshot(),
		OsUser:           p.GetOsUser(),
		CpuQuota:         p.GetCpuQuota(),
		GpuQuota:         p.GetGpuQuota(),
		MemoryQuota:      p.GetMemoryQuota(),
		StorageQuota:     p.GetStorageQuota(),
		Env:              p.GetEnv(),
		Registry:         RegistryToDTO(p.GetRegistry()),
		Entrypoint:       p.GetEntrypoint(),
		Volumes:          VolumesToDTO(p.GetVolumes()),
		NetworkBlockAll:  p.NetworkBlockAll,
		NetworkAllowList: p.NetworkAllowList,
		Metadata:         p.GetMetadata(),
		AuthToken:        p.AuthToken,
		OtelEndpoint:     p.OtelEndpoint,
		SkipStart:        p.SkipStart,
		OrganizationId:   p.OrganizationId,
		RegionId:         p.RegionId,
	}
}

// ResizeSandboxPayloadToDTO converts a parsed ResizeSandboxPayload to the dto
// type expected by docker.Resize.
func ResizeSandboxPayloadToDTO(p *specsgen.ResizeSandboxPayload) dto.ResizeSandboxDTO {
	return dto.ResizeSandboxDTO{
		Cpu:    p.GetCpu(),
		Gpu:    p.GetGpu(),
		Memory: p.GetMemory(),
		Disk:   p.GetDisk(),
	}
}

// RecoverSandboxPayloadToDTO converts a parsed RecoverSandboxPayload to the
// dto type expected by docker.RecoverSandbox.
func RecoverSandboxPayloadToDTO(p *specsgen.RecoverSandboxPayload) dto.RecoverSandboxDTO {
	return dto.RecoverSandboxDTO{
		FromVolumeId:      p.GetFromVolumeId(),
		UserId:            p.GetUserId(),
		Snapshot:          p.Snapshot,
		OsUser:            p.GetOsUser(),
		CpuQuota:          p.GetCpuQuota(),
		GpuQuota:          p.GetGpuQuota(),
		MemoryQuota:       p.GetMemoryQuota(),
		StorageQuota:      p.GetStorageQuota(),
		Env:               p.GetEnv(),
		Volumes:           VolumesToDTO(p.GetVolumes()),
		NetworkBlockAll:   p.NetworkBlockAll,
		NetworkAllowList:  p.NetworkAllowList,
		ErrorReason:       p.GetErrorReason(),
		BackupErrorReason: p.GetBackupErrorReason(),
	}
}

// UpdateNetworkSettingsPayloadToDTO converts a parsed
// UpdateNetworkSettingsPayload to the dto type expected by
// docker.UpdateNetworkSettings.
func UpdateNetworkSettingsPayloadToDTO(p *specsgen.UpdateNetworkSettingsPayload) dto.UpdateNetworkSettingsDTO {
	return dto.UpdateNetworkSettingsDTO{
		NetworkBlockAll:    p.NetworkBlockAll,
		NetworkAllowList:   p.NetworkAllowList,
		NetworkLimitEgress: p.NetworkLimitEgress,
	}
}

// ─── Snapshot payload converters ─────────────────────────────────────────────

// CreateBackupPayloadToDTO converts a parsed CreateBackupPayload to the dto
// type expected by docker.CreateBackup.
func CreateBackupPayloadToDTO(p *specsgen.CreateBackupPayload) dto.CreateBackupDTO {
	return dto.CreateBackupDTO{
		Snapshot: p.GetSnapshot(),
		Registry: RegistryToDTOValue(p.GetRegistry()),
	}
}

// BuildSnapshotPayloadToDTO converts a parsed BuildSnapshotPayload to the dto
// type expected by docker.BuildSnapshot.
func BuildSnapshotPayloadToDTO(p *specsgen.BuildSnapshotPayload) dto.BuildSnapshotRequestDTO {
	result := dto.BuildSnapshotRequestDTO{
		Snapshot:               p.GetSnapshot(),
		Dockerfile:             p.GetDockerfile(),
		OrganizationId:         p.GetOrganizationId(),
		Context:                p.GetContext(),
		PushToInternalRegistry: p.GetPushToInternalRegistry(),
		Registry:               RegistryToDTO(p.GetRegistry()),
	}
	for _, r := range p.GetSourceRegistries() {
		if r != nil {
			result.SourceRegistries = append(result.SourceRegistries, RegistryToDTOValue(r))
		}
	}
	return result
}

// PullSnapshotPayloadToDTO converts a parsed PullSnapshotPayload to the dto
// type expected by docker.PullSnapshot.
func PullSnapshotPayloadToDTO(p *specsgen.PullSnapshotPayload) dto.PullSnapshotRequestDTO {
	return dto.PullSnapshotRequestDTO{
		Snapshot:            p.GetSnapshot(),
		Registry:            RegistryToDTO(p.GetRegistry()),
		DestinationRegistry: RegistryToDTO(p.GetDestinationRegistry()),
		DestinationRef:      p.DestinationRef,
		NewTag:              p.NewTag,
	}
}
