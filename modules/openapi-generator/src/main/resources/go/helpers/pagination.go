package helpers

// PaginatedList is a generic pagination wrapper that replaces 100+ specific types
// like PaginatedDeviceList, PaginatedIPAddressList, etc.
//
// Instead of generating:
//   type PaginatedDeviceList struct { Count int32; Results []Device }
//   type PaginatedIPAddressList struct { Count int32; Results []IPAddress }
//   ... (100+ more)
//
// We use:
//   type PaginatedDeviceList = PaginatedList[Device]
//   type PaginatedIPAddressList = PaginatedList[IPAddress]
//   ... (just type aliases)
type PaginatedList[T any] struct {
	Count    int32   `json:"count"`
	Next     *string `json:"next,omitempty"`
	Previous *string `json:"previous,omitempty"`
	Results  []T     `json:"results"`
}

// GetCount returns the total count of items
func (p *PaginatedList[T]) GetCount() int32 {
	return p.Count
}

// GetResults returns the results slice
func (p *PaginatedList[T]) GetResults() []T {
	return p.Results
}

// HasNext returns true if there's a next page
func (p *PaginatedList[T]) HasNext() bool {
	return p.Next != nil && *p.Next != ""
}

// HasPrevious returns true if there's a previous page
func (p *PaginatedList[T]) HasPrevious() bool {
	return p.Previous != nil && *p.Previous != ""
}
