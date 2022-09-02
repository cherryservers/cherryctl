package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

func readConfig(filename string) (*viper.Viper, error) {

	v := viper.New()

	v.AddConfigPath("$HOME")
	v.SetConfigName(filename)
	v.AutomaticEnv()

	err := v.ReadInConfig()

	return v, err
}

func setDefaultProfile(defaultProfile string) {

	v1, err := readConfig(".cherry")
	if err != nil {
		//panic(fmt.Errorf("Error when reading config: %v", err))
	}

	configDefaultProfile := v1.GetString("default_profile")

	if defaultProfile != configDefaultProfile {

		if v1.IsSet(defaultProfile) {

			v1.Set("default_profile", defaultProfile)
			v1.WriteConfig()
			fmt.Printf("Default profile was set to: %v\n", defaultProfile)

		} else {
			log.Fatalf("There is no such profile in config: %s", defaultProfile)
		}
	}

	getProfileConfig(defaultProfile)

}

func getProfileConfig(defaultProfile string) (teamID int, projectID string) {

	v1, err := readConfig(".cherry")
	if err != nil {
		panic(fmt.Errorf("Error when reading config: %v", err))
	}

	configDefaultProfile := v1.GetString("default_profile")

	if configDefaultProfile != "" {
		profileName := configDefaultProfile

		// Do we have such profile in config?
		if v1.IsSet(profileName) {

			// Do we have team_id set?
			teamKey := fmt.Sprintf("%s.%s", profileName, "team-id")
			if v1.IsSet(teamKey) {
				teamID = v1.GetInt(teamKey)
			} else {
				log.Fatalf("team-id must be specified if configuration file is used")
			}
			// Do we have project_id set?
			projectKey := fmt.Sprintf("%s.%s", profileName, "project-id")
			if v1.IsSet(projectKey) {
				projectID = v1.GetString(projectKey)
			} else {
				log.Fatalf("project-id must be specified if configuration file is used")
			}

			tokenKey := fmt.Sprintf("%s.%s", profileName, "token")
			if v1.IsSet(tokenKey) {
				tokenID := v1.GetString(tokenKey)
				os.Setenv("CHERRY_AUTH_TOKEN", tokenID)
			} else {
				log.Fatalf("token must be specified if configuration file is used")
			}

		} else {
			fmt.Printf("There is no such profile in config: %s", profileName)
		}
	}

	return teamID, projectID
}
