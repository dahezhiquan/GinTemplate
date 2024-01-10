package public

import "strconv"

// CAPTCHA_REDISPREXKEY redis key前缀
const (
	CAPTCHA_REDISPREXKEY = "CAPTCHA:"
)

func AddPrefix(prefix string, v interface{}) string {
	switch v := v.(type) {
	case string:
		return v
	case int:
		return prefix + strconv.Itoa(v)
	case int8:
		return prefix + strconv.FormatInt(int64(v), 10)
	case int16:
		return prefix + strconv.FormatInt(int64(v), 10)
	case int32:
		return prefix + strconv.FormatInt(int64(v), 10)
	case int64:
		return prefix + strconv.FormatInt(v, 10)
	case uint:
		return prefix + strconv.FormatUint(uint64(v), 10)
	case uint8:
		return prefix + strconv.FormatUint(uint64(v), 10)
	case uint16:
		return prefix + strconv.FormatUint(uint64(v), 10)
	case uint32:
		return prefix + strconv.FormatUint(uint64(v), 10)
	case uint64:
		return prefix + strconv.FormatUint(v, 10)
	case float32:
		return prefix + strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return prefix + strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		return prefix + strconv.FormatBool(v)
	default:
		return ""
	}
}

func AddSuffix(v interface{}, suffix string) string {
	switch v := v.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v) + suffix
	case int8:
		return strconv.FormatInt(int64(v), 10) + suffix
	case int16:
		return strconv.FormatInt(int64(v), 10) + suffix
	case int32:
		return strconv.FormatInt(int64(v), 10) + suffix
	case int64:
		return strconv.FormatInt(v, 10) + suffix
	case uint:
		return strconv.FormatUint(uint64(v), 10) + suffix
	case uint8:
		return strconv.FormatUint(uint64(v), 10) + suffix
	case uint16:
		return strconv.FormatUint(uint64(v), 10) + suffix
	case uint32:
		return strconv.FormatUint(uint64(v), 10) + suffix
	case uint64:
		return strconv.FormatUint(v, 10) + suffix
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32) + suffix
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64) + suffix
	case bool:
		return strconv.FormatBool(v) + suffix
	default:
		return ""
	}
}
