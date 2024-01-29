# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres
to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.13] - 2024-01-29

### Added
- Improve push structure
- Subject for clicked messages

## [0.1.12] - 2023-12-20

### Added 
- Flagged field for proposals

## [0.1.11] - 2023-12-05

### Added
- Subject for ens name resolver

## [0.1.10] - 2023-12-04

### Added
- Subject for updated popularity index

## [0.1.8] - 2023-10-13

### Added
- Subject for deleted proposals

## [0.1.7] - 2023-09-12

### Changed
- Replace vote choice from int to json raw message due to multiple choices

## [0.1.5] - 2023-08-23

### Added
- Core timeline subject with payload

## [0.1.4] - 2023-08-14

### Added
- Proposal ends soon subject

## [0.1.3] - 2023-07-25

### Added
- Describe push subject

## [0.1.2] - 2023-07-18

### Changed
- Extend vote model

## [0.1.1] - 2023-07-15

### Fixed
- Added missed const

## [0.1.0] - 2023-07-15

### Added
- Added generic handler
- Added options for creation new consumer

### Changed
- Total refactoring in inbox events

## [0.0.20] - 2023-07-15

### Fixed
- Fixed consuming the exactly required subject

## [0.0.19] - 2023-07-15

### Added
- Added parameter for max ack pending for consumer
- Added hard coded rate limiting for consumer or 3MiB per second

## [0.0.18] - 2023-07-15

### Fixed
- Fixed defining the consumer as durable but without deletion after draining

## [0.0.17] - 2023-07-15

### Fixed
- Disabled unsubscribing from the subject on stopping/closing connection to the jetstream

## [0.0.16] - 2023-07-15

### Fixed
- Fixed strategy.params in core events

## [0.0.15] - 2023-07-14

### Added
- DAO activity_since field

## [0.0.14] - 2023-07-14

### Changed
- Used uuid instead of strings
- Extended feed item structure (timeline)

## [0.0.13] - 2023-07-12

### Changed
- Changed default AckWait for all consumers to 1 minute 

## [0.0.12] - 2023-07-11

### Fixed
- Fixed missed alias field in the DAO object
- Fixed missed field params in the Strategy object 

## [0.0.11] - 2023-07-11

### Fixed
- Fixed linter warnings
- Fixed delivery options for nats consumer and removed unnecessary test

## [0.0.10] - 2023-07-04

### Added
- Added strategy.params field definition

## [0.0.9] - 2023-06-27

### Added
- Added events for inbox

## [0.0.8] - 2023-06-26

### Added
- Inbox dao event
- Inbox proposal event
- Add natsclient publisher

## [0.0.7] - 2023-06-07

### Added
- Callback event

### Changed
- Proposal subject names

## [0.0.6] - 2023-05-23

### Changed
- Core package name

## [0.0.5] - 2023-05-18

### Added
- Vote proposal event

## [0.0.4] - 2023-05-15

### Added
- Register new subjects for internal events

## [0.0.3] - 2023-04-25

### Added
- Core proposal event

## [0.0.2] - 2023-04-21

### Changed
- Split core and aggregator events
- Update code style

## [0.0.1] - 2023-04-17

### Added
- Basic events
- Add layer for message manipulation in the nats jet stream
