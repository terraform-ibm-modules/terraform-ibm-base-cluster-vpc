# IBM Cloud Kubernetes / Red Hat OpenShift cluster on VPC infrastructure module

[![Stable (With quality checks)](https://img.shields.io/badge/Status-Stable%20(With%20quality%20checks)-green)](https://terraform-ibm-modules.github.io/documentation/#/badge-status)
[![latest release](https://img.shields.io/github/v/release/terraform-ibm-modules/terraform-ibm-module-template?logo=GitHub&sort=semver)](https://github.com/terraform-ibm-modules/terraform-ibm-module-template/releases/latest)
[![pre-commit](https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit&logoColor=white)](https://github.com/pre-commit/pre-commit)
[![Renovate enabled](https://img.shields.io/badge/renovate-enabled-brightgreen.svg)](https://renovatebot.com/)
[![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)
[![Terraform Registry](https://img.shields.io/badge/terraform-registry-623CE4?logo=terraform)](https://registry.terraform.io/modules/terraform-ibm-modules/module-template/ibm/latest)
<!--
Add a description of modules in this repo.
Expand on the repo short description in the .github/settings.yml file.

For information, see "Module names and descriptions" at
https://terraform-ibm-modules.github.io/documentation/#/implementation-guidelines?id=module-names-and-descriptions
-->

TODO: Replace this with a description of the modules in this repo.


<!-- The following content is automatically populated by the pre-commit hook -->
<!-- BEGIN OVERVIEW HOOK -->
## Overview
<ul>
  <li><a href="#terraform-ibm-base-cluster-vpc">terraform-ibm-base-cluster-vpc</a></li>
  <li><a href="https://github.com/terraform-ibm-modules/terraform-ibm-base-cluster-vpc/tree/main/modules">Submodules</a>
    <ul>
      <li><a href="https://github.com/terraform-ibm-modules/terraform-ibm-base-cluster-vpc/tree/main/modules/worker-pool">worker-pool</a></li>
    </ul>
  </li>
  <li><a href="https://github.com/terraform-ibm-modules/terraform-ibm-base-cluster-vpc/tree/main/examples">Examples</a>
    <ul>
      <li>
        <a href="https://github.com/terraform-ibm-modules/terraform-ibm-base-cluster-vpc/tree/main/examples/kubernetes">Basic example</a>
        <a href="https://cloud.ibm.com/schematics/workspaces/create?workspace_name=base-cluster-vpc-kubernetes-example&repository=https://github.com/terraform-ibm-modules/terraform-ibm-base-cluster-vpc/tree/main/examples/kubernetes"><img src="https://img.shields.io/badge/Deploy%20with%20IBM%20Cloud%20Schematics-0f62fe?style=flat&logo=ibm&logoColor=white&labelColor=0f62fe" alt="Deploy with IBM Cloud Schematics" style="height: 16px; vertical-align: text-bottom; margin-left: 5px;"></a>
      </li>
      <li>
        <a href="https://github.com/terraform-ibm-modules/terraform-ibm-base-cluster-vpc/tree/main/examples/openshift">Basic single zone cluster with allowed outbound traffic</a>
        <a href="https://cloud.ibm.com/schematics/workspaces/create?workspace_name=base-cluster-vpc-openshift-example&repository=https://github.com/terraform-ibm-modules/terraform-ibm-base-cluster-vpc/tree/main/examples/openshift"><img src="https://img.shields.io/badge/Deploy%20with%20IBM%20Cloud%20Schematics-0f62fe?style=flat&logo=ibm&logoColor=white&labelColor=0f62fe" alt="Deploy with IBM Cloud Schematics" style="height: 16px; vertical-align: text-bottom; margin-left: 5px;"></a>
      </li>
    </ul>
    ℹ️ Ctrl/Cmd+Click or right-click on the Schematics deploy button to open in a new tab.
  </li>
  <li><a href="#known-issues">Known issues</a></li>
  <li><a href="#contributing">Contributing</a></li>
</ul>
<!-- END OVERVIEW HOOK -->


<!-- Replace this heading with the name of the root level module (the repo name) -->
## terraform-ibm-module-template

### Usage

<!--
Add an example of the use of the module in the following code block.

Use real values instead of "var.<var_name>" or other placeholder values
unless real values don't help users know what to change.
-->

```hcl
terraform {
  required_version = ">= 1.9.0"
  required_providers {
    ibm = {
      source  = "IBM-Cloud/ibm"
      version = "X.Y.Z"  # Lock into a provider version that satisfies the module constraints
    }
  }
}

locals {
    region = "us-south"
}

provider "ibm" {
  ibmcloud_api_key = "XXXXXXXXXX"  # replace with apikey value
  region           = local.region
}

module "module_template" {
  source            = "terraform-ibm-modules/<replace>/ibm"
  version           = "X.Y.Z" # Replace "X.Y.Z" with a release version to lock into a specific release
  region            = local.region
  name              = "instance-name"
  resource_group_id = "xxXXxxXXxXxXXXXxxXxxxXXXXxXXXXX" # Replace with the actual ID of resource group to use
}
```

### Required access policies

<!-- PERMISSIONS REQUIRED TO RUN MODULE
If this module requires permissions, uncomment the following block and update
the sample permissions, following the format.
Replace the 'Sample IBM Cloud' service and roles with applicable values.
The required information can usually be found in the services official
IBM Cloud documentation.
To view all available service permissions, you can go in the
console at Manage > Access (IAM) > Access groups and click into an existing group
(or create a new one) and in the 'Access' tab click 'Assign access'.
-->

<!--
You need the following permissions to run this module:

- Service
    - **Resource group only**
        - `Viewer` access on the specific resource group
    - **Sample IBM Cloud** service
        - `Editor` platform access
        - `Manager` service access
-->

<!-- NO PERMISSIONS FOR MODULE
If no permissions are required for the module, uncomment the following
statement instead the previous block.
-->

<!-- No permissions are needed to run this module.-->


<!-- The following content is automatically populated by the pre-commit hook -->
<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
### Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.9.0 |
| <a name="requirement_ibm"></a> [ibm](#requirement\_ibm) | >= 2.1.0, < 3.0.0 |
| <a name="requirement_kubernetes"></a> [kubernetes](#requirement\_kubernetes) | >= 2.16.1, < 4.0.0 |
| <a name="requirement_time"></a> [time](#requirement\_time) | >= 0.9.1, < 1.0.0 |

### Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_attach_sg_to_api_vpe"></a> [attach\_sg\_to\_api\_vpe](#module\_attach\_sg\_to\_api\_vpe) | terraform-ibm-modules/security-group/ibm | 2.10.0 |
| <a name="module_attach_sg_to_lb"></a> [attach\_sg\_to\_lb](#module\_attach\_sg\_to\_lb) | terraform-ibm-modules/security-group/ibm | 2.10.0 |
| <a name="module_attach_sg_to_master_vpe"></a> [attach\_sg\_to\_master\_vpe](#module\_attach\_sg\_to\_master\_vpe) | terraform-ibm-modules/security-group/ibm | 2.10.0 |
| <a name="module_attach_sg_to_registry_vpe"></a> [attach\_sg\_to\_registry\_vpe](#module\_attach\_sg\_to\_registry\_vpe) | terraform-ibm-modules/security-group/ibm | 2.10.0 |
| <a name="module_cbr_rule"></a> [cbr\_rule](#module\_cbr\_rule) | terraform-ibm-modules/cbr/ibm//modules/cbr-rule-module | 1.36.5 |
| <a name="module_existing_secrets_manager_instance_parser"></a> [existing\_secrets\_manager\_instance\_parser](#module\_existing\_secrets\_manager\_instance\_parser) | terraform-ibm-modules/common-utilities/ibm//modules/crn-parser | 1.6.1 |
| <a name="module_worker_pools"></a> [worker\_pools](#module\_worker\_pools) | ./modules/worker-pool | n/a |

### Resources

| Name | Type |
|------|------|
| [ibm_container_addons.addons](https://registry.terraform.io/providers/ibm-cloud/ibm/latest/docs/resources/container_addons) | resource |
| [ibm_container_ingress_instance.instance](https://registry.terraform.io/providers/ibm-cloud/ibm/latest/docs/resources/container_ingress_instance) | resource |
| [ibm_container_vpc_cluster.autoscaling_cluster](https://registry.terraform.io/providers/ibm-cloud/ibm/latest/docs/resources/container_vpc_cluster) | resource |
| [ibm_container_vpc_cluster.autoscaling_cluster_with_upgrade](https://registry.terraform.io/providers/ibm-cloud/ibm/latest/docs/resources/container_vpc_cluster) | resource |
| [ibm_container_vpc_cluster.cluster](https://registry.terraform.io/providers/ibm-cloud/ibm/latest/docs/resources/container_vpc_cluster) | resource |
| [ibm_container_vpc_cluster.cluster_with_upgrade](https://registry.terraform.io/providers/ibm-cloud/ibm/latest/docs/resources/container_vpc_cluster) | resource |
| [ibm_iam_authorization_policy.secrets_manager_iam_auth_policy](https://registry.terraform.io/providers/ibm-cloud/ibm/latest/docs/resources/iam_authorization_policy) | resource |
| [ibm_resource_tag.cluster_access_tag](https://registry.terraform.io/providers/ibm-cloud/ibm/latest/docs/resources/resource_tag) | resource |
| [kubernetes_config_map_v1_data.set_autoscaling](https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/config_map_v1_data) | resource |
| [terraform_data.config_map_status](https://registry.terraform.io/providers/hashicorp/terraform/latest/docs/resources/data) | resource |
| [terraform_data.confirm_network_healthy](https://registry.terraform.io/providers/hashicorp/terraform/latest/docs/resources/data) | resource |
| [terraform_data.install_required_binaries](https://registry.terraform.io/providers/hashicorp/terraform/latest/docs/resources/data) | resource |
| [time_sleep.wait_for_auth_policy](https://registry.terraform.io/providers/hashicorp/time/latest/docs/resources/sleep) | resource |
| [ibm_container_cluster_config.cluster_config](https://registry.terraform.io/providers/ibm-cloud/ibm/latest/docs/data-sources/container_cluster_config) | data source |
| [ibm_container_cluster_versions.cluster_versions](https://registry.terraform.io/providers/ibm-cloud/ibm/latest/docs/data-sources/container_cluster_versions) | data source |
| [ibm_is_lbs.all_lbs](https://registry.terraform.io/providers/ibm-cloud/ibm/latest/docs/data-sources/is_lbs) | data source |
| [ibm_is_virtual_endpoint_gateway.api_vpe](https://registry.terraform.io/providers/ibm-cloud/ibm/latest/docs/data-sources/is_virtual_endpoint_gateway) | data source |
| [ibm_is_virtual_endpoint_gateway.master_vpe](https://registry.terraform.io/providers/ibm-cloud/ibm/latest/docs/data-sources/is_virtual_endpoint_gateway) | data source |
| [ibm_is_virtual_endpoint_gateway.registry_vpe](https://registry.terraform.io/providers/ibm-cloud/ibm/latest/docs/data-sources/is_virtual_endpoint_gateway) | data source |

### Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_access_tags"></a> [access\_tags](#input\_access\_tags) | A list of access tags to apply to the resources created by the module, see https://cloud.ibm.com/docs/account?topic=account-access-tags-tutorial for more details | `list(string)` | `[]` | no |
| <a name="input_additional_lb_security_group_ids"></a> [additional\_lb\_security\_group\_ids](#input\_additional\_lb\_security\_group\_ids) | Additional security groups to add to the load balancers associated with the cluster. Ensure that the `number_of_lbs` is set to the number of LBs associated with the cluster. This comes in addition to the IBM maintained security group. | `list(string)` | `[]` | no |
| <a name="input_additional_vpe_security_group_ids"></a> [additional\_vpe\_security\_group\_ids](#input\_additional\_vpe\_security\_group\_ids) | Additional security groups to add to all existing load balancers. This comes in addition to the IBM maintained security group. | <pre>object({<br/>    master   = optional(list(string), [])<br/>    registry = optional(list(string), [])<br/>    api      = optional(list(string), [])<br/>  })</pre> | `{}` | no |
| <a name="input_addons"></a> [addons](#input\_addons) | Map of cluster add-on versions to install. | <pre>object({<br/>    # Common<br/>    vpc-file-csi-driver = optional(object({<br/>      version         = optional(string)<br/>      parameters_json = optional(string)<br/>    }))<br/>    static-route = optional(object({<br/>      version         = optional(string)<br/>      parameters_json = optional(string)<br/>    }))<br/>    cluster-autoscaler = optional(object({<br/>      version         = optional(string)<br/>      parameters_json = optional(string)<br/>    }))<br/>    vpc-block-csi-driver = optional(object({<br/>      version         = optional(string)<br/>      parameters_json = optional(string)<br/>    }))<br/>    ibm-storage-operator = optional(object({<br/>      version         = optional(string)<br/>      parameters_json = optional(string)<br/>    }))<br/><br/>    # OCP only<br/>    debug-tool = optional(object({<br/>      version         = optional(string)<br/>      parameters_json = optional(string)<br/>    }))<br/>    image-key-synchronizer = optional(object({<br/>      version         = optional(string)<br/>      parameters_json = optional(string)<br/>    }))<br/>    openshift-data-foundation = optional(object({<br/>      version         = optional(string)<br/>      parameters_json = optional(string)<br/>    }))<br/>    openshift-ai = optional(object({<br/>      version         = optional(string)<br/>      parameters_json = optional(string)<br/>    }))<br/><br/>    # iks only<br/>    diagnostics-and-debug-tool = optional(object({<br/>      version         = optional(string)<br/>      parameters_json = optional(string)<br/>    }))<br/>    alb-oauth-proxy = optional(object({<br/>      version         = optional(string)<br/>      parameters_json = optional(string)<br/>    }))<br/>  })</pre> | `{}` | no |
| <a name="input_allow_default_worker_pool_replacement"></a> [allow\_default\_worker\_pool\_replacement](#input\_allow\_default\_worker\_pool\_replacement) | (Advanced users) Set to true to allow the module to recreate a default worker pool. If you wish to make any change to the default worker pool which requires the re-creation of the default pool follow these [steps](https://github.com/terraform-ibm-modules/terraform-ibm-base-ocp-vpc?tab=readme-ov-file#important-considerations-for-terraform-and-default-worker-pool). | `bool` | `false` | no |
| <a name="input_attach_ibm_managed_security_group"></a> [attach\_ibm\_managed\_security\_group](#input\_attach\_ibm\_managed\_security\_group) | Specify whether to attach the IBM-defined default security group (whose name is kube-<clusterid>) to all worker nodes. Only applicable if `custom_security_group_ids` is set. | `bool` | `true` | no |
| <a name="input_cbr_rules"></a> [cbr\_rules](#input\_cbr\_rules) | The context-based restrictions rule to create. Reduce the attack surface with context-based restrictions. Only one rule is allowed. [Learn more](https://cloud.ibm.com/docs/iam?topic=iam-context-restrictions-whatis) | <pre>list(object({<br/>    description = string<br/>    account_id  = string<br/>    rule_contexts = list(object({<br/>      attributes = optional(list(object({<br/>        name  = string<br/>        value = string<br/>    }))) }))<br/>    enforcement_mode = string<br/>    tags = optional(list(object({<br/>      name  = string<br/>      value = string<br/>    })), [])<br/>    operations = optional(list(object({<br/>      api_types = list(object({<br/>        api_type_id = string<br/>      }))<br/>    })))<br/>  }))</pre> | `[]` | no |
| <a name="input_cluster_autoscaler_config"></a> [cluster\_autoscaler\_config](#input\_cluster\_autoscaler\_config) | Cluster Autoscaler configuration parameters controlling scaling behavior of worker pools (scale-up/scale-down decisions, thresholds, and timing), only explicitly provided fields are applied, and unspecified fields use IKS defaults. [Learn more](https://cloud.ibm.com/docs/containers?topic=containers-cluster-scaling-install-addon-enable#ca-configmap). | <pre>object({<br/>    coresTotal                   = optional(string)<br/>    expander                     = optional(string)<br/>    expendablePodsPriorityCutoff = optional(number)<br/>    ignoreDaemonsetsUtilization  = optional(bool)<br/>    imagePullPolicy              = optional(string)<br/><br/>    livenessProbeFailureThreshold = optional(number)<br/>    livenessProbePeriodSeconds    = optional(number)<br/>    livenessProbeTimeoutSeconds   = optional(number)<br/><br/>    logLevel = optional(string)<br/><br/>    maxBulkSoftTaintCount     = optional(number)<br/>    maxBulkSoftTaintTime      = optional(string)<br/>    maxFailingTime            = optional(string)<br/>    maxGracefulTerminationSec = optional(number)<br/>    maxInactivity             = optional(string)<br/>    maxNodeProvisionTime      = optional(string)<br/>    maxRetryGap               = optional(number)<br/>    maxTotalUnreadyPercentage = optional(number)<br/><br/>    memoryTotal         = optional(string)<br/>    minReplicaCount     = optional(number)<br/>    newPodScaleUpDelay  = optional(string)<br/>    okTotalUnreadyCount = optional(number)<br/><br/>    resourcesLimitsCPU      = optional(string)<br/>    resourcesLimitsMemory   = optional(string)<br/>    resourcesRequestsCPU    = optional(string)<br/>    resourcesRequestsMemory = optional(string)<br/><br/>    retryAttempts = optional(number)<br/><br/>    scaleDownCandidatesPoolMinCount  = optional(number)<br/>    scaleDownCandidatesPoolRatio     = optional(number)<br/>    scaleDownDelayAfterAdd           = optional(string)<br/>    scaleDownDelayAfterDelete        = optional(string)<br/>    scaleDownDelayAfterFailure       = optional(string)<br/>    scaleDownEnabled                 = optional(bool)<br/>    scaleDownGPUUtilizationThreshold = optional(number)<br/>    scaleDownNonEmptyCandidatesCount = optional(number)<br/>    scaleDownUnneededTime            = optional(string)<br/>    scaleDownUnreadyTime             = optional(string)<br/>    scaleDownUtilizationThreshold    = optional(number)<br/><br/>    scanInterval = optional(string)<br/><br/>    skipNodesWithLocalStorage = optional(bool)<br/>    skipNodesWithSystemPods   = optional(bool)<br/><br/>    unremovableNodeRecheckTimeout = optional(string)<br/><br/>    # extended fields<br/>    maxNodeGroupBinpackingDuration = optional(string)<br/>    maxNodesPerScaleUp             = optional(number)<br/>    parallelDrain                  = optional(bool)<br/>    maxScaleDownParallelism        = optional(number)<br/>    maxDrainParallelism            = optional(number)<br/>    nodeDeletionBatcherInterval    = optional(string)<br/>    nodeDeleteDelayAfterTaint      = optional(string)<br/>    enforceNodeGroupMinSize        = optional(bool)<br/>    kubeClientBurst                = optional(number)<br/>    kubeClientQPS                  = optional(number)<br/>    scaleDownUnreadyEnabled        = optional(bool)<br/>    maxPodEvictionTime             = optional(string)<br/>    balancingIgnoreLabel           = optional(string)<br/>    OSReservedMemoryGi             = optional(number)<br/>    OSReservedCPUMili              = optional(number)<br/>  })</pre> | `{}` | no |
| <a name="input_cluster_config_endpoint_type"></a> [cluster\_config\_endpoint\_type](#input\_cluster\_config\_endpoint\_type) | Specify which type of endpoint to use for cluster config access: 'default', 'private', 'vpe', 'link'. A 'default' value uses the default endpoint of the cluster. | `string` | `"default"` | no |
| <a name="input_cluster_create_timeout"></a> [cluster\_create\_timeout](#input\_cluster\_create\_timeout) | Timeout duration for cluster creation operations. Specify a duration string (e.g., '3h', '45m'). | `string` | `"3h"` | no |
| <a name="input_cluster_delete_timeout"></a> [cluster\_delete\_timeout](#input\_cluster\_delete\_timeout) | Timeout duration for cluster deletion operations. Specify a duration string (e.g., '2h', '30m'). | `string` | `"2h"` | no |
| <a name="input_cluster_name"></a> [cluster\_name](#input\_cluster\_name) | The name that is assigned to the provisioned cluster. | `string` | n/a | yes |
| <a name="input_cluster_ready_when"></a> [cluster\_ready\_when](#input\_cluster\_ready\_when) | The cluster is ready based on one of the following:: MasterNodeReady (not recommended), OneWorkerNodeReady, Normal, IngressReady | `string` | `"IngressReady"` | no |
| <a name="input_cluster_type"></a> [cluster\_type](#input\_cluster\_type) | The type of cluster to provision. | `string` | `"kubernetes"` | no |
| <a name="input_cluster_update_timeout"></a> [cluster\_update\_timeout](#input\_cluster\_update\_timeout) | Timeout duration for cluster update operations. Specify a duration string (e.g., '3h', '1h30m'). | `string` | `"3h"` | no |
| <a name="input_cluster_version"></a> [cluster\_version](#input\_cluster\_version) | The version of the cluster to provision.<br/><br/>  For Kubernetes clusters, specify a Kubernetes version (for example, 1.33).<br/>  For OpenShift clusters, specify an OpenShift version (for example, 4.19).<br/><br/>  If no value is specified, or if set to "default", the current default version<br/>  for the selected cluster type is used.<br/><br/>  This input is used only during initial cluster provisioning and is ignored for updates. | `string` | `null` | no |
| <a name="input_cos_instance_crn"></a> [cos\_instance\_crn](#input\_cos\_instance\_crn) | crn of the COS instance to provision for OpenShift internal registry storage. | `string` | `null` | no |
| <a name="input_custom_security_group_ids"></a> [custom\_security\_group\_ids](#input\_custom\_security\_group\_ids) | Security groups to add to all worker nodes. This comes in addition to the IBM maintained security group if `attach_ibm_managed_security_group` is set to true. If this variable is set, the default VPC security group is NOT assigned to the worker nodes. | `list(string)` | `null` | no |
| <a name="input_disable_outbound_traffic_protection"></a> [disable\_outbound\_traffic\_protection](#input\_disable\_outbound\_traffic\_protection) | Whether to allow public outbound access from the cluster workers. This is applicable for OCP 4.15 and later. For iks set per your environment's security requirements. | `bool` | `false` | no |
| <a name="input_disable_public_endpoint"></a> [disable\_public\_endpoint](#input\_disable\_public\_endpoint) | Whether access to the public service endpoint is disabled when the cluster is created. Does not affect existing clusters. You can't disable a public endpoint on an existing cluster, so you can't convert a public cluster to a private cluster. To change a public endpoint to private, create another cluster with this input set to `true`. | `bool` | `false` | no |
| <a name="input_enable_cluster_version_upgrade"></a> [enable\_cluster\_version\_upgrade](#input\_enable\_cluster\_version\_upgrade) | When set to true, allows Terraform to manage major cluster version upgrades. This is intended for advanced users who manually control major version upgrades. Defaults to false to avoid unintended drift from IBM-managed patch updates. NOTE: Enabling this on existing clusters requires a one-time terraform state migration. | `bool` | `false` | no |
| <a name="input_enable_ocp_console"></a> [enable\_ocp\_console](#input\_enable\_ocp\_console) | Flag to specify whether to enable or disable the OpenShift console. If set to `null` the module does not modify the current setting on the cluster. Keep in mind that when this input is set to `true` or `false` on a cluster with private only endpoint enabled, the runtime must be able to access the private endpoint. | `bool` | `null` | no |
| <a name="input_enable_secrets_manager_integration"></a> [enable\_secrets\_manager\_integration](#input\_enable\_secrets\_manager\_integration) | Integrate with IBM Cloud Secrets Manager so you can centrally manage Ingress subdomain certificates and other secrets. [Learn more](https://cloud.ibm.com/docs/containers?topic=containers-secrets-mgr) | `bool` | `false` | no |
| <a name="input_existing_secrets_manager_instance_crn"></a> [existing\_secrets\_manager\_instance\_crn](#input\_existing\_secrets\_manager\_instance\_crn) | CRN of the Secrets Manager instance where Ingress certificate secrets are stored. If 'enable\_secrets\_manager\_integration' is set to true then this value is required. | `string` | `null` | no |
| <a name="input_force_delete_storage"></a> [force\_delete\_storage](#input\_force\_delete\_storage) | Flag indicating whether or not to delete attached storage when destroying the cluster - Default: false | `bool` | `false` | no |
| <a name="input_ignore_worker_pool_size_changes"></a> [ignore\_worker\_pool\_size\_changes](#input\_ignore\_worker\_pool\_size\_changes) | Enable if using worker autoscaling. Stops Terraform managing worker count | `bool` | `false` | no |
| <a name="input_image_security_enforcement"></a> [image\_security\_enforcement](#input\_image\_security\_enforcement) | Set to true to enable image security enforcement policies in a cluster. When you enable image security enforcement in your cluster, you install the open-source Portieris Kubernetes project. Then, you can create image policies to prevent pods that don't meet the policies, such as unsigned images, from running in your cluster. For more information, see the Portieris documentation (https://github.com/IBM/portieris). | `bool` | `false` | no |
| <a name="input_install_required_binaries"></a> [install\_required\_binaries](#input\_install\_required\_binaries) | When set to true, a script will run to check if `kubectl` and `jq` exist on the runtime and if not attempt to download them from the public internet and install them to /tmp. Set to false to skip running this script. | `bool` | `true` | no |
| <a name="input_kms_config"></a> [kms\_config](#input\_kms\_config) | Use to attach a Key Protect or Hyper Protect Crypto Service instance to the cluster. If account\_id is not provided, the current account is used. [Learn more](https://cloud.ibm.com/docs/key-protect?topic=key-protect-provision) | <pre>object({<br/>    crk_id           = string               # ID of the customer root key<br/>    instance_id      = string               # GUID of the KMS instance<br/>    private_endpoint = optional(bool, true) # Defaults to true to configure the KMS private service endpoint<br/>    account_id       = optional(string)     # To attach KMS instance from another account<br/>    wait_for_apply   = optional(bool, true) # Defaults to true so terraform will wait until the KMS is applied to the master, ready and deployed<br/>  })</pre> | `null` | no |
| <a name="input_manage_all_addons"></a> [manage\_all\_addons](#input\_manage\_all\_addons) | Instructs Terraform to manage all cluster addons, even if addons were installed outside of the module. If set to 'true' this module destroys any addons that were installed by other sources. | `bool` | `false` | no |
| <a name="input_network_plugin"></a> [network\_plugin](#input\_network\_plugin) | The Container Network Interface (CNI) plugin for the cluster. Requires OpenShift >= 4.20. Supported values are Calico (default) and OVNKubernetes. | `string` | `"Calico"` | no |
| <a name="input_number_of_lbs"></a> [number\_of\_lbs](#input\_number\_of\_lbs) | The number of LBs to associated the `additional_lb_security_group_names` security group with. | `number` | `1` | no |
| <a name="input_ocp_entitlement"></a> [ocp\_entitlement](#input\_ocp\_entitlement) | Value that is applied to the entitlements for OCP cluster provisioning | `string` | `null` | no |
| <a name="input_pod_subnet_cidr"></a> [pod\_subnet\_cidr](#input\_pod\_subnet\_cidr) | Specify a custom subnet CIDR to provide private IP addresses for pods. The subnet must have a CIDR of at least `/23` or larger. Default value is `172.30.0.0/16` when the variable is set to `null`. | `string` | `null` | no |
| <a name="input_region"></a> [region](#input\_region) | The IBM Cloud region where the cluster is provisioned. | `string` | n/a | yes |
| <a name="input_resource_group_id"></a> [resource\_group\_id](#input\_resource\_group\_id) | The ID of an existing IBM Cloud resource group where the cluster is grouped. | `string` | n/a | yes |
| <a name="input_secrets_manager_secret_group_id"></a> [secrets\_manager\_secret\_group\_id](#input\_secrets\_manager\_secret\_group\_id) | Secret group ID where Ingress secrets are stored in the Secrets Manager instance. | `string` | `null` | no |
| <a name="input_service_subnet_cidr"></a> [service\_subnet\_cidr](#input\_service\_subnet\_cidr) | Specify a custom subnet CIDR to provide private IP addresses for services. The subnet must be at least `/24` or larger. Default value is `172.21.0.0/16` when the variable is set to `null`. | `string` | `null` | no |
| <a name="input_skip_secrets_manager_iam_auth_policy"></a> [skip\_secrets\_manager\_iam\_auth\_policy](#input\_skip\_secrets\_manager\_iam\_auth\_policy) | To skip creating auth policy that allows OCP cluster 'Manager' role access in the existing Secrets Manager instance for managing ingress certificates. | `bool` | `false` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | Metadata labels describing this cluster deployment, i.e. test | `list(string)` | `[]` | no |
| <a name="input_verify_worker_network_readiness"></a> [verify\_worker\_network\_readiness](#input\_verify\_worker\_network\_readiness) | By setting this to true, a script runs kubectl commands to verify that all worker nodes can communicate successfully with the master. If the runtime does not have access to the kube cluster to run kubectl commands, set this value to false. | `bool` | `true` | no |
| <a name="input_vpc_id"></a> [vpc\_id](#input\_vpc\_id) | ID of the VPC instance where this cluster is provisioned. | `string` | n/a | yes |
| <a name="input_vpc_subnets"></a> [vpc\_subnets](#input\_vpc\_subnets) | Metadata that describes the VPC's subnets. Obtain this information from the VPC where this cluster is created. | <pre>map(list(object({<br/>    id         = string<br/>    zone       = string<br/>    cidr_block = string<br/>  })))</pre> | n/a | yes |
| <a name="input_worker_pools"></a> [worker\_pools](#input\_worker\_pools) | List of worker pools | <pre>list(object({<br/>    subnet_prefix = optional(string)<br/>    vpc_subnets = optional(list(object({<br/>      id         = string<br/>      zone       = string<br/>      cidr_block = string<br/>    })))<br/>    pool_name         = string<br/>    machine_type      = string<br/>    workers_per_zone  = number<br/>    resource_group_id = optional(string)<br/>    operating_system  = string<br/>    labels            = optional(map(string))<br/>    minSize           = optional(number)<br/>    secondary_storage = optional(string)<br/>    maxSize           = optional(number)<br/>    enableAutoscaling = optional(bool)<br/>    boot_volume_encryption_kms_config = optional(object({<br/>      crk             = string<br/>      kms_instance_id = string<br/>      kms_account_id  = optional(string)<br/>    }))<br/>    additional_security_group_ids = optional(list(string))<br/>  }))</pre> | n/a | yes |
| <a name="input_worker_pools_taints"></a> [worker\_pools\_taints](#input\_worker\_pools\_taints) | Optional, Map of lists containing node taints by node-pool name | `map(list(object({ key = string, value = string, effect = string })))` | `null` | no |

### Outputs

| Name | Description |
|------|-------------|
| <a name="output_api_vpe"></a> [api\_vpe](#output\_api\_vpe) | Info about the api VPE, if it exists. For more info about schema, see https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/is_virtual_endpoint_gateway |
| <a name="output_cluster_crn"></a> [cluster\_crn](#output\_cluster\_crn) | CRN of the cluster |
| <a name="output_cluster_id"></a> [cluster\_id](#output\_cluster\_id) | ID of the cluster |
| <a name="output_cluster_name"></a> [cluster\_name](#output\_cluster\_name) | Name of the cluster |
| <a name="output_cluster_version"></a> [cluster\_version](#output\_cluster\_version) | Version of the cluster |
| <a name="output_cos_crn"></a> [cos\_crn](#output\_cos\_crn) | CRN of the COS instance |
| <a name="output_ingress_hostname"></a> [ingress\_hostname](#output\_ingress\_hostname) | The hostname that was assigned to your Ingress subdomain. |
| <a name="output_kms_config"></a> [kms\_config](#output\_kms\_config) | KMS configuration details |
| <a name="output_kubeconfig_path"></a> [kubeconfig\_path](#output\_kubeconfig\_path) | Path to the administrator kubeconfig file generated for the cluster. |
| <a name="output_master_status"></a> [master\_status](#output\_master\_status) | The status of the Kubernetes master. |
| <a name="output_master_url"></a> [master\_url](#output\_master\_url) | The URL of the Kubernetes master. |
| <a name="output_master_vpe"></a> [master\_vpe](#output\_master\_vpe) | Info about the master, or default, VPE. For more info about schema, see https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/is_virtual_endpoint_gateway |
| <a name="output_operating_system"></a> [operating\_system](#output\_operating\_system) | The operating system of the workers in the default worker pool. |
| <a name="output_private_service_endpoint_url"></a> [private\_service\_endpoint\_url](#output\_private\_service\_endpoint\_url) | Private service endpoint URL |
| <a name="output_public_service_endpoint_url"></a> [public\_service\_endpoint\_url](#output\_public\_service\_endpoint\_url) | Public service endpoint URL |
| <a name="output_region"></a> [region](#output\_region) | Region that the cluster is deployed to |
| <a name="output_registry_vpe"></a> [registry\_vpe](#output\_registry\_vpe) | Info about the registry VPE, if it exists. For more info about schema, see https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/is_virtual_endpoint_gateway |
| <a name="output_resource_group_id"></a> [resource\_group\_id](#output\_resource\_group\_id) | Resource group ID the cluster is deployed in |
| <a name="output_secrets_manager_integration_config"></a> [secrets\_manager\_integration\_config](#output\_secrets\_manager\_integration\_config) | Information about the Secrets Manager instance that is used to store the Ingress certificates. |
| <a name="output_vpc_id"></a> [vpc\_id](#output\_vpc\_id) | ID of the clusters VPC |
| <a name="output_vpe_url"></a> [vpe\_url](#output\_vpe\_url) | The virtual private endpoint URL of the Kubernetes cluster. |
| <a name="output_workerpools"></a> [workerpools](#output\_workerpools) | Worker pools created |
<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->

## Known issues

<!-- Update this if any known issues or limitations -->
There are currently no known issues or limitations at this time.

<!-- Leave this section as is so that your module has a link to local development environment set-up steps for contributors to follow -->
## Contributing

You can report issues and request features for this module in GitHub issues in the module repo. See [Report an issue or request a feature](https://github.com/terraform-ibm-modules/.github/blob/main/.github/SUPPORT.md).

To set up your local development environment, see [Local development setup](https://terraform-ibm-modules.github.io/documentation/#/local-dev-setup) in the project documentation.
