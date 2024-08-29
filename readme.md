# Netsocs Golang SDK

Este SDK es disenado para la interaccion con la API de Netsocs.

## Instalacion

Para poder instalar el SDK de Netsocs, es necesario tener instalado Go en su computadora. Ademas debe tener en cuenta que el SDK de Netsocs es privado y necesitara tener acceso a el.

Ademas debe generar un [token de github](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens) y agregarlo a su configuracion de Go.

1. Primeramente, debera configurar su Go para que pueda acceder al SDK de Netsocs. Para ello, debera agregar el siguiente comando a su terminal:

```bash
vim ~/.netrc
```

2. Luego, debera agregar la siguiente linea al archivo `.netrc`:

```bash
machine github.com login <your_github_username> password <your_github_token>
```

3. Especificar en la variable de entorno `GOPRIVATE` la URL del repositorio privado de Netsocs:

```bash
export GOPRIVATE=github.com/Netsocs-Team/netsocs_go_sdk
```

4. Luego, debera agregar el SDK de Netsocs a su proyecto Go:

```bash
go get github.com/Netsocs-Team/netsocs_go_sdk
```

## Variables de entorno para la configuracion

Las variables de entorno pueden depender de la funcionalidad que se este utilizando. A continuacion se detallan las variables de entorno que se pueden utilizar:

### Device Management

- `DEVICE_MANAGEMENT_API_HOST`: URL de la API de Device Management. Ejemplo: https://device_management.netsocs.com
- `DEVICE_MANAGEMENT_API_USERNAME`: Usuario de la API de Device Management
- `DEVICE_MANAGEMENT_API_PASSWORD`: Password de la API de Device Management

### Configuration Module
- `CONFIGURATION_API_HOST`: URL de la API de Configuration. Ejemplo: https://configuration.netsocs.com