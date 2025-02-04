/*
 * Copyright 2022 VMware, Inc.
 * All Rights Reserved.
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*   http://www.apache.org/licenses/LICENSE-2.0
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*/

package avirest

import (
	"sync"

	"github.com/vmware/load-balancer-and-ingress-services-for-kubernetes/third_party/github.com/vmware/alb-sdk/go/clients"
)

var infraAviClientInstance *clients.AviClient
var ctrlClientOnce sync.Once

func InfraAviClientInstance(c ...*clients.AviClient) *clients.AviClient {
	ctrlClientOnce.Do(func() {
		if len(c) > 0 {
			infraAviClientInstance = c[0]
		}
	})
	return infraAviClientInstance
}
