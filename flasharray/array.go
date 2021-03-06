// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

// ArrayService type creates a service to perform functions for administering
// and querying the flash array itself
type ArrayService struct {
	client *Client
}

// set_console_lock is a helper function used to set the console lock
func (v *ArrayService) setConsoleLock(b string) (*ConsoleLock, error) {

	data := map[string]string{"enabled": b}
	req, err := v.client.NewRequest("PUT", "array/console_lock", nil, data)
	if err != nil {
		return nil, err
	}

	m := &ConsoleLock{}
	_, err = v.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// EnableConsoleLock enables root lockout from the array at the physical console.
// returns A dictionary mapping "console_lock" to "enabled".
func (v *ArrayService) EnableConsoleLock() error {

	_, err := v.setConsoleLock("true")
	if err != nil {
		return err
	}
	return nil
}

// DisableConsoleLock disables root lockout from the array at the physical console.
// returns A dictionary mapping "console_lock" to "disabled".
func (v *ArrayService) DisableConsoleLock() error {

	_, err := v.setConsoleLock("false")
	if err != nil {
		return err
	}
	return nil
}

// GetConsoleLock returns an object giving the console_lock status
func (v *ArrayService) GetConsoleLock() (*ConsoleLock, error) {

	req, err := v.client.NewRequest("GET", "array/console_lock", nil, nil)
	if err != nil {
		return nil, err
	}

	m := &ConsoleLock{}
	_, err = v.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Get returns and object describing the flash array
func (v *ArrayService) Get(data interface{}) (*Array, error) {

	req, err := v.client.NewRequest("GET", "array", nil, data)
	if err != nil {
		return nil, err
	}

	m := &Array{}
	_, err = v.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Set will change the parameter on the array that is passed in the data map
func (v *ArrayService) Set(data interface{}) (*Array, error) {

	req, err := v.client.NewRequest("PUT", "array", nil, data)
	if err != nil {
		return nil, err
	}

	m := &Array{}
	_, err = v.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Rename will change the name of the flash array
func (v *ArrayService) Rename(name string) (*Array, error) {

	data := map[string]string{"name": name}
	m, err := v.Set(data)
	return m, err
}

// Set the phonehome service attributes
func (v *ArrayService) setPhoneHome(data interface{}) (*Phonehome, error) {

	req, err := v.client.NewRequest("PUT", "array/phonehome", nil, data)
	if err != nil {
		return nil, err
	}

	m := &Phonehome{}
	_, err = v.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Set the remote assist service attributes
func (v *ArrayService) setRemoteAssist(data interface{}) (*RemoteAssist, error) {

	req, err := v.client.NewRequest("PUT", "array/remoteassist", nil, data)
	if err != nil {
		return nil, err
	}

	m := &RemoteAssist{}
	_, err = v.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// DisablePhoneHome disables hourly phonehome
func (v *ArrayService) DisablePhoneHome() (*Phonehome, error) {

	data := map[string]bool{"enabled": false}
	m, err := v.setPhoneHome(data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// DisableRemoteAssist disables Remote Assist
func (v *ArrayService) DisableRemoteAssist() (*RemoteAssist, error) {

	data := map[string]string{"action": "disconnect"}
	m, err := v.setRemoteAssist(data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// EnablePhoneHome enables hourly phonehome
func (v *ArrayService) EnablePhoneHome() (*Phonehome, error) {

	data := map[string]bool{"enabled": true}
	m, err := v.setPhoneHome(data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// EnableRemoteAssist enables Remote Assist
func (v *ArrayService) EnableRemoteAssist() (*RemoteAssist, error) {

	data := map[string]string{"action": "connect"}
	m, err := v.setRemoteAssist(data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// GetManualPhoneHome lists manual phone home status
func (v *ArrayService) GetManualPhoneHome() (*Phonehome, error) {

	req, err := v.client.NewRequest("GET", "array/phoneome", nil, nil)
	if err != nil {
		return nil, err
	}

	m := &Phonehome{}
	_, err = v.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// GetPhoneHome lists Phonehome status
func (v *ArrayService) GetPhoneHome() (*Array, error) {

	data := map[string]bool{"phonehome": true}
	m, err := v.Get(data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// GetRemoteAssist lists Remote assist status
func (v *ArrayService) GetRemoteAssist() (*RemoteAssist, error) {

	req, err := v.client.NewRequest("GET", "array/remoteassist", nil, nil)
	if err != nil {
		return nil, err
	}

	m := &RemoteAssist{}
	_, err = v.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Phonehome Manually initiates or cancels phonehome
//
// Parameters
// action
// The timeframe of logs to phonehome or cancel the current phonehome
// action must be one of:
// "send_today", "send_yesterday", "send_all", "cancel"
func (v *ArrayService) Phonehome(action string) (*Phonehome, error) {

	data := map[string]string{"action": action}
	m, err := v.setPhoneHome(data)
	if err != nil {
		return nil, err
	}

	return m, err
}
