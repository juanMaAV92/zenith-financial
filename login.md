Diseña dos pantallas modernas y minimalistas: una de "Login" y otra de "Registro" para la aplicación web financiera "Zenith Financial". Esta aplicación permite a los usuarios gestionar sus portafolios personales de activos financieros (acciones, criptomonedas, divisas, cuentas bancarias, etc.).

🎯 Objetivo del Diseño
Transmitir seguridad, modernidad y profesionalismo.

Debe ser fácil de usar y accesible.

Layout centrado, responsivo y sin elementos visuales innecesarios.

🎨 Paleta de Colores (Zenith Financial)
Fondo principal: #F9FAFB (gris claro)

Color primario: #0D9488 (verde esmeralda oscuro)

Color secundario: #164E63 (azul petróleo oscuro)

Acento: #F59E0B (naranja cálido, para botones destacados)

Texto principal: #111827

Texto secundario: #6B7280

🖌️ Estilo Visual
Tipografía sans-serif moderna: "Inter".

Componentes base de shadcn/ui: Usar Card para el contenedor del formulario, Input para los campos de texto, Button para las acciones, Label para las etiquetas.

Estilo de los componentes: Bordes redondeados (radius 0.5rem), borde sutil en los inputs, botones sólidos con animaciones de hover y focus.

Layout: Centrado vertical y horizontalmente en la página. El formulario debe tener un ancho máximo (max-w-sm o max-w-md).

Iconos: Opcionales, usando lucide-react (integrado en shadcn/ui) para los inputs (ej. mail o candado).

Espaciado: Generoso, interlineado amplio, enfoque minimalista.

📱 Pantalla de Login
Ruta: /login

Título: “Iniciar sesión en Zenith”

Campos:

Email (tipo email)

Contraseña (tipo password)

Botón principal: “Entrar”

Link inferior: “¿No tienes cuenta? Regístrate” (debe ser un <Link> de Next.js).

Manejo de errores: Mostrar mensajes de error en rojo debajo de cada campo si la validación falla.

📝 Pantalla de Registro
Ruta: /register

Título: “Crear cuenta en Zenith”

Campos:

Email

Nombre de usuario

Contraseña

Moneda principal (usar el componente Select de shadcn/ui con opciones: USD, COP, EUR, BTC)

Botón principal: “Crear cuenta”

Link inferior: “¿Ya tienes cuenta? Inicia sesión”

⚙️ Stack Tecnológico y Estructura
Framework: Next.js 14+ con App Router y React.

Lenguaje: TypeScript.

Estilos: Tailwind CSS.

Componentes UI: shadcn/ui (pre-configurado).

Gestión de Formularios: React Hook Form para manejar el estado y la validación.

Validación de Esquema: Zod para definir las reglas de validación.

Estructura de Archivos:

Crear un grupo de rutas (auth) dentro del directorio app.

app/(auth)/layout.tsx: Un layout compartido para las páginas de login y registro que las centre en la pantalla.

app/(auth)/login/page.tsx: La página de Login.

app/(auth)/register/page.tsx: La página de Registro.

Los formularios deben ser Client Components ("use client").

💡 Extras
Añadir un logo textual simple “Zenith” centrado arriba del formulario.

No incluir la funcionalidad de "Olvidé mi contraseña" por ahora.

La interfaz debe ser accesible y mobile-friendly desde el inicio.

🚫 Importante: No incluir funciones de autenticación reales ni llamadas a la API. El objetivo es construir únicamente la UI/UX de los formularios.