/*
Copyright 2019 Citrix Systems, Inc
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

package client_test

import (
	"citrix-istio-adaptor/adsclient"
	"citrix-istio-adaptor/tests/env"
	"testing"
	"time"
)

func init() {
	env.Init()
}

func Test_http_frontend(t *testing.T) {
	t.Log("Http service test start")
	env.ClearNetscalerConfig()
	grpcServer := env.NewGrpcADSServer(1234)
	httpServer, err := env.StartHTTPServer(9000, "/path1", "Hello World!")
	if err != nil {
		t.Errorf("http server creation failed : %v", err)
	}
	discoveryClient, errc := adsclient.NewAdsClient("localhost:1234", "", false, "ads_client_node_1", "test-app", env.GetNetscalerURL(), env.GetNetscalerUser(), env.GetNetscalerPassword(), "nsip", "", "")
	if errc != nil {
		t.Errorf("newAdsClient failed with %v", errc)
	}
	discoveryClient.StartClient()
	route := env.MakeRoute("r1", "*", "c1")
	listener, errl := env.MakeHttpListener("l1", "0.0.0.0", 8000, "r1")
	if errl != nil {
		t.Errorf("makeListener failed with %v", errl)
	}
	cluster := env.MakeCluster("c1")
	endpoint := env.MakeEndpoint("c1", []env.ServiceEndpoint{{env.GetLocalIP(), 9000}})

	err = grpcServer.UpdateSpanshotCache("1", discoveryClient.GetNodeID(), listener, route, cluster, endpoint)
	if errl != nil {
		t.Errorf("updateSpanshotCache failed with %v", err)
	}

	time.Sleep(2 * time.Second)

	code, resp, err1 := env.DoHTTPGet("http://" + env.GetNetscalerIP() + ":8000/path1")
	if err1 != nil {
		t.Errorf("http get returned error: %v", err1)
	}
	t.Logf("HTTPget returned code:%d response:%s", code, resp)
	if code != 200 {
		t.Errorf("Expected 200 OK response, received %d", code)
	}
	discoveryClient.StopClient()
	grpcServer.StopGrpcADSServer()
	env.StopHTTPServer(httpServer)
	t.Log("Http service test stop")
}
