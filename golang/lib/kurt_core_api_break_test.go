package lib

import (
	"fmt"
	"github.com/blang/semver"
	"github.com/kurtosis-tech/kurtosis-core-sdk/api/golang/kurtosis_core_version"
	"github.com/stretchr/testify/require"
	"testing"
)

// NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE
// This test does NOT have a Typescript equivalent, because we only need one test to remind the user and
//  because it means we don't have to implement semver-parsing in Typescript for the same functionality
// NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE

const (
	// !!!!!! BEFORE YOU UPDATE THIS CONSTANT TO FIX THE TEST, ADD A "BREAKING CHANGE" SECTION IN THE CHANGELOG !!!!!!
	// Explanation:
	//  * A breaking API change in Kurt Core must also yield a breaking API change in the module lib, since the
	//    the user will use the KurtosisContext (transitively depended on by the module lib) in their module
	//  * However, we don't want to rely on Kurtosis dev memory to add a "Breaking Change" section to the module
	//    lib's changelog whenever a breaking Core API version change happens
	//  * Therefore, this constant must be manually updated to the X.Y version of the Core version you just
	//    bumped to, which will remind you to update the "Breaking Change" section of the changelog.
	expectedCoreMajorMinorVersion = "0.49"
	// !!!!!! BEFORE YOU UPDATE THIS CONSTANT TO FIX THE TEST, ADD A "BREAKING CHANGE" SECTION IN THE CHANGELOG !!!!!!
)

// This test ensures that when you bump to a Kurt Core version that has an API break, you're reminded to add a "Breaking Changes"
//
//	entry to the engine server's changelog as well (since a Kurt Core API break is an engine server API break)
func TestYouHaveBeenRemindedToAddABreakingChangelogEntryOnKurtCoreAPIBreak(t *testing.T) {
	actualKurtCoreSemver, err := semver.Parse(kurtosis_core_version.KurtosisCoreVersion)
	require.NoError(t, err, "An unexpected error occurred parsing Kurt Core version string '%v'", kurtosis_core_version.KurtosisCoreVersion)
	actualCoreMajorMinorVersion := fmt.Sprintf("%v.%v", actualKurtCoreSemver.Major, actualKurtCoreSemver.Minor)
	require.Equal(t, expectedCoreMajorMinorVersion, actualCoreMajorMinorVersion)
}
