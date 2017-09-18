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

import "github.com/EtixLabs/cameradar/cameradar"
import "log"

func main() {
	streams, err := cmrdr.Discover("localhost", "8554")
	if err != nil {
		log.Fatalf("Fatal Error: %v", err)
	}

	credentials := cmrdr.Credentials{}
	credentials.Usernames = append(credentials.Usernames, "admin")
	credentials.Passwords = append(credentials.Passwords, "12345")

	streams, err = cmrdr.AttackCredentials(streams, credentials)
	if err != nil {
		log.Fatalf("Fatal Error: %v", err)
	}

	routes := cmrdr.Routes{}
	routes = append(routes, "/live.sdp")

	streams, err = cmrdr.AttackRoute(streams, routes)
	if err != nil {
		log.Fatalf("Fatal Error: %v", err)
	}

	for _, stream := range streams {
		log.Printf("Stream: \n%v\n", stream)
	}
}
