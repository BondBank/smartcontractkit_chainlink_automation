// Package gethwrappers provides tools for wrapping solidity contracts with
// golang packages, using abigen.
package gethwrappers

// OCR2VRF
//go:generate go run ../generation/generate/wrap.go ../../../contracts/solc/v0.8.15/DKG.abi ../../../contracts/solc/v0.8.15/DKG.bin DKG dkg
//go:generate go run ../generation/generate/wrap.go ../../../contracts/solc/v0.8.15/VRFCoordinator.abi ../../../contracts/solc/v0.8.15/VRFCoordinator.bin VRFCoordinator vrf_coordinator
//go:generate go run ../generation/generate/wrap.go ../../../contracts/solc/v0.8.15/VRFBeacon.abi ../../../contracts/solc/v0.8.15/VRFBeacon.bin VRFBeacon vrf_beacon
//go:generate go run ../generation/generate/wrap.go ../../../contracts/solc/v0.8.15/BeaconVRFConsumer.abi ../../../contracts/solc/v0.8.15/BeaconVRFConsumer.bin BeaconVRFConsumer vrf_beacon_consumer
//go:generate go run ../generation/generate/wrap.go ../../../contracts/solc/v0.8.15/LoadTestBeaconVRFConsumer.abi ../../../contracts/solc/v0.8.15/LoadTestBeaconVRFConsumer.bin LoadTestBeaconVRFConsumer load_test_beacon_consumer

// Arbitrum specific
//go:generate go run ../generation/generate/wrap.go ../../../contracts/solc/v0.8.15/ArbitrumDKG.abi ../../../contracts/solc/v0.8.15/ArbitrumDKG.bin ArbitrumDKG arbitrum_dkg
//go:generate go run ../generation/generate/wrap.go ../../../contracts/solc/v0.8.15/ArbitrumVRFCoordinator.abi ../../../contracts/solc/v0.8.15/ArbitrumVRFCoordinator.bin ArbitrumVRFCoordinator arbitrum_vrf_coordinator
//go:generate go run ../generation/generate/wrap.go ../../../contracts/solc/v0.8.15/ArbitrumVRFBeacon.abi ../../../contracts/solc/v0.8.15/ArbitrumVRFBeacon.bin ArbitrumVRFBeacon arbitrum_vrf_beacon
//go:generate go run ../generation/generate/wrap.go ../../../contracts/solc/v0.8.15/ArbitrumTestBeaconVRFConsumer.abi ../../../contracts/solc/v0.8.15/ArbitrumTestBeaconVRFConsumer.bin ArbitrumBeaconVRFConsumer arbitrum_beacon_vrf_consumer

// Mock generation
//go:generate mockery --srcpkg github.com/smartcontractkit/chainlink/core/gethwrappers/ocr2vrf/generated/vrf_coordinator --name VRFCoordinatorInterface --output ../../services/ocr2/plugins/ocr2vrf/coordinator/mocks/ --case=underscore --structname VRFCoordinatorInterface --filename vrf_coordinator.go
//go:generate mockery --srcpkg github.com/smartcontractkit/chainlink/core/gethwrappers/ocr2vrf/generated/vrf_beacon --name VRFBeaconInterface --output ../../services/ocr2/plugins/ocr2vrf/coordinator/mocks/ --case=underscore --structname VRFBeaconInterface --filename vrf_beacon.go
