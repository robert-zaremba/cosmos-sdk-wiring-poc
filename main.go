// Copyright 2018 The Wire Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// The greeter binary simulates an event with greeters greeting guests.
package main

import (
	"fmt"
	"os"
)

func main() {
	e, err := InitializeEvent("hi there!")
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}
	e.Start()

	mb, err := InitializeModules("keyA", "keyB")
	if err != nil {
		fmt.Printf("failed to create module: %s\n", err)
		os.Exit(2)
	}
	kb := mb.RegisterKeeper()
	s, err := kb.Slash("robert")
	fmt.Println("Slashing: ", s, err)

}
