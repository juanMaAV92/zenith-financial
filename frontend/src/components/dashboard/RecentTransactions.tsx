'use client';

import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Badge } from '@/components/ui/badge';
import { Button } from '@/components/ui/button';
import Link from 'next/link';
import { Transaction } from '@/types';
import { 
  TrendingUp, 
  TrendingDown, 
  Plus,
  Minus,
  Eye
} from 'lucide-react';

interface RecentTransactionsProps {
  transactions: Transaction[];
}

const transactionIcons: Record<string, React.ElementType> = {
  BUY: Plus,
  SELL: Minus,
  DEPOSIT: TrendingUp,
  WITHDRAW: TrendingDown,
};

const transactionLabels: Record<string, string> = {
  BUY: 'Compra',
  SELL: 'Venta',
  DEPOSIT: 'Depósito',
  WITHDRAW: 'Retiro',
};

const transactionColors: Record<string, string> = {
  BUY: 'text-green-600',
  SELL: 'text-red-600',
  DEPOSIT: 'text-blue-600',
  WITHDRAW: 'text-orange-600',
};

export default function RecentTransactions({ transactions }: RecentTransactionsProps) {
  const formatCurrency = (amount: number, currency = 'USD') => {
    return new Intl.NumberFormat('es-CO', {
      style: 'currency',
      currency: currency,
      minimumFractionDigits: 0,
      maximumFractionDigits: 2,
    }).format(amount);
  };

  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleDateString('es-CO', {
      day: '2-digit',
      month: 'short',
      year: 'numeric',
    });
  };

  return (
    <Card className="col-span-1">
      <CardHeader className="flex flex-row items-center justify-between">
        <CardTitle className="text-lg font-semibold">Transacciones Recientes</CardTitle>
        <Button variant="outline" size="sm" asChild>
          <Link href="/transactions">
            <Eye className="h-4 w-4 mr-2" />
            Ver todas
          </Link>
        </Button>
      </CardHeader>
      <CardContent>
        <div className="space-y-4">
          {transactions.length === 0 ? (
            <div className="text-center py-8 text-muted-foreground">
              <p>No hay transacciones recientes</p>
              <Button variant="outline" size="sm" className="mt-2" asChild>
                <Link href="/portfolio">Agregar transacción</Link>
              </Button>
            </div>
          ) : (
            transactions.map((transaction) => {
              const Icon = transactionIcons[transaction.type];
              const isOutgoing = transaction.type === 'SELL' || transaction.type === 'WITHDRAW';
              
              return (
                <div key={transaction.id} className="flex items-center gap-3 p-3 rounded-lg border">
                  <div className={`p-2 rounded-full bg-muted ${transactionColors[transaction.type]}`}>
                    <Icon className="h-4 w-4" />
                  </div>
                  
                  <div className="flex-1 min-w-0">
                    <div className="flex items-center gap-2">
                      <p className="text-sm font-medium truncate">
                        {transaction.asset?.name || `Asset #${transaction.asset_id}`}
                      </p>
                      <Badge variant="outline" className="text-xs">
                        {transactionLabels[transaction.type]}
                      </Badge>
                    </div>
                    <p className="text-xs text-muted-foreground">
                      {transaction.units} unidades • {formatDate(transaction.created_at)}
                    </p>
                    {transaction.note && (
                      <p className="text-xs text-muted-foreground truncate">
                        {transaction.note}
                      </p>
                    )}
                  </div>
                  
                  <div className="text-right">
                    <p className={`text-sm font-medium ${isOutgoing ? 'text-red-600' : 'text-green-600'}`}>
                      {isOutgoing ? '-' : '+'}
                      {formatCurrency(transaction.total, transaction.currency)}
                    </p>
                    {transaction.fee_total > 0 && (
                      <p className="text-xs text-muted-foreground">
                        Comisión: {formatCurrency(transaction.fee_total, transaction.currency)}
                      </p>
                    )}
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
