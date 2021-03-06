package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/imdario/mergo"
)

var (
	// Set defaults for the config to ensure it works
	// or warns even if a particular setting is missing
	conf *Config
)

// Config represetns the bot configuration loaded from the JSON
// file "./config.json".
type Config struct {
	// Prefix is the string that will prefix all commands
	// which this not will listen for.
	Prefix string `json:"prefix" deepcopier:"field:Prefix"`
	// Token is the Discord bot user token.
	Token string `json:"token" deepcopier:"skip"`
	// HelpChannelID is the channel ID to which help messages from
	// netsoc-admin will be sent.
	HelpChannelID string `json:"helpChannelID" deepcopier:"field:HelpChannelID"`
	// BotHostName is the address which the bot can be reached at
	// over the internet. This is used by netsocadmin to reach the
	// '/help' endpoint.
	BotHostName string `json:"botHostName" deepcopier:"field:BotHostName"`
	// SysAdminTag is the tag which, when included in a discord message,
	// will result in a notification being sent to all SysAdmins so they
	// can be notified of the help message.
	GuildID     string `json:"guildID" deepcopier:"field:GuildID"`
	SysAdminTag string `json:"sysAdminTag" deepcopier:"field:SysAdminTag"`

	// LogFiles dictate where our logs are stored
	LogFiles *LogFiles `json:"logFiles" deepcopier:"field:LogFiles"`

	// Defines which roles can execute commands (if applicable)
	Permissions map[string][]string `json:"permissions" deepcopier:"field:Permissions"`
}

// LogFiles dictate the files/paths of the log files
type LogFiles struct {
	InfoLog  string `json:"info_log" deepcopier:"field:InfoLog"`
	ErrorLog string `json:"error_log" deepcopier:"field:ErrorLog"`
}

// String prints a string representation of the config
func (c Config) String() string {
	conf, _ := json.MarshalIndent(c, "", "  ")
	return fmt.Sprintf("%s", conf)
}

func init() {
	conf = &Config{
		Prefix:        "!",
		Token:         "warn",
		HelpChannelID: "warn",
		BotHostName:   "0.0.0.0:4201",
		GuildID:       "291573897730588684",
		SysAdminTag:   "<@&318907623476822016>",
		LogFiles: &LogFiles{
			InfoLog:  "info.log",
			ErrorLog: "error.log",
		},
		Permissions: map[string][]string{
			"alias": []string{
				"Chairperson",
				"Equipments Officer",
				"Events Officer",
				"Finance Officer",
				"HLM",
				"PRO",
				"Secretary",
				"SysAdmin",
			},
			"config": []string{
				"SysAdmin",
				"HLM",
			},
		},
	}
}

// LoadConfig loads the configuration information found in ./config.json
func LoadConfig() error {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		return fmt.Errorf("Failed to read configuration file: %#v", err)
	}

	if len(file) < 1 {
		return errors.New("Configuration file 'config.json' was empty")
	}

	tmpconf := &Config{}
	if err := json.Unmarshal(file, tmpconf); err != nil {
		return fmt.Errorf("Failed to unmarshal configuration JSON: %s", err)
	}

	if err := mergo.MergeWithOverwrite(conf, tmpconf); err != nil {
		return fmt.Errorf("Failed to merge configuration values: %#v", err)
	}

	return nil
}

// GetConfig gets the loaded configuration
func GetConfig() *Config {
	return conf
}
