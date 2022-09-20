package unittest

import (
	"fmt"
	"testing"
	"time"

	"gitlab.centaline-sc.com/real-db-project/go-pkg/pkg/util"
)

func TestDate(t *testing.T) {
	date := "2022-05-03"
	d := util.ReturnLastDateOfMonth(date)
	fmt.Print(d)
}

//集合对比（交集）
func SliceIntersection(s_data []string, n_data []string) []string {

	if len(s_data) == 0 || len(n_data) == 0 {
		return util.SliceCoincide(s_data, n_data)
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

func TestSilce(t *testing.T) {

	var a = []string{""}
	var b = []string{"3", "4", "5", "6", "7", "8", "9", "10"}

	// fmt.Print(util.SliceCoincide(a, b))
	// fmt.Println()
	fmt.Print(len(SliceIntersection(a, b)))
}

func TestDate2(t *testing.T) {

	now := time.Now()

	fmt.Print(now.AddDate(-1, 0, 0))
}

func TestDate3(t *testing.T) {

	//now := time.Now()
	now, _ := time.Parse("2006-01-02", "2022-05-09")
	fmt.Print(util.ReturnNowMonthkRange(now))

}

func TestDate4(t *testing.T) {
	now := "2022-08-02"
	//now := time.Now()
	//now, _ := time.Parse("2006-01-02", "2022-08-09")
	fmt.Print(util.ReturnDateMonthText(now))

}

func TestMap(t *testing.T) {
	var ms = make([]map[string]int, 0)
	var m = make(map[string]int, 0)
	m["1"] = 1
	m["2"] = 2
	ms = append(ms, m)
	fmt.Print(ms[0])
	fmt.Print(m["2"])
	fmt.Print(m["1"])

}

func TestRound(t *testing.T) {
	var x = 161872381.6
	fmt.Print(util.NumRoundToString(x))

}

func TestTimeWeek(t *testing.T) {

	var d = util.StringToShortDate("2022-04-26")
	fmt.Print(util.ReturnLastWeekRangeByYear(d, 2))
	fmt.Println()

}
