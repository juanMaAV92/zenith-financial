'use client';

import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Badge } from '@/components/ui/badge';
import { 
  TrendingUp, 
  TrendingDown, 
  DollarSign, 
  Target,
  Percent
} from 'lucide-react';
import { PortfolioSummary } from '@/types';

interface SummaryCardsProps {
  summary: PortfolioSummary;
}

export default function SummaryCards({ summary }: SummaryCardsProps) {
  const formatCurrency = (amount: number, currency = 'USD') => {
    return new Intl.NumberFormat('es-CO', {
      style: 'currency',
      currency: currency,
      minimumFractionDigits: 0,
      maximumFractionDigits: 0,
    }).format(amount);
  };

  const formatPercentage = (percentage: number) => {
    return `${percentage >= 0 ? '+' : ''}${percentage.toFixed(2)}%`;
  };

  const isPositive = summary.totalGainLoss >= 0;

  return (
    <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
      {/* Total Value */}
      <Card>
        <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
          <CardTitle className="text-sm font-medium text-muted-foreground">
            Valor Total
          </CardTitle>
          <DollarSign className="h-4 w-4 text-muted-foreground" />
        </CardHeader>
        <CardContent>
          <div className="text-2xl font-bold">
            {formatCurrency(summary.totalValue, summary.currency)}
          </div>
          <p className="text-xs text-muted-foreground">
            Valor actual del portafolio
          </p>
        </CardContent>
      </Card>

      {/* Total Invested */}
      <Card>
        <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
          <CardTitle className="text-sm font-medium text-muted-foreground">
            Total Invertido
          </CardTitle>
          <Target className="h-4 w-4 text-muted-foreground" />
        </CardHeader>
        <CardContent>
          <div className="text-2xl font-bold">
            {formatCurrency(summary.totalInvested, summary.currency)}
          </div>
          <p className="text-xs text-muted-foreground">
            Capital total invertido
          </p>
        </CardContent>
      </Card>

      {/* Gain/Loss Amount */}
      <Card>
        <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
          <CardTitle className="text-sm font-medium text-muted-foreground">
            Ganancia/Pérdida
          </CardTitle>
          {isPositive ? (
            <TrendingUp className="h-4 w-4 text-green-600" />
          ) : (
            <TrendingDown className="h-4 w-4 text-red-600" />
          )}
        </CardHeader>
        <CardContent>
          <div className={`text-2xl font-bold ${isPositive ? 'text-green-600' : 'text-red-600'}`}>
            {formatCurrency(summary.totalGainLoss, summary.currency)}
          </div>
          <p className="text-xs text-muted-foreground">
            Ganancia/pérdida absoluta
          </p>
        </CardContent>
      </Card>

      {/* Gain/Loss Percentage */}
      <Card>
        <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
          <CardTitle className="text-sm font-medium text-muted-foreground">
            Rendimiento
          </CardTitle>
          <Percent className="h-4 w-4 text-muted-foreground" />
        </CardHeader>
        <CardContent>
          <div className="flex items-center space-x-2">
            <span className={`text-2xl font-bold ${isPositive ? 'text-green-600' : 'text-red-600'}`}>
              {formatPercentage(summary.gainLossPercentage)}
            </span>
            <Badge variant={isPositive ? 'default' : 'destructive'} className="text-xs">
              {isPositive ? 'Ganancia' : 'Pérdida'}
            </Badge>
          </div>
          <p className="text-xs text-muted-foreground">
            Rendimiento porcentual
          </p>
        </CardContent>
      </Card>
    </div>
  );
}
