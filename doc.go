/*
Package ghostinspector provides a basic API wrapper around the Ghost Inspector
API.

Note: This package currently uses interfaces to unmarshal JSON in a lot of
places. Version 2 of this package is going to have complete defined structs in
all locations possbile (Suites, Tests, ...) to provide a better SDK to work
with.

Usage:

	import "github.com/lukevers/ghostinspector"

Construct a new `GhostInspector` client. You will need an API key, which you
can get from your dashboard.

	client := ghostinspector.New("api_key")

Once you have a client, you can do things like find a specific suite or list
all suites you have access to.

	suites, err := client.ListSuites()
	if err != nil {
		// Handle...
	}

	suite, err := client.GetSuite("longsuiteid")
	if err != nil {
		// Handle...
	}
*/
package ghostinspector
