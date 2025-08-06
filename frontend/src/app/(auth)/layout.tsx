import { ReactNode } from 'react';
import { DollarSign } from 'lucide-react';

interface AuthLayoutProps {
  children: ReactNode;
}

export default function AuthLayout({ children }: AuthLayoutProps) {
  return (
    <div className="min-h-screen bg-background flex items-center justify-center p-4">
      <div className="w-full max-w-md space-y-8">
        {/* Logo */}
        <div className="flex flex-col items-center space-y-2">
          <div className="flex items-center justify-center space-x-3">
            <div className="flex h-8 w-8 items-center justify-center rounded-xl bg-primary text-primary-foreground">
              <DollarSign className="h-6 w-6" />
            </div>
            <h1 className="text-2xl font-bold text-foreground">Zenith</h1>
          </div>
          <p className="text-sm text-muted-foreground text-center">
            Tu plataforma de gestión financiera personal
          </p>
        </div>
        
        {/* Form Content */}
        {children}
      </div>
    </div>
  );
}
