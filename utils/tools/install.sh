#!/bin/bash
set -o errexit

source ~/.bashrc


# 1./ Install arkade
echo 'Install arkade'

curl -sLS https://get.arkade.dev | sudo sh

arkade version


# 2./ Install kind

echo 'Install kind'


if ! [ -f /usr/local/bin/kind ]; then
  arkade get kind
  sudo mv ${HOME}/.arkade/bin/kind /usr/local/bin/
else
  echo 'kind already installed'
  ls -alh /usr/local/bin/kind
fi;

kind version

# 2./ Install kubectl and helm

echo 'Install kubectl'

if ! [ -f /usr/local/bin/kubectl ]; then
  arkade get kubectl
  sudo mv ${HOME}/.arkade/bin/kubectl /usr/local/bin/
else
  echo 'kubectl already installed'
  ls -alh /usr/local/bin/kubectl
fi;


kubectl version --client


echo 'Install helm'

if ! [ -f /usr/local/bin/helm ]; then
  arkade get helm
  sudo mv ${HOME}/.arkade/bin/helm /usr/local/bin/
else
  echo 'helm already installed'
  ls -alh /usr/local/bin/helm
fi;


helm version
# version.BuildInfo{Version:"v3.14.4", GitCommit:"81c902a123462fd4052bc5e9aa9c513c4c8fc142", GitTreeState:"clean", GoVersion:"go1.21.9"}

export KND_CLUSTER_NAME=${KND_CLUSTER_NAME:-"koperator_play"}
export KND_CLUSTER_CONFIG_FILE=${KND_CLUSTER_CONFIG_FILE:-"./cluster.yaml"}

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




