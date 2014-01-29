//
//  simple_kmeans.go
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
import "github.com/hishboy/gocommons/ml/support"
import "math"

type SimpleKMeans struct {
	points *lang.ArrayList
	numberOfClusters int
	delta float64
}


func NewSimpleKMeans(numberOfClusters int) *SimpleKMeans {
	self := &SimpleKMeans{}
	self.points = lang.NewArrayList()
	self.numberOfClusters = numberOfClusters
	self.delta = 0.001 // default delta
	return self
}


func (self *SimpleKMeans) SetDelta(delta float64) {
	self.delta = delta
}

func (self *SimpleKMeans) AddPoint(point *support.KMeansPoint) {
	self.points.Add(point)
}

func (self *SimpleKMeans) AddPointAsSlice(items []float64) {
	self.points.Add(support.NewKMeansPoint(items))
}

func (self *SimpleKMeans) Cluster() *lang.ArrayList {
	if (self.numberOfClusters == 1) {
		panic("please specify more than one cluster")
	}
	
	clusters := lang.NewArrayList()
	uniqueCenters := lang.NewHashSet()
	for i := 0; i < self.numberOfClusters; i++ {
		randomCenter := self.points.Sample().(*support.KMeansPoint)
		for uniqueCenters.Contains(randomCenter) {
			randomCenter = self.points.Sample().(*support.KMeansPoint)
		}
		uniqueCenters.Add(randomCenter)
		cluster := support.NewKMeansCluster(randomCenter)
		clusters.Add(cluster)
	}
	
	for {
		
		// find nearest cluster and assign point to cluster
		for i := 0; i < self.points.Len(); i++ {
			smallestDistance := math.MaxFloat64
			var nearestCluster *support.KMeansCluster
			
			point := self.points.Get(i).(*support.KMeansPoint)
			for j := 0; j < clusters.Len(); j++ {
				cluster := clusters.Get(j).(*support.KMeansCluster)
				distanceBetweenCenterAndPoint := point.DistanceFromPoint(cluster.Center())
				if  distanceBetweenCenterAndPoint < smallestDistance {
					smallestDistance = distanceBetweenCenterAndPoint
					nearestCluster = cluster
				}
			}
			nearestCluster.Points().Add(point)
		}
		
		// recalculate new center in cluster and check if delta was satisfied
		biggestDeltaDistance := -math.MaxFloat64
		newDeltaDistance := self.delta
		for i := 0; i < clusters.Len(); i++ {
			cluster := clusters.Get(i).(*support.KMeansCluster)						
			newDeltaDistance = cluster.Recenter()
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
				cluster := clusters.Get(i).(*support.KMeansCluster)
				cluster.Points().Clear()
			}
		}
		
	}
	
	return clusters
}
