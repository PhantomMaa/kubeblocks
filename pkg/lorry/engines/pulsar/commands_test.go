/*
Copyright (C) 2022-2024 ApeCloud Co., Ltd

This file is part of KubeBlocks project

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package pulsar

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/apecloud/kubeblocks/pkg/lorry/engines"
)

var _ = Describe("Pulsar Engine", func() {
	It("connection command", func() {
		pulsar := NewProxyCommands()

		Expect(pulsar.ConnectCommand(nil)).ShouldNot(BeNil())
		authInfo := &engines.AuthInfo{
			UserName:   "user-test",
			UserPasswd: "pwd-test",
		}
		Expect(pulsar.ConnectCommand(authInfo)).ShouldNot(BeNil())
	})

	It("proxy connection example", func() {
		pulsar := NewProxyCommands().(*Commands)

		info := &engines.ConnectionInfo{
			User:     "user",
			Host:     "host",
			Password: "*****",
			Port:     "1234",
		}
		for k := range pulsar.examples {
			fmt.Printf("%s Connection Example\n", k.String())
			Expect(pulsar.ConnectExample(info, k.String())).ShouldNot(BeZero())
		}

		Expect(pulsar.ConnectExample(info, "")).ShouldNot(BeZero())
	})

	It("broker connection example", func() {
		pulsar := NewBrokerCommands().(*Commands)

		info := &engines.ConnectionInfo{
			User:     "user",
			Host:     "host",
			Password: "*****",
			Port:     "1234",
		}
		for k := range pulsar.examples {
			fmt.Printf("%s Connection Example\n", k.String())
			Expect(pulsar.ConnectExample(info, k.String())).ShouldNot(BeZero())
		}

		Expect(pulsar.ConnectExample(info, "")).ShouldNot(BeZero())
	})
})
