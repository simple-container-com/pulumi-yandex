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
	"context"

	_ "embed"

	pftfbridge "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"

	yandex "github.com/simple-container-com/pulumi-yandex/provider"
)

//go:embed schema.json
var pulumiSchema []byte

func main() {
	// MainWithMuxer, not the plain tfbridge.Main: Provider().P is a mux over the
	// SDKv2 and plugin-framework halves of terraform-provider-yandex (see
	// provider/resources.go). Plain tfbridge.Main serves that shim directly, and the
	// plugin-framework side of it is schema-only — the binary panics on the engine's
	// very first runtime call with "schemaOnlyProvider does not implement runtime
	// operation InitLogging". MainWithMuxer instead builds a muxed gRPC server that
	// routes runtime operations to whichever half actually owns the resource.
	//
	// The failure mode is worth knowing: tfgen and `go build` are both perfectly
	// happy with the wrong Main, so nothing catches it until a real `pulumi` run.
	pftfbridge.MainWithMuxer(context.Background(), "yandex", yandex.Provider(), pulumiSchema)
}
