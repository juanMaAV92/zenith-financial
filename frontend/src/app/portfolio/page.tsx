'use client';

import MainLayout from '@/components/layout/MainLayout';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Badge } from '@/components/ui/badge';
import { Plus, TrendingUp, TrendingDown } from 'lucide-react';

// Mock data para desarrollo
const mockAssets = [
  {
    id: 1,
    name: 'Apple Inc.',
    symbol: 'AAPL',
    category: 'STOCK',
    units: 50,
    investedTotal: 7500,
    currentValue: 9000,
    currency: 'USD'
  },
  {
    id: 2,
    name: 'Bitcoin',
    symbol: 'BTC',
    category: 'CRYPTO',
    units: 0.5,
    investedTotal: 20000,
    currentValue: 25000,
    currency: 'USD'
  },
  {
    id: 3,
    name: 'Cuenta de Ahorros',
    symbol: 'SAVINGS',
    category: 'SAVINGS_ACCOUNT',
    units: 15000,
    investedTotal: 15000,
    currentValue: 15000,
    currency: 'USD'
  }
];

const categoryLabels: Record<string, string> = {
  STOCK: 'Acción',
  ETF: 'ETF',
  CRYPTO: 'Criptomoneda',
  CASH: 'Efectivo',
  SAVINGS_ACCOUNT: 'Cuenta de Ahorro',
  FIXED_INCOME: 'Renta Fija',
  MUTUAL_FUND: 'Fondo Mutuo',
  COMMODITY: 'Commodity',
  CURRENCIES: 'Divisa',
  OTHER: 'Otro',
};

export default function PortfolioPage() {
  const formatCurrency = (amount: number, currency = 'USD') => {
    return new Intl.NumberFormat('es-CO', {
      style: 'currency',
      currency: currency,
      minimumFractionDigits: 0,
      maximumFractionDigits: 2,
    }).format(amount);
  };

  const calculateGainLoss = (currentValue: number, investedTotal: number) => {
    const gainLoss = currentValue - investedTotal;
    const percentage = investedTotal > 0 ? (gainLoss / investedTotal) * 100 : 0;
    return { gainLoss, percentage };
  };

  return (
    <MainLayout>
      <div className="space-y-6">
        {/* Header */}
        <div className="flex items-center justify-between">
          <div>
            <h1 className="text-3xl font-bold">Portafolio</h1>
            <p className="text-muted-foreground">
              Gestiona tus activos e inversiones
            </p>
          </div>
          <Button>
            <Plus className="h-4 w-4 mr-2" />
            Agregar Activo
          </Button>
        </div>

        {/* Assets Grid */}
        <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
          {mockAssets.map((asset) => {
            const { gainLoss, percentage } = calculateGainLoss(asset.currentValue, asset.investedTotal);
            const isPositive = gainLoss >= 0;

            return (
              <Card key={asset.id} className="hover:shadow-md transition-shadow">
                <CardHeader className="pb-3">
                  <div className="flex items-center justify-between">
                    <CardTitle className="text-lg">{asset.name}</CardTitle>
                    <Badge variant="outline">
                      {categoryLabels[asset.category] || asset.category}
                    </Badge>
                  </div>
                  <p className="text-sm text-muted-foreground">{asset.symbol}</p>
                </CardHeader>
                <CardContent className="space-y-3">
                  <div className="flex justify-between items-center">
                    <span className="text-sm text-muted-foreground">Valor Actual</span>
                    <span className="font-semibold">
                      {formatCurrency(asset.currentValue, asset.currency)}
                    </span>
                  </div>
                  
                  <div className="flex justify-between items-center">
                    <span className="text-sm text-muted-foreground">Invertido</span>
                    <span className="text-sm">
                      {formatCurrency(asset.investedTotal, asset.currency)}
                    </span>
                  </div>
                  
                  <div className="flex justify-between items-center">
                    <span className="text-sm text-muted-foreground">Unidades</span>
                    <span className="text-sm">{asset.units}</span>
                  </div>
                  
                  <div className="flex justify-between items-center pt-2 border-t">
                    <div className="flex items-center gap-1">
                      {isPositive ? (
                        <TrendingUp className="h-4 w-4 text-green-600" />
                      ) : (
                        <TrendingDown className="h-4 w-4 text-red-600" />
                      )}
                      <span className={`text-sm font-medium ${isPositive ? 'text-green-600' : 'text-red-600'}`}>
                        {percentage >= 0 ? '+' : ''}{percentage.toFixed(1)}%
                      </span>
                    </div>
                    <span className={`text-sm font-medium ${isPositive ? 'text-green-600' : 'text-red-600'}`}>
                      {gainLoss >= 0 ? '+' : ''}{formatCurrency(gainLoss, asset.currency)}
                    </span>
                  </div>
                </CardContent>
              </Card>
            );
          })}
        </div>

        {/* Empty State */}
        {mockAssets.length === 0 && (
          <Card>
            <CardContent className="flex flex-col items-center justify-center py-12">
              <div className="text-center space-y-4">
                <div className="mx-auto w-12 h-12 bg-muted rounded-full flex items-center justify-center">
                  <Plus className="h-6 w-6 text-muted-foreground" />
                </div>
                <div>
                  <h3 className="text-lg font-semibold">No hay activos en tu portafolio</h3>
                  <p className="text-muted-foreground">
                    Comienza agregando tu primer activo para empezar a trackear tus inversiones
                  </p>
                </div>
                <Button>
                  <Plus className="h-4 w-4 mr-2" />
                  Agregar tu primer activo
                </Button>
              </div>
            </CardContent>
          </Card>
        )}
      </div>
    </MainLayout>
  );
}
