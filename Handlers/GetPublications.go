package Handlers

import (
	"GORM/Connection"
	"GORM/Models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetPublications(c *gin.Context) {
	conn, _ := Connection.GetConnection()

	var posts []Models.Publicacion

	if err := conn.Preload("Usuario").Find(&posts).Error; err != nil {
		log.Println("Error al obtener las identificaciones:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las identificaciones"})
		return
	}

	var respuesta []map[string]interface{}
	for _, publicacion := range posts {
		publicacionJSON := map[string]interface{}{
			"id_publicacion": publicacion.ID,
			"nombre_usuario": func() string {
				var nombreUsuario string
				if err := conn.Model(&Models.Usuario{}).Where("id = ?", publicacion.UsuarioID).Pluck("nombre", &nombreUsuario).Error; err != nil {
					log.Println("Error al obtener el nombre del usuario:", err)
					return ""
				}
				return nombreUsuario
			}(),
			"id_usuario": func() string {
				var idUsuario string
				if err := conn.Model(&Models.Usuario{}).Where("id = ?", publicacion.UsuarioID).Pluck("id", &idUsuario).Error; err != nil {
					log.Println("Error al obtener el nombre del usuario:", err)
					return ""
				}
				return idUsuario
			}(),
			"likes": func() int64 {
				var sumLikes int64 = 0

				// Primero, obtén las identificaciones asociadas a la publicación
				var identificaciones []Models.Identificacion
				resultIdent := conn.Model(&Models.Identificacion{}).Where("publicacion_id = ?", publicacion.ID).Find(&identificaciones)

				if resultIdent.Error != nil {
					log.Println("Error al obtener las identificaciones:", resultIdent.Error)
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las identificaciones"})

				}

				// Luego, para cada identificación, obtén el número de aprobaciones
				for _, identificacion := range identificaciones {
					var likesIdentificacion int64 = 0

					resultAprob := conn.Model(&Models.Aprobacion{}).Where("identificacion_id = ?", identificacion.ID).Count(&likesIdentificacion)

					if resultAprob.Error != nil {
						log.Println("Error al obtener el número de aprobaciones:", resultAprob.Error)
						c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el número de aprobaciones"})
						continue
					}
					sumLikes += likesIdentificacion
				}
				return sumLikes
			}(),
			"comentarios": func() int64 {
				var countComentarios int64
				result := conn.Model(&Models.Identificacion{}).Where("publicacion_id = ?", publicacion.ID).Count(&countComentarios)
				if result.Error != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el número de aprobaciones"})
					return countComentarios
				}
				return countComentarios
			}(),
			"imagen": publicacion.UrlImagen,
			"nombre_comun": func() string {
				// Obtén las identificaciones asociadas a la publicación
				var identificaciones []Models.Identificacion
				resultIdent := conn.Model(&Models.Identificacion{}).Where("publicacion_id = ?", publicacion.ID).Find(&identificaciones)

				if resultIdent.Error != nil {
					log.Println("Error al obtener las identificaciones:", resultIdent.Error)
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las identificaciones"})
					return ""
				}

				// Crea un mapa para contar la frecuencia de cada nombre común
				nombreComunCount := make(map[string]int)

				// Itera sobre las identificaciones y cuenta la frecuencia de cada nombre común
				for _, identificacion := range identificaciones {
					nombreComunCount[identificacion.NombreComun]++
				}

				// Encuentra el nombre común con la frecuencia más alta
				var maxCount int
				var nombreComunMasRepetido string
				for nombreComun, count := range nombreComunCount {
					if count > maxCount {
						maxCount = count
						nombreComunMasRepetido = nombreComun
					}
				}

				return nombreComunMasRepetido
			}(),
		}
		respuesta = append(respuesta, publicacionJSON)
	}

	c.JSON(http.StatusOK, respuesta)
}
