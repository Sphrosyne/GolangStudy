package main

import (
	"GolangStudy/CloudpodsTest/rbacutils"
	"strings"
	"testing"
	"yunion.io/x/jsonutils"
	"yunion.io/x/log"
)

var (
	allowResult = jsonutils.NewString("allow")
	denyResult  = jsonutils.NewString("deny")
)

type SPolicyData struct {
	Name   string
	Scope  rbacutils.TRbacScope
	Policy jsonutils.JSONObject

	Description   string
	DescriptionCN string
}

// TestGenerateAllPolicies 结果示例：[{sysadmin system {"policy":{"*":"allow"}} System-level full previlliges for any resources 全局任意资源管理权限} {sys-editor system {"policy":{"*":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}},"compute":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"buckets":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","delete":"deny","purge":"deny","upload":"deny"}},"servers":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","create-backup":"deny","create-eip":"deny","createdisk":"deny","delete-backup":"deny","delete-disk":"deny","delete-eip":"deny","purge":"deny","save-image":"deny","snapshot-and-clone":"deny"}}},"k8s":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"kubeclusters":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","add-machines":"deny","clone":"deny","delete-machines":"deny","purge":"deny"}}}}} System-level editor/operator previlliges for any resources 全局任意资源编辑/操作权限} {sys-viewer system {"policy":{"*":{"*":{"*":"deny","get":"allow","list":"allow"}}}} System-level read-only previlliges for any resources 全局任意资源只读权限} {domain-admin domain {"policy":{"*":"allow"}} Domain-level full previlliges for any resources 本域内任意资源管理权限} {domain-editor domain {"policy":{"*":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}},"compute":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"buckets":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","delete":"deny","purge":"deny","upload":"deny"}},"servers":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","create-backup":"deny","create-eip":"deny","createdisk":"deny","delete-backup":"deny","delete-disk":"deny","delete-eip":"deny","purge":"deny","save-image":"deny","snapshot-and-clone":"deny"}}},"k8s":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"kubeclusters":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","add-machines":"deny","clone":"deny","delete-machines":"deny","purge":"deny"}}}}} Domain-level editor/operator previlliges for any resources 本域内任意资源编辑/操作权限} {domain-viewer domain {"policy":{"*":{"*":{"*":"deny","get":"allow","list":"allow"}}}} Domain-level read-only previlliges for any resources 本域内任意资源只读权限} {project-admin project {"policy":{"*":"allow"}} Project-level full previlliges for any resources 本项目内任意资源管理权限} {project-editor project {"policy":{"*":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}},"compute":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"buckets":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","delete":"deny","purge":"deny","upload":"deny"}},"servers":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","create-backup":"deny","create-eip":"deny","createdisk":"deny","delete-backup":"deny","delete-disk":"deny","delete-eip":"deny","purge":"deny","save-image":"deny","snapshot-and-clone":"deny"}}},"k8s":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"kubeclusters":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","add-machines":"deny","clone":"deny","delete-machines":"deny","purge":"deny"}}}}} Project-level editor/operator previlliges for any resources 本项目内任意资源编辑/操作权限} {project-viewer project {"policy":{"*":{"*":{"*":"deny","get":"allow","list":"allow"}}}} Project-level read-only previlliges for any resources 本项目内任意资源只读权限} {sys-dashboard system {"policy":{"compute":{"capabilities":{"*":"deny","list":"allow"},"dashboard":{"*":"deny","get":"allow"},"domain_quotas":{"*":"deny","get":"allow","list":"allow"},"infras_quotas":{"*":"deny","get":"allow","list":"allow"},"project_quotas":{"*":"deny","get":"allow","list":"allow"},"quotas":{"*":"deny","get":"allow","list":"allow"},"region_quotas":{"*":"deny","get":"allow","list":"allow"},"usages":{"*":"deny","get":"allow","list":"allow"},"zone_quotas":{"*":"deny","get":"allow","list":"allow"}},"devtool":{"scriptapplyrecords":{"*":"deny","get":"allow","list":"allow"}},"identity":{"identity_quotas":{"*":"deny","get":"allow","list":"allow"},"projects":{"*":"deny","list":"allow"},"usages":{"*":"deny","get":"allow","list":"allow"}},"image":{"image_quotas":{"*":"deny","get":"allow","list":"allow"},"usages":{"*":"deny","get":"allow","list":"allow"}},"meter":{"bill_conditions":{"*":"deny","list":"allow"}},"monitor":{"alertrecords":{"*":"deny","list":"allow"},"alertresources":{"*":"deny","list":"allow"},"monitorresourcealerts":{"*":"deny","get":"allow","list":"allow"},"nodealerts":{"*":"deny","list":"allow"},"unifiedmonitors":{"*":"deny","perform":"allow"}},"notify":{"notifications":{"*":"deny","get":"allow","list":"allow"},"receivers":{"*":"deny","get":"allow","list":"allow"},"robots":{"*":"deny","get":"allow","list":"allow"}},"yunionconf":{"scopedpolicybindings":{"*":"deny","get":"allow","list":"allow"}}}} System-level previlliges for resources for viewing dashboard 全局控制面板查看相关资源权限} {domain-dashboard domain {"policy":{"compute":{"capabilities":{"*":"deny","list":"allow"},"dashboard":{"*":"deny","get":"allow"},"domain_quotas":{"*":"deny","get":"allow","list":"allow"},"infras_quotas":{"*":"deny","get":"allow","list":"allow"},"project_quotas":{"*":"deny","get":"allow","list":"allow"},"quotas":{"*":"deny","get":"allow","list":"allow"},"region_quotas":{"*":"deny","get":"allow","list":"allow"},"usages":{"*":"deny","get":"allow","list":"allow"},"zone_quotas":{"*":"deny","get":"allow","list":"allow"}},"devtool":{"scriptapplyrecords":{"*":"deny","get":"allow","list":"allow"}},"identity":{"identity_quotas":{"*":"deny","get":"allow","list":"allow"},"projects":{"*":"deny","list":"allow"},"usages":{"*":"deny","get":"allow","list":"allow"}},"image":{"image_quotas":{"*":"deny","get":"allow","list":"allow"},"usages":{"*":"deny","get":"allow","list":"allow"}},"meter":{"bill_conditions":{"*":"deny","list":"allow"}},"monitor":{"alertrecords":{"*":"deny","list":"allow"},"alertresources":{"*":"deny","list":"allow"},"monitorresourcealerts":{"*":"deny","get":"allow","list":"allow"},"nodealerts":{"*":"deny","list":"allow"},"unifiedmonitors":{"*":"deny","perform":"allow"}},"notify":{"notifications":{"*":"deny","get":"allow","list":"allow"},"receivers":{"*":"deny","get":"allow","list":"allow"},"robots":{"*":"deny","get":"allow","list":"allow"}},"yunionconf":{"scopedpolicybindings":{"*":"deny","get":"allow","list":"allow"}}}} Domain-level previlliges for resources for viewing dashboard 本域内控制面板查看相关资源权限} {project-dashboard project {"policy":{"compute":{"capabilities":{"*":"deny","list":"allow"},"dashboard":{"*":"deny","get":"allow"},"domain_quotas":{"*":"deny","get":"allow","list":"allow"},"infras_quotas":{"*":"deny","get":"allow","list":"allow"},"project_quotas":{"*":"deny","get":"allow","list":"allow"},"quotas":{"*":"deny","get":"allow","list":"allow"},"region_quotas":{"*":"deny","get":"allow","list":"allow"},"usages":{"*":"deny","get":"allow","list":"allow"},"zone_quotas":{"*":"deny","get":"allow","list":"allow"}},"devtool":{"scriptapplyrecords":{"*":"deny","get":"allow","list":"allow"}},"identity":{"identity_quotas":{"*":"deny","get":"allow","list":"allow"},"projects":{"*":"deny","list":"allow"},"usages":{"*":"deny","get":"allow","list":"allow"}},"image":{"image_quotas":{"*":"deny","get":"allow","list":"allow"},"usages":{"*":"deny","get":"allow","list":"allow"}},"meter":{"bill_conditions":{"*":"deny","list":"allow"}},"monitor":{"alertrecords":{"*":"deny","list":"allow"},"alertresources":{"*":"deny","list":"allow"},"monitorresourcealerts":{"*":"deny","get":"allow","list":"allow"},"nodealerts":{"*":"deny","list":"allow"},"unifiedmonitors":{"*":"deny","perform":"allow"}},"notify":{"notifications":{"*":"deny","get":"allow","list":"allow"},"receivers":{"*":"deny","get":"allow","list":"allow"},"robots":{"*":"deny","get":"allow","list":"allow"}},"yunionconf":{"scopedpolicybindings":{"*":"deny","get":"allow","list":"allow"}}}} Project-level previlliges for resources for viewing dashboard 本项目内控制面板查看相关资源权限} {sys-compute-admin system {"policy":{"compute":"allow","image":"allow","k8s":"allow"}} System-level full previlliges for resources of computing (cloud servers and containers) 全局计算服务(云主机与容器)相关资源管理权限} {sys-compute-editor system {"policy":{"compute":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"buckets":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","delete":"deny","purge":"deny","upload":"deny"}},"servers":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","create-backup":"deny","create-eip":"deny","createdisk":"deny","delete-backup":"deny","delete-disk":"deny","delete-eip":"deny","purge":"deny","save-image":"deny","snapshot-and-clone":"deny"}}},"image":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}},"k8s":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"kubeclusters":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","add-machines":"deny","clone":"deny","delete-machines":"deny","purge":"deny"}}}}} System-level editor/operator previlliges for resources of computing (cloud servers and containers) 全局计算服务(云主机与容器)相关资源编辑/操作权限} {sys-compute-viewer system {"policy":{"compute":{"*":{"*":"deny","get":"allow","list":"allow"}},"image":{"*":{"*":"deny","get":"allow","list":"allow"}},"k8s":{"*":{"*":"deny","get":"allow","list":"allow"}}}} System-level read-only previlliges for resources of computing (cloud servers and containers) 全局计算服务(云主机与容器)相关资源只读权限} {domain-compute-admin domain {"policy":{"compute":"allow","image":"allow","k8s":"allow"}} Domain-level full previlliges for resources of computing (cloud servers and containers) 本域内计算服务(云主机与容器)相关资源管理权限} {domain-compute-editor domain {"policy":{"compute":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"buckets":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","delete":"deny","purge":"deny","upload":"deny"}},"servers":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","create-backup":"deny","create-eip":"deny","createdisk":"deny","delete-backup":"deny","delete-disk":"deny","delete-eip":"deny","purge":"deny","save-image":"deny","snapshot-and-clone":"deny"}}},"image":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}},"k8s":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"kubeclusters":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","add-machines":"deny","clone":"deny","delete-machines":"deny","purge":"deny"}}}}} Domain-level editor/operator previlliges for resources of computing (cloud servers and containers) 本域内计算服务(云主机与容器)相关资源编辑/操作权限} {domain-compute-viewer domain {"policy":{"compute":{"*":{"*":"deny","get":"allow","list":"allow"}},"image":{"*":{"*":"deny","get":"allow","list":"allow"}},"k8s":{"*":{"*":"deny","get":"allow","list":"allow"}}}} Domain-level read-only previlliges for resources of computing (cloud servers and containers) 本域内计算服务(云主机与容器)相关资源只读权限} {project-compute-admin project {"policy":{"compute":"allow","image":"allow","k8s":"allow"}} Project-level full previlliges for resources of computing (cloud servers and containers) 本项目内计算服务(云主机与容器)相关资源管理权限} {project-compute-editor project {"policy":{"compute":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"buckets":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","delete":"deny","purge":"deny","upload":"deny"}},"servers":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","create-backup":"deny","create-eip":"deny","createdisk":"deny","delete-backup":"deny","delete-disk":"deny","delete-eip":"deny","purge":"deny","save-image":"deny","snapshot-and-clone":"deny"}}},"image":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}},"k8s":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"kubeclusters":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","add-machines":"deny","clone":"deny","delete-machines":"deny","purge":"deny"}}}}} Project-level editor/operator previlliges for resources of computing (cloud servers and containers) 本项目内计算服务(云主机与容器)相关资源编辑/操作权限} {project-compute-viewer project {"policy":{"compute":{"*":{"*":"deny","get":"allow","list":"allow"}},"image":{"*":{"*":"deny","get":"allow","list":"allow"}},"k8s":{"*":{"*":"deny","get":"allow","list":"allow"}}}} Project-level read-only previlliges for resources of computing (cloud servers and containers) 本项目内计算服务(云主机与容器)相关资源只读权限} {sys-server-admin system {"policy":{"compute":{"disks":"allow","eips":"allow","instance_snapshots":"allow","instancegroups":"allow","isolated_devices":{"*":"deny","get":"allow","list":"allow"},"networks":"allow","scalingactivities":"allow","scalinggroups":"allow","scalingpolicies":"allow","secgroupcaches":"allow","secgrouprules":"allow","secgroups":"allow","servers":"allow","servertemplates":"allow","snapshotpolicies":"allow","snapshotpolicycaches":"allow","snapshotpolicydisks":"allow","snapshots":"allow"},"image":"allow"}} System-level full previlliges for resources of cloud servers 全局云主机相关资源管理权限} {sys-server-editor system {"policy":{"compute":{"disks":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"eips":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"instance_snapshots":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"instancegroups":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"isolated_devices":{"*":"deny","get":"allow","list":"allow"},"networks":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"scalingactivities":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"scalinggroups":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"scalingpolicies":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"secgroupcaches":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"secgrouprules":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"secgroups":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"servers":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","create-backup":"deny","create-eip":"deny","createdisk":"deny","delete-backup":"deny","delete-disk":"deny","delete-eip":"deny","purge":"deny","save-image":"deny","snapshot-and-clone":"deny"}},"servertemplates":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"snapshotpolicies":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"snapshotpolicycaches":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"snapshotpolicydisks":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"snapshots":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}},"image":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}}}} System-level editor/operator previlliges for resources of cloud servers 全局云主机相关资源编辑/操作权限} {sys-server-viewer system {"policy":{"compute":{"disks":{"*":"deny","get":"allow","list":"allow"},"eips":{"*":"deny","get":"allow","list":"allow"},"instance_snapshots":{"*":"deny","get":"allow","list":"allow"},"instancegroups":{"*":"deny","get":"allow","list":"allow"},"isolated_devices":{"*":"deny","get":"allow","list":"allow"},"networks":{"*":"deny","get":"allow","list":"allow"},"scalingactivities":{"*":"deny","get":"allow","list":"allow"},"scalinggroups":{"*":"deny","get":"allow","list":"allow"},"scalingpolicies":{"*":"deny","get":"allow","list":"allow"},"secgroupcaches":{"*":"deny","get":"allow","list":"allow"},"secgrouprules":{"*":"deny","get":"allow","list":"allow"},"secgroups":{"*":"deny","get":"allow","list":"allow"},"servers":{"*":"deny","get":"allow","list":"allow"},"servertemplates":{"*":"deny","get":"allow","list":"allow"},"snapshotpolicies":{"*":"deny","get":"allow","list":"allow"},"snapshotpolicycaches":{"*":"deny","get":"allow","list":"allow"},"snapshotpolicydisks":{"*":"deny","get":"allow","list":"allow"},"snapshots":{"*":"deny","get":"allow","list":"allow"}},"image":{"*":{"*":"deny","get":"allow","list":"allow"}}}} System-level read-only previlliges for resources of cloud servers 全局云主机相关资源只读权限} {domain-server-admin domain {"policy":{"compute":{"disks":"allow","eips":"allow","instance_snapshots":"allow","instancegroups":"allow","isolated_devices":{"*":"deny","get":"allow","list":"allow"},"networks":"allow","scalingactivities":"allow","scalinggroups":"allow","scalingpolicies":"allow","secgroupcaches":"allow","secgrouprules":"allow","secgroups":"allow","servers":"allow","servertemplates":"allow","snapshotpolicies":"allow","snapshotpolicycaches":"allow","snapshotpolicydisks":"allow","snapshots":"allow"},"image":"allow"}} Domain-level full previlliges for resources of cloud servers 本域内云主机相关资源管理权限} {domain-server-editor domain {"policy":{"compute":{"disks":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"eips":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"instance_snapshots":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"instancegroups":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"isolated_devices":{"*":"deny","get":"allow","list":"allow"},"networks":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"scalingactivities":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"scalinggroups":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"scalingpolicies":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"secgroupcaches":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"secgrouprules":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"secgroups":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"servers":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","create-backup":"deny","create-eip":"deny","createdisk":"deny","delete-backup":"deny","delete-disk":"deny","delete-eip":"deny","purge":"deny","save-image":"deny","snapshot-and-clone":"deny"}},"servertemplates":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"snapshotpolicies":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"snapshotpolicycaches":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"snapshotpolicydisks":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"snapshots":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}},"image":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}}}} Domain-level editor/operator previlliges for resources of cloud servers 本域内云主机相关资源编辑/操作权限} {domain-server-viewer domain {"policy":{"compute":{"disks":{"*":"deny","get":"allow","list":"allow"},"eips":{"*":"deny","get":"allow","list":"allow"},"instance_snapshots":{"*":"deny","get":"allow","list":"allow"},"instancegroups":{"*":"deny","get":"allow","list":"allow"},"isolated_devices":{"*":"deny","get":"allow","list":"allow"},"networks":{"*":"deny","get":"allow","list":"allow"},"scalingactivities":{"*":"deny","get":"allow","list":"allow"},"scalinggroups":{"*":"deny","get":"allow","list":"allow"},"scalingpolicies":{"*":"deny","get":"allow","list":"allow"},"secgroupcaches":{"*":"deny","get":"allow","list":"allow"},"secgrouprules":{"*":"deny","get":"allow","list":"allow"},"secgroups":{"*":"deny","get":"allow","list":"allow"},"servers":{"*":"deny","get":"allow","list":"allow"},"servertemplates":{"*":"deny","get":"allow","list":"allow"},"snapshotpolicies":{"*":"deny","get":"allow","list":"allow"},"snapshotpolicycaches":{"*":"deny","get":"allow","list":"allow"},"snapshotpolicydisks":{"*":"deny","get":"allow","list":"allow"},"snapshots":{"*":"deny","get":"allow","list":"allow"}},"image":{"*":{"*":"deny","get":"allow","list":"allow"}}}} Domain-level read-only previlliges for resources of cloud servers 本域内云主机相关资源只读权限} {project-server-admin project {"policy":{"compute":{"disks":"allow","eips":"allow","instance_snapshots":"allow","instancegroups":"allow","isolated_devices":{"*":"deny","get":"allow","list":"allow"},"networks":"allow","scalingactivities":"allow","scalinggroups":"allow","scalingpolicies":"allow","secgroupcaches":"allow","secgrouprules":"allow","secgroups":"allow","servers":"allow","servertemplates":"allow","snapshotpolicies":"allow","snapshotpolicycaches":"allow","snapshotpolicydisks":"allow","snapshots":"allow"},"image":"allow"}} Project-level full previlliges for resources of cloud servers 本项目内云主机相关资源管理权限} {project-server-editor project {"policy":{"compute":{"disks":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"eips":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"instance_snapshots":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"instancegroups":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"isolated_devices":{"*":"deny","get":"allow","list":"allow"},"networks":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"scalingactivities":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"scalinggroups":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"scalingpolicies":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"secgroupcaches":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"secgrouprules":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"secgroups":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"servers":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","create-backup":"deny","create-eip":"deny","createdisk":"deny","delete-backup":"deny","delete-disk":"deny","delete-eip":"deny","purge":"deny","save-image":"deny","snapshot-and-clone":"deny"}},"servertemplates":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"snapshotpolicies":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"snapshotpolicycaches":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"snapshotpolicydisks":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"snapshots":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}},"image":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}}}} Project-level editor/operator previlliges for resources of cloud servers 本项目内云主机相关资源编辑/操作权限} {project-server-viewer project {"policy":{"compute":{"disks":{"*":"deny","get":"allow","list":"allow"},"eips":{"*":"deny","get":"allow","list":"allow"},"instance_snapshots":{"*":"deny","get":"allow","list":"allow"},"instancegroups":{"*":"deny","get":"allow","list":"allow"},"isolated_devices":{"*":"deny","get":"allow","list":"allow"},"networks":{"*":"deny","get":"allow","list":"allow"},"scalingactivities":{"*":"deny","get":"allow","list":"allow"},"scalinggroups":{"*":"deny","get":"allow","list":"allow"},"scalingpolicies":{"*":"deny","get":"allow","list":"allow"},"secgroupcaches":{"*":"deny","get":"allow","list":"allow"},"secgrouprules":{"*":"deny","get":"allow","list":"allow"},"secgroups":{"*":"deny","get":"allow","list":"allow"},"servers":{"*":"deny","get":"allow","list":"allow"},"servertemplates":{"*":"deny","get":"allow","list":"allow"},"snapshotpolicies":{"*":"deny","get":"allow","list":"allow"},"snapshotpolicycaches":{"*":"deny","get":"allow","list":"allow"},"snapshotpolicydisks":{"*":"deny","get":"allow","list":"allow"},"snapshots":{"*":"deny","get":"allow","list":"allow"}},"image":{"*":{"*":"deny","get":"allow","list":"allow"}}}} Project-level read-only previlliges for resources of cloud servers 本项目内云主机相关资源只读权限} {sys-host-admin system {"policy":{"compute":{"baremetalagents":"allow","baremetalevents":"allow","baremetalnetworks":"allow","hosts":"allow","hoststorages":"allow","hostwires":"allow","isolated_devices":"allow"}}} System-level full previlliges for resources of hosts and baremetals 全局宿主机和物理机相关资源管理权限} {sys-host-editor system {"policy":{"compute":{"baremetalagents":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"baremetalevents":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"baremetalnetworks":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"hosts":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"hoststorages":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"hostwires":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"isolated_devices":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}}}} System-level editor/operator previlliges for resources of hosts and baremetals 全局宿主机和物理机相关资源编辑/操作权限} {sys-host-viewer system {"policy":{"compute":{"baremetalagents":{"*":"deny","get":"allow","list":"allow"},"baremetalevents":{"*":"deny","get":"allow","list":"allow"},"baremetalnetworks":{"*":"deny","get":"allow","list":"allow"},"hosts":{"*":"deny","get":"allow","list":"allow"},"hoststorages":{"*":"deny","get":"allow","list":"allow"},"hostwires":{"*":"deny","get":"allow","list":"allow"},"isolated_devices":{"*":"deny","get":"allow","list":"allow"}}}} System-level read-only previlliges for resources of hosts and baremetals 全局宿主机和物理机相关资源只读权限} {domain-host-admin domain {"policy":{"compute":{"baremetalagents":"allow","baremetalevents":"allow","baremetalnetworks":"allow","hosts":"allow","hoststorages":"allow","hostwires":"allow","isolated_devices":"allow"}}} Domain-level full previlliges for resources of hosts and baremetals 本域内宿主机和物理机相关资源管理权限} {domain-host-editor domain {"policy":{"compute":{"baremetalagents":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"baremetalevents":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"baremetalnetworks":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"hosts":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"hoststorages":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"hostwires":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"isolated_devices":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}}}} Domain-level editor/operator previlliges for resources of hosts and baremetals 本域内宿主机和物理机相关资源编辑/操作权限} {domain-host-viewer domain {"policy":{"compute":{"baremetalagents":{"*":"deny","get":"allow","list":"allow"},"baremetalevents":{"*":"deny","get":"allow","list":"allow"},"baremetalnetworks":{"*":"deny","get":"allow","list":"allow"},"hosts":{"*":"deny","get":"allow","list":"allow"},"hoststorages":{"*":"deny","get":"allow","list":"allow"},"hostwires":{"*":"deny","get":"allow","list":"allow"},"isolated_devices":{"*":"deny","get":"allow","list":"allow"}}}} Domain-level read-only previlliges for resources of hosts and baremetals 本域内宿主机和物理机相关资源只读权限} {sys-storage-admin system {"policy":{"compute":{"storages":"allow"}}} System-level full previlliges for resources of cloud disk storages 全局云硬盘存储相关资源管理权限} {sys-storage-editor system {"policy":{"compute":{"storages":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}}}} System-level editor/operator previlliges for resources of cloud disk storages 全局云硬盘存储相关资源编辑/操作权限} {sys-storage-viewer system {"policy":{"compute":{"storages":{"*":"deny","get":"allow","list":"allow"}}}} System-level read-only previlliges for resources of cloud disk storages 全局云硬盘存储相关资源只读权限} {domain-storage-admin domain {"policy":{"compute":{"storages":"allow"}}} Domain-level full previlliges for resources of cloud disk storages 本域内云硬盘存储相关资源管理权限} {domain-storage-editor domain {"policy":{"compute":{"storages":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}}}} Domain-level editor/operator previlliges for resources of cloud disk storages 本域内云硬盘存储相关资源编辑/操作权限} {domain-storage-viewer domain {"policy":{"compute":{"storages":{"*":"deny","get":"allow","list":"allow"}}}} Domain-level read-only previlliges for resources of cloud disk storages 本域内云硬盘存储相关资源只读权限} {sys-loadbalancer-admin system {"policy":{"compute":{"loadbalanceracls":"allow","loadbalanceragents":"allow","loadbalancerbackendgroups":"allow","loadbalancerbackends":"allow","loadbalancercertificates":"allow","loadbalancerclusters":"allow","loadbalancerlistenerrules":"allow","loadbalancerlisteners":"allow","loadbalancernetworks":"allow","loadbalancers":"allow","networks":{"*":"deny","get":"allow","list":"allow"}}}} System-level full previlliges for resources of load balancers 全局负载均衡相关资源管理权限} {sys-loadbalancer-editor system {"policy":{"compute":{"loadbalanceracls":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"loadbalanceragents":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"loadbalancerbackendgroups":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"loadbalancerbackends":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"loadbalancercertificates":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"loadbalancerclusters":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"loadbalancerlistenerrules":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"loadbalancerlisteners":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"loadbalancernetworks":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"loadbalancers":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"networks":{"*":"deny","get":"allow","list":"allow"}}}} System-level editor/operator previlliges for resources of load balancers 全局负载均衡相关资源编辑/操作权限} {sys-loadbalancer-viewer system {"policy":{"compute":{"loadbalanceracls":{"*":"deny","get":"allow","list":"allow"},"loadbalanceragents":{"*":"deny","get":"allow","list":"allow"},"loadbalancerbackendgroups":{"*":"deny","get":"allow","list":"allow"},"loadbalancerbackends":{"*":"deny","get":"allow","list":"allow"},"loadbalancercertificates":{"*":"deny","get":"allow","list":"allow"},"loadbalancerclusters":{"*":"deny","get":"allow","list":"allow"},"loadbalancerlistenerrules":{"*":"deny","get":"allow","list":"allow"},"loadbalancerlisteners":{"*":"deny","get":"allow","list":"allow"},"loadbalancernetworks":{"*":"deny","get":"allow","list":"allow"},"loadbalancers":{"*":"deny","get":"allow","list":"allow"},"networks":{"*":"deny","get":"allow","list":"allow"}}}} System-level read-only previlliges for resources of load balancers 全局负载均衡相关资源只读权限} {domain-loadbalancer-admin domain {"policy":{"compute":{"loadbalanceracls":"allow","loadbalanceragents":"allow","loadbalancerbackendgroups":"allow","loadbalancerbackends":"allow","loadbalancercertificates":"allow","loadbalancerclusters":"allow","loadbalancerlistenerrules":"allow","loadbalancerlisteners":"allow","loadbalancernetworks":"allow","loadbalancers":"allow","networks":{"*":"deny","get":"allow","list":"allow"}}}} Domain-level full previlliges for resources of load balancers 本域内负载均衡相关资源管理权限} {domain-loadbalancer-editor domain {"policy":{"compute":{"loadbalanceracls":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"loadbalanceragents":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"loadbalancerbackendgroups":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"loadbalancerbackends":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"loadbalancercertificates":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"loadbalancerclusters":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"loadbalancerlistenerrules":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"loadbalancerlisteners":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"loadbalancernetworks":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"loadbalancers":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"networks":{"*":"deny","get":"allow","list":"allow"}}}} Domain-level editor/operator previlliges for resources of load balancers 本域内负载均衡相关资源编辑/操作权限} {domain-loadbalancer-viewer domain {"policy":{"compute":{"loadbalanceracls":{"*":"deny","get":"allow","list":"allow"},"loadbalanceragents":{"*":"deny","get":"allow","list":"allow"},"loadbalancerbackendgroups":{"*":"deny","get":"allow","list":"allow"},"loadbalancerbackends":{"*":"deny","get":"allow","list":"allow"},"loadbalancercertificates":{"*":"deny","get":"allow","list":"allow"},"loadbalancerclusters":{"*":"deny","get":"allow","list":"allow"},"loadbalancerlistenerrules":{"*":"deny","get":"allow","list":"allow"},"loadbalancerlisteners":{"*":"deny","get":"allow","list":"allow"},"loadbalancernetworks":{"*":"deny","get":"allow","list":"allow"},"loadbalancers":{"*":"deny","get":"allow","list":"allow"},"networks":{"*":"deny","get":"allow","list":"allow"}}}} Domain-level read-only previlliges for resources of load balancers 本域内负载均衡相关资源只读权限} {project-loadbalancer-admin project {"policy":{"compute":{"loadbalanceracls":"allow","loadbalanceragents":"allow","loadbalancerbackendgroups":"allow","loadbalancerbackends":"allow","loadbalancercertificates":"allow","loadbalancerclusters":"allow","loadbalancerlistenerrules":"allow","loadbalancerlisteners":"allow","loadbalancernetworks":"allow","loadbalancers":"allow","networks":{"*":"deny","get":"allow","list":"allow"}}}} Project-level full previlliges for resources of load balancers 本项目内负载均衡相关资源管理权限} {project-loadbalancer-editor project {"policy":{"compute":{"loadbalanceracls":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"loadbalanceragents":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"loadbalancerbackendgroups":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"loadbalancerbackends":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"loadbalancercertificates":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"loadbalancerclusters":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"loadbalancerlistenerrules":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"loadbalancerlisteners":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"loadbalancernetworks":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"loadbalancers":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"networks":{"*":"deny","get":"allow","list":"allow"}}}} Project-level editor/operator previlliges for resources of load balancers 本项目内负载均衡相关资源编辑/操作权限} {project-loadbalancer-viewer project {"policy":{"compute":{"loadbalanceracls":{"*":"deny","get":"allow","list":"allow"},"loadbalanceragents":{"*":"deny","get":"allow","list":"allow"},"loadbalancerbackendgroups":{"*":"deny","get":"allow","list":"allow"},"loadbalancerbackends":{"*":"deny","get":"allow","list":"allow"},"loadbalancercertificates":{"*":"deny","get":"allow","list":"allow"},"loadbalancerclusters":{"*":"deny","get":"allow","list":"allow"},"loadbalancerlistenerrules":{"*":"deny","get":"allow","list":"allow"},"loadbalancerlisteners":{"*":"deny","get":"allow","list":"allow"},"loadbalancernetworks":{"*":"deny","get":"allow","list":"allow"},"loadbalancers":{"*":"deny","get":"allow","list":"allow"},"networks":{"*":"deny","get":"allow","list":"allow"}}}} Project-level read-only previlliges for resources of load balancers 本项目内负载均衡相关资源只读权限} {sys-oss-admin system {"policy":{"compute":{"buckets":"allow"}}} System-level full previlliges for resources of object storages 全局对象存储相关资源管理权限} {sys-oss-editor system {"policy":{"compute":{"buckets":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","delete":"deny","purge":"deny","upload":"deny"}}}}} System-level editor/operator previlliges for resources of object storages 全局对象存储相关资源编辑/操作权限} {sys-oss-viewer system {"policy":{"compute":{"buckets":{"*":"deny","get":"allow","list":"allow"}}}} System-level read-only previlliges for resources of object storages 全局对象存储相关资源只读权限} {domain-oss-admin domain {"policy":{"compute":{"buckets":"allow"}}} Domain-level full previlliges for resources of object storages 本域内对象存储相关资源管理权限} {domain-oss-editor domain {"policy":{"compute":{"buckets":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","delete":"deny","purge":"deny","upload":"deny"}}}}} Domain-level editor/operator previlliges for resources of object storages 本域内对象存储相关资源编辑/操作权限} {domain-oss-viewer domain {"policy":{"compute":{"buckets":{"*":"deny","get":"allow","list":"allow"}}}} Domain-level read-only previlliges for resources of object storages 本域内对象存储相关资源只读权限} {project-oss-admin project {"policy":{"compute":{"buckets":"allow"}}} Project-level full previlliges for resources of object storages 本项目内对象存储相关资源管理权限} {project-oss-editor project {"policy":{"compute":{"buckets":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","delete":"deny","purge":"deny","upload":"deny"}}}}} Project-level editor/operator previlliges for resources of object storages 本项目内对象存储相关资源编辑/操作权限} {project-oss-viewer project {"policy":{"compute":{"buckets":{"*":"deny","get":"allow","list":"allow"}}}} Project-level read-only previlliges for resources of object storages 本项目内对象存储相关资源只读权限} {sys-dbinstance-admin system {"policy":{"compute":{"dbinstance_skus":"allow","dbinstanceaccounts":"allow","dbinstancebackups":"allow","dbinstancedatabases":"allow","dbinstancenetworks":"allow","dbinstanceparameters":"allow","dbinstanceprivileges":"allow","dbinstances":"allow"}}} System-level full previlliges for resources of RDS 全局关系型数据库(MySQL等)相关资源管理权限} {sys-dbinstance-editor system {"policy":{"compute":{"dbinstance_skus":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"dbinstanceaccounts":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"dbinstancebackups":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"dbinstancedatabases":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"dbinstancenetworks":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"dbinstanceparameters":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"dbinstanceprivileges":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"dbinstances":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}}}} System-level editor/operator previlliges for resources of RDS 全局关系型数据库(MySQL等)相关资源编辑/操作权限} {sys-dbinstance-viewer system {"policy":{"compute":{"dbinstance_skus":{"*":"deny","get":"allow","list":"allow"},"dbinstanceaccounts":{"*":"deny","get":"allow","list":"allow"},"dbinstancebackups":{"*":"deny","get":"allow","list":"allow"},"dbinstancedatabases":{"*":"deny","get":"allow","list":"allow"},"dbinstancenetworks":{"*":"deny","get":"allow","list":"allow"},"dbinstanceparameters":{"*":"deny","get":"allow","list":"allow"},"dbinstanceprivileges":{"*":"deny","get":"allow","list":"allow"},"dbinstances":{"*":"deny","get":"allow","list":"allow"}}}} System-level read-only previlliges for resources of RDS 全局关系型数据库(MySQL等)相关资源只读权限} {domain-dbinstance-admin domain {"policy":{"compute":{"dbinstance_skus":"allow","dbinstanceaccounts":"allow","dbinstancebackups":"allow","dbinstancedatabases":"allow","dbinstancenetworks":"allow","dbinstanceparameters":"allow","dbinstanceprivileges":"allow","dbinstances":"allow"}}} Domain-level full previlliges for resources of RDS 本域内关系型数据库(MySQL等)相关资源管理权限} {domain-dbinstance-editor domain {"policy":{"compute":{"dbinstance_skus":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"dbinstanceaccounts":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"dbinstancebackups":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"dbinstancedatabases":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"dbinstancenetworks":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"dbinstanceparameters":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"dbinstanceprivileges":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"dbinstances":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}}}} Domain-level editor/operator previlliges for resources of RDS 本域内关系型数据库(MySQL等)相关资源编辑/操作权限} {domain-dbinstance-viewer domain {"policy":{"compute":{"dbinstance_skus":{"*":"deny","get":"allow","list":"allow"},"dbinstanceaccounts":{"*":"deny","get":"allow","list":"allow"},"dbinstancebackups":{"*":"deny","get":"allow","list":"allow"},"dbinstancedatabases":{"*":"deny","get":"allow","list":"allow"},"dbinstancenetworks":{"*":"deny","get":"allow","list":"allow"},"dbinstanceparameters":{"*":"deny","get":"allow","list":"allow"},"dbinstanceprivileges":{"*":"deny","get":"allow","list":"allow"},"dbinstances":{"*":"deny","get":"allow","list":"allow"}}}} Domain-level read-only previlliges for resources of RDS 本域内关系型数据库(MySQL等)相关资源只读权限} {project-dbinstance-admin project {"policy":{"compute":{"dbinstance_skus":"allow","dbinstanceaccounts":"allow","dbinstancebackups":"allow","dbinstancedatabases":"allow","dbinstancenetworks":"allow","dbinstanceparameters":"allow","dbinstanceprivileges":"allow","dbinstances":"allow"}}} Project-level full previlliges for resources of RDS 本项目内关系型数据库(MySQL等)相关资源管理权限} {project-dbinstance-editor project {"policy":{"compute":{"dbinstance_skus":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"dbinstanceaccounts":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"dbinstancebackups":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"dbinstancedatabases":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"dbinstancenetworks":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"dbinstanceparameters":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"dbinstanceprivileges":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"dbinstances":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}}}} Project-level editor/operator previlliges for resources of RDS 本项目内关系型数据库(MySQL等)相关资源编辑/操作权限} {project-dbinstance-viewer project {"policy":{"compute":{"dbinstance_skus":{"*":"deny","get":"allow","list":"allow"},"dbinstanceaccounts":{"*":"deny","get":"allow","list":"allow"},"dbinstancebackups":{"*":"deny","get":"allow","list":"allow"},"dbinstancedatabases":{"*":"deny","get":"allow","list":"allow"},"dbinstancenetworks":{"*":"deny","get":"allow","list":"allow"},"dbinstanceparameters":{"*":"deny","get":"allow","list":"allow"},"dbinstanceprivileges":{"*":"deny","get":"allow","list":"allow"},"dbinstances":{"*":"deny","get":"allow","list":"allow"}}}} Project-level read-only previlliges for resources of RDS 本项目内关系型数据库(MySQL等)相关资源只读权限} {sys-elasticcache-admin system {"policy":{"compute":{"elasticcacheaccounts":"allow","elasticcacheacls":"allow","elasticcachebackups":"allow","elasticcacheparameters":"allow","elasticcaches":"allow","elasticcacheskus":"allow"}}} System-level full previlliges for resources of elastic caches 全局弹性缓存(Redis等)相关资源管理权限} {sys-elasticcache-editor system {"policy":{"compute":{"elasticcacheaccounts":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"elasticcacheacls":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"elasticcachebackups":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"elasticcacheparameters":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"elasticcaches":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"elasticcacheskus":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}}}} System-level editor/operator previlliges for resources of elastic caches 全局弹性缓存(Redis等)相关资源编辑/操作权限} {sys-elasticcache-viewer system {"policy":{"compute":{"elasticcacheaccounts":{"*":"deny","get":"allow","list":"allow"},"elasticcacheacls":{"*":"deny","get":"allow","list":"allow"},"elasticcachebackups":{"*":"deny","get":"allow","list":"allow"},"elasticcacheparameters":{"*":"deny","get":"allow","list":"allow"},"elasticcaches":{"*":"deny","get":"allow","list":"allow"},"elasticcacheskus":{"*":"deny","get":"allow","list":"allow"}}}} System-level read-only previlliges for resources of elastic caches 全局弹性缓存(Redis等)相关资源只读权限} {domain-elasticcache-admin domain {"policy":{"compute":{"elasticcacheaccounts":"allow","elasticcacheacls":"allow","elasticcachebackups":"allow","elasticcacheparameters":"allow","elasticcaches":"allow","elasticcacheskus":"allow"}}} Domain-level full previlliges for resources of elastic caches 本域内弹性缓存(Redis等)相关资源管理权限} {domain-elasticcache-editor domain {"policy":{"compute":{"elasticcacheaccounts":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"elasticcacheacls":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"elasticcachebackups":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"elasticcacheparameters":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"elasticcaches":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"elasticcacheskus":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}}}} Domain-level editor/operator previlliges for resources of elastic caches 本域内弹性缓存(Redis等)相关资源编辑/操作权限} {domain-elasticcache-viewer domain {"policy":{"compute":{"elasticcacheaccounts":{"*":"deny","get":"allow","list":"allow"},"elasticcacheacls":{"*":"deny","get":"allow","list":"allow"},"elasticcachebackups":{"*":"deny","get":"allow","list":"allow"},"elasticcacheparameters":{"*":"deny","get":"allow","list":"allow"},"elasticcaches":{"*":"deny","get":"allow","list":"allow"},"elasticcacheskus":{"*":"deny","get":"allow","list":"allow"}}}} Domain-level read-only previlliges for resources of elastic caches 本域内弹性缓存(Redis等)相关资源只读权限} {project-elasticcache-admin project {"policy":{"compute":{"elasticcacheaccounts":"allow","elasticcacheacls":"allow","elasticcachebackups":"allow","elasticcacheparameters":"allow","elasticcaches":"allow","elasticcacheskus":"allow"}}} Project-level full previlliges for resources of elastic caches 本项目内弹性缓存(Redis等)相关资源管理权限} {project-elasticcache-editor project {"policy":{"compute":{"elasticcacheaccounts":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"elasticcacheacls":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"elasticcachebackups":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"elasticcacheparameters":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"elasticcaches":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"elasticcacheskus":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}}}} Project-level editor/operator previlliges for resources of elastic caches 本项目内弹性缓存(Redis等)相关资源编辑/操作权限} {project-elasticcache-viewer project {"policy":{"compute":{"elasticcacheaccounts":{"*":"deny","get":"allow","list":"allow"},"elasticcacheacls":{"*":"deny","get":"allow","list":"allow"},"elasticcachebackups":{"*":"deny","get":"allow","list":"allow"},"elasticcacheparameters":{"*":"deny","get":"allow","list":"allow"},"elasticcaches":{"*":"deny","get":"allow","list":"allow"},"elasticcacheskus":{"*":"deny","get":"allow","list":"allow"}}}} Project-level read-only previlliges for resources of elastic caches 本项目内弹性缓存(Redis等)相关资源只读权限} {sys-network-admin system {"policy":{"compute":{"dns_recordsets":"allow","dns_trafficpolicies":"allow","dns_zonecaches":"allow","dns_zones":"allow","dnsrecords":"allow","eips":"allow","globalvpcs":"allow","natdentries":"allow","natgateways":"allow","natsentries":"allow","networkinterfacenetworks":"allow","networkinterfaces":"allow","networks":"allow","reservedips":"allow","route_tables":"allow","vpc_peering_connections":"allow","vpcs":"allow","wires":"allow"}}} System-level full previlliges for resources of networking 全局网络相关资源管理权限} {sys-network-editor system {"policy":{"compute":{"dns_recordsets":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"dns_trafficpolicies":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"dns_zonecaches":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"dns_zones":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"dnsrecords":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"eips":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"globalvpcs":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"natdentries":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"natgateways":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"natsentries":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"networkinterfacenetworks":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"networkinterfaces":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"networks":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"reservedips":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"route_tables":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"vpc_peering_connections":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"vpcs":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"wires":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}}}} System-level editor/operator previlliges for resources of networking 全局网络相关资源编辑/操作权限} {sys-network-viewer system {"policy":{"compute":{"dns_recordsets":{"*":"deny","get":"allow","list":"allow"},"dns_trafficpolicies":{"*":"deny","get":"allow","list":"allow"},"dns_zonecaches":{"*":"deny","get":"allow","list":"allow"},"dns_zones":{"*":"deny","get":"allow","list":"allow"},"dnsrecords":{"*":"deny","get":"allow","list":"allow"},"eips":{"*":"deny","get":"allow","list":"allow"},"globalvpcs":{"*":"deny","get":"allow","list":"allow"},"natdentries":{"*":"deny","get":"allow","list":"allow"},"natgateways":{"*":"deny","get":"allow","list":"allow"},"natsentries":{"*":"deny","get":"allow","list":"allow"},"networkinterfacenetworks":{"*":"deny","get":"allow","list":"allow"},"networkinterfaces":{"*":"deny","get":"allow","list":"allow"},"networks":{"*":"deny","get":"allow","list":"allow"},"reservedips":{"*":"deny","get":"allow","list":"allow"},"route_tables":{"*":"deny","get":"allow","list":"allow"},"vpc_peering_connections":{"*":"deny","get":"allow","list":"allow"},"vpcs":{"*":"deny","get":"allow","list":"allow"},"wires":{"*":"deny","get":"allow","list":"allow"}}}} System-level read-only previlliges for resources of networking 全局网络相关资源只读权限} {domain-network-admin domain {"policy":{"compute":{"dns_recordsets":"allow","dns_trafficpolicies":"allow","dns_zonecaches":"allow","dns_zones":"allow","dnsrecords":"allow","eips":"allow","globalvpcs":"allow","natdentries":"allow","natgateways":"allow","natsentries":"allow","networkinterfacenetworks":"allow","networkinterfaces":"allow","networks":"allow","reservedips":"allow","route_tables":"allow","vpc_peering_connections":"allow","vpcs":"allow","wires":"allow"}}} Domain-level full previlliges for resources of networking 本域内网络相关资源管理权限} {domain-network-editor domain {"policy":{"compute":{"dns_recordsets":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"dns_trafficpolicies":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"dns_zonecaches":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"dns_zones":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"dnsrecords":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"eips":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"globalvpcs":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"natdentries":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"natgateways":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"natsentries":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"networkinterfacenetworks":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"networkinterfaces":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"networks":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"reservedips":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"route_tables":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"vpc_peering_connections":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"vpcs":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"wires":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}}}} Domain-level editor/operator previlliges for resources of networking 本域内网络相关资源编辑/操作权限} {domain-network-viewer domain {"policy":{"compute":{"dns_recordsets":{"*":"deny","get":"allow","list":"allow"},"dns_trafficpolicies":{"*":"deny","get":"allow","list":"allow"},"dns_zonecaches":{"*":"deny","get":"allow","list":"allow"},"dns_zones":{"*":"deny","get":"allow","list":"allow"},"dnsrecords":{"*":"deny","get":"allow","list":"allow"},"eips":{"*":"deny","get":"allow","list":"allow"},"globalvpcs":{"*":"deny","get":"allow","list":"allow"},"natdentries":{"*":"deny","get":"allow","list":"allow"},"natgateways":{"*":"deny","get":"allow","list":"allow"},"natsentries":{"*":"deny","get":"allow","list":"allow"},"networkinterfacenetworks":{"*":"deny","get":"allow","list":"allow"},"networkinterfaces":{"*":"deny","get":"allow","list":"allow"},"networks":{"*":"deny","get":"allow","list":"allow"},"reservedips":{"*":"deny","get":"allow","list":"allow"},"route_tables":{"*":"deny","get":"allow","list":"allow"},"vpc_peering_connections":{"*":"deny","get":"allow","list":"allow"},"vpcs":{"*":"deny","get":"allow","list":"allow"},"wires":{"*":"deny","get":"allow","list":"allow"}}}} Domain-level read-only previlliges for resources of networking 本域内网络相关资源只读权限} {sys-snapshotpolicy-admin system {"policy":{"compute":{"snapshotpolicies":"allow","snapshotpolicydisks":"allow"}}} System-level full previlliges for snapshot policy 全局快照策略管理权限} {sys-snapshotpolicy-editor system {"policy":{"compute":{"snapshotpolicies":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"snapshotpolicydisks":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}}}} System-level editor/operator previlliges for snapshot policy 全局快照策略编辑/操作权限} {sys-snapshotpolicy-viewer system {"policy":{"compute":{"snapshotpolicies":{"*":"deny","get":"allow","list":"allow"},"snapshotpolicydisks":{"*":"deny","get":"allow","list":"allow"}}}} System-level read-only previlliges for snapshot policy 全局快照策略只读权限} {domain-snapshotpolicy-admin domain {"policy":{"compute":{"snapshotpolicies":"allow","snapshotpolicydisks":"allow"}}} Domain-level full previlliges for snapshot policy 本域内快照策略管理权限} {domain-snapshotpolicy-editor domain {"policy":{"compute":{"snapshotpolicies":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"snapshotpolicydisks":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}}}} Domain-level editor/operator previlliges for snapshot policy 本域内快照策略编辑/操作权限} {domain-snapshotpolicy-viewer domain {"policy":{"compute":{"snapshotpolicies":{"*":"deny","get":"allow","list":"allow"},"snapshotpolicydisks":{"*":"deny","get":"allow","list":"allow"}}}} Domain-level read-only previlliges for snapshot policy 本域内快照策略只读权限} {sys-secgroup-admin system {"policy":{"compute":{"secgrouprules":"allow","secgroups":"allow"}}} System-level full previlliges for security group 全局安全组管理权限} {sys-secgroup-editor system {"policy":{"compute":{"secgrouprules":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"secgroups":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}}}} System-level editor/operator previlliges for security group 全局安全组编辑/操作权限} {sys-secgroup-viewer system {"policy":{"compute":{"secgrouprules":{"*":"deny","get":"allow","list":"allow"},"secgroups":{"*":"deny","get":"allow","list":"allow"}}}} System-level read-only previlliges for security group 全局安全组只读权限} {domain-secgroup-admin domain {"policy":{"compute":{"secgrouprules":"allow","secgroups":"allow"}}} Domain-level full previlliges for security group 本域内安全组管理权限} {domain-secgroup-editor domain {"policy":{"compute":{"secgrouprules":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"secgroups":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}}}} Domain-level editor/operator previlliges for security group 本域内安全组编辑/操作权限} {domain-secgroup-viewer domain {"policy":{"compute":{"secgrouprules":{"*":"deny","get":"allow","list":"allow"},"secgroups":{"*":"deny","get":"allow","list":"allow"}}}} Domain-level read-only previlliges for security group 本域内安全组只读权限} {sys-meter-admin system {"policy":{"meter":"allow","notify":{"receivers":"allow"},"suggestion":"allow"}} System-level full previlliges for resources of metering and billing service 全局计费计量分析服务相关资源管理权限} {sys-meter-editor system {"policy":{"meter":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}},"notify":{"receivers":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}},"suggestion":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}}}} System-level editor/operator previlliges for resources of metering and billing service 全局计费计量分析服务相关资源编辑/操作权限} {sys-meter-viewer system {"policy":{"meter":{"*":{"*":"deny","get":"allow","list":"allow"}},"notify":{"receivers":{"*":"deny","get":"allow","list":"allow"}},"suggestion":{"*":{"*":"deny","get":"allow","list":"allow"}}}} System-level read-only previlliges for resources of metering and billing service 全局计费计量分析服务相关资源只读权限} {domain-meter-admin domain {"policy":{"meter":"allow","notify":{"receivers":"allow"},"suggestion":"allow"}} Domain-level full previlliges for resources of metering and billing service 本域内计费计量分析服务相关资源管理权限} {domain-meter-editor domain {"policy":{"meter":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}},"notify":{"receivers":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}},"suggestion":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}}}} Domain-level editor/operator previlliges for resources of metering and billing service 本域内计费计量分析服务相关资源编辑/操作权限} {domain-meter-viewer domain {"policy":{"meter":{"*":{"*":"deny","get":"allow","list":"allow"}},"notify":{"receivers":{"*":"deny","get":"allow","list":"allow"}},"suggestion":{"*":{"*":"deny","get":"allow","list":"allow"}}}} Domain-level read-only previlliges for resources of metering and billing service 本域内计费计量分析服务相关资源只读权限} {project-meter-admin project {"policy":{"meter":"allow","notify":{"receivers":"allow"},"suggestion":"allow"}} Project-level full previlliges for resources of metering and billing service 本项目内计费计量分析服务相关资源管理权限} {project-meter-editor project {"policy":{"meter":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}},"notify":{"receivers":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}},"suggestion":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}}}} Project-level editor/operator previlliges for resources of metering and billing service 本项目内计费计量分析服务相关资源编辑/操作权限} {project-meter-viewer project {"policy":{"meter":{"*":{"*":"deny","get":"allow","list":"allow"}},"notify":{"receivers":{"*":"deny","get":"allow","list":"allow"}},"suggestion":{"*":{"*":"deny","get":"allow","list":"allow"}}}} Project-level read-only previlliges for resources of metering and billing service 本项目内计费计量分析服务相关资源只读权限} {sys-identity-admin system {"policy":{"identity":"allow"}} System-level full previlliges for resources of identity service 全局身份认证(IAM)服务相关资源管理权限} {sys-identity-editor system {"policy":{"identity":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}}}} System-level editor/operator previlliges for resources of identity service 全局身份认证(IAM)服务相关资源编辑/操作权限} {sys-identity-viewer system {"policy":{"identity":{"*":{"*":"deny","get":"allow","list":"allow"}}}} System-level read-only previlliges for resources of identity service 全局身份认证(IAM)服务相关资源只读权限} {domain-identity-admin domain {"policy":{"identity":"allow"}} Domain-level full previlliges for resources of identity service 本域内身份认证(IAM)服务相关资源管理权限} {domain-identity-editor domain {"policy":{"identity":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}}}} Domain-level editor/operator previlliges for resources of identity service 本域内身份认证(IAM)服务相关资源编辑/操作权限} {domain-identity-viewer domain {"policy":{"identity":{"*":{"*":"deny","get":"allow","list":"allow"}}}} Domain-level read-only previlliges for resources of identity service 本域内身份认证(IAM)服务相关资源只读权限} {sys-image-admin system {"policy":{"image":"allow"}} System-level full previlliges for resources of image service 全局镜像服务相关资源管理权限} {sys-image-editor system {"policy":{"image":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}}}} System-level editor/operator previlliges for resources of image service 全局镜像服务相关资源编辑/操作权限} {sys-image-viewer system {"policy":{"image":{"*":{"*":"deny","get":"allow","list":"allow"}}}} System-level read-only previlliges for resources of image service 全局镜像服务相关资源只读权限} {domain-image-admin domain {"policy":{"image":"allow"}} Domain-level full previlliges for resources of image service 本域内镜像服务相关资源管理权限} {domain-image-editor domain {"policy":{"image":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}}}} Domain-level editor/operator previlliges for resources of image service 本域内镜像服务相关资源编辑/操作权限} {domain-image-viewer domain {"policy":{"image":{"*":{"*":"deny","get":"allow","list":"allow"}}}} Domain-level read-only previlliges for resources of image service 本域内镜像服务相关资源只读权限} {project-image-admin project {"policy":{"image":"allow"}} Project-level full previlliges for resources of image service 本项目内镜像服务相关资源管理权限} {project-image-editor project {"policy":{"image":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}}}} Project-level editor/operator previlliges for resources of image service 本项目内镜像服务相关资源编辑/操作权限} {project-image-viewer project {"policy":{"image":{"*":{"*":"deny","get":"allow","list":"allow"}}}} Project-level read-only previlliges for resources of image service 本项目内镜像服务相关资源只读权限} {sys-monitor-admin system {"policy":{"monitor":"allow"}} System-level full previlliges for resources of monitor service 全局监控服务相关资源管理权限} {sys-monitor-editor system {"policy":{"monitor":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}}}} System-level editor/operator previlliges for resources of monitor service 全局监控服务相关资源编辑/操作权限} {sys-monitor-viewer system {"policy":{"monitor":{"*":{"*":"deny","get":"allow","list":"allow"}}}} System-level read-only previlliges for resources of monitor service 全局监控服务相关资源只读权限} {domain-monitor-admin domain {"policy":{"monitor":"allow"}} Domain-level full previlliges for resources of monitor service 本域内监控服务相关资源管理权限} {domain-monitor-editor domain {"policy":{"monitor":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}}}} Domain-level editor/operator previlliges for resources of monitor service 本域内监控服务相关资源编辑/操作权限} {domain-monitor-viewer domain {"policy":{"monitor":{"*":{"*":"deny","get":"allow","list":"allow"}}}} Domain-level read-only previlliges for resources of monitor service 本域内监控服务相关资源只读权限} {project-monitor-admin project {"policy":{"monitor":"allow"}} Project-level full previlliges for resources of monitor service 本项目内监控服务相关资源管理权限} {project-monitor-editor project {"policy":{"monitor":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}}}} Project-level editor/operator previlliges for resources of monitor service 本项目内监控服务相关资源编辑/操作权限} {project-monitor-viewer project {"policy":{"monitor":{"*":{"*":"deny","get":"allow","list":"allow"}}}} Project-level read-only previlliges for resources of monitor service 本项目内监控服务相关资源只读权限} {sys-container-admin system {"policy":{"k8s":"allow"}} System-level full previlliges for resources of container service 全局容器服务相关资源管理权限} {sys-container-editor system {"policy":{"k8s":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"kubeclusters":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","add-machines":"deny","clone":"deny","delete-machines":"deny","purge":"deny"}}}}} System-level editor/operator previlliges for resources of container service 全局容器服务相关资源编辑/操作权限} {sys-container-viewer system {"policy":{"k8s":{"*":{"*":"deny","get":"allow","list":"allow"}}}} System-level read-only previlliges for resources of container service 全局容器服务相关资源只读权限} {domain-container-admin domain {"policy":{"k8s":"allow"}} Domain-level full previlliges for resources of container service 本域内容器服务相关资源管理权限} {domain-container-editor domain {"policy":{"k8s":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"kubeclusters":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","add-machines":"deny","clone":"deny","delete-machines":"deny","purge":"deny"}}}}} Domain-level editor/operator previlliges for resources of container service 本域内容器服务相关资源编辑/操作权限} {domain-container-viewer domain {"policy":{"k8s":{"*":{"*":"deny","get":"allow","list":"allow"}}}} Domain-level read-only previlliges for resources of container service 本域内容器服务相关资源只读权限} {project-container-admin project {"policy":{"k8s":"allow"}} Project-level full previlliges for resources of container service 本项目内容器服务相关资源管理权限} {project-container-editor project {"policy":{"k8s":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"kubeclusters":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","add-machines":"deny","clone":"deny","delete-machines":"deny","purge":"deny"}}}}} Project-level editor/operator previlliges for resources of container service 本项目内容器服务相关资源编辑/操作权限} {project-container-viewer project {"policy":{"k8s":{"*":{"*":"deny","get":"allow","list":"allow"}}}} Project-level read-only previlliges for resources of container service 本项目内容器服务相关资源只读权限} {sys-cloudid-admin system {"policy":{"cloudid":"allow","compute":{"cloudaccounts":"allow","cloudproviders":"allow"},"identity":{"projects":"allow","roles":"allow","users":"allow"}}} System-level full previlliges for resources of service CloudId and IAM 全局云用户及权限管理相关资源管理权限} {sys-cloudid-editor system {"policy":{"cloudid":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}},"compute":{"cloudaccounts":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"cloudproviders":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}},"identity":{"projects":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"roles":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"users":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}}}} System-level editor/operator previlliges for resources of service CloudId and IAM 全局云用户及权限管理相关资源编辑/操作权限} {sys-cloudid-viewer system {"policy":{"cloudid":{"*":{"*":"deny","get":"allow","list":"allow"}},"compute":{"cloudaccounts":{"*":"deny","get":"allow","list":"allow"},"cloudproviders":{"*":"deny","get":"allow","list":"allow"}},"identity":{"projects":{"*":"deny","get":"allow","list":"allow"},"roles":{"*":"deny","get":"allow","list":"allow"},"users":{"*":"deny","get":"allow","list":"allow"}}}} System-level read-only previlliges for resources of service CloudId and IAM 全局云用户及权限管理相关资源只读权限} {domain-cloudid-admin domain {"policy":{"cloudid":"allow","compute":{"cloudaccounts":"allow","cloudproviders":"allow"},"identity":{"projects":"allow","roles":"allow","users":"allow"}}} Domain-level full previlliges for resources of service CloudId and IAM 本域内云用户及权限管理相关资源管理权限} {domain-cloudid-editor domain {"policy":{"cloudid":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}},"compute":{"cloudaccounts":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"cloudproviders":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}},"identity":{"projects":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"roles":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"users":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}}}} Domain-level editor/operator previlliges for resources of service CloudId and IAM 本域内云用户及权限管理相关资源编辑/操作权限} {domain-cloudid-viewer domain {"policy":{"cloudid":{"*":{"*":"deny","get":"allow","list":"allow"}},"compute":{"cloudaccounts":{"*":"deny","get":"allow","list":"allow"},"cloudproviders":{"*":"deny","get":"allow","list":"allow"}},"identity":{"projects":{"*":"deny","get":"allow","list":"allow"},"roles":{"*":"deny","get":"allow","list":"allow"},"users":{"*":"deny","get":"allow","list":"allow"}}}} Domain-level read-only previlliges for resources of service CloudId and IAM 本域内云用户及权限管理相关资源只读权限} {sys-cloudaccount-admin system {"policy":{"compute":{"cloudaccounts":"allow","cloudproviderquotas":"allow","cloudproviderregions":"allow","cloudproviders":"allow"}}} System-level full previlliges for resources for cloud account administration 全局云账号管理相关资源管理权限} {sys-cloudaccount-editor system {"policy":{"compute":{"cloudaccounts":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"cloudproviderquotas":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"cloudproviderregions":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"cloudproviders":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}}}} System-level editor/operator previlliges for resources for cloud account administration 全局云账号管理相关资源编辑/操作权限} {sys-cloudaccount-viewer system {"policy":{"compute":{"cloudaccounts":{"*":"deny","get":"allow","list":"allow"},"cloudproviderquotas":{"*":"deny","get":"allow","list":"allow"},"cloudproviderregions":{"*":"deny","get":"allow","list":"allow"},"cloudproviders":{"*":"deny","get":"allow","list":"allow"}}}} System-level read-only previlliges for resources for cloud account administration 全局云账号管理相关资源只读权限} {domain-cloudaccount-admin domain {"policy":{"compute":{"cloudaccounts":"allow","cloudproviderquotas":"allow","cloudproviderregions":"allow","cloudproviders":"allow"}}} Domain-level full previlliges for resources for cloud account administration 本域内云账号管理相关资源管理权限} {domain-cloudaccount-editor domain {"policy":{"compute":{"cloudaccounts":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"cloudproviderquotas":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"cloudproviderregions":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"cloudproviders":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}}}} Domain-level editor/operator previlliges for resources for cloud account administration 本域内云账号管理相关资源编辑/操作权限} {domain-cloudaccount-viewer domain {"policy":{"compute":{"cloudaccounts":{"*":"deny","get":"allow","list":"allow"},"cloudproviderquotas":{"*":"deny","get":"allow","list":"allow"},"cloudproviderregions":{"*":"deny","get":"allow","list":"allow"},"cloudproviders":{"*":"deny","get":"allow","list":"allow"}}}} Domain-level read-only previlliges for resources for cloud account administration 本域内云账号管理相关资源只读权限} {sys-projectresource-admin system {"policy":{"compute":{"project_quotas":"allow","quotas":"allow","region_quotas":"allow","zone_quotas":"allow"},"identity":{"policies":"allow","projects":"allow","roles":"allow"},"image":{"image_quotas":"allow"}}} System-level full previlliges for resources for project administration 全局项目管理相关资源管理权限} {sys-projectresource-editor system {"policy":{"compute":{"project_quotas":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"quotas":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"region_quotas":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"zone_quotas":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}},"identity":{"policies":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"projects":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"roles":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}},"image":{"image_quotas":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}}}} System-level editor/operator previlliges for resources for project administration 全局项目管理相关资源编辑/操作权限} {sys-projectresource-viewer system {"policy":{"compute":{"project_quotas":{"*":"deny","get":"allow","list":"allow"},"quotas":{"*":"deny","get":"allow","list":"allow"},"region_quotas":{"*":"deny","get":"allow","list":"allow"},"zone_quotas":{"*":"deny","get":"allow","list":"allow"}},"identity":{"policies":{"*":"deny","get":"allow","list":"allow"},"projects":{"*":"deny","get":"allow","list":"allow"},"roles":{"*":"deny","get":"allow","list":"allow"}},"image":{"image_quotas":{"*":"deny","get":"allow","list":"allow"}}}} System-level read-only previlliges for resources for project administration 全局项目管理相关资源只读权限} {domain-projectresource-admin domain {"policy":{"compute":{"project_quotas":"allow","quotas":"allow","region_quotas":"allow","zone_quotas":"allow"},"identity":{"policies":"allow","projects":"allow","roles":"allow"},"image":{"image_quotas":"allow"}}} Domain-level full previlliges for resources for project administration 本域内项目管理相关资源管理权限} {domain-projectresource-editor domain {"policy":{"compute":{"project_quotas":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"quotas":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"region_quotas":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"zone_quotas":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}},"identity":{"policies":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"projects":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"roles":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}},"image":{"image_quotas":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}}}} Domain-level editor/operator previlliges for resources for project administration 本域内项目管理相关资源编辑/操作权限} {domain-projectresource-viewer domain {"policy":{"compute":{"project_quotas":{"*":"deny","get":"allow","list":"allow"},"quotas":{"*":"deny","get":"allow","list":"allow"},"region_quotas":{"*":"deny","get":"allow","list":"allow"},"zone_quotas":{"*":"deny","get":"allow","list":"allow"}},"identity":{"policies":{"*":"deny","get":"allow","list":"allow"},"projects":{"*":"deny","get":"allow","list":"allow"},"roles":{"*":"deny","get":"allow","list":"allow"}},"image":{"image_quotas":{"*":"deny","get":"allow","list":"allow"}}}} Domain-level read-only previlliges for resources for project administration 本域内项目管理相关资源只读权限} {domainresource-admin system {"policy":{"compute":{"domain_quotas":"allow","infras_quotas":"allow"},"identity":{"domains":"allow","groups":"allow","identity_quotas":"allow","policies":"allow","projects":"allow","roles":"allow","users":"allow"}}} System-level full previlliges for resources for domain administration 全局域管理相关资源管理权限} {domainresource-editor system {"policy":{"compute":{"domain_quotas":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"infras_quotas":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}},"identity":{"domains":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"groups":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"identity_quotas":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"policies":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"projects":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"roles":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}},"users":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}}}} System-level editor/operator previlliges for resources for domain administration 全局域管理相关资源编辑/操作权限} {domainresource-viewer system {"policy":{"compute":{"domain_quotas":{"*":"deny","get":"allow","list":"allow"},"infras_quotas":{"*":"deny","get":"allow","list":"allow"}},"identity":{"domains":{"*":"deny","get":"allow","list":"allow"},"groups":{"*":"deny","get":"allow","list":"allow"},"identity_quotas":{"*":"deny","get":"allow","list":"allow"},"policies":{"*":"deny","get":"allow","list":"allow"},"projects":{"*":"deny","get":"allow","list":"allow"},"roles":{"*":"deny","get":"allow","list":"allow"},"users":{"*":"deny","get":"allow","list":"allow"}}}} System-level read-only previlliges for resources for domain administration 全局域管理相关资源只读权限} {notify-admin system {"policy":{"notify":"allow"}} System-level full previlliges for resources of notify service 全局通知服务相关资源管理权限} {notify-editor system {"policy":{"notify":{"*":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","purge":"deny"}}}}} System-level editor/operator previlliges for resources of notify service 全局通知服务相关资源编辑/操作权限} {notify-viewer system {"policy":{"notify":{"*":{"*":"deny","get":"allow","list":"allow"}}}} System-level read-only previlliges for resources of notify service 全局通知服务相关资源只读权限} {sys-opsadmin system {"policy":{"*":{"events":{"*":"deny","list":{"*":"deny","splitable":"deny"}}},"compute":{"*":{"*":"allow"},"dynamicschedtags":{"*":"deny"},"events":{"*":"deny","list":{"*":"deny","splitable":"deny"}},"recyclebins":{"*":"deny"},"secgroups":{"*":"deny","get":"allow","list":"allow"},"servers":{"*":"allow","perform":{"*":"allow","add-secgroup":"deny","assign-admin-secgroup":"deny","assign-secgroup":"deny","change-owner":"deny","revoke-admin-secgroup":"deny","revoke-secgroup":"deny","set-secgroup":"deny","start":"deny","stop":"deny"}}},"identity":{"*":{"*":"deny","get":"allow","list":"allow"},"events":{"*":"deny"}},"image":{"*":{"*":"allow"},"events":{"*":"deny"}},"log":{"actions":{"list":{"*":"deny","splitable":"deny"}}},"monitor":{"*":{"*":"allow"},"events":{"*":"deny"}}}} System-wide operation manager 全局系统管理员权限} {sys-secadmin system {"policy":{"*":{"events":{"*":"deny"}},"compute":{"*":{"*":"deny","get":"allow","list":"allow"},"disks":{"*":"deny","delete":"allow","get":"allow","list":"allow","perform":{"*":"deny","change-owner":"allow","purge":"allow"}},"dynamicschedtags":{"*":"allow"},"events":{"*":"deny"},"recyclebins":{"*":"allow","get":"allow","list":"allow"},"schedpolicies":{"*":"deny"},"schedtags":{"*":"deny"},"secgrouprules":{"*":"allow"},"secgroups":{"*":"allow"},"servers":{"*":"deny","delete":"allow","get":{"*":"allow","vnc":"deny"},"list":"allow","perform":{"*":"deny","add-secgroup":"allow","assign-admin-secgroup":"allow","assign-secgroup":"allow","change-owner":"allow","purge":"allow","revoke-admin-secgroup":"allow","revoke-secgroup":"allow","set-secgroup":"allow"}}},"identity":{"*":{"*":"allow"},"events":{"*":"deny"}},"image":{"*":{"*":"deny","delete":"allow","get":"allow","list":"allow","perform":{"*":"deny","change-owner":"allow","purge":"allow"}},"events":{"*":"deny"}},"log":{"*":{"*":"deny","get":"allow","list":"allow","perform":{"*":"deny","purge-splitable":"allow"}}},"notify":{"*":{"*":"allow"},"events":{"*":"deny"}},"yunionconf":{"*":{"*":"allow"}}}} System-wide security manager 全局安全管理员权限} {sys-adtadmin system {"policy":{"*":{"*":{"*":"deny"},"events":{"*":"allow"}},"identity":{"*":{"*":"deny"}},"log":{"*":{"*":"deny","get":"allow","list":"allow","perform":{"*":"deny","purge-splitable":"allow"}}}}} System-wide audit manager 全局审计管理员权限} {domain-opsadmin domain {"policy":{"*":{"events":{"*":"deny","list":{"*":"deny","splitable":"deny"}}},"compute":{"*":{"*":"allow"},"dynamicschedtags":{"*":"deny"},"events":{"*":"deny","list":{"*":"deny","splitable":"deny"}},"recyclebins":{"*":"deny"},"secgroups":{"*":"deny","get":"allow","list":"allow"},"servers":{"*":"allow","perform":{"*":"allow","add-secgroup":"deny","assign-admin-secgroup":"deny","assign-secgroup":"deny","change-owner":"deny","revoke-admin-secgroup":"deny","revoke-secgroup":"deny","set-secgroup":"deny","start":"deny","stop":"deny"}}},"identity":{"*":{"*":"deny","get":"allow","list":"allow"},"events":{"*":"deny"}},"image":{"*":{"*":"allow"},"events":{"*":"deny"}},"log":{"actions":{"list":{"*":"deny","splitable":"deny"}}},"monitor":{"*":{"*":"allow"},"events":{"*":"deny"}}}} Domain-wide operation manager 组织系统管理员权限} {domain-secadmin domain {"policy":{"*":{"events":{"*":"deny"}},"compute":{"*":{"*":"deny","get":"allow","list":"allow"},"disks":{"*":"deny","delete":"allow","get":"allow","list":"allow","perform":{"*":"deny","change-owner":"allow","purge":"allow"}},"dynamicschedtags":{"*":"allow"},"events":{"*":"deny"},"recyclebins":{"*":"allow","get":"allow","list":"allow"},"schedpolicies":{"*":"deny"},"schedtags":{"*":"deny"},"secgrouprules":{"*":"allow"},"secgroups":{"*":"allow"},"servers":{"*":"deny","delete":"allow","get":{"*":"allow","vnc":"deny"},"list":"allow","perform":{"*":"deny","add-secgroup":"allow","assign-admin-secgroup":"allow","assign-secgroup":"allow","change-owner":"allow","purge":"allow","revoke-admin-secgroup":"allow","revoke-secgroup":"allow","set-secgroup":"allow"}}},"identity":{"*":{"*":"allow"},"events":{"*":"deny"}},"image":{"*":{"*":"deny","delete":"allow","get":"allow","list":"allow","perform":{"*":"deny","change-owner":"allow","purge":"allow"}},"events":{"*":"deny"}},"log":{"*":{"*":"deny","get":"allow","list":"allow","perform":{"*":"deny","purge-splitable":"allow"}}},"notify":{"*":{"*":"allow"},"events":{"*":"deny"}},"yunionconf":{"*":{"*":"allow"}}}} Domain-wide security manager 组织安全管理员权限} {domain-adtadmin domain {"policy":{"*":{"*":{"*":"deny"},"events":{"*":"allow"}},"identity":{"*":{"*":"deny"}},"log":{"*":{"*":"deny","get":"allow","list":"allow","perform":{"*":"deny","purge-splitable":"allow"}}}}} Domain-wide audit manager 组织审计管理员权限} {normal-user project {"policy":{"compute":{"*":{"*":"deny","get":"allow","list":"allow"},"servers":{"*":"allow","delete":"deny","perform":{"change-bandwidth":"deny","change-config":"deny","change-disk-storage":"deny","change-ipaddr":"deny","change-owner":"deny","clone":"deny","purge":"deny","snapshot-and-clone":"deny"}}},"image":{"images":{"*":"deny","get":"allow","list":"allow"}}}} Default policy for normal user 普通用户默认权限}]
func TestGenerateAllPolicies(t *testing.T) {
	ret := make([]SPolicyData, 0)
	for i := range policyDefinitons {
		def := policyDefinitons[i]
		for _, scope := range []rbacutils.TRbacScope{
			rbacutils.ScopeSystem,
			rbacutils.ScopeDomain,
			rbacutils.ScopeProject,
		} {
			if scope.HigherEqual(def.Scope) {
				ps := generatePolicies(scope, def)
				ret = append(ret, ps...)
			}
		}
	}
	ret = append(ret, predefinedPolicyData...)
	for _, p := range ret {
		t.Logf("name %s description %s scope %s policy %s", p.Name, p.Description, p.Scope, p.Policy)
	}
	t.Logf("total: %d", len(ret))
}

