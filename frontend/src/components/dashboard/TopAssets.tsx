'use client';

import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Badge } from '@/components/ui/badge';
import { Button } from '@/components/ui/button';
import Link from 'next/link';
import { Asset } from '@/types';
import { TrendingUp, TrendingDown, Eye } from 'lucide-react';

interface TopAssetsProps {
  assets: Asset[];
}

const categoryLabels: Record<string, string> = {
  STOCK: 'Acción',
  ETF: 'ETF',
  CRYPTO: 'Crypto',
  CASH: 'Efectivo',
  SAVINGS_ACCOUNT: 'Ahorro',
  FIXED_INCOME: 'Renta Fija',
  MUTUAL_FUND: 'Fondo',
  COMMODITY: 'Commodity',
  CURRENCIES: 'Divisa',
  OTHER: 'Otro',
};

export default function TopAssets({ assets }: TopAssetsProps) {
  const formatCurrency = (amount: number, currency = 'USD') => {
    return new Intl.NumberFormat('es-CO', {
      style: 'currency',
      currency: currency,
      minimumFractionDigits: 0,
      maximumFractionDigits: 2,
    }).format(amount);
  };

  const calculateCurrentValue = (asset: Asset) => {
    // Si tiene valor manual, usarlo, sino calcular basado en precio actual
    if (asset.current_value !== null && asset.current_value !== undefined) {
      return asset.current_value * asset.total_units;
    }
    // Por ahora usamos el valor invertido como placeholder
    return asset.invested_total;
  };

  const calculateGainLoss = (asset: Asset) => {
    const currentValue = calculateCurrentValue(asset);
    const gainLoss = currentValue - asset.invested_total;
    const percentage = asset.invested_total > 0 ? (gainLoss / asset.invested_total) * 100 : 0;
    return { gainLoss, percentage };
  };

  return (
    <Card className="col-span-1">
      <CardHeader className="flex flex-row items-center justify-between">
        <CardTitle className="text-lg font-semibold">Principales Activos</CardTitle>
        <Button variant="outline" size="sm" asChild>
          <Link href="/portfolio">
            <Eye className="h-4 w-4 mr-2" />
            Ver portafolio
          </Link>
        </Button>
      </CardHeader>
      <CardContent>
        <div className="space-y-4">
          {assets.length === 0 ? (
            <div className="text-center py-8 text-muted-foreground">
              <p>No hay activos en el portafolio</p>
              <Button variant="outline" size="sm" className="mt-2" asChild>
                <Link href="/portfolio">Agregar activo</Link>
              </Button>
            </div>
          ) : (
            assets.map((asset) => {
              const currentValue = calculateCurrentValue(asset);
              const { gainLoss, percentage } = calculateGainLoss(asset);
              const isPositive = gainLoss >= 0;
              
              return (
                <div key={asset.id} className="flex items-center gap-3 p-3 rounded-lg border">
                  <div className="flex-1 min-w-0">
                    <div className="flex items-center gap-2">
                      <p className="text-sm font-medium truncate">
                        {asset.name}
                      </p>
                      <Badge variant="outline" className="text-xs">
                        {asset.symbol}
                      </Badge>
                    </div>
                    <div className="flex items-center gap-2 mt-1">
                      <p className="text-xs text-muted-foreground">
                        {categoryLabels[asset.category?.name || 'OTHER'] || 'Otro'}
                      </p>
                      <span className="text-xs text-muted-foreground">•</span>
                      <p className="text-xs text-muted-foreground">
                        {asset.total_units} unidades
                      </p>
                    </div>
                  </div>
                  
                  <div className="text-right">
                    <p className="text-sm font-medium">
                      {formatCurrency(currentValue, asset.currency)}
                    </p>
                    <div className="flex items-center gap-1 justify-end">
                      {isPositive ? (
                        <TrendingUp className="h-3 w-3 text-green-600" />
                      ) : (
                        <TrendingDown className="h-3 w-3 text-red-600" />
                      )}
                      <span className={`text-xs ${isPositive ? 'text-green-600' : 'text-red-600'}`}>
                        {percentage >= 0 ? '+' : ''}{percentage.toFixed(1)}%
                      </span>
                    </div>
                  </div>
                </div>
              );
            })
          )}
        </div>
      </CardContent>
    </Card>
  );
}
