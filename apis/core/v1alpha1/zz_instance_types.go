/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AgentConfigObservation struct {
}

type AgentConfigParameters struct {

	// (Updatable) Whether Oracle Cloud Agent can run all the available plugins. This includes the management and monitoring plugins.
	// +kubebuilder:validation:Optional
	AreAllPluginsDisabled *bool `json:"areAllPluginsDisabled,omitempty" tf:"are_all_plugins_disabled,omitempty"`

	// (Updatable) Whether Oracle Cloud Agent can run all the available management plugins. Default value is false (management plugins are enabled).
	// +kubebuilder:validation:Optional
	IsManagementDisabled *bool `json:"isManagementDisabled,omitempty" tf:"is_management_disabled,omitempty"`

	// (Updatable) Whether Oracle Cloud Agent can gather performance metrics and monitor the instance using the monitoring plugins. Default value is false (monitoring plugins are enabled).
	// +kubebuilder:validation:Optional
	IsMonitoringDisabled *bool `json:"isMonitoringDisabled,omitempty" tf:"is_monitoring_disabled,omitempty"`

	// (Updatable) The configuration of plugins associated with this instance.
	// +kubebuilder:validation:Optional
	PluginsConfig []PluginsConfigParameters `json:"pluginsConfig,omitempty" tf:"plugins_config,omitempty"`
}

type AvailabilityConfigObservation struct {
}

type AvailabilityConfigParameters struct {

	// (Updatable) Whether to live migrate supported VM instances to a healthy physical VM host without disrupting running instances during infrastructure maintenance events. If null, Oracle chooses the best option for migrating the VM during infrastructure maintenance events.
	// +kubebuilder:validation:Optional
	IsLiveMigrationPreferred *bool `json:"isLiveMigrationPreferred,omitempty" tf:"is_live_migration_preferred,omitempty"`

	// (Updatable) The lifecycle state for an instance when it is recovered after infrastructure maintenance.
	// +kubebuilder:validation:Optional
	RecoveryAction *string `json:"recoveryAction,omitempty" tf:"recovery_action,omitempty"`
}

type CreateVnicDetailsIpv6AddressIpv6SubnetCidrPairDetailsObservation struct {
}

type CreateVnicDetailsIpv6AddressIpv6SubnetCidrPairDetailsParameters struct {

	// +kubebuilder:validation:Optional
	Ipv6Address *string `json:"ipv6address,omitempty" tf:"ipv6address,omitempty"`

	// +kubebuilder:validation:Optional
	Ipv6SubnetCidr *string `json:"ipv6subnetCidr,omitempty" tf:"ipv6subnet_cidr,omitempty"`
}

type CreateVnicDetailsObservation struct {
}

