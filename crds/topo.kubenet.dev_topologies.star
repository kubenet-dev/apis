
#Joins a list of items with a delimiter, filtering out any None values
def join(delimiter, items):
  return delimiter.join(items)

#Returns the value from dict or defaults based on key
def get_value_or_default(dict, defaults, key):
  return dict.get(key, defaults.get(key, None))

# Constructs a node resource from topology node data.
def build_node(self, topo_node):
  defaults = self.get("spec", {}).get("defaults", "")
  partition = self.get("metadata", {}).get("name", "")
  region = get_value_or_default(topo_node, defaults, "region")
  site = get_value_or_default(topo_node, defaults, "site")
  node_name = topo_node.get("name", "")

  name_parts = [partition, region, site, node_name]

  node = {
    "apiVersion": "infra.kuid.dev/v1alpha1",
    "kind": "Node",
    "metadata": {
      "namespace": self.get("metadata", {}).get("namespace", ""),
      "name": join(".", name_parts) 
    },
    "spec": {
      "partition": partition,
      "region": region,
      "site": site,
      "node": node_name,
      "provider": get_value_or_default(topo_node, defaults, "provider"),
      "platformType": get_value_or_default(topo_node, defaults, "platformType"),
    }
  }
  optional_fields = ["rack", "position", "location"]
  for field in optional_fields:
    value = get_value_or_default(topo_node, defaults, field)
    if value != None:
      node["spec"][field] = value
  
  return node

# Constructs a link resource from topology link data.
def build_link(self, topo_link, nodes):
  link_parts = []
  endpoints = []
  index = 0
  for endpoint in topo_link.get("endpoints", []):
    node_name = endpoint.get("node", "")
    node = nodes.get(node_name, {})
    if not node:
      return None, "node name" + node_name + "not present in link endpoint"
    node_spec = node.get("spec", {})
    partition = node_spec.get("partition", "")
    region = node_spec.get("region", "")
    site = node_spec.get("site", "")

    port = endpoint.get("port", 0)
    ep_id = endpoint.get("endpoint", 0)
    endpoints.append(
      {
        "partition": partition,
        "region": region,
        "site": site,
        "node": node_name,
        "port": int(port),
        "endpoint": int(ep_id),
        "adaptor": endpoint.get("adaptor", ""),
        #"module": endpoint.get("module", 0),
        #"moduleBay": endpoint.get("moduleBay", 0),
      }
    )

    if index == 0:
      link_parts.append(partition)
      link_parts.append(region)
      link_parts.append(site)   

    link_parts.append(join(".", [node_name, str(port), str(ep_id)]) )
    index += 1

  link = {
    "apiVersion": "infra.kuid.dev/v1alpha1",
    "kind": "Link",
    "metadata": {
      "namespace": self.get("metadata", {}).get("namespace", ""),
      "name": join(".", link_parts),
    },
    "spec": {
      #"internal": True,
      "endpoints": endpoints,
    }
  }
  return link, None

