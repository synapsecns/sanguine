package metrics

import "gopkg.in/DataDog/dd-trace-go.v1/profiler"

// AllProfilesString exports allProfileString for testing.
func AllProfilesString() (res string) {
	return allProfilesString()
}

// AllProfileTypes epxorts allPRofileTypes for testing.
func AllProfileTypes() map[string]profiler.ProfileType {
	return allProfileTypes
}

// GetProfileTypesFromEnv returns all profiles.
func GetProfileTypesFromEnv() []profiler.ProfileType {
	return getProfileTypesFromEnv()
}
