package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StaticIPs describes a StaticIPs resource
type StaticIPs struct {
	// TypeMeta为各资源通用元信息, 包括kind和apiVersion.
	metav1.TypeMeta `json:",inline"`
	// ObjectMeta为特定类型的元信息, 包括name, namespace, selfLink, labels等.
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// spec字段
	Spec StaticIPsSpec `json:"spec"`
	// status字段
	Status StaticIPsStatus `json:"status"`
}

// StaticIPsSpec is the spec for a MyResource resource
type StaticIPsSpec struct {
	Namespace string `json:"namespace"`
	OwnerKind string `json:"ownerKind"`
	// 格式可为 "192.168.1.1/24,192.168.1.2/24"
	IPPool  string `json:"ipPool"`
	Gateway string `json:"gateway"`

	// IPMap key 为 192.168.1.1/24 这种点分十进制字符串
	// val 为 OwnerPod 对象, 表示此 IP 的拥有者
	IPMap map[string]*OwnerPod `json:"ipmap"`
}

// OwnerPod ...
type OwnerPod struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	UID       string `json:"uid"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StaticIPsList is a list of StaticIPs resources
type StaticIPsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []StaticIPs `json:"items"`
}

// StaticIPsStatus is the status for a StaticIPsStatus resource
type StaticIPsStatus struct{}
