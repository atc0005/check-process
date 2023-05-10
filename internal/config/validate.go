// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/check-process
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

import (
	"fmt"

	"github.com/atc0005/check-process/internal/textutils"
)

// validate verifies all Config struct fields have been provided acceptable
// values.
func (c Config) validate(_ AppType) error {

	// Validate the specified logging level
	supportedLogLevels := supportedLogLevels()
	if !textutils.InList(c.LoggingLevel, supportedLogLevels, true) {
		return fmt.Errorf(
			"%w: invalid logging level;"+
				" got %v, expected one of %v",
			ErrUnsupportedOption,
			c.LoggingLevel,
			supportedLogLevels,
		)
	}

	// 	switch {
	// 	case appType.Inspector:
	//
	// 	case appType.Plugin:
	//
	// 	}

	// Optimist
	return nil

}
