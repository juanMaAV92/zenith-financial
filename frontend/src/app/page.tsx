'use client';

import { useEffect } from 'react';
import { useRouter } from 'next/navigation';
import Link from 'next/link';
import { Button } from '@/components/ui/button';
import { DollarSign } from 'lucide-react';

export default function Home() {
  const router = useRouter();

  useEffect(() => {
    // Verificar si hay un token de autenticación
    const token = localStorage.getItem('authToken');
    
    if (token) {
      // Si hay token, redirigir al dashboard
      router.replace('/dashboard');
    }
  }, [router]);

  return (
    <div className="min-h-screen bg-background flex items-center justify-center p-4">
      <div className="text-center space-y-8 max-w-md">
        {/* Logo */}
        <div className="flex flex-col items-center space-y-4">
          <div className="flex h-16 w-16 items-center justify-center rounded-xl bg-primary text-primary-foreground">
            <DollarSign className="h-8 w-8" />
          </div>
          <div>
            <h1 className="text-3xl font-bold text-foreground">Zenith Financial</h1>
            <p className="text-muted-foreground mt-2">
              Gestiona tu portafolio de inversiones de manera inteligente
            </p>
          </div>
        </div>

        {/* Features */}
        <div className="space-y-4 text-sm text-muted-foreground">
          <div className="grid grid-cols-1 gap-3">
            <div className="flex items-center justify-center space-x-2">
              <div className="w-2 h-2 bg-primary rounded-full"></div>
              <span>Trackea acciones, criptomonedas y más</span>
            </div>
            <div className="flex items-center justify-center space-x-2">
              <div className="w-2 h-2 bg-primary rounded-full"></div>
              <span>Analiza el rendimiento de tus inversiones</span>
            </div>
            <div className="flex items-center justify-center space-x-2">
              <div className="w-2 h-2 bg-primary rounded-full"></div>
              <span>Mantén un registro completo de transacciones</span>
            </div>
          </div>
        </div>

        {/* Action Buttons */}
        <div className="space-y-3">
          <Button asChild className="w-full">
            <Link href="/login">
              Iniciar Sesión
            </Link>
          </Button>
          <Button asChild variant="outline" className="w-full">
            <Link href="/register">
              Crear Cuenta
            </Link>
          </Button>
        </div>

        {/* Demo Info */}
        <div className="p-3 rounded-md bg-blue-50 border border-blue-200">
          <p className="text-xs text-blue-600">
            <strong>Demo disponible:</strong> demo@zenith.com / password123
          </p>
        </div>
      </div>
    </div>
  );
}
