'use client';

import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Badge } from '@/components/ui/badge';
import { AssetAllocation } from '@/types';

interface AssetAllocationChartProps {
  allocations: AssetAllocation[];
}

const categoryColors: Record<string, string> = {
  STOCK: 'bg-blue-500',
  ETF: 'bg-green-500',
  CRYPTO: 'bg-orange-500',
  CASH: 'bg-gray-500',
  SAVINGS_ACCOUNT: 'bg-emerald-500',
  FIXED_INCOME: 'bg-purple-500',
  MUTUAL_FUND: 'bg-pink-500',
  COMMODITY: 'bg-yellow-500',
  CURRENCIES: 'bg-indigo-500',
  OTHER: 'bg-slate-500',
};

const categoryLabels: Record<string, string> = {
  STOCK: 'Acciones',
  ETF: 'ETFs',
  CRYPTO: 'Criptomonedas',
  CASH: 'Efectivo',
  SAVINGS_ACCOUNT: 'Cuentas de Ahorro',
  FIXED_INCOME: 'Renta Fija',
  MUTUAL_FUND: 'Fondos Mutuos',
  COMMODITY: 'Commodities',
  CURRENCIES: 'Divisas',
  OTHER: 'Otros',
};

export default function AssetAllocationChart({ allocations }: AssetAllocationChartProps) {
  const formatCurrency = (amount: number) => {
    return new Intl.NumberFormat('es-CO', {
      style: 'currency',
      currency: 'USD',
      minimumFractionDigits: 0,
      maximumFractionDigits: 0,
    }).format(amount);
  };

  const sortedAllocations = [...allocations].sort((a, b) => b.value - a.value);
  const totalValue = allocations.reduce((sum, allocation) => sum + allocation.value, 0);

  return (
    <Card className="col-span-1">
      <CardHeader>
        <CardTitle className="text-lg font-semibold">Distribución por Categoría</CardTitle>
      </CardHeader>
      <CardContent>
        <div className="space-y-4">
          {/* Progress bars */}
          <div className="space-y-3">
            {sortedAllocations.map((allocation) => (
              <div key={allocation.category} className="space-y-1">
                <div className="flex items-center justify-between text-sm">
                  <div className="flex items-center gap-2">
                    <div 
                      className={`h-3 w-3 rounded-full ${categoryColors[allocation.category] || 'bg-gray-500'}`}
                    />
                    <span className="font-medium">
                      {categoryLabels[allocation.category] || allocation.category}
                    </span>
                    <Badge variant="outline" className="text-xs">
                      {allocation.count}
                    </Badge>
                  </div>
                  <span className="text-muted-foreground">
                    {allocation.percentage.toFixed(1)}%
                  </span>
                </div>
                <div className="h-2 rounded-full bg-muted overflow-hidden">
                  <div
                    className={`h-full transition-all duration-300 ${categoryColors[allocation.category] || 'bg-gray-500'}`}
                    style={{ width: `${allocation.percentage}%` }}
                  />
                </div>
                <div className="text-xs text-muted-foreground">
                  {formatCurrency(allocation.value)}
                </div>
              </div>
            ))}
          </div>

          {/* Summary */}
          <div className="border-t pt-4">
            <div className="flex justify-between items-center text-sm font-medium">
              <span>Total</span>
              <span>{formatCurrency(totalValue)}</span>
            </div>
            <div className="flex justify-between items-center text-xs text-muted-foreground mt-1">
              <span>Activos</span>
              <span>{allocations.reduce((sum, allocation) => sum + allocation.count, 0)}</span>
            </div>
          </div>
        </div>
      </CardContent>
    </Card>
  );
}