func generatePolicies(scope rbacutils.TRbacScope, def sPolicyDefinition) []SPolicyData {
	level := ""
	switch scope {
	case rbacutils.ScopeSystem:
		level = "sys"
		if def.Scope == rbacutils.ScopeSystem {
			level = ""
		}
	case rbacutils.ScopeDomain:
		level = "domain"
	case rbacutils.ScopeProject:
		level = "project"
	}

	type sRoleConf struct {
		name       string
		policyFunc func(services map[string][]string) jsonutils.JSONObject
		fullName   string
		fullNameCN string
	}

	var roleConfs []sRoleConf
	if len(def.Services) > 0 {
		roleConfs = []sRoleConf{
			{
				name:       "admin",
				policyFunc: getAdminPolicy,
				fullNameCN: "管理",
				fullName:   "full",
			},
			{
				name:       "editor",
				policyFunc: getEditorPolicy,
				fullNameCN: "编辑/操作",
				fullName:   "editor/operator",
			},
			{
				name:       "viewer",
				policyFunc: getViewerPolicy,
				fullNameCN: "只读",
				fullName:   "read-only",
			},
		}
	} else {
		roleConfs = []sRoleConf{
			{
				name:       "",
				policyFunc: nil,
				fullNameCN: "",
				fullName:   "",
			},
		}
	}

	ret := make([]SPolicyData, 0)
	for _, role := range roleConfs {
		nameSegs := make([]string, 0)
		if len(level) > 0 {
			nameSegs = append(nameSegs, level)
		}
		if len(def.Name) > 0 {
			nameSegs = append(nameSegs, def.Name)
		}
		if len(role.name) > 0 {
			nameSegs = append(nameSegs, role.name)
		}
		name := strings.Join(nameSegs, "-")
		if name == "sys-admin" {
			name = "sysadmin"
		}
		var policy jsonutils.JSONObject
		if def.Services != nil {
			policy = role.policyFunc(def.Services)
		} else {
			policy = jsonutils.NewDict()
		}
		policy = addExtraPolicy(policy.(*jsonutils.JSONDict), def.Extra)
		desc := ""
		descCN := ""
		switch scope {
		case rbacutils.ScopeSystem:
			descCN += "全局"
			desc += "System-level"
		case rbacutils.ScopeDomain:
			descCN += "本域内"
			desc += "Domain-level"
		case rbacutils.ScopeProject:
			descCN += "本项目内"
			desc += "Project-level"
		}
		if len(role.fullName) > 0 {
			desc += " " + role.fullName
		}
		desc += " previlliges for"
		if len(def.Desc) > 0 {
			desc += " " + def.Desc
		}
		if len(def.DescCN) > 0 {
			descCN += def.DescCN
		}
		if len(role.fullNameCN) > 0 {
			descCN += role.fullNameCN
		}
		descCN += "权限"
		policyJson := jsonutils.NewDict()
		policyJson.Add(policy, "policy")
		ret = append(ret, SPolicyData{
			Name:   name,
			Scope:  scope,
			Policy: policyJson,

			Description:   strings.TrimSpace(desc),
			DescriptionCN: strings.TrimSpace(descCN),
		})
	}
	return ret
}

