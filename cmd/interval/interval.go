// Copyright © 2019 VMware, INC
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

package interval

import (
	"github.com/edgexfoundry/edgex-cli/cmd/interval/add"
	"github.com/edgexfoundry/edgex-cli/cmd/interval/list"
	"github.com/edgexfoundry/edgex-cli/cmd/interval/rm"
	"github.com/edgexfoundry/edgex-cli/cmd/interval/update"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "interval",
		Short: "Interval command",
		Long:  `Actions related to intervals (scheduler).`,
		RunE:  list.Handler,
	}
	cmd.AddCommand(add.NewCommand())
	cmd.AddCommand(rm.NewCommand())
	cmd.AddCommand(update.NewCommand())
	cmd.AddCommand(list.NewCommand())
	return cmd
}
