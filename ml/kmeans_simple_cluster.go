//
//  kmeans_simple_cluster.go
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

package ml
import "github.com/hishboy/gocommons/lang"
import "math"
import "fmt"

type KMeansPoint struct {
	items *lang.ArrayList
}

func NewKMeansPoint(items []float64) *KMeansPoint {
	self := &KMeansPoint{}
	self.items = lang.NewArrayList()
	for i :=0; i < len(items); i++ {
		self.items.Add(items[i])
	}
	return self
}

func (self *KMeansPoint) DistanceFromPoint(otherPoint *KMeansPoint) float64 {
	// FIXME: hicham - should return error if array size doesn't match
	
	total := 0.0 
	for i := 0; i < self.items.Len(); i++ {
		thisCoordinate := self.items.Get(i).(float64)
		otherCoordinate := otherPoint.items.Get(i).(float64)
		total = total + math.Pow(thisCoordinate-otherCoordinate, 2)
	}
	
	return math.Sqrt(total)
}

// kmeans_cluster class

type kmeans_cluster struct {
	center *KMeansPoint
	points *lang.ArrayList
}

func new_kmeans_cluster(center *KMeansPoint) *kmeans_cluster {
	self := &kmeans_cluster{}
	self.center = center
	self.points = lang.NewArrayList()
	return self
}

func (self *kmeans_cluster) recenter() float64 {
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


// NewKMeansSimpleCluster struct

type KMeansSimpleCluster struct {
	points *lang.ArrayList
	numberOfClusters int
	delta float64
}


func NewKMeansSimpleCluster(numberOfClusters int) *KMeansSimpleCluster {
	self := &KMeansSimpleCluster{}
	self.points = lang.NewArrayList()
	self.numberOfClusters = numberOfClusters
	self.delta = 0.001 // default delta
	return self
}


func (self *KMeansSimpleCluster) SetDelta(delta float64) {
	self.delta = delta
}

func (self *KMeansSimpleCluster) AddPoint(point *KMeansPoint) {
	self.points.Add(point)
}

func (self *KMeansSimpleCluster) AddPointAsSlice(items []float64) {
	self.points.Add(NewKMeansPoint(items))
}

func (self *KMeansSimpleCluster) Cluster() {
	// FIXME: hicham - don't allow single cluster
	
	clusters := lang.NewArrayList()
	uniqueCenters := lang.NewHashSet()
	for i := 0; i < self.numberOfClusters; i++ {
		randomCenter := self.points.Sample().(*KMeansPoint)
		for uniqueCenters.Contains(randomCenter) {
			randomCenter = self.points.Sample().(*KMeansPoint)
		}
		uniqueCenters.Add(randomCenter)
		cluster := new_kmeans_cluster(randomCenter)
		clusters.Add(cluster)
	}
	
	for {
		
		// find nearest cluster and assign point to cluster
		for i := 0; i < self.points.Len(); i++ {
			smallestDistance := math.MaxFloat64
			var nearestCluster *kmeans_cluster
			
			point := self.points.Get(i).(*KMeansPoint)
			for j := 0; j < clusters.Len(); j++ {
				cluster := clusters.Get(j).(*kmeans_cluster)
				distanceBetweenCenterAndPoint := point.DistanceFromPoint(cluster.center)
				if  distanceBetweenCenterAndPoint < smallestDistance {
					smallestDistance = distanceBetweenCenterAndPoint
					nearestCluster = cluster
				}
			}
			nearestCluster.points.Add(point)
		}
		
		// recalculate new center in cluster and check if delta was satisfied
		biggestDeltaDistance := -math.MaxFloat64
		newDeltaDistance := self.delta
		for i := 0; i < clusters.Len(); i++ {
			cluster := clusters.Get(i).(*kmeans_cluster)						
			newDeltaDistance = cluster.recenter()
			if newDeltaDistance > biggestDeltaDistance {
				biggestDeltaDistance = newDeltaDistance
			}
		}
		
		// quit if delta was satisfied
		if newDeltaDistance < self.delta {
			break;
		} else {
			// otherwise clear cluster and try again
			for i := 0; i < clusters.Len(); i++ {
				cluster := clusters.Get(i).(*kmeans_cluster)
				cluster.points.Clear()
			}
		}
		
	}
	
	
	// otherwise clear cluster and try again
	for i := 0; i < clusters.Len(); i++ {
		cluster := clusters.Get(i).(*kmeans_cluster)
		fmt.Println("cluster:",cluster.center.items, "bucket:", cluster.points)
	}
}
