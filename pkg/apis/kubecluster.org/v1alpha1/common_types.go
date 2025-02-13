package v1alpha1

type ReplicaType string

const (
	ReplicaTypeNameNode        ReplicaType = "namenode"
	ReplicaTypeDataNode        ReplicaType = "datanode"
	ReplicaTypeResourcemanager ReplicaType = "resourcemanager"
	ReplicaTypeNodemanager     ReplicaType = "nodemanager"

	ReplicaTypeConfigMap ReplicaType = "configmap"
)
