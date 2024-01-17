package datastructures

type Index[T any] map[string]T

func NewIndex[T any](items []T, getID func(T) string) Index[T] {
	idx := make(Index[T], len(items))
	for _, item := range items {
		idx[getID(item)] = item
	}
	return idx
}

func Keys[T any](idx Index[T]) []string {
	keys := make([]string, len(idx))
	i := 0
	for k := range idx {
		keys[i] = k
		i++
	}
	return keys
}

func AnyExists[T any](idx Index[T], keys ...string) bool {
	for _, id := range keys {
		if _, ok := idx[id]; ok {
			return true
		}
	}
	return false
}

type DiffResult struct {
	Added   []string
	Deleted []string
	Rest    []string
}

func Diff[C, P any](current Index[C], previous Index[P]) DiffResult {
	processed := make(map[string]struct{})
	var added, deleted, rest []string

	for k := range current {
		if _, ok := previous[k]; !ok {
			added = append(added, k)
		} else {
			processed[k] = struct{}{}
		}
	}
	for k := range previous {
		if _, ok := current[k]; !ok {
			deleted = append(deleted, k)
		} else {
			processed[k] = struct{}{}
		}
	}
	for k := range processed {
		rest = append(rest, k)
	}

	return DiffResult{
		Added:   added,
		Deleted: deleted,
		Rest:    rest,
	}
}
