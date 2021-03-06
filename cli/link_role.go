/*
  Copyright (C) 2016 H2O.ai, Inc. <http://h2o.ai/>

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU Affero General Public License as
  published by the Free Software Foundation, either version 3 of the
  License, or (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU Affero General Public License for more details.

  You should have received a copy of the GNU Affero General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package cli

// FIXME use ByName remote calls

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var linkRoleHelp = `
role [roleName] [permissionIds]
Add permissions to a role. 
Examples:

	$ steam link role engineer ViewCluster ViewModel ViewWorkgroup
`

func linkRole(c *context) *cobra.Command {
	cmd := newCmd(c, linkRoleHelp, func(c *context, args []string) {
		if len(args) < 2 {
			log.Fatalln("Invalid usage. See 'steam help link role'.")
		}

		// -- Args --

		roleName := args[0]
		permissions := args[1:len(args)]
		fmt.Println(roleName, permissions)

		// -- Execution --

		role, err := c.remote.GetRoleByName(roleName)
		if err != nil {
			log.Fatalln(err) //FIXME format error
		}

		permissionIds, err := getPermissionIds(c, permissions...)
		if err != nil {
			log.Fatalln(err) //FIXME format error
		}

		c.remote.LinkRoleAndPermissions(role.Id, permissionIds)
		fmt.Println("Role", roleName, "linked to permissions:", permissions)
	})

	return cmd
}
