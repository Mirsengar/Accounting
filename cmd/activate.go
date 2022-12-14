package cmd

import (
	"Accounting/helpers"
	"Accounting/models"
	"github.com/spf13/cobra"
	"log"
)

var logger *log.Logger = helpers.GetLoggerInstace()

var activateCmd = &cobra.Command{
	Use:   "activate",
	Short: "Activate a particular user using their userID",
	Long: `Activating a user will overwrite previously logged in user.
		Activated data will be saved in credentials.json file.
	`,
	Run: func(cmd *cobra.Command, args []string) {

		userID, err := cmd.Flags().GetUint("user")
		if err != nil {
			logger.Fatalln("Failed to get `user` flag \n", err)
		}
		user := models.User{}
		sqlDb.GetUserDetails(userID, &user)
		helpers.SaveUser(&user)
		logger.Printf("User %d is activated", user.ID)
	},
}

func init() {
	rootCmd.AddCommand(activateCmd)
	activateCmd.Flags().Uint("user", 0, "activate --user 1")
}
