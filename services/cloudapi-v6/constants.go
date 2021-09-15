package cloudapi_v6

// Resources
const (
	DatacenterResource = "datacenter"
	ServerResource     = "server"
	VolumeResource     = "volume"
	IpBlockResource    = "ipblock"
	SnapshotResource   = "snapshot"
)

// Flags
const (
	ArgCols                  = "cols"
	ArgUserData              = "user-data"
	ArgFirstName             = "first-name"
	ArgLastName              = "last-name"
	ArgEmail                 = "email"
	ArgEmailShort            = "e"
	ArgPassword              = "password"
	ArgPasswordShort         = "p"
	ArgAdmin                 = "admin"
	ArgName                  = "name"
	ArgVolumeName            = "volume-name"
	ArgNameShort             = "n"
	ArgApiSubnets            = "api-subnets"
	ArgDescription           = "description"
	ArgDescriptionShort      = "d"
	ArgLocation              = "location"
	ArgLocationShort         = "l"
	ArgAvailabilityZone      = "availability-zone"
	ArgAvailabilityZoneShort = "z"
	ArgDirection             = "direction"
	ArgDirectionShort        = "d"
	ArgAction                = "action"
	ArgActionShort           = "a"
	ArgS3Bucket              = "s3bucket"
	ArgS3BucketShort         = "b"
	ArgCores                 = "cores"
	ArgRam                   = "ram"
	ArgCPUFamily             = "cpu-family"
	ArgSize                  = "size"
	ArgBus                   = "bus"
	ArgLicenceType           = "licence-type"
	ArgType                  = "type"
	ArgSshKeys               = "ssh-keys"
	ArgPublic                = "public"
	ArgIps                   = "ips"
	ArgIp                    = "ip"
	ArgDhcp                  = "dhcp"
	ArgNetwork               = "network"
	ArgListenerLan           = "listener-lan"
	ArgListenerIp            = "listener-ip"
	ArgListenerPort          = "listener-port"
	ArgAlgorithm             = "algorithm"
	ArgTargetLan             = "target-lan"
	ArgRetries               = "retries"
	ArgClientTimeout         = "client-timeout"
	ArgConnectionTimeout     = "connection-timeout"
	ArgTargetTimeout         = "target-timeout"
	ArgCheck                 = "check"
	ArgCheckInterval         = "check-interval"
	ArgMaintenance           = "maintenance"
	ArgFirewallActive        = "firewall-active"
	ArgFirewallType          = "firewall-type"
	ArgCpuHotPlug            = "cpu-hot-plug"
	ArgCpuHotUnplug          = "cpu-hot-unplug"
	ArgRamHotPlug            = "ram-hot-plug"
	ArgRamHotUnplug          = "ram-hot-unplug"
	ArgNicHotPlug            = "nic-hot-plug"
	ArgNicHotUnplug          = "nic-hot-unplug"
	ArgDiscVirtioHotPlug     = "disc-virtio-hot-plug"
	ArgDiscVirtioHotUnplug   = "disc-virtio-hot-unplug"
	ArgDiscScsiHotPlug       = "disc-scsi-hot-plug"
	ArgDiscScsiHotUnplug     = "disc-scsi-hot-unplug"
	ArgSecAuthProtection     = "sec-auth-protection"
	ArgImageAlias            = "image-alias"
	ArgProtocol              = "protocol"
	ArgProtocolShort         = "p"
	ArgSourceSubnet          = "source-subnet"
	ArgTargetSubnet          = "target-subnet"
	ArgSourceMac             = "source-mac"
	ArgSourceIp              = "source-ip"
	ArgTargetIp              = "target-ip"
	ArgTargetPort            = "target-port"
	ArgWeight                = "weight"
	ArgIcmpCode              = "icmp-code"
	ArgIcmpType              = "icmp-type"
	ArgPortRangeStart        = "port-range-start"
	ArgPortRangeEnd          = "port-range-end"
	ArgLabelUrn              = "label-urn"
	ArgLabelKey              = "label-key"
	ArgLabelValue            = "label-value"
	ArgResourceLimits        = "resource-limits"
	ArgResourceType          = "resource-type"
	ArgForceSecAuth          = "force-secure-auth"
	ArgCreateDc              = "create-dc"
	ArgCreateSnapshot        = "create-snapshot"
	ArgReserveIp             = "reserve-ip"
	ArgAccessLog             = "access-logs"
	ArgS3Privilege           = "s3privilege"
	ArgCreateBackUpUnit      = "create-backup"
	ArgCreatePcc             = "create-pcc"
	ArgCreateNic             = "create-nic"
	ArgCreateK8s             = "create-k8s"
	ArgCreateFlowLog         = "create-flowlog"
	ArgAccessMonitoring      = "access-monitoring"
	ArgAccessCerts           = "access-certs"
	ArgEditPrivilege         = "edit-privilege"
	ArgSharePrivilege        = "share-privilege"
	ArgS3KeyActive           = "s3key-active"
	ArgK8sVersion            = "k8s-version"
	ArgK8sNodeCount          = "node-count"
	ArgCpuFamily             = "cpu-family"
	ArgStorageType           = "storage-type"
	ArgStorageSize           = "storage-size"
	ArgK8sMinNodeCount       = "min-node-count"
	ArgK8sMaxNodeCount       = "max-node-count"
	ArgK8sMaintenanceDay     = "maintenance-day"
	ArgK8sMaintenanceTime    = "maintenance-time"
	ArgK8sAnnotationKey      = "annotation-key"
	ArgK8sAnnotationValue    = "annotation-value"
	ArgPublicIps             = "public-ips"
	ArgPrivateIps            = "private-ips"
	ArgGatewayIp             = "gateway-ip"
	ArgLatest                = "latest"
	ArgMethod                = "method"
)

