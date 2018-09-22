Library of statistical functions.

# Base
Basic statistical functions.

## Min
Min returns the minimum value in the data. 

```
data := []float64{15.01, 51.0, 32.00015, 15.02}
min, err := gostats.Min(data)
if err != nil {
	fmt.Println(err)
}

fmt.Println(min)

// 15.01
```

## Max
Max returns the maximum value in the data.

```
data := []float64{882.99, 51.0, 32.00015, 15.02, 883.0}
max, err := gostats.Max(data)
if err != nil {
	fmt.Println(err)
}

fmt.Println(max)

// 883
```

## Range
Range returns range between minimum and maximum value in the data.

```
data := []float64{48.0743, 1.01, 98.1, 99.99}
r, err := gostats.Range(data)
if err != nil {
	fmt.Println(err)
}

fmt.Println(r)

// 98.97999999999999
```

## Sum
Sum returns the sum of the values in the data.

```
data := []float64{2.0, 2.0, 20.1, 0.9}
sum, err := gostats.Sum(data)
if err != nil {
	fmt.Println(err)
}

fmt.Println(sum)

// 25
```

## ArithmeticMean
ArithmeticMean returns the arithmetic mean of the values in the data.

```
data := []float64{5.2, 2.0, 2.0, 20.1, 0.9}
amean, err := gostats.ArithmeticMean(data)
if err != nil {
	fmt.Println(err)
}
fmt.Println(amean)

// 6.04
```

## Median
Median returns the median of the values in the data.

With one value:

```
data := []float64{7.0}
median, err := gostats.Median(data)
if err != nil {
	fmt.Println(err)
}
fmt.Println(median)

// 7
```

Even number of values:

```
data := []float64{1.0, 51.01, 83.91, 44.00}
median, err := gostats.Median(data)
if err != nil {
	fmt.Println(err)
}
fmt.Println(median)

// 47.504999999999995
```

Odd number of values:

```
data := []float64{77.31, 1.0, 51.01, 83.91, 44.00}
median, err := gostats.Median(data)
if err != nil {
	fmt.Println(err)
}
fmt.Println(median)

// 51.01
```

## Mode
Mode returns an array of the modal values in the data.

```
data := []float64{7.0, 10.2, 29.15, 0.0001, 7.0, 43.27, 88.1, 29.15}
mode, err := gostats.Mode(data)
if err != nil {
	fmt.Println(err)
}
fmt.Println(mode)

// [7 29.15]
```

No the modal values:

```
data := []float64{7.0, 10.2, 0.0001, 43.27, 88.1}
mode, err := gostats.Mode(data)
if err != nil {
	fmt.Println(err)
}
fmt.Println(mode)

// []
```

With one modal value:

```
data := []float64{6.7, 6.7}
mode, err := gostats.Mode(data)
if err != nil {
	fmt.Println(err)
}
fmt.Println(mode)

// [6.7]
```

Quantity of each value in the data is equel:

```
data := []float64{10.99, 10.99, 11.99, 11.99, 12.99, 12.99}
mode, err := gostats.Mode(data)
if err != nil {
	fmt.Println(err)
}
fmt.Println(mode)

// []
```

## Variance
Variance returns the variance of the values in the data.

```
data := []float64{4.0, 5.0, 0.5, 0.5}
v, err := gostats.Variance(data)
if err != nil {
	fmt.Println(err)
}
fmt.Println(v)

// 5.5
```

## SD
SD returns the standard deviation of the values in the data.

```
data := []float64{4.0, 5.0, 0.5, 0.5}
sd, err := gostats.SD(data)
if err != nil {
	fmt.Println(err)
}
fmt.Println(sd)

// 2.345207879911715
```

# TODO

* Add correlation functions: Kendall, Spearman, Pearson.