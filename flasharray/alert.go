// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

import (
	"fmt"
)

// AlertService is a struct for the alert endpoints
type AlertService struct {
	client *Client
}

// ListAlerts Lists the email recipients that are designated to recieve Purity alert messages
func (a *AlertService) ListAlerts(params map[string]string) ([]Alert, error) {

	req, err := a.client.NewRequest("GET", "alert", params, nil)
	if err != nil {
		return nil, err
	}

	m := []Alert{}
	if _, err = a.client.Do(req, m, false); err != nil {
		return nil, err
	}

	return m, err
}

// GetAlert Lists the information about the specified email recipient
func (a *AlertService) GetAlert(name string) (*Alert, error) {

	path := fmt.Sprintf("alert/%s", name)
	req, err := a.client.NewRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	m := &Alert{}
	if _, err = a.client.Do(req, m, false); err != nil {
		return nil, err
	}

	return m, err
}

// CreateAlert Designates and valid email address to receive Purity alert messages
// Up to 20 addresses can be designated in an array.
func (a *AlertService) CreateAlert(alert string, data interface{}) (*Alert, error) {

	path := fmt.Sprintf("alert/%s", alert)
	req, err := a.client.NewRequest("POST", path, nil, data)
	if err != nil {
		return nil, err
	}

	m := &Alert{}
	if _, err = a.client.Do(req, m, false); err != nil {
		return nil, err
	}

	return m, err
}

// TestAlert Tests the ability of the array to send alert messages to all of the designated email addresses.
func (a *AlertService) TestAlert() (*Alert, error) {

	req, err := a.client.NewRequest("PUT", "alert", nil, nil)
	if err != nil {
		return nil, err
	}

	m := &Alert{}
	if _, err = a.client.Do(req, m, false); err != nil {
		return nil, err
	}

	return m, err
}

// SetAlert Modifies a alert
func (a *AlertService) SetAlert(alert string, data interface{}) (*Alert, error) {

	path := fmt.Sprintf("alert/%s", alert)
	req, err := a.client.NewRequest("PUT", path, nil, data)
	if err != nil {
		return nil, err
	}

	m := &Alert{}
	if _, err = a.client.Do(req, m, false); err != nil {
		return nil, err
	}

	return m, err
}

// EnableAlert enable the transmission of alert messages to the specified email address
func (a *AlertService) EnableAlert(address string) (*Alert, error) {

	data := map[string]bool{"enabled": true}
	m, err := a.SetAlert(address, data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// DisableAlert disable the transmission of alert messages to the specified email address
func (a *AlertService) DisableAlert(address string) (*Alert, error) {

	data := map[string]bool{"enabled": false}
	m, err := a.SetAlert(address, data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// DeleteAlert deletes a alert
func (a *AlertService) DeleteAlert(address string) (*Alert, error) {

	path := fmt.Sprintf("alert/%s", address)
	req, err := a.client.NewRequest("DELETE", path, nil, nil)
	if err != nil {
		return nil, err
	}

	m := &Alert{}
	if _, err = a.client.Do(req, m, false); err != nil {
		return nil, err
	}

	return m, err
}