func addExtraPolicy(policy *jsonutils.JSONDict, extra map[string]map[string][]string) jsonutils.JSONObject {
	for s, resources := range extra {
		var resourcePolicy *jsonutils.JSONDict
		resPolicy, _ := policy.Get(s)
		if resPolicy != nil {
			resourcePolicy = resPolicy.(*jsonutils.JSONDict)
		} else {
			resourcePolicy = jsonutils.NewDict()
		}
		for r, actions := range resources {
			var actionPolicy *jsonutils.JSONDict
			actPolicy, _ := resourcePolicy.Get(r)
			if actPolicy != nil {
				actionPolicy = actPolicy.(*jsonutils.JSONDict)
			} else {
				actionPolicy = jsonutils.NewDict()
			}
			for i := range actions {
				actionPolicy.Add(allowResult, actions[i])
			}
			actionPolicy.Add(denyResult, "*")
			resourcePolicy.Add(actionPolicy, r)
		}
		policy.Add(resourcePolicy, s)
	}
	return policy
}

type sPolicyDefinition struct {
	Name     string
	DescCN   string
	Desc     string
	Scope    rbacutils.TRbacScope
	Services map[string][]string
	Extra    map[string]map[string][]string
}

type SRoleDefiniton struct {
	Name        string
	Description string
	Policies    []string
	Project     string
	IsPublic    bool

	DescriptionCN string
}

