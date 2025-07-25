package entities

import (
	"encoding/json"
	"fmt"
	"runtime"
	"time"
)

// Processing represents the result of a data operation, tracking metadata and excluding non-JSON marshaled data.
type Processing struct {
	// Data doesn't marshal to JSON well
	Data     [][]int            `json:"-"`
	Metadata ProcessingMetadata `json:"metadata"`
}

// ProcessingMetadata holds metadata about the processing of data, including rows, filters, performance, and memory usage.
type ProcessingMetadata struct {
	SourceTotalRows       int                `json:"sourceTotalRows"`
	FilteredTotalRows     int                `json:"filterTotalRows"`
	AppliedFilters        []string           `json:"appliedFilters"`
	PerformedAggregations []string           `json:"performedAggregations"`
	PerformedMerges       []string           `json:"performedMerges"`
	ProcessingTime        time.Duration      `json:"processingTime"`
	StartTime             time.Time          `json:"startTime"`
	EndTime               time.Time          `json:"endTime"`
	ConfigName            string             `json:"configName"`
	DataSource            string             `json:"dataSource"`
	MemoryStats           MemoryStats        `json:"memoryStats"`
	StepPerformance       []PerformanceEntry `json:"stepPerformance"`
}

// MemoryStats represents memory statistics during program execution.
// It includes details about peak memory usage, allocations, GC count, and memory usage growth percentage.
type MemoryStats struct {
	PeakAllocBytes        uint64  `json:"peakAllocBytes"`        // Peak allocation in bytes
	PeakSysBytes          uint64  `json:"peakSysBytes"`          // Peak system memory obtained from OS
	TotalAllocBytes       uint64  `json:"totalAllocBytes"`       // Total bytes allocated (even if freed)
	FinalAllocBytes       uint64  `json:"finalAllocBytes"`       // Bytes allocated and not yet freed
	NumGC                 uint32  `json:"numGC"`                 // Number of garbage collections
	MemoryIncreasePercent float64 `json:"memoryIncreasePercent"` // Percentage increase in memory usage
}

// PerformanceEntry represents a record of performance metrics for a specific processing step.
// It includes details about timing, input/output rows, and memory usage.
type PerformanceEntry struct {
	StepName         string        `json:"stepName"`
	StartTime        time.Time     `json:"startTime"`
	EndTime          time.Time     `json:"endTime"`
	Duration         time.Duration `json:"duration"`
	InputRows        int           `json:"inputRows"`
	OutputRows       int           `json:"outputRows"`
	MemoryUsageBytes uint64        `json:"memoryUsageBytes"`
}

// NewProcessing initializes a new Processing instance with provided data and configuration name.
// It records initial memory statistics, start time, and sets up metadata for tracking processing operations.
func NewProcessing(data [][]int, configName string) *Processing {
	var initMemStats runtime.MemStats
	runtime.ReadMemStats(&initMemStats)

	totalRows := len(data)

	return &Processing{
		Data: data,
		Metadata: ProcessingMetadata{
			SourceTotalRows:       totalRows,
			AppliedFilters:        make([]string, 0),
			PerformedAggregations: make([]string, 0),
			PerformedMerges:       make([]string, 0),
			StartTime:             time.Now(),
			ConfigName:            configName,
			MemoryStats: MemoryStats{
				PeakAllocBytes:  initMemStats.Alloc,
				PeakSysBytes:    initMemStats.Sys,
				TotalAllocBytes: initMemStats.TotalAlloc,
				NumGC:           initMemStats.NumGC,
			},
			StepPerformance: make([]PerformanceEntry, 0),
		},
	}
}

// AddFilter appends a filter expression to the list of applied filters in the metadata of the Processing instance.
func (p *Processing) AddFilter(filterExp string) {
	p.Metadata.AppliedFilters = append(p.Metadata.AppliedFilters, filterExp)
}

// AddAggregation appends an aggregation expression to the list of performed aggregations in the metadata of the Processing instance.
func (p *Processing) AddAggregation(aggregationMethod, targetColumn string) {
	aggregation := fmt.Sprintf("%s(%s)", aggregationMethod, targetColumn)
	p.Metadata.PerformedAggregations = append(p.Metadata.PerformedAggregations, aggregation)
}

// AddMerge appends a merge description to the list of performed merges in the metadata of the Processing instance.
func (p *Processing) AddMerge(firstColumn, secondColumn, strategy string) {
	merge := fmt.Sprintf("%s + %s (%s)", firstColumn, secondColumn, strategy)
	p.Metadata.PerformedMerges = append(p.Metadata.PerformedMerges, merge)
}

// UpdateRows updates the total rows before and after filtering in the metadata of the Processing instance.
func (p *Processing) UpdateRows(originalRows, filteredRows int) {
	// TODO: Confirm if the SourceTotalRows needs to update
	p.Metadata.SourceTotalRows = originalRows
	p.Metadata.FilteredTotalRows = filteredRows
}

// SetDataSourceInfo updates the data source information in the metadata of the Processing instance.
func (p *Processing) SetDataSourceInfo(info string) {
	p.Metadata.DataSource = info
}

// CompleteProcess finalizes processing by capturing end time, calculating processing time, and updating memory usage statistics.
func (p *Processing) CompleteProcess() {
	p.Metadata.EndTime = time.Now()
	p.Metadata.ProcessingTime = p.Metadata.EndTime.Sub(p.Metadata.StartTime)

	// Capture final memory stats
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	p.Metadata.MemoryStats.FinalAllocBytes = memStats.Alloc
	p.Metadata.MemoryStats.TotalAllocBytes = memStats.TotalAlloc
	p.Metadata.MemoryStats.NumGC = memStats.NumGC

	// Calculate memory increase percentage
	if p.Metadata.MemoryStats.PeakAllocBytes > 0 {
		p.Metadata.MemoryStats.MemoryIncreasePercent =
			float64(memStats.Alloc-p.Metadata.MemoryStats.PeakAllocBytes) / float64(p.Metadata.MemoryStats.PeakAllocBytes) * 100
	}

	// Update peak values if final values are higher
	if memStats.Alloc > p.Metadata.MemoryStats.PeakAllocBytes {
		p.Metadata.MemoryStats.PeakAllocBytes = memStats.Alloc
	}

	if memStats.Sys > p.Metadata.MemoryStats.PeakSysBytes {
		p.Metadata.MemoryStats.PeakSysBytes = memStats.Sys
	}
}

// ToJSON converts the Processing instance into a formatted JSON string and returns it. Returns an error if marshaling fails.
func (p *Processing) ToJSON() (string, error) {
	data, err := json.MarshalIndent(p, "", "    ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal Processing to JSON: %w", err)
	}

	return string(data), nil
}

// HasData checks if the Processing instance contains any data by verifying the length of the `Data` field. Returns true if data exists.
func (p *Processing) HasData() bool {
	return len(p.Data) > 0
}

// GetRowCount returns the number of rows in the Data field of the Processing instance. Returns 0 if Data is nil.
func (p *Processing) GetRowCount() int {
	if p.Data == nil {
		return 0
	}

	return len(p.Data)
}

// GetColumnCount returns the number of columns in the Data field of the Processing instance. Returns 0 if Data is nil.
func (p *Processing) GetColumnCount() int {
	if p.Data == nil {
		return 0
	}

	return len(p.Data[0])
}

// GetColumnNames returns a slice of strings representing the names of the columns in the Data field of the Processing instance.
func (p *Processing) GetColumnNames() []string {
	panic("implement me")
}
