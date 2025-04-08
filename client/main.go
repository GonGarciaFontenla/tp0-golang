package main

import (
	"client/globals"
	"client/utils"
	"log"
)

func main() {
	utils.ConfigurarLogger()

	// loggear "Hola soy un log" usando la biblioteca log
	log.Println("Hola soy un log")

	globals.ClientConfig = utils.IniciarConfiguracion("config.json")
	// validar que la config este cargada correctamente
	if globals.ClientConfig == nil {
		log.Fatalf("Error al inicializar el archivo config")
	}
	// loggeamos el valor de la config
	mensaje := globals.ClientConfig.Mensaje
	log.Println(mensaje)

	// ADVERTENCIA: Antes de continuar, tenemos que asegurarnos que el servidor esté corriendo para poder conectarnos a él

	// enviar un mensaje al servidor con el valor de la config
	ip := globals.ClientConfig.Ip
	puerto := globals.ClientConfig.Puerto

	utils.EnviarMensaje(ip, puerto, globals.ClientConfig.Mensaje)

	// generamos un paquete y lo enviamos al servidor
	utils.GenerarYEnviarPaquete()
}
