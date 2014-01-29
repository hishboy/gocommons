//
//  kmeans_cluster.go
//
//  Created by Hicham Bouabdallah
//  Copyright (c) 2012 SimpleRocket LLC
//
//  Permission is hereby granted, free of charge, to any person
//  obtaining a copy of this software and associated documentation
//  files (the "Software"), to deal in the Software without
//  restriction, including without limitation the rights to use,
//  copy, modify, merge, publish, distribute, sublicense, and/or sell
//  copies of the Software, and to permit persons to whom the
//  Software is furnished to do so, subject to the following
//  conditions:
//
//  The above copyright notice and this permission notice shall be
//  included in all copies or substantial portions of the Software.
//
//  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
//  EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
//  OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
//  NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
//  HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
//  WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
//  FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
//  OTHER DEALINGS IN THE SOFTWARE.
//

package support

import "github.com/hishboy/gocommons/lang"

type KMeansCluster struct {
	center *KMeansPoint
	points *lang.ArrayList
}

func NewKMeansCluster(center *KMeansPoint) *KMeansCluster {
	self := &KMeansCluster{}
	self.center = center
	self.points = lang.NewArrayList()
	return self
}

func (self *KMeansCluster) Points() *lang.ArrayList {
	return self.points
}

func (self *KMeansCluster) Center() *KMeansPoint {
	return self.center
}

func (self *KMeansCluster) Recenter() float64 {
	totalPoint := self.points.Len()
	firstPoint := self.points.Get(0).(*KMeansPoint)
	totalCoordinates := firstPoint.items.Len()
	
	// sum up the points
	totals := make([]float64, totalCoordinates)
	for i := 0; i < totalPoint; i++ {
		point := self.points.Get(i).(*KMeansPoint)
		for j := 0; j < totalCoordinates; j++ {
			totals[j] = totals[j] + point.items.Get(j).(float64)
		}
	}
	
	// average out
	averages := make([]float64, totalCoordinates)
	for i := 0; i < totalCoordinates; i++ {
		averages[i] = totals[i] / float64(totalPoint)
	}
	
	// calculate the distance between old and new center
	newCenter := NewKMeansPoint(averages)
	distance := self.center.DistanceFromPoint(newCenter)
	self.center = newCenter
	
	return distance
}