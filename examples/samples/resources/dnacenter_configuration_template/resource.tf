
terraform {
  required_providers {
    dnacenter = {
      version = "1.1.7-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_configuration_template" "example" {
  provider = dnacenter
  parameters {

    author    = "Test"
    composite = "false"
    containing_templates {
      /*
        composite = "false"
        description = "string"
        device_types {
          
          product_family = "Routers"
          #product_series = "string"
          #product_type = "string"
        }
        language = "string"
        name = "string"
        project_name = "string"
        rollback_template_params {
          
          binding = "string"
          custom_order = 1
          data_type = "string"
          default_value = "string"
          description = "string"
          display_name = "string"
          group = "string"
          id = "string"
          instruction_text = "string"
          key = "string"
          not_param = "false"
          order = 1
          param_array = "false"
          parameter_name = "string"
          provider = "string"
          range {
            
            id = "string"
            max_value = 1
            min_value = 1
          }
          required = "false"
          selection {
            
            default_selected_values = ["string"]
            id = "string"
            selection_type = "string"
            selection_values = ["string"]
          }
        }
        tags {
          
          id = "string"
          name = "string"
        }
        template_content = "string"
        template_params {
          
          binding = "string"
          custom_order = 1
          data_type = "string"
          default_value = "string"
          description = "string"
          display_name = "string"
          group = "string"
          id = "string"
          instruction_text = "string"
          key = "string"
          not_param = "false"
          order = 1
          param_array = "false"
          parameter_name = "string"
          provider = "string"
          range {
            
            id = "string"
            max_value = 1
            min_value = 1
          }
          required = "false"
          selection {
            
            default_selected_values = ["string"]
            id = "string"
            selection_type = "string"
            selection_values = ["string"]
          }
        }
        version = "string"
        */
    }
    #create_time = 1
    custom_params_order = "false"
    description         = "Test Terraform by ID Update"
    device_types {

      product_family = "Routers"
      #product_series = "string"
      #product_type = "string"
    }
    failure_policy = "CONTINUE_ON_ERROR"
    #id             = "fe2bd8b9-2cf0-4b73-b7dc-755ff0f26363"
    language = "VELOCITY"
    #last_update_time = 1
    #latest_version_time = 1
    name                      = "Saludo2"
    parent_template_id        = "fe2bd8b9-2cf0-4b73-b7dc-755ff0f26363"
    project_id                = "035c73c4-ec57-4a1a-b1d3-4cb20643ce13"
    project_name              = "Cloud Test Template 4"
    rollback_template_content = "string"
    # rollback_template_params {
    #   /*
    #     binding = "string"
    #     custom_order = 1
    #     data_type = "string"
    #     default_value = "string"
    #     description = "string"
    #     display_name = "string"
    #     group = "string"
    #     id = "string"
    #     instruction_text = "string"
    #     key = "string"
    #     not_param = "false"
    #     order = 1
    #     param_array = "false"
    #     parameter_name = "string"
    #     provider = "string"
    #     range {

    #       id = "string"
    #       max_value = 1
    #       min_value = 1
    #     }
    #     required = "false"
    #     selection {

    #       default_selected_values = ["string"]
    #       id = "string"
    #       selection_type = "string"
    #       selection_values = ["string"]
    #     }
    #     */
    # }
    software_type    = "IOS"
    software_variant = "XE"
    # template_errors = ["string"]
    #software_version = "string"
    tags {
    }
    #template_content = "string"
    # template_id = "fe2bd8b9-2cf0-4b73-b7dc-755ff0f26363"
    template_params {
      /*
        binding = "string"
        custom_order = 1
        data_type = "string"
        default_value = "string"
        description = "string"
        display_name = "string"
        group = "string"
        id = "string"
        instruction_text = "string"
        key = "string"
        not_param = "false"
        order = 1
        param_array = "false"
        parameter_name = "string"
        provider = "string"
        range {
          
          id = "string"
          max_value = 1
          min_value = 1
        }
        required = "false"
        selection {
          
          default_selected_values = ["string"]
          id = "string"
          selection_type = "string"
          selection_values = ["string"]
        }
        */
    }
    validation_errors {

      rollback_template_errors = ""
      template_errors          = ""
      # template_id              = "fe2bd8b9-2cf0-4b73-b7dc-755ff0f26363"
      template_version = null
    }
    version = "1.0"
    # comments = "Test tf"
  }
}

output "dnacenter_configuration_template_example" {
  value = dnacenter_configuration_template.example
}