const (
	RoleAdmin         = "admin"
	RoleFA            = "fa"
	RoleDomainFA      = "domainfa"
	RoleProjectFA     = "projectfa"
	RoleSA            = "sa"
	RoleProjectOwner  = "project_owner"
	RoleDomainAdmin   = "domainadmin"
	RoleDomainEditor  = "domain_editor"
	RoleDomainViewer  = "domain_viewer"
	RoleProjectEditor = "project_editor"
	RoleProjectViewer = "project_viewer"

	RoleMember = "member"
)

var (
	policyDefinitons = []sPolicyDefinition{
		{
			Name:   "",
			DescCN: "任意资源",
			Desc:   "any resources",
			Scope:  rbacutils.ScopeProject,
			Services: map[string][]string{
				"*": nil,
			},
		},
		{
			Name:   "dashboard",
			DescCN: "控制面板查看相关资源",
			Desc:   "resources for viewing dashboard",
			Scope:  rbacutils.ScopeProject,
			Extra: map[string]map[string][]string{
				"compute": {
					"dashboard": {
						"get",
					},
					"capabilities": {
						"list",
					},
					"usages": {
						"list",
						"get",
					},
					"quotas": {
						"list",
						"get",
					},
					"zone_quotas": {
						"list",
						"get",
					},
					"region_quotas": {
						"list",
						"get",
					},
					"project_quotas": {
						"list",
						"get",
					},
					"domain_quotas": {
						"list",
						"get",
					},
					"infras_quotas": {
						"list",
						"get",
					},
				},
				"image": {
					"usages": {
						"list",
						"get",
					},
					"image_quotas": {
						"list",
						"get",
					},
				},
				"identity": {
					"usages": {
						"list",
						"get",
					},
					"identity_quotas": {
						"list",
						"get",
					},
					"projects": {
						"list",
					},
				},
				"meter": {
					"bill_conditions": {
						"list",
					},
				},
				"monitor": {
					"alertrecords": {
						"list",
					},
					"alertresources": {
						"list",
					},
					"unifiedmonitors": {
						"perform",
					},
					"monitorresourcealerts": {
						"list",
						"get",
					},
					"nodealerts": {
						"list",
					},
				},
				"notify": {
					"notifications": {
						"list",
						"get",
					},
					"robots": {
						"list",
						"get",
					},
					"receivers": {
						"list",
						"get",
					},
				},
				"devtool": {
					"scriptapplyrecords": {
						"list",
						"get",
					},
				},
				"yunionconf": {
					"scopedpolicybindings": {
						"list",
						"get",
					},
				},
			},
		},
		{
			Name:   "compute",
			DescCN: "计算服务(云主机与容器)相关资源",
			Desc:   "resources of computing (cloud servers and containers)",
			Scope:  rbacutils.ScopeProject,
			Services: map[string][]string{
				"compute": nil,
				"image":   nil,
				"k8s":     nil,
			},
		},
		{
			Name:   "server",
			DescCN: "云主机相关资源",
			Desc:   "resources of cloud servers",
			Scope:  rbacutils.ScopeProject,
			Services: map[string][]string{
				"compute": {
					"servers",
					"servertemplates",
					"instancegroups",
					"scalinggroups",
					"scalingactivities",
					"scalingpolicies",
					"disks",
					"networks",
					"eips",
					"snapshotpolicies",
					"snapshotpolicycaches",
					"snapshotpolicydisks",
					"snapshots",
					"instance_snapshots",
					"snapshotpolicies",
					"secgroupcaches",
					"secgrouprules",
					"secgroups",
				},
				"image": nil,
			},
			Extra: map[string]map[string][]string{
				"compute": {
					"isolated_devices": {
						"get",
						"list",
					},
				},
			},
		},
		{
			Name:   "host",
			DescCN: "宿主机和物理机相关资源",
			Desc:   "resources of hosts and baremetals",
			Scope:  rbacutils.ScopeDomain,
			Services: map[string][]string{
				"compute": {
					"hosts",
					"isolated_devices",
					"hostwires",
					"hoststorages",
					"baremetalagents",
					"baremetalnetworks",
					"baremetalevents",
				},
			},
		},
		{
			Name:   "storage",
			DescCN: "云硬盘存储相关资源",
			Desc:   "resources of cloud disk storages",
			Scope:  rbacutils.ScopeDomain,
			Services: map[string][]string{
				"compute": {
					"storages",
				},
			},
		},
		{
			Name:   "loadbalancer",
			DescCN: "负载均衡相关资源",
			Desc:   "resources of load balancers",
			Scope:  rbacutils.ScopeProject,
			Services: map[string][]string{
				"compute": {
					"loadbalanceracls",
					"loadbalanceragents",
					"loadbalancerbackendgroups",
					"loadbalancerbackends",
					"loadbalancercertificates",
					"loadbalancerclusters",
					"loadbalancerlistenerrules",
					"loadbalancerlisteners",
					"loadbalancernetworks",
					"loadbalancers",
				},
			},
			Extra: map[string]map[string][]string{
				"compute": {
					"networks": {
						"get",
						"list",
					},
				},
			},
		},
		{
			Name:   "oss",
			DescCN: "对象存储相关资源",
			Desc:   "resources of object storages",
			Scope:  rbacutils.ScopeProject,
			Services: map[string][]string{
				"compute": {
					"buckets",
				},
			},
		},
		{
			Name:   "dbinstance",
			DescCN: "关系型数据库(MySQL等)相关资源",
			Desc:   "resources of RDS",
			Scope:  rbacutils.ScopeProject,
			Services: map[string][]string{
				"compute": {
					"dbinstance_skus",
					"dbinstanceaccounts",
					"dbinstancebackups",
					"dbinstancedatabases",
					"dbinstancenetworks",
					"dbinstanceparameters",
					"dbinstanceprivileges",
					"dbinstances",
				},
			},
		},
		{
			Name:   "elasticcache",
			DescCN: "弹性缓存(Redis等)相关资源",
			Desc:   "resources of elastic caches",
			Scope:  rbacutils.ScopeProject,
			Services: map[string][]string{
				"compute": {
					"elasticcacheaccounts",
					"elasticcacheacls",
					"elasticcachebackups",
					"elasticcacheparameters",
					"elasticcaches",
					"elasticcacheskus",
				},
			},
		},
		{
			Name:   "network",
			DescCN: "网络相关资源",
			Desc:   "resources of networking",
			Scope:  rbacutils.ScopeDomain,
			Services: map[string][]string{
				"compute": {
					"vpcs",
					"wires",
					"natdentries",
					"natgateways",
					"natsentries",
					"networkinterfacenetworks",
					"networkinterfaces",
					"networks",
					"reservedips",
					"route_tables",
					"globalvpcs",
					"vpc_peering_connections",
					"eips",
					"dns_recordsets",
					"dns_trafficpolicies",
					"dns_zonecaches",
					"dns_zones",
					"dnsrecords",
				},
			},
		},
		{
			Name:   "snapshotpolicy",
			DescCN: "快照策略",
			Desc:   "snapshot policy",
			Scope:  rbacutils.ScopeDomain,
			Services: map[string][]string{
				"compute": {
					"snapshotpolicies",
					"snapshotpolicydisks",
				},
			},
		},
		{
			Name:   "secgroup",
			DescCN: "安全组",
			Desc:   "security group",
			Scope:  rbacutils.ScopeDomain,
			Services: map[string][]string{
				"compute": {
					"secgroups",
					"secgrouprules",
				},
			},
		},
		{
			Name:   "meter",
			DescCN: "计费计量分析服务相关资源",
			Desc:   "resources of metering and billing service",
			Scope:  rbacutils.ScopeProject,
			Services: map[string][]string{
				"meter":      nil,
				"suggestion": nil,
				"notify": {
					"receivers",
				},
			},
		},
		{
			Name:   "identity",
			DescCN: "身份认证(IAM)服务相关资源",
			Desc:   "resources of identity service",
			Scope:  rbacutils.ScopeDomain,
			Services: map[string][]string{
				"identity": nil,
			},
		},
		{
			Name:   "image",
			DescCN: "镜像服务相关资源",
			Desc:   "resources of image service",
			Scope:  rbacutils.ScopeProject,
			Services: map[string][]string{
				"image": nil,
			},
		},
		{
			Name:   "monitor",
			DescCN: "监控服务相关资源",
			Desc:   "resources of monitor service",
			Scope:  rbacutils.ScopeProject,
			Services: map[string][]string{
				"monitor": nil,
			},
		},
		{
			Name:   "container",
			DescCN: "容器服务相关资源",
			Desc:   "resources of container service",
			Scope:  rbacutils.ScopeProject,
			Services: map[string][]string{
				"k8s": nil,
			},
		},
		{
			Name:   "cloudid",
			DescCN: "云用户及权限管理相关资源",
			Desc:   "resources of service CloudId and IAM",
			Scope:  rbacutils.ScopeDomain,
			Services: map[string][]string{
				"compute": {
					"cloudaccounts",
					"cloudproviders",
				},
				"identity": {
					"users",
					"projects",
					"roles",
				},
				"cloudid": nil,
			},
		},
		{
			Name:   "cloudaccount",
			DescCN: "云账号管理相关资源",
			Desc:   "resources for cloud account administration",
			Scope:  rbacutils.ScopeDomain,
			Services: map[string][]string{
				"compute": {
					"cloudaccounts",
					"cloudproviderquotas",
					"cloudproviderregions",
					"cloudproviders",
				},
			},
		},
		{
			Name:   "projectresource",
			DescCN: "项目管理相关资源",
			Desc:   "resources for project administration",
			Scope:  rbacutils.ScopeDomain,
			Services: map[string][]string{
				"compute": {
					"project_quotas",
					"quotas",
					"region_quotas",
					"zone_quotas",
				},
				"image": {
					"image_quotas",
				},
				"identity": {
					"projects",
					"roles",
					"policies",
				},
			},
		},
		{
			Name:   "domainresource",
			DescCN: "域管理相关资源",
			Desc:   "resources for domain administration",
			Scope:  rbacutils.ScopeSystem,
			Services: map[string][]string{
				"compute": {
					"domain_quotas",
					"infras_quotas",
				},
				"identity": {
					"domains",
					"identity_quotas",
					"projects",
					"roles",
					"policies",
					"users",
					"groups",
				},
			},
		},
		{
			Name:   "notify",
			DescCN: "通知服务相关资源",
			Desc:   "resources of notify service",
			Scope:  rbacutils.ScopeSystem,
			Services: map[string][]string{
				"notify": nil,
			},
		},
	}

	adminPerformActions = map[string]map[string][]string{
		"compute": map[string][]string{
			"servers": []string{
				"snapshot-and-clone",
				"createdisk",
				"create-eip",
				"create-backup",
				"save-image",
				"delete-disk",
				"delete-eip",
				"delete-backup",
			},
			"buckets": []string{
				"upload",
				"delete",
			},
		},
		"k8s": map[string][]string{
			"kubeclusters": []string{
				"add-machines",
				"delete-machines",
			},
		},
	}

	RoleDefinitions = []SRoleDefiniton{
		{
			Name:          RoleAdmin,
			DescriptionCN: "系统管理员",
			Description:   "System administrator",
			Policies: []string{
				"sysadmin",
			},
			Project:  "system",
			IsPublic: false,
		},
		{
			Name:          RoleDomainAdmin,
			DescriptionCN: "域管理员",
			Description:   "Domain administrator",
			Policies: []string{
				"domain-admin",
			},
			IsPublic: true,
		},
		{
			Name:          RoleProjectOwner,
			DescriptionCN: "项目主管",
			Description:   "Project owner",
			Policies: []string{
				"project-admin",
			},
			IsPublic: true,
		},
		{
			Name:          RoleFA,
			DescriptionCN: "系统财务管理员",
			Description:   "System finance administrator",
			Policies: []string{
				"sys-meter-admin",
				"sys-dashboard",
			},
			IsPublic: false,
		},
		{
			Name:          RoleDomainFA,
			DescriptionCN: "域财务管理员",
			Description:   "Domain finance administrator",
			Policies: []string{
				"domain-meter-admin",
				"domain-dashboard",
			},
			IsPublic: true,
		},
		{
			Name:          RoleProjectFA,
			DescriptionCN: "项目财务管理员",
			Description:   "Project finance administrator",
			Policies: []string{
				"project-meter-admin",
				"project-dashboard",
			},
			IsPublic: true,
		},
		{
			Name:          RoleDomainEditor,
			DescriptionCN: "域操作员",
			Description:   "Domain operation administrator",
			Policies: []string{
				"domain-editor",
				"domain-dashboard",
			},
			IsPublic: true,
		},
		{
			Name:          RoleProjectEditor,
			DescriptionCN: "项目操作员",
			Description:   "Project operator",
			Policies: []string{
				"project-editor",
				"project-dashboard",
			},
			IsPublic: true,
		},
		{
			Name:          RoleDomainViewer,
			DescriptionCN: "域只读管理员",
			Description:   "Domain read-only administrator",
			Policies: []string{
				"domain-viewer",
				"domain-dashboard",
			},
			IsPublic: true,
		},
		{
			Name:          RoleProjectViewer,
			DescriptionCN: "项目只读成员",
			Description:   "Project read-only member",
			Policies: []string{
				"project-viewer",
				"project-dashboard",
			},
			IsPublic: true,
		},
		{
			Name:          "sys_opsadmin",
			DescriptionCN: "全局系统管理员",
			Description:   "System-wide operation manager",
			Policies: []string{
				"sys-opsadmin",
			},
			IsPublic: true,
		},
		{
			Name:          "sys_secadmin",
			DescriptionCN: "全局安全管理员",
			Description:   "System-wide security manager",
			Policies: []string{
				"sys-secadmin",
			},
			IsPublic: true,
		},
		{
			Name:          "sys_adtadmin",
			DescriptionCN: "全局审计管理员",
			Description:   "System-wide audit manager",
			Policies: []string{
				"sys-adtadmin",
			},
			IsPublic: true,
		},
		{
			Name:          "domain_opsadmin",
			DescriptionCN: "组织系统管理员",
			Description:   "Domain-wide operation manager",
			Policies: []string{
				"domain-opsadmin",
			},
			IsPublic: true,
		},
		{
			Name:          "domain_secadmin",
			DescriptionCN: "组织安全管理员",
			Description:   "Domain-wide security manager",
			Policies: []string{
				"domain-secadmin",
			},
			IsPublic: true,
		},
		{
			Name:          "domain_adtadmin",
			DescriptionCN: "组织审计管理员",
			Description:   "Domain-wide audit manager",
			Policies: []string{
				"domain-adtadmin",
			},
			IsPublic: true,
		},
		{
			Name:          "normal_user",
			DescriptionCN: "缺省普通用户角色",
			Description:   "Default normal user role",
			Policies: []string{
				"normal-user",
			},
			IsPublic: true,
		},
	}
)

