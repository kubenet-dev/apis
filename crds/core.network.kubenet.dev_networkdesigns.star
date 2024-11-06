def get_network_design(name, namespace):
  r = get_resource("core.network.kubenet.dev/v1alpha1", "NetworkDesign")
  rsp = client_get(name, namespace, r["resource"])
  if rsp["error"] != None:
    return None, rsp["error"]
  return rsp["resource"], None

def get_ipindex(self):
  namespace = self.get("metadata", {}).get("namespace", "")
  partition = self.get("metadata", {}).get("name", "")

  prefixes = []
  interface_loopback = self.get("spec", {}).get("interfaces", {}).get("loopback", {})
  for prefix in interface_loopback.get("prefixes", []):
    prefixes.append({
      "prefixType": "pool",
      "prefix": prefix.get("prefix", {}),
      "labels": {"infra.kuid.dev/purpose": "loopback"},
    })

  interface_underlay = self.get("spec", {}).get("interfaces", {}).get("underlay", {})
  for prefix in interface_underlay.get("prefixes", []):
    prefixes.append({
      "prefix": prefix.get("prefix", ""),
      "labels": {"infra.kuid.dev/purpose": "underlay"},
    })

  return {
    "apiVersion": "ipam.be.kuid.dev/v1alpha1",
    "kind": "IPIndex",
    "metadata": {
      "namespace": namespace,
      "name": ".".join([partition, "default"]),
    },
    "spec": {
      "prefixes": prefixes,
    },
  }

def get_asindex(self):
  return {
    "apiVersion": "as.be.kuid.dev/v1alpha1",
    "kind": "ASIndex",
    "metadata": {
      "namespace": self.get("metadata", {}).get("namespace", ""),
      "name": self.get("metadata", {}).get("name", ""),
    },
    "spec": {},
  }

def get_genidindex(self):
  return {
    "apiVersion": "genid.be.kuid.dev/v1alpha1",
    "kind": "GENIDIndex",
    "metadata": {
      "namespace": self.get("metadata", {}).get("namespace", ""),
      "name": self.get("metadata", {}).get("name", ""),
    },
    "spec": {
      "minID": 0,
      "maxID": 16383,
      "type": "16bit",
    },
  }

def get_node_ipclaims(self, node):
  # self is netowrk design
  partition = self.get("metadata", {}).get("name", "")
  namespace = self.get("metadata", {}).get("namespace", "")
  node_name = node.get("metadata", {}).get("name", "")
  index = ".".join([partition, "default"])
  ip_claims = []
  interface_loopback = self.get("spec", {}).get("interfaces", {}).get("loopback", {})
  af = "ipv4"
  if interface_loopback.get("addressing") == "dualstack" or interface_loopback.get("addressing") == "ipv4numbered":
      ip_claims.append(get_ipclaim_pool(".".join([node_name, af]), namespace, index,af))
  else:
      ip_claims.append(get_ipclaim_pool(".".join([node_name, "routerid"]), namespace, index, af))
  af = "ipv6"
  if interface_loopback.get("addressing") == "dualstack" or interface_loopback.get("addressing") == "ipv6numbered":
      ip_claims.append(get_ipclaim_pool(".".join([node_name, af]), namespace, index, af))
  return ip_claims


def get_ipclaim_pool(name, namespace, index, af):
  return {
    "apiVersion": "ipam.be.kuid.dev/v1alpha1",
    "kind": "IPClaim",
    "metadata": {
      "namespace": namespace,
      "name": name,
    },
    "spec": {
      "index": index,
      "prefixType": "pool",
      "selector": {
        "matchLabels": {
          "infra.kuid.dev/purpose": "loopback",
          "ipam.be.kuid.dev/address-family": af,
        },
      },
    },
  }

def get_asclaims(self):
  # self is network design
  partition = self.get("metadata", {}).get("name", "")
  namespace = self.get("metadata", {}).get("namespace", "")
  
  as_claims = []
  protocols = self.get("spec", {}).get("protocols", {})
  # is ibgp enabled
  ibgp = protocols.get("ibgp", None)
  if ibgp != None:
    as_claim_name = partition + "." + "ibgp"
    spec = {"index": partition,"id": ibgp.get("as", 0)}
    as_claims.append(get_asclaim(as_claim_name, namespace, spec)) 
  # is ebgp enabled
  ebgp = protocols.get("ebgp", None)
  if ebgp != None:
    as_claim_name = partition + "." + "aspool"
    spec = {"index": partition,"range": ebgp.get("asPool", "")}
    as_claims.append(get_asclaim(as_claim_name, namespace, spec)) 
  return as_claims

def get_node_asclaims(self, node):
  # self is network design
  partition = self.get("metadata", {}).get("name", "")
  namespace = self.get("metadata", {}).get("namespace", "")
  node_name = node.get("metadata", {}).get("name", "")

  as_claims = []
  protocols = self.get("spec", {}).get("protocols", {})
  # is ebgp enabled
  ebgp = protocols.get("ebgp", None)
  if ebgp != None:
    as_claim_name = partition + "." + "aspool"
    spec = {
      "index": partition,
      "selector": {
        "matchLabels": {
          "be.kuid.dev/claim-name": partition + "." + "aspool"
        },
      }
    }
    as_claims.append(get_asclaim(node_name, namespace, spec)) 
  
def get_asclaim(name, namespace, spec):
  return {
    "apiVersion": "as.be.kuid.dev/v1alpha1",
    "kind": "ASClaim",
    "metadata": {
        "name": name,
        "namespace": namespace
    },
    "spec": spec,
  }