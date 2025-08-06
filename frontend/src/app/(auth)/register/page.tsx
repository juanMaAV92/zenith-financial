'use client';

import { useState } from 'react';
import Link from 'next/link';
import { useRouter } from 'next/navigation';
import { useForm, Controller } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import axios from 'axios';
import { toast } from 'sonner';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { Mail, Lock, User, Eye, EyeOff, DollarSign } from 'lucide-react';
import { RegisterFormData, registerSchema, currencyOptions, RegisterSuccessResponse, APIErrorResponse } from '@/types/auth';

export default function RegisterPage() {
  const router = useRouter();
  const [showPassword, setShowPassword] = useState(false);
  const [isLoading, setIsLoading] = useState(false);

  const {
    register,
    handleSubmit,
    formState: { errors },
    control,
  } = useForm<RegisterFormData>({
    resolver: zodResolver(registerSchema),
    defaultValues: {
      email: '',
      user_name: '',
      password: '',
      currency: '',
    },
  });

  const onSubmit = async (data: RegisterFormData) => {
    try {
      setIsLoading(true);
      
      // Usar proxy temporal para evitar problemas de CORS
      const apiUrl = '/api/backend';
      
      const response = await axios.post<RegisterSuccessResponse>(
        `${apiUrl}/v1/users/register`,
        {
          user_name: data.user_name,
          password: data.password,
          email: data.email,
          currency: data.currency,
        },
        {
          headers: {
            'Content-Type': 'application/json',
          },
        }
      );

      // En caso de éxito, la respuesta contiene los datos del usuario creado
      console.log('Usuario registrado exitosamente:', response.data);
      
      // Mostrar notificación de éxito
      toast.success('¡Cuenta creada exitosamente! Redirigiendo al login...');
      
      // Éxito - redirigir al login con un pequeño delay para que se vea el toast
      setTimeout(() => {
        router.push('/login');
      }, 1500);
      
    } catch (error) {
      if (axios.isAxiosError(error) && error.response) {
        // Error de la API
        const errorData = error.response.data as APIErrorResponse;
        if (errorData.messages && errorData.messages.length > 0) {
          toast.error(errorData.messages[0]);
        } else {
          toast.error('Error al crear la cuenta. Inténtalo de nuevo.');
        }
      } else {
        // Error de red u otro tipo
        toast.error('Error de conexión. Verifica tu conexión a internet.');
      }
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <Card className="w-full max-w-md mx-auto">
      <CardHeader className="space-y-1 text-center">
        <CardTitle className="text-2xl font-bold">Crear cuenta</CardTitle>
        <CardDescription>
          Completa el formulario para crear tu cuenta y comenzar a gestionar tu portafolio
        </CardDescription>
      </CardHeader>
      <CardContent>
        <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
          {/* Email Field */}
          <div className="space-y-2">
            <Label htmlFor="email">Email</Label>
            <div className="relative">
              <Mail className="absolute left-3 top-1/2 transform -translate-y-1/2 h-4 w-4 text-muted-foreground" />
              <Input
                id="email"
                type="email"
                placeholder="tu@email.com"
                className="pl-10"
                {...register('email')}
                disabled={isLoading}
              />
            </div>
            {errors.email && (
              <p className="text-sm text-red-600">{errors.email.message}</p>
            )}
          </div>

          {/* Username Field */}
          <div className="space-y-2">
            <Label htmlFor="user_name">Nombre de usuario</Label>
            <div className="relative">
              <User className="absolute left-3 top-1/2 transform -translate-y-1/2 h-4 w-4 text-muted-foreground" />
              <Input
                id="user_name"
                type="text"
                placeholder="tu_usuario"
                className="pl-10"
                {...register('user_name')}
                disabled={isLoading}
              />
            </div>
            {errors.user_name && (
              <p className="text-sm text-red-600">{errors.user_name.message}</p>
            )}
          </div>

          {/* Password Field */}
          <div className="space-y-2">
            <Label htmlFor="password">Contraseña</Label>
            <div className="relative">
              <Lock className="absolute left-3 top-1/2 transform -translate-y-1/2 h-4 w-4 text-muted-foreground" />
              <Input
                id="password"
                type={showPassword ? 'text' : 'password'}
                placeholder="Tu contraseña segura"
                className="pl-10 pr-10"
                {...register('password')}
                disabled={isLoading}
              />
              <button
                type="button"
                onClick={() => setShowPassword(!showPassword)}
                className="absolute right-3 top-1/2 transform -translate-y-1/2 text-muted-foreground hover:text-foreground"
                disabled={isLoading}
              >
                {showPassword ? (
                  <EyeOff className="h-4 w-4" />
                ) : (
                  <Eye className="h-4 w-4" />
                )}
              </button>
            </div>
            {errors.password && (
              <p className="text-sm text-red-600">{errors.password.message}</p>
            )}
            <div className="text-xs text-muted-foreground space-y-1">
              <p>La contraseña debe contener:</p>
              <ul className="list-disc list-inside space-y-0.5 ml-2">
                <li>Al menos 8 caracteres</li>
                <li>Una letra mayúscula</li>
                <li>Una letra minúscula</li>
                <li>Un número</li>
              </ul>
            </div>
          </div>

          {/* Currency Field */}
          <div className="space-y-2">
            <Label htmlFor="currency">Moneda principal</Label>
            <Controller
              name="currency"
              control={control}
              render={({ field }) => (
                <Select
                  onValueChange={field.onChange}
                  value={field.value}
                  disabled={isLoading}
                >
                  <SelectTrigger className="w-full">
                    <div className="flex items-center">
                      <DollarSign className="h-4 w-4 text-muted-foreground mr-2" />
                      <SelectValue placeholder="Selecciona tu moneda principal" />
                    </div>
                  </SelectTrigger>
                  <SelectContent>
                    {currencyOptions.map((option) => (
                      <SelectItem key={option.value} value={option.value}>
                        {option.label}
                      </SelectItem>
                    ))}
                  </SelectContent>
                </Select>
              )}
            />
            {errors.currency && (
              <p className="text-sm text-red-600">{errors.currency.message}</p>
            )}
          </div>

          {/* Submit Button */}
          <Button 
            type="submit" 
            className="w-full" 
            disabled={isLoading}
          >
            {isLoading ? (
              <>
                <div className="animate-spin rounded-full h-4 w-4 border-b-2 border-white mr-2" />
                Creando cuenta...
              </>
            ) : (
              'Crear cuenta'
            )}
          </Button>
        </form>

        {/* Login Link */}
        <div className="mt-6 text-center">
          <p className="text-sm text-muted-foreground">
            ¿Ya tienes cuenta?{' '}
            <Link 
              href="/login" 
              className="font-medium text-primary hover:text-primary/80 transition-colors"
            >
              Inicia sesión
            </Link>
          </p>
        </div>
      </CardContent>
    </Card>
  );
}
