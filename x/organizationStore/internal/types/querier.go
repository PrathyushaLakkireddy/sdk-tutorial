package types

import "strings"

// QueryResOrgs Queries Result Payload for a orgs query
type QueryResOrgs []string

// implement fmt.Stringer
func (n QueryResOrgs) String() string {
	return strings.Join(n[:], "\n")
}