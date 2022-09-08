package main

type testCsv struct {
	TestStr     string
	TestInt32   int32
	TestInt64   int64
	TestFloat32 float32
	TestFloat64 float64
}

type testCsvWithTags struct {
	TestStr     string
	TestInt64   int32 `csv:"2"`
	TestFloat32 float32
}
