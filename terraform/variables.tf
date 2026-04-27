variable "location" {
  description = "Azure region where all resources will be deployed."
  type        = string
  default     = "spaincentral"
}

variable "resource_group_name" {
  description = "Name of the Azure Resource Group."
  type        = string
  default     = "devops-lecture-rg"
}

variable "cluster_name" {
  description = "Name of the AKS cluster."
  type        = string
  default     = "devops-lecture-aks"
}

variable "node_count" {
  description = "Number of nodes in the default node pool."
  type        = number
  default     = 1
}

variable "vm_size" {
  description = "VM size for AKS nodes. Standard_D2s_v3: 2 vCPU, 8 GB RAM — smallest available in spaincentral with student subscription quota."
  type        = string
  default     = "Standard_D2s_v3"
}
