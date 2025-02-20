package collect

import (
	"encoding/json"

	"github.com/pkg/errors"
	troubleshootv1beta2 "github.com/replicatedhq/troubleshoot/pkg/apis/troubleshoot/v1beta2"
	"github.com/shirou/gopsutil/mem"
)

type MemoryInfo struct {
	Total uint64 `json:"total"`
}

type CollectHostMemory struct {
	hostCollector *troubleshootv1beta2.Memory
}

func (c *CollectHostMemory) Title() string {
	return hostCollectorTitleOrDefault(c.hostCollector.HostCollectorMeta, "Amount of Memory")
}

func (c *CollectHostMemory) IsExcluded() (bool, error) {
	return isExcluded(c.hostCollector.Exclude)
}

func (c *CollectHostMemory) Collect(progressChan chan<- interface{}) (map[string][]byte, error) {
	memoryInfo := MemoryInfo{}

	vmstat, err := mem.VirtualMemory()
	if err != nil {
		return nil, errors.Wrap(err, "failed to read virtual memory")
	}
	memoryInfo.Total = vmstat.Total

	b, err := json.Marshal(memoryInfo)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal memory info")
	}

	return map[string][]byte{
		"system/memory.json": b,
	}, nil
}