type CreateVnicDetailsParameters struct {

	// Whether to allocate an IPv6 address at instance and VNIC creation from an IPv6 enabled subnet. Default: False. When provided you may optionally provide an IPv6 prefix (ipv6SubnetCidr) of your choice to assign the IPv6 address from. If ipv6SubnetCidr is not provided then an IPv6 prefix is chosen for you.
	// +kubebuilder:validation:Optional
	AssignIpv6Ip *bool `json:"assignIpv6Ip,omitempty" tf:"assign_ipv6ip,omitempty"`

	// Whether the VNIC should be assigned a DNS record. If set to false, there will be no DNS record registration for the VNIC. If set to true, the DNS record will be registered. The default value is true.
	// If you specify a hostnameLabel, the assignPrivateDnsRecord is require to be set to true.
	// +kubebuilder:validation:Optional
	AssignPrivateDNSRecord *bool `json:"assignPrivateDnsRecord,omitempty" tf:"assign_private_dns_record,omitempty"`

	// (Updatable) Whether the VNIC should be assigned a public IP address. Defaults to whether the subnet is public or private. If not set and the VNIC is being created in a private subnet (that is, where prohibitPublicIpOnVnic = true in the Subnet), then no public IP address is assigned. If not set and the subnet is public (prohibitPublicIpOnVnic = false), then a public IP address is assigned. If set to true and prohibitPublicIpOnVnic = true, an error is returned.
	// +kubebuilder:validation:Optional
	AssignPublicIP *string `json:"assignPublicIp,omitempty" tf:"assign_public_ip,omitempty"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.  Example: {"Operations.CostCenter": "42"}
	// +kubebuilder:validation:Optional
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags.  Example: {"Department": "Finance"}
	// +kubebuilder:validation:Optional
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// (Updatable) The hostname for the VNIC's primary private IP. Used for DNS. The value is the hostname portion of the primary private IP's fully qualified domain name (FQDN) (for example, bminstance1 in FQDN bminstance1.subnet123.vcn1.oraclevcn.com). Must be unique across all VNICs in the subnet and comply with RFC 952 and RFC 1123. The value appears in the Vnic object and also the PrivateIp object returned by ListPrivateIps and GetPrivateIp.
	// +kubebuilder:validation:Optional
	HostnameLabel *string `json:"hostnameLabel,omitempty" tf:"hostname_label,omitempty"`

	// A list of IPv6 prefix ranges from which the VNIC should be assigned an IPv6 address. You can provide only the prefix ranges from which Oracle Cloud Infrastructure will select an available address from the range. You can optionally choose to leave the prefix range empty and instead provide the specific IPv6 address that should be used from within that range.
	// +kubebuilder:validation:Optional
	Ipv6AddressIpv6SubnetCidrPairDetails []CreateVnicDetailsIpv6AddressIpv6SubnetCidrPairDetailsParameters `json:"ipv6addressIpv6SubnetCidrPairDetails,omitempty" tf:"ipv6address_ipv6subnet_cidr_pair_details,omitempty"`

	// (Updatable) A list of the OCIDs of the network security groups (NSGs) to add the VNIC to. For more information about NSGs, see NetworkSecurityGroup.
	// +crossplane:generate:reference:type=NetworkSecurityGroup
	// +kubebuilder:validation:Optional
	NsgIds []*string `json:"nsgIds,omitempty" tf:"nsg_ids,omitempty"`

	// References to NetworkSecurityGroup to populate nsgIds.
	// +kubebuilder:validation:Optional
	NsgIdsRefs []v1.Reference `json:"nsgIdsRefs,omitempty" tf:"-"`

	// Selector for a list of NetworkSecurityGroup to populate nsgIds.
	// +kubebuilder:validation:Optional
	NsgIdsSelector *v1.Selector `json:"nsgIdsSelector,omitempty" tf:"-"`

	// A private IP address of your choice to assign to the VNIC. Must be an available IP address within the subnet's CIDR. If you don't specify a value, Oracle automatically assigns a private IP address from the subnet. This is the VNIC's primary private IP address. The value appears in the Vnic object and also the PrivateIp object returned by ListPrivateIps and GetPrivateIp.
	// +kubebuilder:validation:Optional
	PrivateIP *string `json:"privateIp,omitempty" tf:"private_ip,omitempty"`

	// (Updatable) Whether the source/destination check is disabled on the VNIC. Defaults to false, which means the check is performed. For information about why you would skip the source/destination check, see Using a Private IP as a Route Target.
	// +kubebuilder:validation:Optional
	SkipSourceDestCheck *bool `json:"skipSourceDestCheck,omitempty" tf:"skip_source_dest_check,omitempty"`

	// The OCID of the subnet to create the VNIC in. When launching an instance, use this subnetId instead of the deprecated subnetId in LaunchInstanceDetails. At least one of them is required; if you provide both, the values must match.
	// +crossplane:generate:reference:type=Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// Provide this attribute only if you are an Oracle Cloud VMware Solution customer and creating a secondary VNIC in a VLAN. The value is the OCID of the VLAN. See Vlan.
	// +crossplane:generate:reference:type=Vlan
	// +kubebuilder:validation:Optional
	VlanID *string `json:"vlanId,omitempty" tf:"vlan_id,omitempty"`

	// Reference to a Vlan to populate vlanId.
	// +kubebuilder:validation:Optional
	VlanIDRef *v1.Reference `json:"vlanIdRef,omitempty" tf:"-"`

	// Selector for a Vlan to populate vlanId.
	// +kubebuilder:validation:Optional
	VlanIDSelector *v1.Selector `json:"vlanIdSelector,omitempty" tf:"-"`
}

type InstanceLaunchOptionsObservation struct {
}

type InstanceLaunchOptionsParameters struct {

	// (Updatable) Emulation type for the boot volume.
	// +kubebuilder:validation:Optional
	BootVolumeType *string `json:"bootVolumeType,omitempty" tf:"boot_volume_type,omitempty"`

	// Firmware used to boot VM. Select the option that matches your operating system.
	// +kubebuilder:validation:Optional
	Firmware *string `json:"firmware,omitempty" tf:"firmware,omitempty"`

	// Whether to enable consistent volume naming feature. Defaults to false.
	// +kubebuilder:validation:Optional
	IsConsistentVolumeNamingEnabled *bool `json:"isConsistentVolumeNamingEnabled,omitempty" tf:"is_consistent_volume_naming_enabled,omitempty"`

	// Whether to enable in-transit encryption for the data volume's paravirtualized attachment. The default value is false. Use this field only during create. To update use is_pv_encryption_in_transit_enabled under launch_options instead.
	// +kubebuilder:validation:Optional
	IsPvEncryptionInTransitEnabled *bool `json:"isPvEncryptionInTransitEnabled,omitempty" tf:"is_pv_encryption_in_transit_enabled,omitempty"`

	// (Updatable) Emulation type for the physical network interface card (NIC).
	// +kubebuilder:validation:Optional
	NetworkType *string `json:"networkType,omitempty" tf:"network_type,omitempty"`

	// Emulation type for volume.
	// +kubebuilder:validation:Optional
	RemoteDataVolumeType *string `json:"remoteDataVolumeType,omitempty" tf:"remote_data_volume_type,omitempty"`
}

type InstanceObservation struct {

	// The OCID of the attached boot volume. If the source_type is bootVolume, this will be the same OCID as the source_id.
	BootVolumeID *string `json:"bootVolumeId,omitempty" tf:"boot_volume_id,omitempty"`

	// The OCID of the instance.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether the instance’s OCPUs and memory are distributed across multiple NUMA nodes.
	IsCrossNumaNode *bool `json:"isCrossNumaNode,omitempty" tf:"is_cross_numa_node,omitempty"`

	// Specifies the configuration mode for launching virtual machine (VM) instances. The configuration modes are:
	LaunchMode *string `json:"launchMode,omitempty" tf:"launch_mode,omitempty"`

	// A private IP address of your choice to assign to the VNIC. Must be an available IP address within the subnet's CIDR. If you don't specify a value, Oracle automatically assigns a private IP address from the subnet. This is the VNIC's primary private IP address. The value appears in the Vnic object and also the PrivateIp object returned by ListPrivateIps and GetPrivateIp.
	PrivateIP *string `json:"privateIp,omitempty" tf:"private_ip,omitempty"`

	// The public IP address of instance VNIC (if enabled).
	PublicIP *string `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	// The region that contains the availability domain the instance is running in.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// (Updatable) The shape configuration requested for the instance.
	// +kubebuilder:validation:Optional
	ShapeConfig []ShapeConfigObservation `json:"shapeConfig,omitempty" tf:"shape_config,omitempty"`

	// System tags for this resource. Each key is predefined and scoped to a namespace. Example: {"foo-namespace.bar-key": "value"}
	SystemTags map[string]*string `json:"systemTags,omitempty" tf:"system_tags,omitempty"`

	// The date and time the instance was created, in the format defined by RFC3339.  Example: 2016-08-25T21:10:29.600Z
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	// The date and time the instance is expected to be stopped / started,  in the format defined by RFC3339. After that time if instance hasn't been rebooted, Oracle will reboot the instance within 24 hours of the due time. Regardless of how the instance was stopped, the flag will be reset to empty as soon as instance reaches Stopped state. Example: 2018-05-25T21:10:29.600Z
	TimeMaintenanceRebootDue *string `json:"timeMaintenanceRebootDue,omitempty" tf:"time_maintenance_reboot_due,omitempty"`
}

