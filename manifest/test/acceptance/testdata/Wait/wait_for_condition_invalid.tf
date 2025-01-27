resource "kubernetes_manifest" "test" {
  manifest = {
    apiVersion = "v1"
    kind       = "Namespace"

    metadata = {
      name      = var.name
    }
  }

  wait {
    condition {
      type = "Ready"
      status = "True"
    }
  }

  timeouts {
    create = "3s"
  }
}
