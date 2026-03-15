#!/bin/bash
set -o errexit

source ~/.bashrc



sudo mkdir -p /usr/local/bin/

export KND_CLUSTER_NAME=${KND_CLUSTER_NAME:-"koperator-play"}
export KND_CLUSTER_CONFIG_FILE=${KND_CLUSTER_CONFIG_FILE:-"./utils/cluster/cluster.yaml"}

# create kind kubernetes cluster, with a local docker registry

echo 'Create kubernetes cluster'
echo "Just before creating cluster: KND_CLUSTER_NAME=[${KND_CLUSTER_NAME}]"
kind delete clusters --all
# 2. Create kind cluster with containerd registry config dir enabled
# TODO: kind will eventually enable this by default and this patch will
# be unnecessary.
#
# See:
# https://github.com/kubernetes-sigs/kind/issues/2875
# https://github.com/containerd/containerd/blob/main/docs/cri/config.md#registry-configuration
# See: https://github.com/containerd/containerd/blob/main/docs/hosts.md
# https://kind.sigs.k8s.io/docs/user/configuration#extra-mounts
# extraMounts of /etc/localtime and /etc/timezone are there for later deployment of https://github.com/k8tz/k8tz
# ---
#  to change KND_CLUSTER_CONFIG_FILE: edit the 'cluster.yaml' in the same folder than this script


# -- about /etc/localtime, I may get that error, which i want to avoid:
# 
#   docker: Error response from daemon: failed to create task for container: failed to create shim task: OCI runtime create failed: runc create failed: unable to start container process: error during container init: error mounting "/etc/localtime" to rootfs at "/etc/localtime": mount /etc/localtime:/etc/localtime (via /proc/self/fd/7), flags: 0x5000: not a directory: unknown: Are you trying to mount a directory onto a file (or vice-versa)? Check if the specified host path exists and is the expected type.
# -
# if that error happens, that would be because the /etc/localtime file does not exist on the system, so i have to test that

if [ -f /etc/localtime ]; then
  echo "Okay, /etc/localtime does exists"
  ls -alh /etc/localtime
  export TARGET_OF_LOCALTIME_SYMLINK=$(sudo ls -alh /etc/localtime | awk '{ print $NF }')
  echo " TARGET_OF_LOCALTIME_SYMLINK=[${TARGET_OF_LOCALTIME_SYMLINK}]"
  ls -alh ${TARGET_OF_LOCALTIME_SYMLINK}
fi;


# ---
#  First we create all the folders mounte in the kind cluser nodes, which are subfolders of the ${OPENEBS_CSI_DISK_MOUNT_POINT} Folder
ls -alh ${KND_CLUSTER_CONFIG_FILE}

# cat ${KND_CLUSTER_CONFIG_FILE} | yq '.nodes[].extraMounts.[]  | select( .hostPath | contains("OPENEBS_CSI_DISK_MOUNT_POINT"))' | grep hostPath | awk -F ': ' '{ print $NF }' > ./list.of.path.to.mount.on.each.cluster.nodes.list
cat ${KND_CLUSTER_CONFIG_FILE}

# ---
# And now the cluster is ready to create

kind create cluster -n "${KND_CLUSTER_NAME}" --wait 25m --config=${KND_CLUSTER_CONFIG_FILE}

# Then we wait a bit to let the cluster get up'n running cosily
sleep 60s

kubectl cluster-info









