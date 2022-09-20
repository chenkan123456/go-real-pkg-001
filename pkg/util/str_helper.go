package util

import (
	"encoding/json"
	"fmt"
	"math"
	"runtime"
	"strconv"
	"strings"
)

func NumRoundToString(x float64) string {
	var num = int(math.Floor(x + 0.5))
	return fmt.Sprint(num)
}

func ValidationArrByString(k string) []string {
	arr := make([]string, 0)
	if k != "" {
		arr = strings.Split(k, ",")
	}
	return arr
}

func ConvertFloatByString(k string) float64 {
	value, err := strconv.ParseFloat(k, 64)
	if err != nil {
		return 0
	}
	return value
}

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

func StringDecimal(value string) string {
	pf, err := strconv.ParseFloat(value, 64)
	if err != nil || pf == 0 {
		return "0.00"
	} else {
		return fmt.Sprintf("%.2f", pf)
	}
}

//集合对比（重合）
func SliceCoincide(s_data []string, n_data []string) []string {

	for _, v := range n_data {
		var notSave = false
		for _, vv := range s_data {

			if v == vv {
				notSave = true
				break

			}
		}
		if !notSave {
			s_data = append(s_data, v)
		}
	}
	return s_data

}

//集合对比（交集）
func SliceIntersection(s_data []string, n_data []string) []string {

	if len(s_data) == 0 || len(n_data) == 0 {
		return SliceCoincide(s_data, n_data)
	} else if (len(s_data) == 1 && s_data[0] == "") || (len(n_data) == 1 && n_data[0] == "") {
		return []string{""}
	} else {
		var data []string
		for _, v := range n_data {
			var notSave = false
			for _, vv := range s_data {

				if v == vv {
					notSave = true
					break

				}
			}
			if notSave {
				data = append(data, v)
			}
		}
		return data
	}
}

func RemoveRep(slc []string) []string {
	if len(slc) < 1024 {
		// 切片长度小于1024的时候，循环来过滤
		return removeRepByLoop(slc)
	} else {
		// 大于的时候，通过map来过滤
		return removeRepByMap(slc)
	}
}

// 通过两重循环过滤重复元素
func removeRepByLoop(slc []string) []string {
	result := []string{} // 存放结果
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			result = append(result, slc[i])
		}
	}
	return result
}

// 通过map主键唯一的特性过滤重复元素
func removeRepByMap(slc []string) []string {
	result := []string{}
	tempMap := map[string]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}

// 数组排序
func SliceASC(arr []string) []string {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			a := ConvertFloatByString(arr[j])
			b := ConvertFloatByString(arr[i])
			if a < b {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}

	}
	return arr

}

func SliceDSEC(arr []string) []string {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			a := ConvertFloatByString(arr[j])
			b := ConvertFloatByString(arr[i])
			if a > b {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}

	}
	return arr

}

func StringDateTimeFormat(time string) string {
	str := time[0:10]
	str = strings.ReplaceAll(str, "/", "-")
	return str

}
func StringTrimRight(s, k string) string {
	if s == "" {
		return s
	} else {
		return strings.TrimRight(s, k)
	}

}

func GoRuntineID() int {

	var buf [64]byte

	n := runtime.Stack(buf[:], false)

	// 得到id字符串

	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]

	id, err := strconv.Atoi(idField)

	if err != nil {

		panic(fmt.Sprintf("cannot get goroutine id: %v", err))

	}

	return id

}

func Strval(value interface{}) string {
	var key string
	if value == nil {
		return key
	}
	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}
	return key
}
