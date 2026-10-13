variable "project_name" {
  description = "Nazwa projektu (prefiks zasobów)"
  type        = string
  default     = "devops-projekt"
}

variable "region" {
  description = "Region AWS"
  type        = string
  default     = "eu-central-1" # Frankfurt
}

variable "instance_type" {
  description = "Typ instancji EC2 (t3.small wystarcza na app + monitoring)"
  type        = string
  default     = "t3.small"
}

variable "ssh_public_key_path" {
  description = "Ścieżka do klucza publicznego SSH"
  type        = string
  default     = "~/.ssh/devops-projekt.pub"
}

variable "admin_cidr" {
  description = "Twój adres IP w notacji CIDR (np. 1.2.3.4/32) – dostęp do SSH, Grafany i Prometheusa"
  type        = string
}
