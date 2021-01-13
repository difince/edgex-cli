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

package list

import (
	"context"
	"github.com/edgexfoundry/edgex-cli/config"
	"github.com/edgexfoundry/edgex-cli/pkg/formatters"

	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/scheduler"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/urlclient/local"
	"github.com/edgexfoundry/go-mod-core-contracts/models"

	"github.com/spf13/cobra"
)

const intervalTemplate = "Interval ID\tName\tStart\tEnd\tFrequency\tCron\tRunOnce\n" +
	"{{range .}}" +
	"{{.ID}}\t{{.Name}}\t{{.Start}}\t{{.End}}\t{{.Frequency}}\t{{.Cron}}\t{{.RunOnce}}\n" +
	"{{end}}"

//Handler list Intervals and display them
func Handler(cmd *cobra.Command, args []string) (err error) {
	client := scheduler.NewIntervalClient(
		local.New(config.Conf.Clients["Scheduler"].Url() + clients.ApiIntervalRoute),
	)

	var intervals []models.Interval
	if len(args) == 0 {
		intervals, err = client.Intervals(cmd.Context())
	} else {
		intervals, err = getInterval(cmd.Context(), client, args[0])
	}
	if err != nil {
		return err
	}

	formatter := formatters.NewFormatter(intervalTemplate, nil)
	err = formatter.Write(intervals)
	return
}

func getInterval(ctx context.Context, client scheduler.IntervalClient, id string) ([]models.Interval, error) {
	interval, err := client.Interval(ctx, id)
	if err != nil {
		return nil, err
	}
	return []models.Interval{interval}, nil
}

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "A list of all intervals",
		Long:  `Return a list of all intervals or retrieve an interval by id`,
		Args:  cobra.MaximumNArgs(1),
		RunE:  Handler,
	}
	return cmd
}
