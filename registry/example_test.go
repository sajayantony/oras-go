// Copyright The ORAS Authors.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package registry_test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"oras.land/oras-go/v2/registry/remote"
)

func TestMain(m *testing.M) {
	fmt.Println("Setup registry")
	// Run the http test server with test content preloaded.
}

func Example_resolveTag() {

	ctx := context.Background()
	tagName := "v1"
	registryUri := "localhost:5000"
	repositoryName := "test"

	registry, err := remote.NewRegistry(registryUri)
	if err != nil {
		log.Fatal(err)
	}

	repository, err := registry.Repository(ctx, repositoryName)
	if err != nil {
		log.Fatal(err)
	}

	desc, err := repository.Resolve(ctx, tagName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Digest: %s\n", desc.Digest)
	fmt.Printf("Size: %d\n", desc.Size)
	fmt.Printf("Mediatype: %s\n", desc.MediaType)

	// Output:
	// Digest: sha256:92c7f9c92844bbbb5d0a101b22f7c2a7949e40f8ea90c8b3bc396879d95e899a
	// Size: 524
	// Mediatype: application/vnd.docker.distribution.manifest.v2+json
}

func Example_getBlob() {
	// Show how to download blob
}

func Example_copy() {
	// Copy a single artifact
}

func Example_copyWithReferences() {
	// Copy a graph of references
}

func ExampleRepositories() {
	// Show how to return the repositories here -
}

func ExampleTags() {
	// Show how to return the tags for a repository
}
