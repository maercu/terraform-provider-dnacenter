terraform {
  required_providers {
    dnacenter = {
      versions = ["0.2"]
      source   = "hashicorp.com/edu/dnacenter"
    }
  }
}

provider "dnacenter" {
}

resource "dna_template_project" "project_1" {
  provider = dnacenter
  item {
    is_deletable = true
    name         = "Cloud Test Template 2"
  }
}
output "dna_template_project_1" {
  value = dna_template_project.project_1
}
output "dna_template_project_1_id" {
  value = dna_template_project.project_1.item.0.id
}

resource "dna_template" "template_1" {
  provider   = dnacenter
  depends_on = [dna_template_project.project_1]
  item {
    project_id = dna_template_project.project_1.item.0.id
    name       = "DMVPN Spoke for Branch Router - System Default for Test Project"
    device_types {
      product_family = "Routers"
    }
    software_type = "IOS-XE"
  }
}

output "dna_template_1" {
  value = dna_template.template_1
}
