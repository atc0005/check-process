// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/check-process
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package process

import (
	"fmt"
	"strconv"
)

// setProcessProps coordinates parsing select process property string values
// as more specific type values. These values are used to set "primary" fields
// for direct access.
func (p *Process) setProcessProps() error {

	// Name, State, Pid, PPid, Threads, VMSwap
	const requiredPropsIndexCount = 6

	if len(p.AllProperties) < requiredPropsIndexCount {
		return fmt.Errorf(
			"error setting process field values: %w",
			ErrInvalidProcessPropertiesIndexCount,
		)
	}

	// 	type setter func() error
	// 	setters := []setter{
	// 		p.setNameField,
	// 		p.setStateField,
	// 		p.setPidField,
	// 		p.setPPidField,
	// 		p.setThreadsField,
	// 		p.setVMSwapField,
	// 	}
	//
	// 	for _, setter := range setters{
	// 		if err := setter(); err != nil {
	// 			return err
	// 		}
	// 	}

	if err := p.setNameField(); err != nil {
		return err
	}

	if err := p.setStateField(); err != nil {
		return err
	}

	if err := p.setPidField(); err != nil {
		return err
	}

	if err := p.setPPidField(); err != nil {
		return err
	}

	if err := p.setThreadsField(); err != nil {
		return err
	}

	if err := p.setVMSwapField(); err != nil {
		return err
	}

	return nil
}

// setNameField retrieves the process property name value and sets the
// "primary" Name field for direct access.
func (p *Process) setNameField() error {
	name, ok := p.AllProperties[ProcessNameField]
	if !ok {
		return fmt.Errorf(
			"process property %s: %w",
			ProcessNameField,
			ErrMissingProcessPropertiesIndexEntry,
		)
	}
	p.Name = name

	return nil
}

// setStateField retrieves the process property state value and sets the
// "primary" State field for direct access.
func (p *Process) setStateField() error {
	state, ok := p.AllProperties[ProcessStateField]
	if !ok {
		return fmt.Errorf(
			"process property %s: %w",
			ProcessStateField,
			ErrMissingProcessPropertiesIndexEntry,
		)
	}
	p.State = state

	return nil
}

// setPidField retrieves the process property pid value and sets the
// "primary" Pid field for direct access.
func (p *Process) setPidField() error {
	pidStr, ok := p.AllProperties[ProcessPidField]
	if !ok {
		return fmt.Errorf(
			"process property %s: %w",
			ProcessPidField,
			ErrMissingProcessPropertiesIndexEntry,
		)
	}
	n, err := strconv.Atoi(pidStr)
	if err != nil {
		return fmt.Errorf(
			"failed to convert value %s for property %s as number: %w",
			pidStr,
			ProcessPidField,
			ErrInvalidProcStatusLineValue,
		)
	}
	p.Pid = n

	return nil
}

// setPPidField retrieves the process property ppid value and sets the
// "primary" PPid field for direct access.
func (p *Process) setPPidField() error {
	ppidStr, ok := p.AllProperties[ProcessPPidField]
	if !ok {
		return fmt.Errorf(
			"process property %s: %w",
			ProcessPPidField,
			ErrMissingProcessPropertiesIndexEntry,
		)
	}
	n, err := strconv.Atoi(ppidStr)
	if err != nil {
		return fmt.Errorf(
			"failed to convert value %s for property %s as number: %w",
			ppidStr,
			ProcessPPidField,
			ErrInvalidProcStatusLineValue,
		)
	}
	p.PPid = n

	return nil
}

// setThreadsField retrieves the process property threads value and sets the
// "primary" Threads field for direct access.
func (p *Process) setThreadsField() error {
	threadsStr, ok := p.AllProperties[ProcessThreadsField]
	if !ok {
		return fmt.Errorf(
			"process property %s: %w",
			ProcessThreadsField,
			ErrMissingProcessPropertiesIndexEntry,
		)
	}
	n, err := strconv.Atoi(threadsStr)
	if err != nil {
		return fmt.Errorf(
			"failed to convert value %s for property %s as number: %w",
			threadsStr,
			ProcessThreadsField,
			ErrInvalidProcStatusLineValue,
		)
	}
	p.Threads = n

	return nil
}

// setVMSwapField retrieves the process property vmswap value and sets the
// "primary" VMSwap field for direct access.
//
// NOTE: This value only appears to be set if a process is actually using swap
// memory. Because this value may not be available we use it if it is,
// otherwise we do *not* return an error.
func (p *Process) setVMSwapField() error {
	vmswap, ok := p.AllProperties[ProcessVMSwapField]
	// if !ok {
	// 	return fmt.Errorf(
	// 		"process property %s: %w",
	// 		ProcessVMSwapField,
	// 		ErrMissingProcessPropertiesIndexEntry,
	// 	)
	// }
	if ok {
		p.VMSwap = vmswap
	}

	return nil
}
