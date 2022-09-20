package util

import "math"

type Point struct {
	Lng float64 `json:"lng"` //经度(180°W-180°E) -180 - 0 - 180
	Lat float64 `json:"lat"` //纬度（90°S-90°N）-90 -0 -90
}

type Area struct {
	Points []Point //多边形坐标 顺时针方向
}

//实例化一个面区域
func NewArea(points []Point) *Area {

	return &Area{
		Points: points,
	}
}

//判断一个点是否在一个面内
//如果点位于多边形的顶点或边上，也算做点在多边形内，直接返回true
func (a *Area) PointInArea(point Point) bool {
	pointNum := len(a.Points) //点个数
	intersectCount := 0       //cross points count of x
	precision := 2e-10        //浮点类型计算时候与0比较时候的容差
	p1 := Point{}             //neighbour bound vertices
	p2 := Point{}
	p := point //测试点

	p1 = a.Points[0] //left vertex
	for i := 0; i < pointNum; i++ {
		if p.Lng == p1.Lng && p.Lat == p1.Lat {
			return true
		}
		p2 = a.Points[i%pointNum]
		if p.Lat < math.Min(p1.Lat, p2.Lat) || p.Lat > math.Max(p1.Lat, p2.Lat) {
			p1 = p2
			continue //next ray left point
		}

		if p.Lat > math.Min(p1.Lat, p2.Lat) && p.Lat < math.Max(p1.Lat, p2.Lat) {
			if p.Lng <= math.Max(p1.Lng, p2.Lng) { //x is before of ray
				if p1.Lat == p2.Lat && p.Lng >= math.Min(p1.Lng, p2.Lng) {
					return true
				}

				if p1.Lng == p2.Lng { //ray is vertical
					if p1.Lng == p.Lng { //overlies on a vertical ray
						return true
					} else { //before ray
						intersectCount++
					}
				} else { //cross point on the left side
					xinters := (p.Lat-p1.Lat)*(p2.Lng-p1.Lng)/(p2.Lat-p1.Lat) + p1.Lng
					if math.Abs(p.Lng-xinters) < precision {
						return true
					}

					if p.Lng < xinters { //before ray
						intersectCount++
					}
				}
			}
		} else { //special case when ray is crossing through the vertex
			if p.Lat == p2.Lat && p.Lng <= p2.Lng { //p crossing over p2
				p3 := a.Points[(i+1)%pointNum]
				if p.Lat >= math.Min(p1.Lat, p3.Lat) && p.Lat <= math.Max(p1.Lat, p3.Lat) {
					intersectCount++
				} else {
					intersectCount += 2
				}
			}
		}
		p1 = p2 //next ray left point
	}
	if intersectCount%2 == 0 { //偶数在多边形外
		return false
	} else { //奇数在多边形内
		return true
	}
}