type InstanceOptionsObservation struct {
}

type InstanceOptionsParameters struct {

	// (Updatable) Whether to disable the legacy (/v1) instance metadata service endpoints. Customers who have migrated to /v2 should set this to true for added security. Default is false.
	// +kubebuilder:validation:Optional
	AreLegacyImdsEndpointsDisabled *bool `json:"areLegacyImdsEndpointsDisabled,omitempty" tf:"are_legacy_imds_endpoints_disabled,omitempty"`
}

type InstanceParameters struct {

	// (Updatable) Configuration options for the Oracle Cloud Agent software running on the instance.
	// +kubebuilder:validation:Optional
	AgentConfig []AgentConfigParameters `json:"agentConfig,omitempty" tf:"agent_config,omitempty"`

	// The default value is false.
	// +kubebuilder:validation:Optional
	Async *bool `json:"async,omitempty" tf:"async,omitempty"`

	// (Updatable) Options for VM migration during infrastructure maintenance events and for defining the availability of a VM instance after a maintenance event that impacts the underlying hardware.
	// +kubebuilder:validation:Optional
	AvailabilityConfig []AvailabilityConfigParameters `json:"availabilityConfig,omitempty" tf:"availability_config,omitempty"`

	// The availability domain of the instance.  Example: Uocm:PHX-AD-1
	// +kubebuilder:validation:Required
	AvailabilityDomain *string `json:"availabilityDomain" tf:"availability_domain,omitempty"`

	// (Updatable) The OCID of the compute capacity reservation this instance is launched under. You can opt out of all default reservations by specifying an empty string as input for this field. For more information, see Capacity Reservations.
	// +kubebuilder:validation:Optional
	CapacityReservationID *string `json:"capacityReservationId,omitempty" tf:"capacity_reservation_id,omitempty"`

	// (Updatable) The OCID of the compartment.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment
	// +kubebuilder:validation:Optional
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// Reference to a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDRef *v1.Reference `json:"compartmentIdRef,omitempty" tf:"-"`

	// Selector for a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDSelector *v1.Selector `json:"compartmentIdSelector,omitempty" tf:"-"`

	// The OCID of the compute cluster that the instance will be created in.
	// +kubebuilder:validation:Optional
	ComputeClusterID *string `json:"computeClusterId,omitempty" tf:"compute_cluster_id,omitempty"`

	// (Updatable) Contains properties for a VNIC. You use this object when creating the primary VNIC during instance launch or when creating a secondary VNIC. For more information about VNICs, see Virtual Network Interface Cards (VNICs).
	// +kubebuilder:validation:Optional
	CreateVnicDetails []CreateVnicDetailsParameters `json:"createVnicDetails,omitempty" tf:"create_vnic_details,omitempty"`

	// (Updatable) The OCID of the dedicated virtual machine host to place the instance on.
	// +crossplane:generate:reference:type=DedicatedVMHost
	// +kubebuilder:validation:Optional
	DedicatedVMHostID *string `json:"dedicatedVmHostId,omitempty" tf:"dedicated_vm_host_id,omitempty"`

	// Reference to a DedicatedVMHost to populate dedicatedVmHostId.
	// +kubebuilder:validation:Optional
	DedicatedVMHostIDRef *v1.Reference `json:"dedicatedVmHostIdRef,omitempty" tf:"-"`

	// Selector for a DedicatedVMHost to populate dedicatedVmHostId.
	// +kubebuilder:validation:Optional
	DedicatedVMHostIDSelector *v1.Selector `json:"dedicatedVmHostIdSelector,omitempty" tf:"-"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.  Example: {"Operations.CostCenter": "42"}
	// +kubebuilder:validation:Optional
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (Updatable) Additional metadata key/value pairs that you provide. They serve the same purpose and functionality as fields in the metadata object.
	// +kubebuilder:validation:Optional
	ExtendedMetadata map[string]*string `json:"extendedMetadata,omitempty" tf:"extended_metadata,omitempty"`

	// (Updatable) A fault domain is a grouping of hardware and infrastructure within an availability domain. Each availability domain contains three fault domains. Fault domains let you distribute your instances so that they are not on the same physical hardware within a single availability domain. A hardware failure or Compute hardware maintenance that affects one fault domain does not affect instances in other fault domains.
	// +kubebuilder:validation:Optional
	FaultDomain *string `json:"faultDomain,omitempty" tf:"fault_domain,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags.  Example: {"Department": "Finance"}
	// +kubebuilder:validation:Optional
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// (Updatable) The hostname for the VNIC's primary private IP. Used for DNS. The value is the hostname portion of the primary private IP's fully qualified domain name (FQDN) (for example, bminstance1 in FQDN bminstance1.subnet123.vcn1.oraclevcn.com). Must be unique across all VNICs in the subnet and comply with RFC 952 and RFC 1123. The value appears in the Vnic object and also the PrivateIp object returned by ListPrivateIps and GetPrivateIp.
	// +kubebuilder:validation:Optional
	HostnameLabel *string `json:"hostnameLabel,omitempty" tf:"hostname_label,omitempty"`

	// Deprecated. Use sourceDetails with InstanceSourceViaImageDetails source type instead. If you specify values for both, the values must match.
	// +kubebuilder:validation:Optional
	Image *string `json:"image,omitempty" tf:"image,omitempty"`

	// The OCID of the Instance Configuration containing instance launch details. Any other fields supplied in this instance launch request will override the details stored in the Instance Configuration for this instance launch.
	// +kubebuilder:validation:Optional
	InstanceConfigurationID *string `json:"instanceConfigurationId,omitempty" tf:"instance_configuration_id,omitempty"`

	// (Updatable) Optional mutable instance options
	// +kubebuilder:validation:Optional
	InstanceOptions []InstanceOptionsParameters `json:"instanceOptions,omitempty" tf:"instance_options,omitempty"`

	// This is an advanced option.
	// +kubebuilder:validation:Optional
	IpxeScript *string `json:"ipxeScript,omitempty" tf:"ipxe_script,omitempty"`

	// Whether to enable in-transit encryption for the data volume's paravirtualized attachment. The default value is false. Use this field only during create. To update use is_pv_encryption_in_transit_enabled under launch_options instead.
	// +kubebuilder:validation:Optional
	IsPvEncryptionInTransitEnabled *bool `json:"isPvEncryptionInTransitEnabled,omitempty" tf:"is_pv_encryption_in_transit_enabled,omitempty"`

	// (Updatable) Options for tuning the compatibility and performance of VM shapes. The values that you specify override any default values.
	// +kubebuilder:validation:Optional
	LaunchOptions []InstanceLaunchOptionsParameters `json:"launchOptions,omitempty" tf:"launch_options,omitempty"`

	// fields in that these can be nested JSON objects (whereas metadata fields are string/string maps only).
	// +kubebuilder:validation:Optional
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// (Updatable only for VM's) The platform configuration requested for the instance.
	// +kubebuilder:validation:Optional
	PlatformConfig []PlatformConfigParameters `json:"platformConfig,omitempty" tf:"platform_config,omitempty"`

	// Configuration options for preemptible instances.
	// +kubebuilder:validation:Optional
	PreemptibleInstanceConfig []PreemptibleInstanceConfigParameters `json:"preemptibleInstanceConfig,omitempty" tf:"preemptible_instance_config,omitempty"`

	// Whether to preserve the boot volume that was used to launch the preemptible instance when the instance is terminated. Defaults to false if not specified.
	// +kubebuilder:validation:Optional
	PreserveBootVolume *bool `json:"preserveBootVolume,omitempty" tf:"preserve_boot_volume,omitempty"`

	// that you specify. If you don't provide the parameter, the default values for the shape are used.
	// +kubebuilder:validation:Optional
	Shape *string `json:"shape,omitempty" tf:"shape,omitempty"`

	// (Updatable) The shape configuration requested for the instance.
	// +kubebuilder:validation:Optional
	ShapeConfig []ShapeConfigParameters `json:"shapeConfig,omitempty" tf:"shape_config,omitempty"`

	// (Updatable)
	// +kubebuilder:validation:Optional
	SourceDetails []InstanceSourceDetailsParameters `json:"sourceDetails,omitempty" tf:"source_details,omitempty"`

	// (Updatable) The target state for the instance. Could be set to RUNNING or STOPPED.
	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The OCID of the subnet to create the VNIC in. When launching an instance, use this subnetId instead of the deprecated subnetId in LaunchInstanceDetails. At least one of them is required; if you provide both, the values must match.
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	UpdateOperationConstraint *string `json:"updateOperationConstraint,omitempty" tf:"update_operation_constraint,omitempty"`
}

