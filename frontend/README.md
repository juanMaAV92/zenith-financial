# Zenith Financial - Frontend

Frontend moderno para la aplicación de gestión de portafolio de inversiones Zenith Financial, construido con Next.js 14, TypeScript y Tailwind CSS.

## 🚀 Características

- **Next.js 14** con App Router para un rendimiento óptimo
- **TypeScript** para desarrollo type-safe
- **Tailwind CSS** para estilos modernos y responsivos
- **shadcn/ui** para componentes de UI accesibles y personalizables
- **Zustand** para gestión de estado global
- **Axios** para comunicación con la API REST
- **Lucide React** para iconos consistentes
- **Fuente Inter** para una tipografía limpia y moderna

## 🎨 Diseño

### Paleta de Colores
- **Fondo**: `#F9FAFB` (Gris claro)
- **Primario**: `#0D9488` (Verde esmeralda oscuro)
- **Secundario**: `#164E63` (Azul petróleo oscuro)
- **Acento**: `#F59E0B` (Naranja cálido)
- **Texto principal**: `#111827`
- **Texto secundario**: `#6B7280`

### Inspiración
Dashboards financieros modernos como CoinMarketCap, Binance Lite, Fintual y Personal Capital.

## 📁 Estructura del Proyecto

```
src/
├── app/                    # App Router de Next.js
│   ├── dashboard/         # Página del dashboard principal
│   ├── portfolio/         # Página de gestión de portafolio
│   ├── globals.css        # Estilos globales con variables CSS
│   ├── layout.tsx         # Layout raíz con configuración de fuentes
│   └── page.tsx           # Página de inicio (redirige a dashboard)
├── components/            # Componentes reutilizables
│   ├── dashboard/         # Componentes específicos del dashboard
│   ├── layout/           # Componentes de layout (navegación, etc.)
│   ├── portfolio/        # Componentes del portafolio
│   └── ui/               # Componentes base de shadcn/ui
├── lib/                  # Utilidades y configuraciones
│   ├── api/             # Cliente y servicios de API
│   └── utils.ts         # Utilidades compartidas
├── store/               # Estado global con Zustand
├── types/               # Definiciones de tipos TypeScript
└── ...
```

## 🛠️ Instalación

1. **Navegar al directorio del frontend:**
   ```bash
   cd frontend
   ```

2. **Instalar dependencias:**
   ```bash
   npm install
   ```

3. **Configurar variables de entorno:**
   - Copiar `.env.local` y ajustar la URL de la API según sea necesario
   - Por defecto apunta a `http://localhost:8080`

4. **Iniciar el servidor de desarrollo:**
   ```bash
   npm run dev
   ```

5. **Abrir en el navegador:**
   - http://localhost:3000

## 📋 Scripts Disponibles

- `npm run dev` - Inicia el servidor de desarrollo
- `npm run build` - Construye la aplicación para producción
- `npm run start` - Inicia el servidor de producción
- `npm run lint` - Ejecuta el linter de ESLint

## 🔗 Integración con Backend

El frontend está diseñado para consumir la API REST del backend de Go. Los servicios de API se encuentran en `src/lib/api/services.ts` e incluyen:

- **Health Check**: Verificación del estado del backend
- **User Service**: Gestión de datos de usuario
- **Asset Service**: CRUD de activos de inversión
- **Transaction Service**: CRUD de transacciones
- **Dashboard Service**: Datos agregados para el dashboard

## 📱 Responsividad

El diseño sigue un enfoque **Mobile-first** que escala elegantemente:
- **Mobile**: Navegación con menú hamburguesa
- **Tablet**: Vistas optimizadas para pantallas medianas
- **Desktop**: Sidebar permanente y layouts de múltiples columnas

## 🧩 Componentes Principales

### Dashboard
- `SummaryCards`: Tarjetas de resumen financiero
- `AssetAllocationChart`: Distribución de activos por categoría
- `RecentTransactions`: Lista de transacciones recientes
- `TopAssets`: Principales activos del portafolio

### Layout
- `MainLayout`: Layout principal con navegación lateral
- Navegación responsive con sheet para móviles

## 🚀 Próximas Funcionalidades

El scaffolding está preparado para agregar fácilmente:
- Gráficos interactivos de rendimiento temporal
- Formularios para agregar/editar activos y transacciones
- Sistema de alertas y notificaciones
- Integración con APIs de precios en tiempo real
- Exportación de reportes
- Simuladores de inversión
- Integración con IA para recomendaciones

## 🔧 Desarrollo

### Agregar Nuevos Componentes shadcn/ui
```bash
npx shadcn@latest add [component-name]
```

### Estructura de Datos
Los tipos TypeScript están definidos en `src/types/index.ts` y reflejan la estructura de la base de datos PostgreSQL del backend.

### Estado Global
Se utiliza Zustand para el estado global. El store principal está en `src/store/index.ts`.

## 📄 Licencia

Este proyecto es parte del sistema Zenith Financial.
