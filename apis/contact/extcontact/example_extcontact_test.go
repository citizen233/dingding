// Copyright 2020 FastWeGo
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

package extcontact_test

import (
	"fmt"

	"github.com/fastwego/dingding"
	"github.com/fastwego/dingding/apis/contact/extcontact"
)

func ExampleListLabelGroups() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := extcontact.ListLabelGroups(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleList() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := extcontact.List(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGet() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := extcontact.Get(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleCreate() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := extcontact.Create(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUpdate() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := extcontact.Update(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDelete() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := extcontact.Delete(ctx, payload)

	fmt.Println(resp, err)
}
