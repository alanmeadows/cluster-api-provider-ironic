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

package testing

import (
	"fmt"
	"github.com/alanmeadows/cluster-api-provider-ironic/pkg/cloud/ironic/actuators/machine"
	"github.com/gophercloud/gophercloud"
	th "github.com/gophercloud/gophercloud/testhelper"
	"testing"
)

func TestNewBareMetalService(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	// mock out ironic endpoint so we can list nodes
	HandleNodeListSuccessfully(t)

	options := gophercloud.AuthOptions{
		Username:         "me",
		Password:         "secret",
		DomainName:       "default",
		TenantName:       "project",
		IdentityEndpoint: th.Endpoint(),
	}

	bms, err := machine.NewBareMetalService(options)
	if err != nil {
		t.Fatalf("Couldn't create new baremetal service instance: %v", err)
	}

	opts := &machine.BareMetalListOpts{}
	fmt.Printf("Trying to do baremetal list\n")
	bareMetalNodes, err := bms.GetBareMetalList(opts)
	if err != nil {
		t.Errorf("Couldn't list baremetal servers: %v", err)
	}
	if len(bareMetalNodes) != 3 {
		t.Errorf("Expected to find 3 baremetal nodes in the test data, but received: %v", len(bareMetalNodes))
	}

}