//获取多服务的超级管理员权限
func getAdminPolicy(services map[string][]string) jsonutils.JSONObject {
	policy := jsonutils.NewDict()
	for k := range services {
		resList := services[k]
		if len(resList) == 0 {
			//如果服务的权限策略数组为空，就给予服务准许权限
			policy.Add(allowResult, k)
		} else {
			resPolicy := jsonutils.NewDict()
			for i := range resList {
				//如果服务的权限策略数组已存在，就更改原全部策略为准许
				resPolicy.Add(allowResult, resList[i])
			}
			policy.Add(resPolicy, k)
		}
	}
	return policy
}

//获取编辑动作权限
//结果示例：{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","create-backup":"deny","create-eip":"deny","createdisk":"deny","delete-backup":"deny","delete-disk":"deny","delete-eip":"deny","purge":"deny","save-image":"deny","snapshot-and-clone":"deny"}}
func getEditActionPolicy(service, resource string) jsonutils.JSONObject {
	p := jsonutils.NewDict()
	p.Add(denyResult, "create")
	p.Add(denyResult, "delete")
	perform := jsonutils.NewDict()
	perform.Add(denyResult, "purge")
	perform.Add(denyResult, "clone")
	if resActions, ok := adminPerformActions[service]; ok {
		if actions, ok := resActions[resource]; ok {
			for _, action := range actions {
				//服务-》资源-》动作，对指定服务下的资源下的动作全部禁止
				perform.Add(denyResult, action)
			}
		}
	}
	perform.Add(allowResult, "*")
	p.Add(perform, "perform")
	p.Add(allowResult, "*")
	return p
}

