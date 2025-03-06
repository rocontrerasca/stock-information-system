# *Stock Information System*

Este proyecto es un sistema para obtener, almacenar y mostrar información de acciones (stocks) desde una API externa. El sistema consta de un backend en *Golang* y un frontend en *Vue 3* con *TypeScript, **Pinia* y *Tailwind CSS. Los datos se almacenan en una base de datos **CockroachDB*.

---

## *Tabla de Contenidos*

1. [Requisitos](#requisitos)
2. [Instalación](#instalación)
3. [Configuración](#configuración)
   - [Backend](#backend)
   - [Frontend](#frontend)
4. [Ejecución](#ejecución)
5. [Estructura del Proyecto](#estructura-del-proyecto)
6. [Algoritmo de Recomendaciones](#algoritmo-de-recomendaciones)
7. [Contribución](#contribución)
8. [Licencia](#licencia)

---

## *Requisitos*

- *Node.js 22+* (para el frontend).
- *Go 1.20+* (para el backend).
- *CockroachDB* (como base de datos).
- *Git* (para clonar el repositorio).

---

## *Instalación*

1. Clona el repositorio:
   bash
   git clone https://github.com/rocontrerasca/stock-information-system.git
   cd stock-information-system
   

2. Configura el backend:
   bash
   cd backend
   go mod tidy
   

3. Configura el frontend:
   bash
   cd ../frontend
   npm install
   

---

## *Configuración*

### *Backend*

1. *CockroachDB*:
   - Inicia un clúster local:
     bash
     cockroach start-single-node --insecure --listen-addr=localhost
     
   - Crea la base de datos y la tabla:
     sql
     CREATE DATABASE stocks;
     USE stocks;
     CREATE TABLE stock_data (
         ticker STRING PRIMARY KEY,
         target_from STRING,
         target_to STRING,
         company STRING,
         action STRING,
         brokerage STRING,
         rating_from STRING,
         rating_to STRING,
         time TIMESTAMP
     );
     

2. *Archivo .env*:
   - En la raíz del proyecto backend, crea un archivo .env con el siguiente contenido:
     env
     # Configuración de la base de datos
     DB_URL=postgresql://root@localhost:26257/stocks?sslmode=disable

     # Configuración de la API externa
     API_URL=https://8j5baasof2.execute-api.us-west-2.amazonaws.com/production/swechallenge/list
     API_TOKEN=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdHRlbXB0cyI6MjUsImVtYWlsIjoicm9jb250cmVyYXNjYUBnbWFpbC5jb20iLCJleHAiOjE3NDExODkyNDksImlkIjoiMCIsInBhc3N3b3JkIjoiJyBPUiAnMScgPSAnMSJ9.WGkPJNZ-P8V5pzaZS0I6aLrw-SB5xc2BQViouriHgdA
     
     # CORS
     CORS_ORIGIN=http://localhost:5173
     CORS_METHODS=GET,PUT,POST,DELETE

3. *Instalar godotenv*:
   - Ejecuta el siguiente comando para instalar el paquete godotenv:
     bash
     go get github.com/joho/godotenv
     

### *Frontend*

1. *Tailwind CSS*:
   - Configura Tailwind CSS en tailwind.config.js y postcss.config.js.
   - Importa los estilos en src/assets/main.css.

---

## *Ejecución*

1. *Backend*:
   bash
   cd backend
   go run main.go
   

2. *Frontend*:
   bash
   cd ../frontend
   npm run dev
   

3. Accede a la aplicación en http://localhost:5173.

---

## *Estructura del Proyecto*


stock-information-system/
├── backend/
│   ├── controllers/
│   ├── models/
│   ├── repositories/
│   ├── services/
│   ├── database/
│   ├── main.go
│   ├── .env
├── frontend/
│   ├── public/
│   ├── src/
│   │   ├── assets/
│   │   ├── components/
│   │   ├── stores/
│   │   ├── views/
│   │   ├── App.vue
│   │   ├── main.ts
│   │   ├── router/
│   ├── tailwind.config.js
│   ├── package.json


---

## *Algoritmo de Recomendaciones*

El sistema utiliza un algoritmo mejorado para recomendar acciones basado en:

1. *Cambio en el precio objetivo*.
2. *Mejora en la calificación*.
3. *Tiempo de la recomendación*.
4. *Credibilidad de la correduría*.

Las acciones se puntúan y ordenan para proporcionar las 5 mejores recomendaciones.

---

## *Contribución*

¡Las contribuciones son bienvenidas! Si deseas contribuir al proyecto, sigue estos pasos:

1. Haz un fork del repositorio.
2. Crea una rama para tu feature o corrección:
   bash
   git checkout -b nombre-de-tu-rama
   
3. Realiza tus cambios y haz commit:
   bash
   git commit -m "Descripción de tus cambios"
   
4. Envía un pull request.

---

## *Licencia*

Este proyecto está bajo la licencia *MIT*. Para más detalles.

---

## *Contacto*

Si tienes alguna pregunta o sugerencia, no dudes en contactarme:

- *Nombre*: Richar Contreras
- *Email*: rocontrerasca@gmail.com
- *GitHub*: https://github.com/rocontrerasca
