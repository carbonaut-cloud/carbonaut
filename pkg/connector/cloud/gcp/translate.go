// Copyright 2022 The Carbonaut Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gcp

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"carbonaut.cloud/carbonaut/pkg/data/models"
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

// TODO: translate gcp data to unified carbonaut data schema (implement spec connector/translator.go)

type translateLookup struct {
	Emissions *models.Emissions
	Error     error
}

func translateRecords(r []*Record) ([]*models.Emissions, error) {
	// Worker
	requestData := func(done <-chan interface{}, records []*Record) <-chan translateLookup {
		lookups := make(chan translateLookup)
		go func() {
			defer close(lookups)
			for i := range records {
				r, err := translateRecord(records[i])
				select {
				case <-done:
					return
				case lookups <- translateLookup{
					Emissions: r,
					Error:     err,
				}:
				}
			}
		}()
		return lookups
	}

	done := make(chan interface{})
	defer close(done)
	e := []*models.Emissions{}
	var err error

	// Collect data from buffered channel
	for lookups := range requestData(done, r) {
		if lookups.Error != nil {
			err = multierror.Append(err, errors.Wrapf(lookups.Error, "error translating record %v", lookups.Emissions))
		} else {
			e = append(e, lookups.Emissions)
		}
	}
	return e, err
}

func translateRecord(r *Record) (*models.Emissions, error) {
	providerVersion, err := strconv.Atoi(r.CarbonModelVersion)
	if err != nil {
		return nil, fmt.Errorf("could not translate gcp record, invalid provider version: %s", r.CarbonModelVersion)
	}
	mtc02e, err := strconv.ParseFloat(r.CarbonFootprintKgCO2e, 64)
	if err != nil {
		return nil, fmt.Errorf("could not translate gcp record, invalid carbon footprint kgco2e: %s", r.CarbonFootprintKgCO2e)
	}
	return &models.Emissions{
		ResourceName:    r.ServiceDescription,
		ResourceOwner:   r.ProjectNumber,
		ResourceType:    "",
		Provider:        string(Provider.Name),
		MTCO2e:          mtc02e,
		Location:        r.LocationLocation,
		ProviderVersion: providerVersion,
		Timestamp:       dateFormat(r.UsageMonth),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}, nil
}

const (
	timestampRegexLayout = `(\d{2})\/(\d{2})\/(\d{4})`
	// group 0: entire timestamp -> "DD/MM/YYYY"
	timestampDayGroup   = 1
	timestampMonthGroup = 2
	timestampYearGroup  = 3
)

var (
	timestampRegex = regexp.MustCompile(timestampRegexLayout)
	emptyDate      = time.Date(0, 0, 0, 0, 0, 0, 0, &time.Location{})
)

func dateFormat(d string) time.Time {
	t := timestampRegex.FindStringSubmatch(d)
	// it should match to this -> [FullTimestamp, Day, Month, Year]
	if len(t) != 4 {
		return emptyDate
	}
	year, err := strconv.Atoi(t[timestampYearGroup])
	if err != nil || year <= 0 {
		return emptyDate
	}
	month, err := strconv.Atoi(t[timestampMonthGroup])
	if err != nil || month <= 0 {
		return emptyDate
	}
	day, err := strconv.Atoi(t[timestampDayGroup])
	if err != nil || day <= 0 {
		return emptyDate
	}
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, &time.Location{})
}