func getViewerActionPolicy() jsonutils.JSONObject {
	p := jsonutils.NewDict()
	p.Add(allowResult, "get")
	p.Add(allowResult, "list")
	p.Add(denyResult, "*")
	return p
}

//获取服务下的资源的编辑权限
//结果实例：{"compute":{"servers":{"*":"allow","create":"deny","delete":"deny","perform":{"*":"allow","clone":"deny","create-backup":"deny","create-eip":"deny","createdisk":"deny","delete-backup":"deny","delete-disk":"deny","delete-eip":"deny","purge":"deny","save-image":"deny","snapshot-and-clone":"deny"}}}}
func getEditorPolicy(services map[string][]string) jsonutils.JSONObject {
	policy := jsonutils.NewDict()
	if len(services) == 1 {
		for k := range services {
			if k == "*" {
				ns := make(map[string][]string)
				ns[k] = services[k]
				// expand adminPerformActions
				for s, resActions := range adminPerformActions {
					resList := make([]string, 0, len(resActions)+1)
					resList = append(resList, "*")
					for res := range resActions {
						resList = append(resList, res)
					}
					ns[s] = resList
				}
				services = ns
			}
		}
	}
	for k, resList := range services {
		resPolicy := jsonutils.NewDict()
		if len(resList) == 0 {
			resList = []string{"*"}
		}
		if len(resList) == 1 && resList[0] == "*" {
			if resActions, ok := adminPerformActions[k]; ok {
				for res := range resActions {
					resList = append(resList, res)
				}
			}
		}
		for i := range resList {
			resPolicy.Add(getEditActionPolicy(k, resList[i]), resList[i])
		}
		policy.Add(resPolicy, k)
	}
	return policy
}

