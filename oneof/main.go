// Copyright © 2020 Krishna Iyer Easwaran.
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
	"fmt"

	"go.krishnaiyer.dev/go-snippets/oneof/pbgen"
)

func main() {

	userWithName := &pbgen.User{
		Identifier: &pbgen.User_Username{
			Username: "user1",
		},
		Password: "test",
	}
	fmt.Println(userWithName)

	userWithEmail := &pbgen.User{
		Identifier: &pbgen.User_Email{
			Email: "user1@email.com",
		},
		Password: "test",
	}
	fmt.Println(userWithEmail)

	userWithUserID := &pbgen.User{
		Identifier: &pbgen.User_UserId{
			UserId: 1234,
		},
		Password: "test",
	}
	fmt.Println(userWithUserID)

	user := pbgen.User{}

	switch id := user.Identifier.(type) {
	case *pbgen.User_Username:
		fmt.Println(id.Username)
	case *pbgen.User_Email:
		fmt.Println(id.Email)
	case *pbgen.User_UserId:
		fmt.Println(id.UserId)
	default:
		// It's fine to panic here since this branch should not be reached.
		panic(fmt.Sprintf("proto: unexpected type %T", id))
	}

	sampleMessage1 := &pbgen.SampleMessage{
		TestOneof: &pbgen.SampleMessage_SubMessage{
			SubMessage: &pbgen.SubMessage{
				Id: "test",
			},
		},
	}
	fmt.Println(sampleMessage1)

	sampleMessage2 := &pbgen.SampleMessage{
		TestOneof: &pbgen.SampleMessage_Name{
			Name: "test",
		},
	}
	fmt.Println(sampleMessage2)

	sampleMessage := &pbgen.SampleMessage{}

	switch msg := sampleMessage.TestOneof.(type) {
	case *pbgen.SampleMessage_SubMessage:
		if msg.SubMessage != nil {
			fmt.Println(msg.SubMessage.Id)
		}
	case *pbgen.SampleMessage_Name:
		fmt.Println(msg.Name)
	default:
		panic(fmt.Sprintf("proto: unexpected type %T", msg))
	}

}
