## Prompt para Agente de IA: Creación de Frontend con Next.js

---

# Título del Proyecto
**Frontend para "Zenith Financial" con Next.js**

## Resumen del Proyecto
Quiero crear el frontend de una aplicación llamada **Zenith Financial**, una herramienta web para gestionar y visualizar mi portafolio de inversiones personales. La app ya cuenta con un backend en Go y una API REST, y utiliza una base de datos PostgreSQL con activos como acciones, criptomonedas, cuentas de ahorro, CDT, dólares y ETFs.

La aplicación está orientada a un solo usuario (por ahora) que desea tener control y visibilidad sobre sus inversiones.

---

## 🎨 Diseño Visual Esperado
- **Estética:** Moderna, minimalista, clara, con enfoque en datos financieros.
- **Inspiración:** Dashboards como CoinMarketCap, Binance Lite, Fintual, Personal Capital.
- **Tipografía:** Sans serif limpia como Inter (preferida), Roboto o Manrope.

### Paleta de colores

| Elemento         | Color      | Descripción              |
|------------------|------------|--------------------------|
| Fondo            | `#F9FAFB`  | Gris claro               |
| Primario         | `#0D9488`  | Verde esmeralda oscuro   |
| Secundario       | `#164E63`  | Azul petróleo oscuro     |
| Acento           | `#F59E0B`  | Naranja cálido           |
| Texto principal  | `#111827`  |                          |
| Texto secundario | `#6B7280`  |                          |

---

## 🛠️ Stack Técnico del Frontend
- **Framework:** Next.js 14+ con React y el App Router
- **Lenguaje:** TypeScript
- **Styling:** Tailwind CSS
- **Componentes UI:** shadcn/ui (construido sobre Radix UI) para una base de componentes accesibles y personalizables
- **Estado Global:** Zustand o React Context para una gestión de estado global ligera y moderna
- **Consultas a la API REST:** Axios
- **Autenticación (futuro):** 
- **Responsivo:** Diseño Mobile-first, que escale elegantemente a tablet y desktop
- **Internacionalización:** No requerida inicialmente

---

## ✅ Objetivo
Entregar el scaffolding de una aplicación web moderna y funcional usando Next.js. El frontend debe estar desacoplado y consumir el backend mediante la API REST.

El proyecto debe aprovechar las fortalezas de Next.js, como los **Server Components** para el renderizado inicial y el fetching de datos, y los **Client Components** (`"use client"`) solo cuando sea necesaria la interactividad del usuario (ej: botones, formularios, gráficos interactivos).

Se prioriza el rendimiento, una excelente UX y una arquitectura de código que facilite agregar nuevas funcionalidades en el futuro (como alertas, IA o simuladores).

> **Nota:** No crees ninguna lógica de negocio compleja, solo el scaffolding de la aplicación con las vistas principales y componentes reutilizables para ir agregando funcionalidades posteriormente.