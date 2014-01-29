//
//  main.go
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

package main

import "fmt"
import "github.com/hishboy/gocommons/lang"
import "github.com/hishboy/gocommons/ml"
import "github.com/hishboy/gocommons/ml/support"

func main() { 
	fmt.Println("*** Queue ***")
	queue := lang.NewQueue()
	queue.Push("Hello")
	queue.Push(4)
	queue.Push(5)
	queue.Push(8)
	queue.Push(6)
	queue.Push("World")
	fmt.Println("Total items before poll: ", queue.Len())
	fmt.Println("peek before poll:", queue.Peek())
	fmt.Println(queue.Poll(), queue.Poll(), queue.Poll(), queue.Poll(), queue.Poll(), queue.Poll(), queue.Poll())
	fmt.Println("Total items after poll: ", queue.Len())
	fmt.Println("peek: ", queue.Peek())
	
	fmt.Println("\n")
	fmt.Println("*** Stack ***")
	stack := lang.NewStack()
	stack.Push("World")
	stack.Push(4)
	stack.Push(5)
	stack.Push(8)
	stack.Push(6)
	stack.Push("Hello")
	fmt.Println("Total items before pop: ", stack.Len())
	fmt.Println("peek before popping:", stack.Peek())
	fmt.Println(stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop())
	fmt.Println("Total items after pop: ", stack.Len())
	fmt.Println("peek: ", stack.Peek())
	
	fmt.Println("\n")
	fmt.Println("*** ArrayList ***")
	array := lang.NewArrayList()
	fmt.Println("random item in empty array:", array.Sample())
	array.Add("hello")
	array.Add("world")
	array.Add("Hicham")
	fmt.Println("total items in array:", array.Len())
	fmt.Println("array:", array)
	fmt.Println("array[1]:", array.Get(1))
	fmt.Println("array.indexOf(world):", array.IndexOf("world"))
	fmt.Println("array.Contains(world):", array.Contains("world"), array.Contains("random"))
	fmt.Println("random item in array:", array.Sample())
	fmt.Println("random item in array:", array.Sample())
	
	for i := 0; i < array.Len(); i++ {
		fmt.Println("item(", i, "):", array.Get(i))
	}
	
	fmt.Println("array.Remove(world):", array.Remove("world"), array)
	fmt.Println("array.Remove(Hicham):", array.Remove("Hicham"), array)
	fmt.Println("array.IsEmpty():", array.IsEmpty(), array.Remove("hello"), array.IsEmpty(), array.Len())
	
	for i := 0; i < 30; i++ {
		array.Add(i)
	}
	
	fmt.Println("array.indexOf(5):", array.IndexOf(5), array.Len())
	array.Clear()
	array.Add("one", "two", "three")
	
	array2 := lang.NewArrayList()
	array2.Add("four", "five", "six")
	
	array.AddFromArrayList(array2)
	
	fmt.Println("array:", array, array.Len())
	fmt.Println("array.First/Last:", array.First(), array.Last())
	fmt.Println("array.ToSlice:", array.ToSlice())
	
	fmt.Println("\n")
	fmt.Println("*** HashSet ***")	
	set := lang.NewHashSet()
	set.Add("hello", "world", "hello")
	fmt.Println("total items in set:", set.Len(), set)
	setToSlice := set.ToSlice()
	fmt.Println("set.ToSlice:", setToSlice, len(setToSlice))
	
	fmt.Println("\n")
	fmt.Println("*** KMeansSimpleCluster ***")
	kMeansCluster := ml.NewSimpleKMeans(4)
	kMeansCluster.AddPointAsSlice([]float64{48.2641334571,86.4516903905})
	kMeansCluster.AddPointAsSlice([]float64{0.114004262656,35.8368597414})
	kMeansCluster.AddPointAsSlice([]float64{97.4319168245,92.8009240744})
	kMeansCluster.AddPointAsSlice([]float64{24.4614031388,18.3292584382})
	kMeansCluster.AddPointAsSlice([]float64{36.2367675367,32.8294024271})
	kMeansCluster.AddPointAsSlice([]float64{75.5836860736,68.30729977})
	kMeansCluster.AddPointAsSlice([]float64{38.6577034445,25.7701728584})
	kMeansCluster.AddPointAsSlice([]float64{28.2607136287,64.4493377817})
	kMeansCluster.AddPointAsSlice([]float64{61.5358486771,61.2195232194})
	kMeansCluster.AddPointAsSlice([]float64{1.52352224798,38.5083779618})
	kMeansCluster.AddPointAsSlice([]float64{11.6392182793,68.2369021579})
	kMeansCluster.AddPointAsSlice([]float64{53.9486870607,53.9136556533})
	kMeansCluster.AddPointAsSlice([]float64{14.6671651772,26.0132534731})
	kMeansCluster.AddPointAsSlice([]float64{65.9506725878,82.5639317581})
	kMeansCluster.AddPointAsSlice([]float64{58.3682872339,51.6414580337})
	kMeansCluster.AddPointAsSlice([]float64{12.6918921252,2.28888447759})
	kMeansCluster.AddPointAsSlice([]float64{31.7587852231,18.1368234166})
	kMeansCluster.AddPointAsSlice([]float64{63.6631115204,24.933301389})
	kMeansCluster.AddPointAsSlice([]float64{29.1652289905,34.456759171})
	kMeansCluster.AddPointAsSlice([]float64{44.3830953085,70.4813875779})
	kMeansCluster.AddPointAsSlice([]float64{47.0571691145,65.3507625811})
	kMeansCluster.AddPointAsSlice([]float64{74.0584537502,98.2271944247})
	kMeansCluster.AddPointAsSlice([]float64{55.8929146157,86.6196265477})
	kMeansCluster.AddPointAsSlice([]float64{20.4744253473,12.0025149302})
	kMeansCluster.AddPointAsSlice([]float64{14.2867767281,40.2850440995})
	kMeansCluster.AddPointAsSlice([]float64{40.43551369,94.5410407116})
	kMeansCluster.AddPointAsSlice([]float64{87.6178871195,12.4700151639})
	kMeansCluster.AddPointAsSlice([]float64{47.2703048197,93.0636237124})
	kMeansCluster.AddPointAsSlice([]float64{59.7895104175,69.2621288413})
	kMeansCluster.AddPointAsSlice([]float64{80.8612333922,42.9183411179})
	kMeansCluster.AddPointAsSlice([]float64{31.1271795535,55.6669044656})
	kMeansCluster.AddPointAsSlice([]float64{78.9671049353,65.833739365})
	kMeansCluster.AddPointAsSlice([]float64{39.8324533414,63.0343115139})
	kMeansCluster.AddPointAsSlice([]float64{79.126343548,14.9128874133})
	kMeansCluster.AddPointAsSlice([]float64{65.8152400306,77.5202358013})
	kMeansCluster.AddPointAsSlice([]float64{75.2762752704,42.4858435609})
	kMeansCluster.AddPointAsSlice([]float64{29.6475948493,61.2068411763})
	kMeansCluster.AddPointAsSlice([]float64{67.421857106,54.8955604259})
	kMeansCluster.AddPointAsSlice([]float64{10.4652931501,29.7954139372})
	kMeansCluster.AddPointAsSlice([]float64{32.0272462745,99.5422900971})
	kMeansCluster.AddPointAsSlice([]float64{80.1520927001,84.2710379142})
	kMeansCluster.AddPointAsSlice([]float64{2.27240208403,41.2138854089})
	kMeansCluster.AddPointAsSlice([]float64{44.4601509555,1.72563901513})
	kMeansCluster.AddPointAsSlice([]float64{16.8676021068,35.3415636277})
	kMeansCluster.AddPointAsSlice([]float64{58.1977544121,29.2752085455})
	kMeansCluster.AddPointAsSlice([]float64{24.6119080085,39.9440735137})
	kMeansCluster.AddPointAsSlice([]float64{63.0759798755,60.9841014448})
	kMeansCluster.AddPointAsSlice([]float64{30.9289119657,95.0173219502})
	kMeansCluster.AddPointAsSlice([]float64{8.54972950047,41.7384441737})
	kMeansCluster.AddPointAsSlice([]float64{61.2606910793,4.06738902059})
	kMeansCluster.AddPointAsSlice([]float64{83.2302091964,11.6373312879})
	kMeansCluster.AddPointAsSlice([]float64{89.4443065362,42.5694882801})
	kMeansCluster.AddPointAsSlice([]float64{24.5619318152,97.7947977804})
	kMeansCluster.AddPointAsSlice([]float64{50.3134024475,40.6429336223})
	kMeansCluster.AddPointAsSlice([]float64{58.1422402033,36.1112632557})
	kMeansCluster.AddPointAsSlice([]float64{32.0668520827,29.9924151435})
	kMeansCluster.AddPointAsSlice([]float64{89.6057447137,84.9532177777})
	kMeansCluster.AddPointAsSlice([]float64{9.8876440816,18.2540486261})
	kMeansCluster.AddPointAsSlice([]float64{17.9670383961,47.596032257})
	kMeansCluster.AddPointAsSlice([]float64{50.2977668282,93.6851189223})
	kMeansCluster.AddPointAsSlice([]float64{98.0700386253,86.5816924579})
	kMeansCluster.AddPointAsSlice([]float64{10.8175290981,26.4344732252})
	kMeansCluster.AddPointAsSlice([]float64{34.7463851288,24.4154447141})
	kMeansCluster.AddPointAsSlice([]float64{92.5470100593,17.3595513748})
	kMeansCluster.AddPointAsSlice([]float64{79.0426629356,4.59850018907})
	kMeansCluster.AddPointAsSlice([]float64{89.9791366918,29.523946842})
	kMeansCluster.AddPointAsSlice([]float64{3.89920214563,91.3650215111})
	kMeansCluster.AddPointAsSlice([]float64{35.4669861576,62.1865368798})
	kMeansCluster.AddPointAsSlice([]float64{2.78150918086,24.5280230552})
	kMeansCluster.AddPointAsSlice([]float64{50.0390951889,57.0414421682})
	kMeansCluster.AddPointAsSlice([]float64{64.4521660758,48.4962172448})
	kMeansCluster.AddPointAsSlice([]float64{94.4915452316,56.6508179406})
	kMeansCluster.AddPointAsSlice([]float64{47.1655534769,15.8292055671})
	kMeansCluster.AddPointAsSlice([]float64{94.2027011374,45.6802385454})
	kMeansCluster.AddPointAsSlice([]float64{30.5846324871,54.783635876})
	kMeansCluster.AddPointAsSlice([]float64{57.7043252948,0.286661610381})
	kMeansCluster.AddPointAsSlice([]float64{41.7908674949,14.7206014023})
	kMeansCluster.AddPointAsSlice([]float64{59.6689465934,64.8849831965})
	kMeansCluster.AddPointAsSlice([]float64{92.2553335495,55.9096460272})
	kMeansCluster.AddPointAsSlice([]float64{48.493467262,69.4766837809})
	kMeansCluster.AddPointAsSlice([]float64{23.1837859581,71.4406867443})
	kMeansCluster.AddPointAsSlice([]float64{29.0737623652,66.9391416961})
	kMeansCluster.AddPointAsSlice([]float64{95.7442323112,89.4677505059})
	kMeansCluster.AddPointAsSlice([]float64{68.7707275828,40.9900140055})
	kMeansCluster.AddPointAsSlice([]float64{84.5445737133,32.1707309618})
	kMeansCluster.AddPointAsSlice([]float64{67.4126251988,56.6710579117})
	kMeansCluster.AddPointAsSlice([]float64{10.688352016,28.1745892928})
	kMeansCluster.AddPointAsSlice([]float64{56.7620324155,18.3034334207})
	kMeansCluster.AddPointAsSlice([]float64{50.6751320678,86.6916908032})
	kMeansCluster.AddPointAsSlice([]float64{74.6185482896,34.022483532})
	kMeansCluster.AddPointAsSlice([]float64{20.7011996002,32.855295357})
	kMeansCluster.AddPointAsSlice([]float64{11.479054664,1.59204297586})
	kMeansCluster.AddPointAsSlice([]float64{51.6805387648,25.4063026358})
	kMeansCluster.AddPointAsSlice([]float64{84.4109522357,47.237632645})
	kMeansCluster.AddPointAsSlice([]float64{90.6395051745,57.7917166935})
	kMeansCluster.AddPointAsSlice([]float64{58.6159601042,84.1226173848})
	kMeansCluster.AddPointAsSlice([]float64{46.2184509277,28.559934585})
	kMeansCluster.AddPointAsSlice([]float64{97.0302485783,41.3135022812})
	kMeansCluster.AddPointAsSlice([]float64{31.3144587058,87.2459910122})
	kMeansCluster.AddPointAsSlice([]float64{5.93357833962,95.6812831872})
	clusters := kMeansCluster.Cluster()
	
	for i := 0; i < clusters.Len(); i++ {
		cluster := clusters.Get(i).(*support.KMeansCluster)
		fmt.Println("cluster:",cluster.Center().Items().ToSlice())
		for j := 0; j < cluster.Points().Len(); j++ {
			point := cluster.Points().Get(j).(*support.KMeansPoint)
			fmt.Println("--", point.Items().ToSlice())
		}
	}
	
	
}