package sets

type sorted_set[T comparable] struct {
	sorted []T
	set    Set[T]
}


