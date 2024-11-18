package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/montanaflynn/stats"
)

func main() {
	B := 100
	fmt.Printf("\nRunning study with %d bootstrap samples\n", B)

	popMean := 100.0
	popSD := 10.0
	studySampleSizes := []int{25, 100, 225, 400}

	fmt.Printf("\nStudy conditions:\n  Population mean: %.2f SD: %.2f\n", popMean, popSD)

	rand.Seed(9999)

	var studyResults []map[string]float64

	for iteration := 0; iteration < 100; iteration++ {
		for _, n := range studySampleSizes {
			thisSample := generateSample(n, popMean, popSD)
			thisSampleMean, _ := stats.Mean(thisSample)

			var bootstrapSampleResults []map[string]float64
			for b := 0; b < B; b++ {
				thisBootstrapSample := bootstrapSample(thisSample, n)
				thisBootstrapMean, _ := stats.Mean(thisBootstrapSample)
				thisBootstrapMedian, _ := stats.Median(thisBootstrapSample)

				bootstrapSampleResults = append(bootstrapSampleResults, map[string]float64{
					"bootMean":   thisBootstrapMean,
					"bootMedian": thisBootstrapMedian,
				})
			}

			bootMean := calculateMean(bootstrapSampleResults, "bootMean")
			bootMedian := calculateMean(bootstrapSampleResults, "bootMedian")

			studyResults = append(studyResults, map[string]float64{
				"n":          float64(n),
				"sampleMean": thisSampleMean,
				"bootMean":   bootMean,
				"bootMedian": bootMedian,
			})
		}
	}

	fmt.Printf("\nEstimated standard errors using %d bootstrap samples\n", B)
	for _, n := range studySampleSizes {
		var sampleMeans, bootMeans, bootMedians []float64
		for _, result := range studyResults {
			if int(result["n"]) == n {
				sampleMeans = append(sampleMeans, result["sampleMean"])
				bootMeans = append(bootMeans, result["bootMean"])
				bootMedians = append(bootMedians, result["bootMedian"])
			}
		}

		seMeanCLT := popSD / math.Sqrt(float64(n))
		seMeanSample, _ := stats.StandardDeviation(sampleMeans)
		seMeanBoot, _ := stats.StandardDeviation(bootMeans)
		seMedianBoot, _ := stats.StandardDeviation(bootMedians)

		fmt.Printf("\nSamples of size n = %d\n", n)
		fmt.Printf("  SE Mean from Central Limit Theorem: %.2f\n", seMeanCLT)
		fmt.Printf("  SE Mean from Samples: %.2f\n", seMeanSample)
		fmt.Printf("  SE Mean from Bootstrap Samples: %.2f\n", seMeanBoot)
		fmt.Printf("  SE Median from Bootstrap Samples: %.2f\n", seMedianBoot)
	}
}

func generateSample(n int, mean, sd float64) []float64 {
	sample := make([]float64, n)
	for i := 0; i < n; i++ {
		sample[i] = rand.NormFloat64()*sd + mean
	}
	return sample
}

func bootstrapSample(sample []float64, n int) []float64 {
	bootstrap := make([]float64, n)
	for i := 0; i < n; i++ {
		index := rand.Intn(len(sample))
		bootstrap[i] = sample[index]
	}
	return bootstrap
}

func calculateMean(data []map[string]float64, field string) float64 {
	sum := 0.0
	for _, item := range data {
		sum += item[field]
	}
	return sum / float64(len(data))
}
