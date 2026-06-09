package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"gorm.io/gorm"
)

type HealthHandler struct {
	DB *gorm.DB
}

func (h *HealthHandler) Check(c *gin.Context) {
    sqlDB, err := h.DB.DB()
    dbStatus := "UP"
    if err != nil || sqlDB.Ping() != nil {
        dbStatus = "DOWN"
    }
    cpuPercent, _ := cpu.Percent(time.Second, false)
    vm, _ := mem.VirtualMemory()

    response := gin.H{
        "status":    "UP",
        "version":   "v1.0.1",           // ← bump this each test push
        "timestamp": time.Now().Format(time.RFC3339),
        "services": gin.H{
            "database": dbStatus,
        },
        "system": gin.H{
            "cpu_usage_percent": cpuPercent[0],
            "ram_usage_percent": vm.UsedPercent,
            "ram_total_gb":      vm.Total / 1024 / 1024 / 1024,
        },
    }

    if dbStatus == "DOWN" {
        c.JSON(http.StatusServiceUnavailable, response)
        return
    }
    c.JSON(http.StatusOK, response)
}s