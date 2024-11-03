def get_node_templates(name, namespace):
  r = get_resource("device.network.kubenet.dev/v1alpha1", "NodeTemplate")
  rsp = client_get(name, namespace, r["resource"])
  if rsp["error"] != None:
    return None, rsp["error"]
  return rsp["resource"], None

def build_port(self, portID):
  print("build_port portID", portID)
  # self = node
  """Constructs a node resource from topology node data."""
  node = {
    "apiVersion": "infra.kuid.dev/v1alpha1",
    "kind": "Port",
    "metadata": {
      "name": ".".join([self.get("metadata", {}).get("name", ""), str(portID)]),
      "namespace": self.get("metadata", {}).get("namespace", ""),
    },
    "spec": {
      "partition": self.get("spec", {}).get("partition", ""),
      "region": self.get("spec", {}).get("region", ""),
      "site": self.get("spec", {}).get("partition", ""),
      "node": self.get("spec", {}).get("node", ""),
      "port": int(portID),
    }
  }
  return node

def build_adaptor(self, portID, adaptor):
  # self = node
  """Constructs a node resource from topology node data."""
  adaptor_name = adaptor.get("name", "")
  print("build_port portID", portID)
  print("build_port adaptor_name", adaptor_name)

  node = {
    "apiVersion": "infra.kuid.dev/v1alpha1",
    "kind": "Adaptor",
    "metadata": {
      "name": ".".join([self.get("metadata", {}).get("name", ""), str(portID), adaptor_name]),
      "namespace": self.get("metadata", {}).get("namespace", ""),
    },
    "spec": {
      "partition": self.get("spec", {}).get("partition", ""),
      "region": self.get("spec", {}).get("region", ""),
      "site": self.get("spec", {}).get("partition", ""),
      "node": self.get("spec", {}).get("node", ""),
      "port": int(portID),
      "adaptor": adaptor_name,
    }
  }
  return node

def build_endpoint(self, portID, adaptor, connectorID):
  # self = node
  """Constructs a node resource from topology node data."""
  adaptor_name = adaptor.get("name", "")
  print("build_port portID", portID)
  print("build_port adaptor_name", adaptor_name)
  print("build_port connectorID", connectorID)

  node = {
    "apiVersion": "infra.kuid.dev/v1alpha1",
    "kind": "Endpoint",
    "metadata": {
      "name": ".".join([self.get("metadata", {}).get("name", ""), str(portID), adaptor_name, str(connectorID)]),
      "namespace": self.get("metadata", {}).get("namespace", ""),
    },
    "spec": {
      "partition": self.get("spec", {}).get("partition", ""),
      "region": self.get("spec", {}).get("region", ""),
      "site": self.get("spec", {}).get("partition", ""),
      "node": self.get("spec", {}).get("node", ""),
      "port": int(portID),
      "adaptor": adaptor_name,
      "endpoint": int(connectorID),
      "speed": adaptor.get("speed", "",)
    }
  }
  return node