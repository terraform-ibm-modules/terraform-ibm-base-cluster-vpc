##############################################################################
# Outputs
##############################################################################
output "cluster_id" {
  description = "ID of the cluster"
  value       = local.cluster_id
  depends_on  = [null_resource.confirm_network_healthy]
}
