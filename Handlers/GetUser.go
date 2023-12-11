package Handlers

import (
	"GORM/Connection"
	"GORM/Models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	conn, _ := Connection.GetConnection()

	// Obtén el ID del usuario desde los parámetros de la URL
	UserID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario no válido"})
		return
	}

	// Consulta la base de datos para obtener las identificaciones asociadas a la publicación
	var usuario Models.Usuario
	result := conn.Preload("Rol").Where("id = ?", UserID).Find(&usuario)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las identificaciones"})
		return
	}

	var posts []Models.Publicacion

	if err := conn.Preload("Usuario").Where("usuario_id = ?", usuario.ID).Find(&posts).Error; err != nil {
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

	publicacionJSON := map[string]interface{}{
		"programa": usuario.Programa,
		"nombre":   usuario.Nombre,
		"correo":   usuario.CorreoUsuario,
		"semestre": usuario.Semestre,
		"num_publicaciones": func() int64 {
			var countpub int64
			result := conn.Model(&Models.Publicacion{}).Where("usuario_id = ?", usuario.ID).Count(&countpub)
			if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el número de aprobaciones"})
				return countpub
			}
			return countpub
		}(),
		"num_identificaciones": func() int64 {
			var countiden int64
			result := conn.Model(&Models.Identificacion{}).Where("usuario_id = ?", usuario.ID).Count(&countiden)
			if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el número de aprobaciones"})
				return countiden
			}
			return countiden
		}(),
		"num_aprobaciones": func() int64 {
			var countapro int64
			result := conn.Model(&Models.Aprobacion{}).Where("usuario_id = ?", usuario.ID).Count(&countapro)
			if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el número de aprobaciones"})
				return countapro
			}
			return countapro
		}(),
		"publicaciones": respuesta,
	}

	c.JSON(http.StatusOK, publicacionJSON)
}
