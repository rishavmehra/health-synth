terraform {
  required_providers {
    kestra = {
      source  = "kestra-io/kestra" 
      version = "~> 0.18.2"
    }
  }
}

provider "kestra" {
  url = "http://localhost:8080"
}

resource "kestra_flow" "flows" {
  for_each  = fileset(path.module, "flows/*/*.yml")
  flow_id   = yamldecode(templatefile(each.value, {}))["id"]
  namespace = yamldecode(templatefile(each.value, {}))["namespace"]
  content   = templatefile(each.value, {})
}

resource "kestra_namespace_file" "data_gen" {
  for_each = fileset(path.module, "code/data_gen.go")
  namespace = "healthsynth.hospital"
  filename = "/${each.value}"
  content = file(each.value)
}

resource "kestra_namespace_file" "records_gen" {
  for_each = fileset(path.module, "code/data_gen.go")
  namespace = "healthsynth.records"
  filename = "/${each.value}"
  content = file(each.value)
}


