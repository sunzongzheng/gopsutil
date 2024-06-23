//go:build darwin && ios
// +build darwin,ios

package host

import (
	"context"

	"github.com/shirou/gopsutil/v3/internal/common"
)

func SensorsTemperaturesWithContext(ctx context.Context) ([]TemperatureStat, error) {
	return []TemperatureStat{}, common.ErrNotImplementedError
}
