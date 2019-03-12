package testing

import (
	"fmt"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"net/http"
	"testing"
)

const ID = "cbc36478b0bd8e67e89469c7749d4127"

// NodeListBody contains the canned body of a nodes.List response, without detail.
const NodeListBody = `
 {
  "nodes": [
    {
      "instance_uuid": "abcd12345",
      "links": [
        {
          "href": "http://ironic.example.com:6385/v1/nodes/d2630783-6ec8-4836-b556-ab427c4b581e",
          "rel": "self"
        },
        {
          "href": "http://ironic.example.com:6385/nodes/d2630783-6ec8-4836-b556-ab427c4b581e",
          "rel": "bookmark"
        }
      ],
      "maintenance": false,
      "name": "foo",
      "power_state": null,
      "provision_state": "enroll"
    },
    {
      "instance_uuid": "abcd54321",
      "links": [
        {
          "href": "http://ironic.example.com:6385/v1/nodes/08c84581-58f5-4ea2-a0c6-dd2e5d2b3662",
          "rel": "self"
        },
        {
          "href": "http://ironic.example.com:6385/nodes/08c84581-58f5-4ea2-a0c6-dd2e5d2b3662",
          "rel": "bookmark"
        }
      ],
      "maintenance": false,
      "name": "bar",
      "power_state": null,
      "provision_state": "enroll"
    },
    {
      "instance_uuid": null,
      "links": [
        {
          "href": "http://ironic.example.com:6385/v1/nodes/c9afd385-5d89-4ecb-9e1c-68194da6b474",
          "rel": "self"
        },
        {
          "href": "http://ironic.example.com:6385/nodes/c9afd385-5d89-4ecb-9e1c-68194da6b474",
          "rel": "bookmark"
        }
      ],
      "maintenance": false,
      "name": "baz",
      "power_state": null,
      "provision_state": "enroll"
    }
  ]
}
`

// KeystoneRoot is meant to represent the root of a keystone endpoint.
const KeystoneRootBody = `
      {
        "versions": {
          "values": [
            {
              "status": "stable",
              "id": "v3.0",
              "links": [
                { "href": "%s", "rel": "self" }
              ]
            },
            {
              "status": "stable",
              "id": "v2.0",
              "links": [
                { "href": "%s", "rel": "self" }
              ]
            }
          ]
        }
      }
    `

// KeystoneToken represents a token request response from Keystone.
const KeystoneTokenBody = `
  {
    "token": {
        "audit_ids": ["VcxU2JYqT8OzfUVvrjEITQ", "qNUTIJntTzO1-XUk5STybw"],
        "catalog": [
            {
                "endpoints": [
                    {
                        "id": "39dc322ce86c4111b4f06c2eeae0841b",
                        "interface": "public",
                        "region": "RegionOne",
                        "url": "http://localhost:5000"
                    },
                    {
                        "id": "ec642f27474842e78bf059f6c48f4e99",
                        "interface": "internal",
                        "region": "RegionOne",
                        "url": "http://localhost:5000"
                    },
                    {
                        "id": "c609fc430175452290b62a4242e8a7e8",
                        "interface": "admin",
                        "region": "RegionOne",
                        "url": "http://localhost:35357"
                    }
                ],
                "id": "4363ae44bdf34a3981fde3b823cb9aa2",
                "type": "identity",
                "name": "keystone"
            },
            {
                "endpoints": [
                    {
                        "id": "49dc322ce86c4111b4f06c2eeae0841b",
                        "interface": "public",
                        "region": "RegionOne",
                        "url": "%s"
                    },
                    {
                        "id": "bc642f27474842e78bf059f6c48f4e99",
                        "interface": "internal",
                        "region": "RegionOne",
                        "url": "%s"
                    },
                    {
                        "id": "a609fc430175452290b62a4242e8a7e8",
                        "interface": "admin",
                        "region": "RegionOne",
                        "url": "%s"
                    }
                ],
                "id": "1363ae44bdf34a3981fde3b823cb9aa2",
                "type": "baremetal",
                "name": "ironic"
            }       
        ],
        "expires_at": "2013-02-27T18:30:59.999999Z",
        "is_domain": false,
        "issued_at": "2013-02-27T16:30:59.999999Z",
        "methods": [
            "password"
        ],
        "project": {
            "domain": {
                "id": "1789d1",
                "name": "example.com"
            },
            "id": "263fd9",
            "name": "project-x"
        },
        "roles": [
            {
                "id": "76e72a",
                "name": "admin"
            },
            {
                "id": "f4f392",
                "name": "member"
            }
        ],
        "service_providers": [
            {
                "auth_url":"https://example.com:5000/v3/OS-FEDERATION/identity_providers/acme/protocols/saml2/auth",
                "id": "sp1",
                "sp_url": "https://example.com:5000/Shibboleth.sso/SAML2/ECP"
            },
            {
                "auth_url":"https://other.example.com:5000/v3/OS-FEDERATION/identity_providers/acme/protocols/saml2/auth",
                "id": "sp2",
                "sp_url": "https://other.example.com:5000/Shibboleth.sso/SAML2/ECP"
            }
        ],
        "user": {
            "domain": {
                "id": "1789d1",
                "name": "example.com"
            },
            "id": "0ca8f6",
            "name": "Joe",
            "password_expires_at": "2016-11-06T15:32:17.000000"
        }
    }
}
`

// This is a helper function to be used by other tests to support stubbing keystone basics.
func keystoneHelper(t *testing.T) {

	// handle root of keystone endpoint
	th.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, KeystoneRootBody, th.Endpoint()+"v3/", th.Endpoint()+"v2.0/")
	})

	// handle token response
	th.Mux.HandleFunc("/v3/auth/tokens", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Subject-Token", ID)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, KeystoneTokenBody, th.Endpoint(), th.Endpoint(), th.Endpoint())
	})

}

// HandleNodeListSuccessfully sets up the test server to respond to a server List request.
func HandleNodeListSuccessfully(t *testing.T) {

	// setup stubs for keystone to work properly
	keystoneHelper(t)

	// stub for /nodes
	th.Mux.HandleFunc("/nodes", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		r.ParseForm()

		marker := r.Form.Get("marker")
		switch marker {
		case "":
			fmt.Fprintf(w, NodeListBody)
		case "9e5476bd-a4ec-4653-93d6-72c93aa682ba":
			fmt.Fprintf(w, `{ "servers": [] }`)
		default:
			t.Fatalf("/nodes invoked with unexpected marker=[%s]", marker)
		}
	})
}
