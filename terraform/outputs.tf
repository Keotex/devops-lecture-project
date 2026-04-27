output "resource_group_name" {
  value = azurerm_resource_group.main.name
}

output "cluster_name" {
  value = azurerm_kubernetes_cluster.main.name
}

output "kube_config" {
  description = "Raw kubeconfig — pipe into a file or use with KUBECONFIG env var."
  value       = azurerm_kubernetes_cluster.main.kube_config_raw
  sensitive   = true
}
