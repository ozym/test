package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	formatsFile = "formats.csv"
	modelsFile  = "models.csv"
	usersFile   = "users.csv"
	mountsFile  = "mounts.csv"
)

// BuildConfig extracts information from delta meta data and local csv files.
func BuildConfig(base, input string) (*Config, error) {

	marks, err := ReadMarks(base)
	if err != nil {
		return nil, err
	}

	receivers, err := ReadDeployedReceivers(base)
	if err != nil {
		return nil, err
	}

	formats, err := ReadFormats(filepath.Join(input, formatsFile))
	if err != nil {
		return nil, err
	}

	models, err := ReadModels(filepath.Join(input, modelsFile))
	if err != nil {
		return nil, err
	}

	mounts, err := ReadMounts(filepath.Join(input, mountsFile))
	if err != nil {
		return nil, err
	}

	users, err := ReadUsers(filepath.Join(input, usersFile))
	if err != nil {
		return nil, err
	}

	var config Config

	for _, u := range users {
		config.Users = append(config.Users, UserConfig{
			Username: u.Username,
			Password: u.Password,
		})
	}

	groups := make(map[string][]string)
	for _, u := range users {
		for _, g := range u.Groups {
			groups[g] = append(groups[g], u.Username)
		}
	}
	for k, v := range groups {
		config.Groups = append(config.Groups, GroupConfig{
			Group: k,
			Users: v,
		})
	}

	for _, m := range mounts {
		config.ClientMounts = append(config.ClientMounts, ClientMountConfig{
			Mount:  m.Mount,
			Groups: m.Groups,
		})
	}

	for _, m := range mounts {
		mark, ok := marks[m.Mark]
		if !ok {
			continue
		}

		receiver, ok := receivers[m.Mark]
		if !ok {
			continue
		}

		model, ok := models[receiver]
		if !ok {
			return nil, fmt.Errorf("unknown model: %s", receiver)
		}

		details, ok := formats[m.Details]
		if !ok {
			return nil, fmt.Errorf("unknown format: %s", m.Details)
		}

		config.Mounts = append(config.Mounts, MountConfig{
			Mount:      m.Mount,
			Mark:       mark.Code,
			Name:       mark.Name,
			Country:    m.Country,
			Latitude:   strconv.FormatFloat(mark.Latitude, 'f', 2, 64),
			Longitude:  strconv.FormatFloat(mark.Longitude, 'f', 2, 64),
			Format:     m.Format,
			Details:    strings.Join(details, ","),
			Navigation: m.Navigation,
			Model:      model,
			Address:    m.Address,
		})
	}

	return &config, nil
}
