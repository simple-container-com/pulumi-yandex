// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	pftfgen "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfgen"

	yandex "github.com/simple-container-com/pulumi-yandex/provider"
)

func main() {
	// MainWithMuxer must pair with MainWithMuxer in cmd/pulumi-resource-yandex: it is
	// what writes the "mux" dispatch table into bridge-metadata.json, recording which
	// half of the muxed terraform-provider-yandex owns each resource. The plain
	// tfgen.Main emits a schema that looks complete but no mux mapping, and the
	// provider binary then dies at startup with "Missing precomputed mapping. Did you
	// run `make tfgen`?".
	pftfgen.MainWithMuxer("yandex", yandex.Provider())
}
