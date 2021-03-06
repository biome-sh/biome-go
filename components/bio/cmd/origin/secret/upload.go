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

package secret

import (
	"github.com/spf13/cobra"
)

// UploadCmd ...
var UploadCmd = &cobra.Command{
	Use:   "upload <KEY_NAME> <SECRET>",
	Short: "Create and upload a secret for your origin.",
	Args:  cobra.ExactArgs(2),
}
