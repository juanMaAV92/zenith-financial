import { z } from 'zod';

// Login form schema
export const loginSchema = z.object({
  email: z.email('Ingresa un email válido'),
  password: z
    .string()
    .min(1, 'La contraseña es obligatoria')
    .min(6, 'La contraseña debe tener al menos 6 caracteres'),
});

// Register form schema
export const registerSchema = z.object({
  email: z.email('Ingresa un email válido'),
  user_name: z
    .string()
    .min(1, 'El nombre de usuario es obligatorio')
    .min(3, 'El nombre de usuario debe tener al menos 3 caracteres')
    .max(30, 'El nombre de usuario no puede tener más de 30 caracteres')
    .regex(/^[a-zA-Z0-9_]+$/, 'Solo se permiten letras, números y guiones bajos'),
  password: z
    .string()
    .min(1, 'La contraseña es obligatoria')
    .min(8, 'La contraseña debe tener al menos 8 caracteres')
    .regex(/[A-Z]/, 'La contraseña debe contener al menos una mayúscula')
    .regex(/[a-z]/, 'La contraseña debe contener al menos una minúscula')
    .regex(/[0-9]/, 'La contraseña debe contener al menos un número'),
  currency: z
    .string()
    .min(1, 'Selecciona tu moneda principal'),
});

// Type definitions
export type LoginFormData = z.infer<typeof loginSchema>;
export type RegisterFormData = z.infer<typeof registerSchema>;

// Currency options for the select
export const currencyOptions = [
  { value: 'USD', label: 'USD - Dólar Estadounidense' },
  { value: 'COP', label: 'COP - Peso Colombiano' },
  { value: 'EUR', label: 'EUR - Euro' },
] as const;

// API Response types
export interface RegisterSuccessResponse {
  code: string;
  user_name: string;
  email: string;
  currency: string;
  created_at: string;
}

export interface LoginSuccessResponse {
  code: string;
  user_name: string;
  email: string;
  currency: string;
  created_at: string;
  access_token: string;
  refresh_token: string;
}

export interface APIErrorResponse {
  code: string;
  messages: string[];
}
