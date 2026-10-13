terraform {
  required_version = ">= 1.6"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
    local = {
      source  = "hashicorp/local"
      version = "~> 2.5"
    }
  }

  # Opcjonalnie: zdalny stan w S3 (odkomentuj po utworzeniu bucketu)
  # backend "s3" {
  #   bucket = "moj-bucket-tfstate"
  #   key    = "devops-projekt/terraform.tfstate"
  #   region = "eu-central-1"
  # }
}

provider "aws" {
  region = var.region

  default_tags {
    tags = {
      Project   = var.project_name
      ManagedBy = "terraform"
    }
  }
}
