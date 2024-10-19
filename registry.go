/*
 * Copyright 2021 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package polaris

import (
	polariskitex "github.com/cloudwego-contrib/cwgo-pkg/registry/polaris/polariskitex"
)

// Registry is extension interface of Kitex registry.Registry.
type Registry = polariskitex.Registry

// NewPolarisRegistry creates a polaris based registry.
func NewPolarisRegistry(so ServerOptions, configFile ...string) (Registry, error) {
	return polariskitex.NewPolarisRegistry(so, configFile...)
}
