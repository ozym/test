package main

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/GeoNet/delta/meta"
)

// Read Mark meta information from delta csv files.
func ReadMarks(base string) (map[string]meta.Mark, error) {
	var marks meta.MarkList
	if err := meta.LoadList(filepath.Join(base, "network", "marks.csv"), &marks); err != nil {
		return nil, fmt.Errorf("unable to load delta marks: %v", err)
	}
	res := make(map[string]meta.Mark)
	for _, m := range marks {
		res[m.Code] = m
	}
	return res, nil
}

// Read Deployed Receivers meta information from delta csv files.
func ReadDeployedReceivers(base string) (map[string]string, error) {
	var receivers meta.DeployedReceiverList
	if err := meta.LoadList(filepath.Join(base, "install", "receivers.csv"), &receivers); err != nil {
		return nil, fmt.Errorf("unable to load delta marks: %v", err)
	}

	res := make(map[string]string)
	for _, r := range receivers {
		if t := time.Now().UTC(); t.After(r.End) {
			continue
		}
		res[r.Mark] = r.Model
	}

	return res, nil
}
