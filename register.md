
# Registro de Usuario - Zenith Financial

## Requisitos Técnicos y de Implementación

### Archivo a Modificar

- `app/(auth)/register/page.tsx`

### Gestión de la URL del API

- La URL base del backend debe ser manejada como una variable de entorno.
- Crea un archivo `.env.local` en la raíz del proyecto.
- Añade la siguiente variable:

```env
NEXT_PUBLIC_API_URL=http://localhost:8080/zenith-financial
```

### Lógica de Envío del Formulario (`onSubmit`)

1. Utiliza **axios** para realizar la llamada a la API.
2. Implementa un estado de carga (`isLoading`) usando `useState` para dar feedback visual al usuario durante la petición.
3. El botón "Crear cuenta" debe mostrar un spinner y estar deshabilitado mientras `isLoading` sea `true`.

### Manejo de Respuestas del API

- **En caso de Éxito (HTTP 2xx):**
  - El backend responderá con los datos del usuario creado.
  - Redirige al usuario a la página de login (`/login`) utilizando el hook `useRouter` de `next/navigation`.

- **En caso de Error (HTTP 4xx/5xx):**
  - El backend responderá con un JSON que contiene un mensaje de error.
  - Utiliza el sistema de notificaciones **Toast** de shadcn/ui. Muestra el mensaje de error (`messages[0]`) en una notificación de tipo "destructivo" (rojo).

> **Importante:** Asegúrate de que el componente `<Toaster />` esté añadido en el layout raíz (`app/layout.tsx`) para que las notificaciones funcionen.

---

## Especificaciones del Endpoint

- **Método:** POST
- **Endpoint:** `/v1/users/register` (se concatena a la variable de entorno)
- **URL Completa de Ejemplo:**

  ```
  http://localhost:8080/zenith-financial/v1/users/register
  ```

- **Headers:**
  - `'Content-Type': 'application/json'`

- **Cuerpo de la Petición (Request Body):**

  ```json
  {
    "user_name": "juan22",
    "password": "12345677",
    "email": "test21@mail.com",
    "currency": "USD"
  }
  ```

---

## Formatos de Respuesta a Manejar

### Respuesta Exitosa

```json
{
  "code": "f4a654ff-6f9b-40fa-b3e1-f0b806db1d13",
  "user_name": "juan2a2",
  "email": "test21s@mail.com",
  "currency": "USD",
  "created_at": "2024-08-06T15:37:04.900188Z"
}
```

### Respuesta de Error (Ej: Usuario ya existe)

```json
{
  "code": "USER_EXISTS",
  "messages": [
    "User already exists"
  ]
}
```