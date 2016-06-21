//
// Copyright (c) 2016 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package uuid

import "testing"

func TestUUID(t *testing.T) {
	testUUIDs := []string{
		"f81d4fae-7dec-11d0-a765-00a0c91e6bf6",
		"30dedd5c-48d9-45d3-8b44-f973e4f35e48",
		"69e84267-ed01-4738-b15f-b47de06b62e7",
		"e35ed972-c46c-4aad-a1e7-ef103ae079a2",
		"eba04826-62a5-48bd-876f-9119667b1487",
		"ca957444-fa46-11e5-94f9-38607786d9ec",
		"ab68111c-03a6-11e6-87de-001320fb6e31",
	}

	for _, s := range testUUIDs {
		uuid, err := Parse(s)
		if err != nil {
			t.Fatalf("Unable to parse %s: %s", s, err)
		}
		s2 := uuid.String()
		if s != s2 {
			t.Fatalf("%s and %s do not match", s, s2)
		}
	}
}

func TestGenUUID(t *testing.T) {
	for i := 0; i < 100; i++ {
		u := Generate()
		s := u.String()
		if s[14] != '4' {
			t.Fatalf("Invalid UUID.  Version number is incorrect")
		}
		u2, err := Parse(s)
		if err != nil {
			t.Fatalf("Failed to parse UUID %s : %s", s, err)
		}
		if u != u2 {
			t.Fatalf("Generated and Parsed UUIDs are not equal")
		}
	}
}
