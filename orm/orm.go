package orm

import (
	"fmt"
)

func FormatEqString(m map[string]string) map[string]interface{} {
	ret := make(map[string]interface{})
	for k, v := range m {
		if v == "" {
			continue
		}
		ret[fmt.Sprintf("%s = ?", k)] = v
	}
	return ret
}

func FormatEqInt(m map[string]int) map[string]interface{} {
	ret := make(map[string]interface{})
	for k, v := range m {
		if v == 0 {
			continue
		}
		ret[fmt.Sprintf("%s = ?", k)] = v
	}
	return ret
}

func FormatGtString(m map[string]string) map[string]interface{} {
	ret := make(map[string]interface{})
	for k, v := range m {
		if v == "" {
			continue
		}
		ret[fmt.Sprintf("%s > ?", k)] = v
	}
	return ret
}

func FormatGtInt(m map[string]int) map[string]interface{} {
	ret := make(map[string]interface{})
	for k, v := range m {
		if v == 0 {
			continue
		}
		ret[fmt.Sprintf("%s > ?", k)] = v
	}
	return ret
}

func FormatLtString(m map[string]string) map[string]interface{} {
	ret := make(map[string]interface{})
	for k, v := range m {
		if v == "" {
			continue
		}
		ret[fmt.Sprintf("%s < ?", k)] = v
	}
	return ret
}

func FormatLtInt(m map[string]int) map[string]interface{} {
	ret := make(map[string]interface{})
	for k, v := range m {
		if v == 0 {
			continue
		}
		ret[fmt.Sprintf("%s < ?", k)] = v
	}
	return ret
}

func FormatGteString(m map[string]string) map[string]interface{} {
	ret := make(map[string]interface{})
	for k, v := range m {
		if v == "" {
			continue
		}
		ret[fmt.Sprintf("%s >= ?", k)] = v
	}
	return ret
}

func FormatGteInt(m map[string]int) map[string]interface{} {
	ret := make(map[string]interface{})
	for k, v := range m {
		if v == 0 {
			continue
		}
		ret[fmt.Sprintf("%s >=", k)] = v
	}
	return ret
}

func FormatLteString(m map[string]string) map[string]interface{} {
	ret := make(map[string]interface{})
	for k, v := range m {
		if v == "" {
			continue
		}
		ret[fmt.Sprintf("%s <=", k)] = v
	}
	return ret
}

func FormatLteInt(m map[string]int) map[string]interface{} {
	ret := make(map[string]interface{})
	for k, v := range m {
		if v == 0 {
			continue
		}
		ret[fmt.Sprintf("%s <=", k)] = v
	}
	return ret
}

func FormatLike(m map[string]string) map[string]interface{} {
	ret := make(map[string]interface{})
	for k, v := range m {
		if v == "" {
			continue
		}
		ret[fmt.Sprintf("%s Like ?", k)] = v
	}
	return ret
}

func FormatInString(m map[string]interface{}) map[string]interface{} {
	ret := make(map[string]interface{})
	for k, v := range m {
		if v == nil {
			continue
		}
		ret[fmt.Sprintf("%s in (?)", k)] = v
	}
	return ret
}

func FormatInInt(m map[string]interface{}) map[string]interface{} {
	ret := make(map[string]interface{})
	for k, v := range m {
		if v == nil {
			continue
		}
		ret[fmt.Sprintf("%s in (?)", k)] = v
	}
	return ret
}

func FormatWhereQuery(intQueryMap map[string]int, stringQueryMap map[string]string, likeQueryMap map[string]string,
	intSliceQueryMap map[string][]int, stringSliceQueryMap map[string][]string) (map[string]int, map[string]string, map[string]string, map[string][]int, map[string][]string) {

	daoIntQueryMap := make(map[string]int)
	for k, v := range intQueryMap {
		if v == 0 {
			continue
		}
		query := fmt.Sprintf("%s = ?", k)
		daoIntQueryMap[query] = v
	}

	daoStringQueryMap := make(map[string]string)
	for k, v := range stringQueryMap {
		if v == "" {
			continue
		}

		query := fmt.Sprintf("%s = ?", k)
		daoStringQueryMap[query] = v
	}

	daoLikequeryMap := make(map[string]string)
	for k, v := range likeQueryMap {
		if v == "" {
			continue
		}
		query := fmt.Sprintf("%s Like ?", k)
		daoLikequeryMap[query] = v
	}

	daoIntSliceQueryMap := make(map[string][]int)
	for k, v := range intSliceQueryMap {
		if v == nil {
			continue
		}
		daoIntSliceQueryMap[fmt.Sprintf("%s in (?)", k)] = v
	}

	daoStringSliceQueryMap := make(map[string][]string)
	for k, v := range stringSliceQueryMap {
		if v == nil {
			continue
		}
		daoStringSliceQueryMap[fmt.Sprintf("%s in (?)", k)] = v
	}

	return daoIntQueryMap, daoStringQueryMap, daoLikequeryMap, daoIntSliceQueryMap, daoStringSliceQueryMap
}