type InstanceSourceDetailsObservation struct {
}

type InstanceSourceDetailsParameters struct {

	// (Applicable when source_type=image) (Updatable) The size of the boot volume in GBs. Minimum value is 50 GB and maximum value is 32,768 GB (32 TB).
	// +kubebuilder:validation:Optional
	BootVolumeSizeInGbs *string `json:"bootVolumeSizeInGbs,omitempty" tf:"boot_volume_size_in_gbs,omitempty"`

	// (Applicable when source_type=image) The number of volume performance units (VPUs) that will be applied to this volume per GB, representing the Block Volume service's elastic performance options. See Block Volume Performance Levels for more information.
	// +kubebuilder:validation:Optional
	BootVolumeVpusPerGb *string `json:"bootVolumeVpusPerGb,omitempty" tf:"boot_volume_vpus_per_gb,omitempty"`

	// (Applicable when source_type=image) These are the criteria for selecting an image. This is required if imageId is not specified.
	// +kubebuilder:validation:Optional
	InstanceSourceImageFilterDetails []InstanceSourceImageFilterDetailsParameters `json:"instanceSourceImageFilterDetails,omitempty" tf:"instance_source_image_filter_details,omitempty"`

	// (Applicable when source_type=image) The OCID of the Vault service key to assign as the master encryption key for the boot volume.
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// The OCID of an image or a boot volume to use, depending on the value of source_type.
	// +crossplane:generate:reference:type=Image
	// +kubebuilder:validation:Optional
	SourceID *string `json:"sourceId,omitempty" tf:"source_id,omitempty"`

	// Reference to a Image to populate sourceId.
	// +kubebuilder:validation:Optional
	SourceIDRef *v1.Reference `json:"sourceIdRef,omitempty" tf:"-"`

	// Selector for a Image to populate sourceId.
	// +kubebuilder:validation:Optional
	SourceIDSelector *v1.Selector `json:"sourceIdSelector,omitempty" tf:"-"`

	// The source type for the instance. Use image when specifying the image OCID. Use bootVolume when specifying the boot volume OCID.
	// +kubebuilder:validation:Required
	SourceType *string `json:"sourceType" tf:"source_type,omitempty"`
}

