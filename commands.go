package main

import (
	"errors"
	"fmt"
	"github.com/danryan/go-pagerduty/pagerduty"
	"os"
)

func init() {
	CmdRunner.Register(cmdHelp)
	CmdRunner.Register(cmdIncidents)
	CmdRunner.Register(cmdSchedules)
	CmdRunner.Register(cmdUsers)
}

var cmdHelp = &Command{
	Name:        "help",
	Description: "Show usage",
	Run: func(args *Args) error {
		fmt.Println("usage: duty <command>")
		fmt.Println()
		fmt.Println("Available commands are:")
		for _, cmd := range CmdRunner.All() {
			if cmd.Name != "" {
				fmt.Printf("  %-12s%s\n", cmd.Name, cmd.Description)
			}
		}
		return nil
	},
}

var cmdIncidents = &Command{
	Name:        "incidents",
	Description: "List incidents",
	Run: func(args *Args) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		incidents, _, err := client.Incidents.List(&pagerduty.IncidentsOptions{})
		if err != nil {
			return err
		}

		for i, incident := range incidents {
			if i != 0 {
				fmt.Println()
			}
			fmt.Println(incident.ID)
			fmt.Println("Number:      ", incident.IncidentNumber)
			fmt.Println("Status:      ", incident.Status)
			fmt.Println("Assigned to: ", incident.AssignedToUser.Name)
			fmt.Println()
			fmt.Printf("    %s\n", incident.Summary.Subject)
			fmt.Printf("    %s\n", incident.Summary.Description)
		}
		return nil
	},
}

var cmdSchedules = &Command{
	Name:        "schedules",
	Description: "List schedules",
	Run: func(args *Args) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		schedules, _, err := client.Schedules.List(&pagerduty.SchedulesOptions{})
		if err != nil {
			return err
		}

		for i, schedule := range schedules.Schedules {
			if i != 0 {
				fmt.Println()
			}
			fmt.Println(schedule.ID)
			fmt.Println("Name:     ", schedule.Name)
			fmt.Println("Timezone: ", schedule.Timezone)
		}
		return nil
	},
}

var cmdUsers = &Command{
	Name:        "users",
	Description: "List users",
	Run: func(args *Args) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		users, _, err := client.Users.List(&pagerduty.UsersOptions{})
		if err != nil {
			return err
		}

		for i, user := range users {
			if i != 0 {
				fmt.Println()
			}
			fmt.Println(user.ID)
			fmt.Println("Name:     ", user.Name)
			fmt.Println("Email:    ", user.Email)
			fmt.Println("Role:     ", user.Role)
			fmt.Println("Timezone: ", user.TimeZone)
		}
		return nil
	},
}

func newClient() (*pagerduty.Client, error) {
	subdomain := os.Getenv("PAGERDUTY_SUBDOMAIN")
	apiKey := os.Getenv("PAGERDUTY_API_KEY")

	if subdomain == "" {
		return nil, errors.New("PAGERDUTY_SUBDOMAIN environment variable not set")
	}
	if apiKey == "" {
		return nil, errors.New("PAGERDUTY_API_KEY environment variable not set")
	}

	return pagerduty.New(subdomain, apiKey), nil
}