// IDs Flags
const (
	ArgIdShort               = "i"
	ArgDataCenterId          = "datacenter-id"
	ArgServerId              = "server-id"
	ArgNatGatewayId          = "natgateway-id"
	ArgNetworkLoadBalancerId = "networkloadbalancer-id"
	ArgNicId                 = "nic-id"
	ArgLanId                 = "lan-id"
	ArgLanIds                = "lan-ids"
	ArgLocationId            = "location-id"
	ArgVolumeId              = "volume-id"
	ArgLoadBalancerId        = "loadbalancer-id"
	ArgRequestId             = "request-id"
	ArgSnapshotId            = "snapshot-id"
	ArgImageId               = "image-id"
	ArgIpBlockId             = "ipblock-id"
	ArgFirewallRuleId        = "firewallrule-id"
	ArgFlowLogId             = "flowlog-id"
	ArgUserId                = "user-id"
	ArgGroupId               = "group-id"
	ArgResourceId            = "resource-id"
	ArgRuleId                = "rule-id"
	ArgS3KeyId               = "s3key-id"
	ArgBackupUnitId          = "backupunit-id"
	ArgPccId                 = "pcc-id"
	ArgK8sClusterId          = "cluster-id"
	ArgK8sNodePoolId         = "nodepool-id"
	ArgK8sNodeId             = "node-id"
	ArgCdromId               = "cdrom-id"
	ArgTemplateId            = "template-id"
)

// Descriptions for Flags Resources
const (
	DatacenterId          = "The unique Data Center Id"
	LanId                 = "The unique LAN Id"
	LoadBalancerId        = "The unique Load Balancer Id"
	NicId                 = "The unique NIC Id"
	RequestId             = "The unique Request Id"
	ServerId              = "The unique Server Id"
	VolumeId              = "The unique Volume Id"
	SnapshotId            = "The unique Snapshot Id"
	ImageId               = "The unique Image Id"
	IpBlockId             = "The unique IpBlock Id"
	FirewallRuleId        = "The unique FirewallRule Id"
	LocationId            = "The unique Location Id"
	LabelKey              = "The unique Label Key"
	LabelValue            = "The unique Label Value"
	UserId                = "The unique User Id"
	GroupId               = "The unique Group Id"
	ResourceId            = "The unique Resource Id"
	S3KeyId               = "The unique User S3Key Id"
	BackupUnitId          = "The unique BackupUnit Id"
	PccId                 = "The unique Private Cross-Connect Id"
	K8sClusterId          = "The unique K8s Cluster Id"
	K8sNodePoolId         = "The unique K8s Node Pool Id"
	K8sNodeId             = "The unique K8s Node Id"
	CdromId               = "The unique Cdrom Id"
	TemplateId            = "The unique Template Id"
	FlowLogId             = "The unique FlowLog Id"
	NatGatewayId          = "The unique NatGateway Id"
	RuleId                = "The unique Rule Id"
	NetworkLoadBalancerId = "The unique NetworkLoadBalancer Id"
	ForwardingRuleId      = "The unique ForwardingRule Id"
)

// Default values
const (
	DefaultOutputFormat    = "text"
	DefaultWait            = false
	DefaultPublic          = false
	DefaultDhcp            = true
	DefaultFirewallActive  = false
	DefaultTimeoutSeconds  = int(60)
	NlbTimeoutSeconds      = int(300)
	K8sTimeoutSeconds      = int(600)
	DefaultServerCores     = 2
	DefaultVolumeSize      = 10
	DefaultNicLanId        = 1
	DefaultServerCPUFamily = "AMD_OPTERON"
)