type InstanceSourceImageFilterDetailsObservation struct {
}

type InstanceSourceImageFilterDetailsParameters struct {

	// (Updatable) The OCID of the compartment.
	// +kubebuilder:validation:Required
	CompartmentID *string `json:"compartmentId" tf:"compartment_id,omitempty"`

	// (Applicable when source_type=image) Filter based on these defined tags. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.
	// +kubebuilder:validation:Optional
	DefinedTagsFilter map[string]*string `json:"definedTagsFilter,omitempty" tf:"defined_tags_filter,omitempty"`

	// (Applicable when source_type=image) The image's operating system.  Example: Oracle Linux
	// +kubebuilder:validation:Optional
	OperatingSystem *string `json:"operatingSystem,omitempty" tf:"operating_system,omitempty"`

	// (Applicable when source_type=image) The image's operating system version.  Example: 7.2
	// +kubebuilder:validation:Optional
	OperatingSystemVersion *string `json:"operatingSystemVersion,omitempty" tf:"operating_system_version,omitempty"`
}

type PlatformConfigObservation struct {
}

type PlatformConfigParameters struct {

	// (Applicable when type=AMD_MILAN_BM | AMD_MILAN_BM_GPU | AMD_ROME_BM | AMD_ROME_BM_GPU | GENERIC_BM) Whether virtualization instructions are available. For example, Secure Virtual Machine for AMD shapes or VT-x for Intel shapes.
	// +kubebuilder:validation:Optional
	AreVirtualInstructionsEnabled *bool `json:"areVirtualInstructionsEnabled,omitempty" tf:"are_virtual_instructions_enabled,omitempty"`

	// (Applicable when type=AMD_MILAN_BM | AMD_MILAN_BM_GPU | AMD_ROME_BM | AMD_ROME_BM_GPU | GENERIC_BM | INTEL_ICELAKE_BM | INTEL_SKYLAKE_BM) Instance Platform Configuration Configuration Map for flexible setting input.
	// +kubebuilder:validation:Optional
	ConfigMap map[string]*string `json:"configMap,omitempty" tf:"config_map,omitempty"`

	// (Applicable when type=AMD_MILAN_BM | AMD_MILAN_BM_GPU | AMD_ROME_BM | AMD_ROME_BM_GPU | GENERIC_BM) Whether the Access Control Service is enabled on the instance. When enabled, the platform can enforce PCIe device isolation, required for VFIO device pass-through.
	// +kubebuilder:validation:Optional
	IsAccessControlServiceEnabled *bool `json:"isAccessControlServiceEnabled,omitempty" tf:"is_access_control_service_enabled,omitempty"`

	// (Applicable when type=AMD_MILAN_BM | AMD_MILAN_BM_GPU | AMD_ROME_BM | AMD_ROME_BM_GPU | GENERIC_BM | INTEL_ICELAKE_BM | INTEL_SKYLAKE_BM) Whether the input-output memory management unit is enabled.
	// +kubebuilder:validation:Optional
	IsInputOutputMemoryManagementUnitEnabled *bool `json:"isInputOutputMemoryManagementUnitEnabled,omitempty" tf:"is_input_output_memory_management_unit_enabled,omitempty"`

	// Whether the Measured Boot feature is enabled on the instance.
	// +kubebuilder:validation:Optional
	IsMeasuredBootEnabled *bool `json:"isMeasuredBootEnabled,omitempty" tf:"is_measured_boot_enabled,omitempty"`

	// Whether the instance is a confidential instance. If this value is true, the instance is a confidential instance. The default value is false.
	// +kubebuilder:validation:Optional
	IsMemoryEncryptionEnabled *bool `json:"isMemoryEncryptionEnabled,omitempty" tf:"is_memory_encryption_enabled,omitempty"`

	// Whether Secure Boot is enabled on the instance.
	// +kubebuilder:validation:Optional
	IsSecureBootEnabled *bool `json:"isSecureBootEnabled,omitempty" tf:"is_secure_boot_enabled,omitempty"`

	// (Applicable when type=AMD_MILAN_BM | AMD_MILAN_BM_GPU | AMD_ROME_BM | AMD_ROME_BM_GPU | AMD_VM | GENERIC_BM | INTEL_ICELAKE_BM | INTEL_SKYLAKE_BM | INTEL_VM) (Updatable only for INTEL_VM and AMD_VM) Whether symmetric multithreading is enabled on the instance. Symmetric multithreading is also called simultaneous multithreading (SMT) or Intel Hyper-Threading.
	// +kubebuilder:validation:Optional
	IsSymmetricMultiThreadingEnabled *bool `json:"isSymmetricMultiThreadingEnabled,omitempty" tf:"is_symmetric_multi_threading_enabled,omitempty"`

	// Whether the Trusted Platform Module (TPM) is enabled on the instance.
	// +kubebuilder:validation:Optional
	IsTrustedPlatformModuleEnabled *bool `json:"isTrustedPlatformModuleEnabled,omitempty" tf:"is_trusted_platform_module_enabled,omitempty"`

	// (Applicable when type=AMD_MILAN_BM | AMD_MILAN_BM_GPU | AMD_ROME_BM | AMD_ROME_BM_GPU | GENERIC_BM | INTEL_ICELAKE_BM | INTEL_SKYLAKE_BM) The number of NUMA nodes per socket (NPS).
	// +kubebuilder:validation:Optional
	NumaNodesPerSocket *string `json:"numaNodesPerSocket,omitempty" tf:"numa_nodes_per_socket,omitempty"`

	// (Applicable when type=AMD_MILAN_BM | AMD_ROME_BM | GENERIC_BM | INTEL_ICELAKE_BM | INTEL_SKYLAKE_BM) The percentage of cores enabled. Value must be a multiple of 25%. If the requested percentage results in a fractional number of cores, the system rounds up the number of cores across processors and provisions an instance with a whole number of cores.
	// +kubebuilder:validation:Optional
	PercentageOfCoresEnabled *float64 `json:"percentageOfCoresEnabled,omitempty" tf:"percentage_of_cores_enabled,omitempty"`

	// The type of platform being configured.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type PluginsConfigObservation struct {
}