//获取服务下的资源的只读权限
//结果示例：{"compute":{"servers":{"*":"deny","get":"allow","list":"allow"}}}
func getViewerPolicy(services map[string][]string) jsonutils.JSONObject {
	policy := jsonutils.NewDict()
	for k, resList := range services {
		if len(resList) == 0 {
			resList = []string{"*"}
		}
		resPolicy := jsonutils.NewDict()
		for i := range resList {
			resPolicy.Add(getViewerActionPolicy(), resList[i])
		}
		policy.Add(resPolicy, k)
	}
	return policy
}

var opsAdminPolicy = `
policy:
  '*':
    events:
      '*': deny
      list:
        '*': deny
        splitable: deny
  compute:
    '*':
      '*': allow
    events:
      '*': deny
      list:
        '*': deny
        splitable: deny
    dynamicschedtags:
      '*': deny
    recyclebins:
      '*': deny
    secgroups:
      '*': deny
      list: allow
      get: allow
    servers:
      '*': allow
      perform:
        '*': allow
        start: deny
        stop: deny
        change-owner: deny
        add-secgroup: deny
        set-secgroup: deny
        revoke-secgroup: deny
        revoke-admin-secgroup: deny
        assign-secgroup: deny
        assign-admin-secgroup: deny
  image:
    '*':
      '*': allow
    events:
      '*': deny
  identity:
    '*':
      '*': deny
      list: allow
      get: allow
    events:
      '*': deny
  monitor:
    events:
      '*': deny
    '*':
      '*': allow
  log:
    actions:
      list:
        '*': deny
        splitable: deny
`

