# Carbonaut

![carbonaut-banner](./assets/Carbonaut_Banner.png)

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/carbonaut-cloud/carbonaut.svg)](https://github.com/carbonaut-cloud/carbonaut)
[![Go Report Card](https://goreportcard.com/badge/carbonaut-cloud/carbonaut)](https://goreportcard.com/report/carbonaut-cloud/carbonaut)
[![Coverage Status](https://coveralls.io/repos/github/carbonaut-cloud/carbonaut/badge.svg?branch=main)](https://coveralls.io/github/carbonaut-cloud/carbonaut?branch=main)
[![Slack](https://img.shields.io/badge/Slack-%23general-blueviolet)](https://carbonautgroup.slack.com/archives/C03B9P2T3AB)
[![CircleCI](https://circleci.com/gh/carbonaut-cloud/carbonaut/tree/main.svg?style=svg)](https://circleci.com/gh/carbonaut-cloud/carbonaut/tree/main)

Carbonaut is a open source tool to measure your carbon emissions, analyze your resource consumptions and support you in optimizing your green house gas footprint.

Carbonaut targets any ICT infrastructure, also in the first phases of development public cloud provider and IaaS provider are the main target. The system will also integrate with Kubernetes and other tools on the market which are able to manage and predict resource utilization.

Our target is to provide precises insights which are not based on estimations (where possible).

## Motivation
Carbonaut is not targeting for the easy 5% cloud workload, but for the full 100%. This will be a long and tough journey, but only then we are able to get a clear picture of the ICT caused carbon emissions. And we expect to see in future further KPIs like water usage, heat exchange and others to gain in relevance. 

## Feature Milestones

Our nearest features to be provided are:

- [ ] Integrate with major CSP
    - [ ] AWS
    - [ ] GCP
    - [ ] Azure
- [ ] Provide a Grafana dashboard for GHG and energy consumption monitoring
- [ ] Support enterprise structures (Products, Projects, Departments, ...)
- [ ] Provide KPIs for carbon emissions and custom KPIs
- [ ] Data exporter/importer
- [ ] GHG optimization recommendation
- [ ] API to be able to integrated with other tools

However, to fulfill our mission we also want to:
- [ ] Integrate with Kubernetes HPA/VPA/ClusterAutoscaler
- [ ] Integrate with IaC tools to provide forecasts
- [ ] Integration with hypervisor platforms
- [ ] Integrate with baremetal meassurements
- [ ] Integrate with different power supply and energy mix datasets

## Getting Started
LINK To DOCS

## Contribution
If you like to contribute, get in contact with us and follow [this short guide](https://github.com/carbonaut-cloud/community/blob/main/CONTRIBUTING.md).

## Code of Conduct
In order to contribute to Carbonaut you agree to follow our [Code of Conduct](https://github.com/carbonaut-cloud/community/blob/main/CODE_OF_CONDUCT). The code of conduct applies at any place where the contributors of Carbonaut interact with each other, virtually or in person.

## Development

To develop `Carbonaut` you need to install [golang](https://golang.org/doc/install).

**To verify the code**
```bash
make verify
```

Feel free to jump on unassigned issues and open up PRs. If you have any questions feel free to reach out over [Slack](https://join.slack.com/t/carbonautgroup/shared_invite/zt-17d78zd8j-qa0KvMS_be21E3hH9fpdYQ).