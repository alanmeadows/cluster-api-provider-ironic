/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package machine

import (
	//"encoding/base64"
	"fmt"
	"time"

	"gopkg.in/yaml.v2"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/baremetal/v1/nodes"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/tokens"
	"github.com/gophercloud/utils/openstack/clientconfig"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	CloudsSecretKey = "clouds.yaml"

	TimeoutTrunkDelete       = 3 * time.Minute
	RetryIntervalTrunkDelete = 5 * time.Second

	TimeoutPortDelete       = 3 * time.Minute
	RetryIntervalPortDelete = 5 * time.Second
)

type BareMetalService struct {
	provider        *gophercloud.ProviderClient
	baremetalClient *gophercloud.ServiceClient
	identityClient  *gophercloud.ServiceClient
}

type BareMetal struct {
	nodes.Node
}

type BareMetalListOpts struct {
	// just the basics for now

	InstanceUUID string `q:"instance_uuid"`

	ChassisUUID string `q:"chassis_uuid"`
}

func GetCloudFromSecret(kubeClient kubernetes.Interface, namespace string, secretName string, cloudName string) (clientconfig.Cloud, error) {
	emptyCloud := clientconfig.Cloud{}

	if secretName == "" {
		return emptyCloud, nil
	}

	if secretName != "" && cloudName == "" {
		return emptyCloud, fmt.Errorf("Secret name set to %v but no cloud was specified. Please set cloud_name in your machine spec.", secretName)
	}

	secret, err := kubeClient.CoreV1().Secrets(namespace).Get(secretName, metav1.GetOptions{})
	if err != nil {
		return emptyCloud, err
	}

	content, ok := secret.Data[CloudsSecretKey]
	if !ok {
		return emptyCloud, fmt.Errorf("OpenStack credentials secret %v did not contain key %v",
			secretName, CloudsSecretKey)
	}
	var clouds clientconfig.Clouds
	err = yaml.Unmarshal(content, &clouds)
	if err != nil {
		return emptyCloud, fmt.Errorf("failed to unmarshal clouds credentials stored in secret %v: %v", secretName, err)
	}

	return clouds.Clouds[cloudName], nil
}

func NewBareMetalService(clientOpts gophercloud.AuthOptions) (*BareMetalService, error) {

	clientOpts.AllowReauth = true

	provider, err := openstack.AuthenticatedClient(clientOpts)
	if err != nil {
		return nil, fmt.Errorf("Create providerClient err: %v", err)
	}

	identityClient, err := openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{
		Region: "",
	})
	if err != nil {
		return nil, fmt.Errorf("Create identityClient err: %v", err)
	}
	baremetalClient, err := openstack.NewBareMetalV1(provider, gophercloud.EndpointOpts{
		//TODO
		//Region: clientOpts.RegionName,
		Region: "RegionOne",
	})
	if err != nil {
		return nil, fmt.Errorf("Create baremetalClient err: %v", err)
	}

	return &BareMetalService{
		provider:        provider,
		identityClient:  identityClient,
		baremetalClient: baremetalClient,
	}, nil
}

// UpdateToken to update token if need.
func (is *BareMetalService) UpdateToken() error {
	token := is.provider.Token()
	result, err := tokens.Validate(is.identityClient, token)
	if err != nil {
		return fmt.Errorf("Validate token err: %v", err)
	}
	if result {
		return nil
	}
	klog.V(2).Infof("Token is out of date, getting new token.")
	reAuthFunction := is.provider.ReauthFunc
	if reAuthFunction() != nil {
		return fmt.Errorf("reAuth err: %v", err)
	}
	return nil
}

// Return a list of BareMetal nodes from ironic
func (is *BareMetalService) GetBareMetalList(opts *BareMetalListOpts) ([]*BareMetal, error) {

	var listOpts nodes.ListOpts
	if opts != nil {
		listOpts = nodes.ListOpts{
			InstanceUUID: opts.InstanceUUID,
		}
	} else {
		listOpts = nodes.ListOpts{}
	}

	allPages, err := nodes.List(is.baremetalClient, listOpts).AllPages()
	if err != nil {
		return nil, fmt.Errorf("Get service list err: %v", err)
	}
	nodeList, err := nodes.ExtractNodes(allPages)
	if err != nil {
		return nil, fmt.Errorf("Extract services list err: %v", err)
	}
	var bareMetalList []*BareMetal
	for _, node := range nodeList {
		bareMetalList = append(bareMetalList, nodeToBareMetal(&node))
	}
	return bareMetalList, nil
}

// Translate a node to a Baremetal instance
func nodeToBareMetal(node *nodes.Node) *BareMetal {
	return &BareMetal{*node}
}
