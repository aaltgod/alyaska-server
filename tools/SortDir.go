package tools

import (
	"os"
	"sort"
)

func SortDir(dir []os.FileInfo) []os.FileInfo {

	var (
		fileInfoIdx = make(map[int]os.FileInfo)
		names       []string
	)

	for i, v := range dir {
		name := v.Name()
		names = append(names, name)
		fileInfoIdx[i] = v
	}

	sort.Slice(names, func(p, q int) bool {
		return names[p] < names[q]
	})

	for _, fii := range fileInfoIdx {
		for idx, v := range names {
			if fii.Name() == v {
				dir[idx] = fii
				break
			}
			continue
		}
	}
	return dir
}
