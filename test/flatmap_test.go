package test

import (
	"encoding/json"
	"fmt"
	"github.com/wencaiwulue/stream4go/stream"
	"sigs.k8s.io/yaml"
	"testing"
)

func TestFlat(t *testing.T) {
	s := "apiEndpoints:\n  docker-desktop:\n    advertiseAddress: 192.168.65.3\n    bindPort: 6443\napiVersion: kubeadm.k8s.io/v1beta2\nkind: ClusterStatus"
	b, err := yaml.YAMLToJSON([]byte(s))
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("json: %v\n", string(b))
	status := ClusterStatus{}
	_ = json.Unmarshal(b, &status)
	ss := stream.ObjectStream.Of(status.APIEndpoints).FlatMap().MapToString(func(i interface{}) string {
		return i.(*stream.Entry).Value.Interface().(APIEndpoint).AdvertiseAddress
	}).ToSlice()
	for _, endpoint := range ss {
		fmt.Printf("endpoint: %v\n", endpoint)
	}
}

type ClusterStatus struct {
	APIEndpoints map[string]APIEndpoint `json:"apiEndpoints"`
}
type APIEndpoint struct {
	AdvertiseAddress string `json:"advertiseAddress"`
	BindPort         int32  `json:"bindPort"`
}
