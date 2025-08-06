
# ¡Excelente! Vamos con la pantalla de login

Tu pregunta sobre cómo manejar los tokens es fundamental, y es la pieza clave para una autenticación segura.

---

## ¿Cómo se manejan los tokens? (La Guía Rápida)

### **access_token** (Token de Acceso)

- **¿Qué es?** Es una "llave" de corta duración (ej. 15 minutos, 1 hora).
- **¿Para qué sirve?** Se envía en el encabezado (Header) de cada petición a las rutas protegidas de tu backend (ej. para ver tu portafolio, añadir un activo, etc.). El backend lo verifica y, si es válido, te da acceso.
- **¿Cómo se envía?** Se añade un header `Authorization: Bearer <access_token>`.

### **refresh_token** (Token de Refresco)

- **¿Qué es?** Es una "llave maestra" de larga duración (ej. 7 días, 1 mes).
- **¿Para qué sirve?** Su único propósito es obtener un nuevo access_token cuando el actual expira. Nunca se usa para acceder a datos.
- **¿Cómo funciona?** Cuando el backend te dice "tu token de acceso ha expirado", el frontend hace una llamada a un endpoint especial (ej. `/v1/token/refresh`) enviando el refresh_token. El backend lo valida y te devuelve un nuevo access_token y, a veces, un nuevo refresh_token.

---

## ¿Dónde se guardan de forma segura?

La mejor práctica es guardarlos en **cookies HttpOnly**.

> **¿Por qué HttpOnly?** Esto significa que la cookie no puede ser accedida por JavaScript en el navegador. Esto te protege contra ataques XSS (Cross-Site Scripting). Si un atacante lograra inyectar un script en tu página, no podría robar los tokens.

---

## ¿Cómo se hace en Next.js?

La página de login (que es un Client Component) no puede establecer una cookie HttpOnly. Por eso, creamos una ruta de API intermedia en Next.js. El flujo es:

1. El formulario de login envía los datos al backend de Go.
2. El backend de Go devuelve los tokens al formulario.
3. El formulario envía inmediatamente los tokens a una ruta de API de Next.js (ej. `/api/auth/login`).
4. Esta ruta de API (que se ejecuta en el servidor) recibe los tokens y los establece como cookies HttpOnly en la respuesta al navegador.

---

## Prompt para el agente de programación

### **Título del Prompt:** Implementación de Autenticación en Login y Manejo Seguro de Tokens JWT

---

## Contexto del Proyecto

Continuamos con el desarrollo de **"Zenith Financial"**. Ya tenemos la UI de la página de login (`app/(auth)/login/page.tsx`). El siguiente paso es implementar la lógica de autenticación, conectarla al backend y manejar de forma segura los tokens JWT (`access_token` y `refresh_token`) recibidos.

---

## Objetivo Principal

Modificar el componente de login para que realice una llamada **POST** al backend. En caso de éxito, los tokens JWT deben ser almacenados de forma segura en cookies HttpOnly, y el usuario debe ser redirigido al dashboard principal. En caso de error, se debe notificar al usuario.

---

## Requisitos Técnicos y de Implementación

**Archivo a Modificar:**

- `app/(auth)/login/page.tsx`

**Lógica de Envío del Formulario (`onSubmit`):**

- Como en el registro, utiliza **axios** y un estado `isLoading` para gestionar la llamada a la API y el feedback visual en el botón.

**Manejo de Tokens (Flujo de Autenticación Segura):**

1. **Llamada al Backend de Go:** Al recibir una respuesta exitosa del endpoint `/v1/login`, el componente de login recibirá los tokens.
2. **Llamada a la API Route de Next.js:** Inmediatamente después, el componente debe hacer una llamada **POST** a una nueva ruta de API interna de Next.js (`/api/auth/login`) enviando los tokens recibidos en el cuerpo de la petición.
3. **Almacenamiento en Cookies HttpOnly:** La nueva ruta de API de Next.js se encargará de establecer el `access_token` y el `refresh_token` como cookies HttpOnly. Esto es **CRÍTICO** para la seguridad.
4. **Redirección:** Una vez que la ruta de API de Next.js responda con éxito, redirige al usuario al dashboard principal (crea una nueva página en `app/dashboard/page.tsx` si no existe).


---

## Manejo de Errores

Si el endpoint `/v1/login` del backend de Go devuelve un error, muestra el mensaje de error usando el componente **Toast**, igual que en el registro.

---

## Especificaciones de Endpoints y Payloads

### Endpoint de Login (Backend Go):

- **Método:** `POST`
- **URL:** `http://localhost:8080/zenith-financial/v1/login`

**Cuerpo (Request):**

```json
{
    "email": "testa21s@mail.com",
    "password" : "12345677"
}
```

**Respuesta Exitosa (Response):**

```json
{
    {
    "code": "398609dd-17a2-4cb8-a57e-bd0d3fa6b723",
    "user_name": "juana2a2",
    "email": "testa21s@mail.com",
    "currency": "USD",
    "created_at": "2025-08-06T11:30:25.543186-05:00",
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTQ1MDI2NjcsImlhdCI6MTc1NDUwMTc2NywiaXNzIjoiemVuaXRoLWZpbmFuY2lhbCIsInR5cGUiOiJhY2Nlc3MiLCJ1c2VyX2NvZGUiOiIzOTg2MDlkZC0xN2EyLTRjYjgtYTU3ZS1iZDBkM2ZhNmI3MjMifQ.9OEjtsUjeoN6TxEqIO5_wXMdj5NkoQt8yuKHjT6FtCM",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTQ1ODgxNjksImlhdCI6MTc1NDUwMTc2OSwiaXNzIjoiemVuaXRoLWZpbmFuY2lhbCIsInR5cGUiOiJhY2Nlc3MiLCJ1c2VyX2NvZGUiOiIzOTg2MDlkZC0xN2EyLTRjYjgtYTU3ZS1iZDBkM2ZhNmI3MjMifQ.l-8v8ROJaxb1wxokK7ggb4UbNvqIMg0eKJTdzl5ZHhQ"
}
}
```

**Respuesta de Error:** Formato estándar ya definido (código y array de mensajes).