var secAdminPolicy = `
policy:
  '*':
    events:
      '*': deny
  compute:
    events:
      '*': deny
    '*':
      '*': deny
      get: allow
      list: allow
    disks:
      '*': deny
      delete: allow
      get: allow
      list: allow
      perform:
        '*': deny
        change-owner: allow
        purge: allow
    dynamicschedtags:
      '*': allow
    recyclebins:
      '*': allow
      get: allow
      list: allow
    schedpolicies:
      '*': deny
    schedtags:
      '*': deny
    secgroups:
      '*': allow
    secgrouprules:
      '*': allow
    servers:
      '*': deny
      delete: allow
      get:
        '*': allow
        vnc: deny
      list: allow
      perform:
        '*': deny
        change-owner: allow
        add-secgroup: allow
        set-secgroup: allow
        revoke-secgroup: allow
        revoke-admin-secgroup: allow
        assign-secgroup: allow
        assign-admin-secgroup: allow
        purge: allow
  notify:
    events:
      '*': deny
    '*':
      '*': allow
  identity:
    events:
      '*': deny
    '*':
      '*': allow
  image:
    events:
      '*': deny
    '*':
      '*': deny
      delete: allow
      get: allow
      list: allow
      perform:
        '*': deny
        change-owner: allow
        purge: allow
  log:
    '*':
      '*': deny
      get: allow
      list: allow
      perform:
        '*': deny
        purge-splitable: allow
  yunionconf:
    '*':
      '*': allow
`

var adtAdminPolicy = `
policy:
  '*':
    '*':
      '*': deny
    events:
      '*': allow
  log:
    '*':
      '*': deny
      get: allow
      list: allow
      perform:
        '*': deny
        purge-splitable: allow
  identity:
    '*':
      '*': deny
`

var normalUserPolicy = `
policy:
  compute:
    '*':
      '*': deny
      list: allow
      get: allow
    servers:
      '*': allow
      delete: deny
      perform:
        clone: deny
        snapshot-and-clone: deny
        purge: deny
        change-ipaddr: deny
        change-bandwidth: deny
        change-config: deny
        change-owner: deny
        change-disk-storage: deny
  image:
    images:
      '*': deny
      list: allow
      get: allow
`

func toJson(yamlDef string) jsonutils.JSONObject {
	yaml, err := jsonutils.ParseYAML(yamlDef)
	if err != nil {
		log.Errorf("fail to parse %s: %s", yamlDef, err)
	}
	return yaml
}

var predefinedPolicyData = []SPolicyData{
	{
		Name:          "sys-opsadmin",
		Scope:         rbacutils.ScopeSystem,
		Policy:        toJson(opsAdminPolicy),
		Description:   "System-wide operation manager",
		DescriptionCN: "全局系统管理员权限",
	},
	{
		Name:          "sys-secadmin",
		Scope:         rbacutils.ScopeSystem,
		Policy:        toJson(secAdminPolicy),
		Description:   "System-wide security manager",
		DescriptionCN: "全局安全管理员权限",
	},
	{
		Name:          "sys-adtadmin",
		Scope:         rbacutils.ScopeSystem,
		Policy:        toJson(adtAdminPolicy),
		Description:   "System-wide audit manager",
		DescriptionCN: "全局审计管理员权限",
	},
	{
		Name:          "domain-opsadmin",
		Scope:         rbacutils.ScopeDomain,
		Policy:        toJson(opsAdminPolicy),
		Description:   "Domain-wide operation manager",
		DescriptionCN: "组织系统管理员权限",
	},
	{
		Name:          "domain-secadmin",
		Scope:         rbacutils.ScopeDomain,
		Policy:        toJson(secAdminPolicy),
		Description:   "Domain-wide security manager",
		DescriptionCN: "组织安全管理员权限",
	},
	{
		Name:          "domain-adtadmin",
		Scope:         rbacutils.ScopeDomain,
		Policy:        toJson(adtAdminPolicy),
		Description:   "Domain-wide audit manager",
		DescriptionCN: "组织审计管理员权限",
	},
	{
		Name:          "normal-user",
		Scope:         rbacutils.ScopeProject,
		Policy:        toJson(normalUserPolicy),
		Description:   "Default policy for normal user",
		DescriptionCN: "普通用户默认权限",
	},
}
