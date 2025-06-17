# xVirt [WIP]

xVirt is a Kubernetes-native virtualization solution that integrates QEMU and libvirt with Kubernetes, leveraging the power of Kubernetes' scheduling, monitoring, and declarative APIs while preserving the flexibility of QEMU/libvirt. Unlike traditional virtualization approaches, xVirt runs as a lightweight plugin on each cluster node and uses a Kubernetes Operator to manage virtual machines (VMs) through Custom Resource Definitions (CRDs).

## Goals

- **Seamless Integration**: Bring QEMU/libvirt-based virtualization into Kubernetes clusters.
- **Kubernetes Ecosystem**: xVirt is a Go-based reimplementation of a Python project, using Kubernetes Operators for improved performance, resource management, and observability.
- **Flexibility**: Retain QEMU/libvirt's rich feature set, avoiding the limitations imposed by containerizing libvirt.
- **Performance**: Minimize overhead by running plugins directly on nodes, bypassing unnecessary abstractions.

## Why Not KubeVirt?

While KubeVirt is a robust solution for Kubernetes virtualization, xVirt addresses specific use cases and limitations:

**Avoiding Containerization Constraints**:

- KubeVirt encapsulates libvirt within pods, fully aligning with Kubernetes' container model but introducing restrictions, particularly in networking. KubeVirt's networking relies heavily on CNI plugins, which may not offer the same flexibility as QEMU/libvirt's native networking capabilities outside containers.
- xVirt runs libvirt directly on the node, preserving its full feature set and flexibility, especially for advanced networking scenarios.

## Architecture

xVirt consists of two main components:

1. **xVirt Plugin**:
   - Runs on each Kubernetes node.
   - Directly connects to the node's QEMU/libvirt instance.
   - Handles the creation, updating, and deletion of VMs based on requests forwarded from the controller.
   - Ensures low-latency and direct access to virtualization capabilities without containerization overhead.

2. **xVirt Operator and Controller**:
   - Manages CRDs for VM definitions.
   - Forwards VM-related requests to the appropriate node's xVirt plugin.
   - Integrates with Kubernetes' scheduling and monitoring ecosystem for declarative VM management.


## Use Cases

- Running VMs with complex networking requirements that exceed standard CNI plugin capabilities.
- Migrating legacy QEMU/libvirt-based workloads to Kubernetes without sacrificing functionality.
- Building hybrid clusters with both containerized and VM workloads, managed declaratively.

## Roadmap

TODO