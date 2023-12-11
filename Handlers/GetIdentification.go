package Handlers

import (
	"GORM/Connection"
	"GORM/Models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sort"
	"strconv"
)

type IdentificacionWithAprobaciones struct {
	Identificacion Models.Identificacion
	Aprobaciones   int64 `json:"aprobaciones"`
}

func GetIdentication(c *gin.Context) {
	conn, _ := Connection.GetConnection()

	// Obtén el ID de la publicación desde los parámetros de la URL
	publicacionID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de publicación no válido"})
		return
	}

	// Consulta la base de datos para obtener las identificaciones asociadas a la publicación
	var identificaciones []Models.Identificacion
	result := conn.Preload("Usuario").Preload("Publicacion").Where("publicacion_id = ?", publicacionID).Find(&identificaciones)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las identificaciones"})
		return
	}

	// Crea una lista de identificaciones con aprobaciones
	var identificacionesConAprobaciones []IdentificacionWithAprobaciones

	// Itera sobre las identificaciones para obtener la cantidad de aprobaciones y almacenarla
	for _, identificacion := range identificaciones {
		var aprobaciones int64 = 0

		resultAprob := conn.Model(&Models.Aprobacion{}).Where("identificacion_id = ?", identificacion.ID).Count(&aprobaciones)

		if resultAprob.Error != nil {
			log.Println("Error al obtener el número de aprobaciones:", resultAprob.Error)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el número de aprobaciones"})
			continue
		}

		// Agrega la identificación con aprobaciones a la lista
		identificacionConAprobaciones := IdentificacionWithAprobaciones{
			Identificacion: identificacion,
			Aprobaciones:   aprobaciones,
		}

		identificacionesConAprobaciones = append(identificacionesConAprobaciones, identificacionConAprobaciones)
	}

	// Ordena la lista por la cantidad de aprobaciones de forma descendente
	sort.Slice(identificacionesConAprobaciones, func(i, j int) bool {
		return identificacionesConAprobaciones[i].Aprobaciones > identificacionesConAprobaciones[j].Aprobaciones
	})

	c.JSON(http.StatusOK, identificacionesConAprobaciones)
}
