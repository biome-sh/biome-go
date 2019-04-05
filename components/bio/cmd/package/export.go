// Copyright © 2018 Elliott Davis <elliott@excellent.io>
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

package pkg

import "github.com/spf13/cobra"

// ExportCmd ...
var ExportCmd = &cobra.Command{
	Use:   "export <FORMAT> <PKG_IDENT>",
	Short: "Exports the package to the specified format",
	Args:  cobra.ExactArgs(2),
}

func init() {
	ExportCmd.Flags().StringP("url", "u", "https://bldr.habitat.sh/v1", `Specify an alternate Builder endpoint. If not specified, the value will be taken from the
HAB_BLDR_URL environment variable if defined.`)
	ExportCmd.Flags().StringP("channel", "c", "stable", "Retrieve the container's package from the specified release channel")
}
