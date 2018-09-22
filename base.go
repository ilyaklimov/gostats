package gostats

import (
    "fmt"
    "sort"
    "math"

    "github.com/ilyaklimov/godata"
)

// Min returns the minimum value in the data. 
func Min(data []float64) (float64, error) {
    if len(data) == 0 {
        return 0.0, errDataIsEmpty
    }
    min := data[0]
    for i := 0; i < len(data); i++ {
        if data[i] < min {
            min = data[i]
        }
    }
    return min, nil
}

// Max returns the maximum value in the data.
func Max(data []float64) (float64, error) {
    if len(data) == 0 {
        return 0.0, errDataIsEmpty
    }
    max := data[0]
    for i := 0; i < len(data); i++ {
        if data[i] > max {
            max = data[i]
        }
    }
    return max, nil
}

// Range returns range between minimum and maximum value in the data.
func Range(data []float64) (float64, error) {
    switch len(data) {
    case 0:
        return 0.0, errDataIsEmpty
    case 1:
        return 0.0, errInsufficientData
    }
    min, err := Min(data)
    if err != nil {
        return 0.0, fmt.Errorf("cannot calculate min: %s", err)
    }
    max, err := Max(data)
    if err != nil {
        return 0.0, fmt.Errorf("cannot calculate max: %s", err)
    }
    return max - min, nil
}

// Sum returns the sum of the values in the data.
func Sum(data []float64) (float64, error) {
    if len(data) == 0 {
        return 0.0, errDataIsEmpty
    }
    var sum float64 
    for i := 0; i < len(data); i++ {
        sum += data[i]
    }
    return sum, nil
}

// ArithmeticMean returns the arithmetic mean of the values in the data.
func ArithmeticMean(data []float64) (float64, error) {
    if len(data) == 0 {
        return 0.0, errDataIsEmpty
    }
    sum, err := Sum(data)
    if err != nil {
        return 0.0, fmt.Errorf("cannot calculate sum: %s", err)
    }
    return sum / float64(len(data)), nil
}

// Median returns the median of the values in the data.
func Median(data []float64) (float64, error) {
    switch len(data) {
    case 0:
        return 0.0, errDataIsEmpty
    case 1:
        return data[0], nil
    }
    sort.Float64s(data)
    i := int(math.Floor(float64(len(data)) / 2))
    if len(data) % 2 == 1 {
        return data[i], nil
    } else {
        return (data[i-1] + data[i]) / 2, nil
    }
}

// Mode returns an array of the modal values in the data.
func Mode(data []float64) ([]float64, error) {
    switch len(data) {
    case 0:
        return []float64{}, errDataIsEmpty
    case 1:
        return []float64{}, errInsufficientData
    }

    ds := godata.NewFloat64List(data).Duplicates()

    switch len(ds) {
    case 0:
        return []float64{}, nil
    case 1:
        return []float64{ds[0].Value}, nil
    }

    // If the quantity of each value in the data is equel, then there are no the modal values.
    // Example: data = [10.99, 10.99, 11.99, 11.99, 12.99, 12.99], ds = [{10.99, 2}, {11.99, 2}, {12.99, 2}] => mode = [].
    if ((ds[0].Count == ds[len(ds)-1].Count) && (ds[0].Count * len(ds) == len(data))) {
        return []float64{}, nil
    }

    ds.SortByCountReverse()

    mode := make([]float64, 1, len(ds))
    mode[0] = ds[0].Value

    for i := 1; i < len(ds); i++ {
        if ds[0].Count == ds[i].Count {
            mode = append(mode, ds[i].Value)
        }
    }

    return mode, nil    
}

// Variance returns the variance of the values in the data.
func Variance(data []float64) (float64, error) {
    switch len(data) {
    case 0:
        return 0.0, errDataIsEmpty
    case 1:
        return 0.0, errInsufficientData
    }
    mean, err := ArithmeticMean(data)
    if err != nil {
        return 0.0, fmt.Errorf("cannot calculate mean: %v", err)
    }
    var sum float64 
    for _, value := range data {
        sum += math.Pow(value - mean, 2)
    }
    return sum / float64(len(data) - 1), nil
}

// SD returns the standard deviation of the values in the data.
func SD(data []float64) (float64, error) {
    switch len(data) {
    case 0:
        return 0.0, errDataIsEmpty
    case 1:
        return 0.0, errInsufficientData
    }
    variance, err := Variance(data)
    if err != nil {
        return 0.0, fmt.Errorf("cannot calculate variance: %v", err)
    }
    return math.Sqrt(variance), nil
}