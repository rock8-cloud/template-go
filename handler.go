package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SpaceTarget struct {
	ID                 int64   `json:"id"`
	Name               string  `json:"name"`
	Type               string  `json:"type"`
	DistanceLightYears float64 `json:"distance_light_years"`
}

// listTargets godoc
// @Summary      Returns a list of all space targets
// @Tags         Space Targets
// @Produce      json
// @Success      200  {array}   SpaceTarget
// @Router       /api/targets [get]
func listTargets(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := pool.Query(context.Background(),
			"SELECT id, name, type, distance_light_years FROM space_target ORDER BY id")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var targets []SpaceTarget
		for rows.Next() {
			var t SpaceTarget
			if err := rows.Scan(&t.ID, &t.Name, &t.Type, &t.DistanceLightYears); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			targets = append(targets, t)
		}
		c.JSON(http.StatusOK, targets)
	}
}