type PluginsConfigParameters struct {

	// (Updatable) Whether the plugin should be enabled or disabled.
	// +kubebuilder:validation:Required
	DesiredState *string `json:"desiredState" tf:"desired_state,omitempty"`

	// (Updatable) The plugin name. To get a list of available plugins, use the ListInstanceagentAvailablePlugins operation in the Oracle Cloud Agent API. For more information about the available plugins, see Managing Plugins with Oracle Cloud Agent.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type PreemptibleInstanceConfigObservation struct {
}

type PreemptibleInstanceConfigParameters struct {

	// The action to run when the preemptible instance is interrupted for eviction.
	// +kubebuilder:validation:Required
	PreemptionAction []PreemptionActionParameters `json:"preemptionAction" tf:"preemption_action,omitempty"`
}

type PreemptionActionObservation struct {
}

type PreemptionActionParameters struct {

	// Whether to preserve the boot volume that was used to launch the preemptible instance when the instance is terminated. Defaults to false if not specified.
	// +kubebuilder:validation:Optional
	PreserveBootVolume *bool `json:"preserveBootVolume,omitempty" tf:"preserve_boot_volume,omitempty"`

	// The type of platform being configured.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type ShapeConfigObservation struct {

	// A short description of the instance's graphics processing unit (GPU).
	GpuDescription *string `json:"gpuDescription,omitempty" tf:"gpu_description,omitempty"`

	// The number of GPUs available to the instance.
	Gpus *float64 `json:"gpus,omitempty" tf:"gpus,omitempty"`

	// A short description of the local disks available to this instance.
	LocalDiskDescription *string `json:"localDiskDescription,omitempty" tf:"local_disk_description,omitempty"`

	// The number of local disks available to the instance.
	LocalDisks *float64 `json:"localDisks,omitempty" tf:"local_disks,omitempty"`

	// The aggregate size of all local disks, in gigabytes.
	LocalDisksTotalSizeInGbs *float64 `json:"localDisksTotalSizeInGbs,omitempty" tf:"local_disks_total_size_in_gbs,omitempty"`

	// The maximum number of VNIC attachments for the instance.
	MaxVnicAttachments *float64 `json:"maxVnicAttachments,omitempty" tf:"max_vnic_attachments,omitempty"`

	// The networking bandwidth available to the instance, in gigabits per second.
	NetworkingBandwidthInGbps *float64 `json:"networkingBandwidthInGbps,omitempty" tf:"networking_bandwidth_in_gbps,omitempty"`

	// A short description of the instance's processor (CPU).
	ProcessorDescription *string `json:"processorDescription,omitempty" tf:"processor_description,omitempty"`
}

