package models

// Volume represents a volume to be attached to a microvm machine.
type Volume struct {
	// ID is the uinique identifier of the volume.
	ID string `json:"id"`
	// IsRoot specifies that the volume is to be used as the root volume. A machine
	// must have a root volume.
	IsRoot bool `json:"is_root"`
	// IsReadOnly specifies that the volume is to be mounted readonly.
	IsReadOnly bool `json:"is_read_only,omitempty"`
	// MountPoint is the mount point for the volume in the microvm.
	MountPoint string `json:"mount_point" validate:"required,startswith=/"`
	// Source is where the volume will be sourced from.
	Source VolumeSource `json:"source"`
	// PartitionID is the uuid of the boot partition.
	PartitionID string `json:"partition_id,omitempty"`
	// Size is the size to resize this volume to.
	Size int32 `json:"size,omitempty"`
	// TODO: add rate limiting.
}

// Volumes represents a collection of volumes.
type Volumes []Volume

// GetByID will get a volume status by id.
func (v Volumes) GetByID(id string) *Volume {
	for _, vol := range v {
		if vol.ID == id {
			return &vol
		}
	}

	return nil
}

// VolumeSource is the source of a volume. Based loosely on the volumes in Kubernetes Pod specs.
type VolumeSource struct {
	// Container is used to specify a source of a volume as a OCI container.
	Container *ContainerVolumeSource `json:"container,omitempty"`
	// TODO: add CSI.
}

// ContainerDriveSource represents the details of a volume coming from a OCI image.
type ContainerVolumeSource struct {
	// Image is the OCI image to use.
	Image ContainerImage `json:"image"`
}

// Mount represents a volume mount point.
type Mount struct {
	// Type specifies the type of the mount (e.g. device or directory).
	Type MountType
	// Source is the location of the mounted volume.
	Source string
}

// MountType is a type representing the type of mount.
type MountType string

const (
	// MountTypeDev represents a mount point that is a block device.
	MountTypeDev MountType = "dev"
	// MountTypeHostPath represents a mount point that is a directory on the host.
	MountTypeHostPath MountType = "hostpath"
)

// ImageUse is a type representing the how an image will be used.
type ImageUse string

const (
	// ImageUseVolume represents the usage of af an image for a volume.
	ImageUseVolume ImageUse = "volume"
	// ImageUseKernel represents the usage of af an image for a kernel.
	ImageUseKernel ImageUse = "kernel"
	// ImageUseKernel represents the usage of af an image for a initial ramdisk.
	ImageUseInitrd ImageUse = "initrd"
)

// VolumeStatus holds status information about the volumes.
type VolumeStatus struct {
	// Mount is the mount point information for the volume.
	Mount Mount `json:"mount"`
}

// VolumeStatuses represents a collection of volume status.
type VolumeStatuses map[string]*VolumeStatus