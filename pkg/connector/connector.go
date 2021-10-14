package connector

// Connector is the interface to implement to fetch the
// changelog from a provider.
type Connector interface {
	// FetchChangelog gathers the changelog from a specific software name with a specific
	// version.
	// It returns the content of the changelog and the source URL.
	FetchChangelog(string, string) (string, string, error)
}