type ShapeConfigParameters struct {

	// (Updatable) The baseline OCPU utilization for a subcore burstable VM instance. Leave this attribute blank for a non-burstable instance, or explicitly specify non-burstable with BASELINE_1_1.
	// +kubebuilder:validation:Optional
	BaselineOcpuUtilization *string `json:"baselineOcpuUtilization,omitempty" tf:"baseline_ocpu_utilization,omitempty"`

	// (Updatable) The total amount of memory available to the instance, in gigabytes.
	// +kubebuilder:validation:Optional
	MemoryInGbs *float64 `json:"memoryInGbs,omitempty" tf:"memory_in_gbs,omitempty"`

	// (Updatable) The number of NVMe drives to be used for storage. A single drive has 6.8 TB available.
	// +kubebuilder:validation:Optional
	Nvmes *float64 `json:"nvmes,omitempty" tf:"nvmes,omitempty"`

	// (Updatable) The total number of OCPUs available to the instance.
	// +kubebuilder:validation:Optional
	Ocpus *float64 `json:"ocpus,omitempty" tf:"ocpus,omitempty"`

	// (Updatable) The total number of VCPUs available to the instance. This can be used instead of OCPUs, in which case the actual number of OCPUs will be calculated based on this value and the actual hardware. This must be a multiple of 2.
	// +kubebuilder:validation:Optional
	Vcpus *float64 `json:"vcpus,omitempty" tf:"vcpus,omitempty"`
}

// InstanceSpec defines the desired state of Instance
type InstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceParameters `json:"forProvider"`
}

// InstanceStatus defines the observed state of Instance.
type InstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Instance is the Schema for the Instances API. Provides the Instance resource in Oracle Cloud Infrastructure Core service
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec"`
	Status            InstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceList contains a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instance `json:"items"`
}

// Repository type metadata.
var (
	Instance_Kind             = "Instance"
	Instance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Instance_Kind}.String()
	Instance_KindAPIVersion   = Instance_Kind + "." + CRDGroupVersion.String()
	Instance_GroupVersionKind = CRDGroupVersion.WithKind(Instance_Kind)
)

func init() {
	SchemeBuilder.Register(&Instance{}, &InstanceList{})
}
