'use client';

import { useEffect, useState } from 'react';
import MainLayout from '@/components/layout/MainLayout';
import SummaryCards from '@/components/dashboard/SummaryCards';
import AssetAllocationChart from '@/components/dashboard/AssetAllocationChart';
import RecentTransactions from '@/components/dashboard/RecentTransactions';
import TopAssets from '@/components/dashboard/TopAssets';
import { Button } from '@/components/ui/button';
import { RefreshCw } from 'lucide-react';
import { DashboardData } from '@/types';
import { dashboardService } from '@/lib/api/services';
import useAppStore from '@/store';

// Mock data para desarrollo
const mockDashboardData: DashboardData = {
  portfolioSummary: {
    totalValue: 125000,
    totalInvested: 100000,
    totalGainLoss: 25000,
    gainLossPercentage: 25.0,
    currency: 'USD'
  },
  assetAllocations: [
    { category: 'STOCK', value: 60000, percentage: 48, count: 8 },
    { category: 'ETF', value: 30000, percentage: 24, count: 4 },
    { category: 'CRYPTO', value: 20000, percentage: 16, count: 5 },
    { category: 'SAVINGS_ACCOUNT', value: 15000, percentage: 12, count: 2 }
  ],
  recentTransactions: [
    {
      id: 1,
      code: 'tx-001',
      asset_id: 1,
      type: 'BUY',
      units: 10,
      total: 1500,
      fee_total: 5,
      currency: 'USD',
      note: 'Compra mensual programada',
      created_at: new Date().toISOString(),
      asset: {
        id: 1,
        code: 'asset-001',
        user_id: 1,
        name: 'Apple Inc.',
        symbol: 'AAPL',
        ticker: 'AAPL',
        currency: 'USD',
        category_id: 4,
        total_units: 50,
        invested_total: 7500,
        auto_pricing_enabled: true,
        price_source: 'yahoo',
        created_at: new Date().toISOString(),
        updated_at: new Date().toISOString()
      }
    },
    {
      id: 2,
      code: 'tx-002',
      asset_id: 2,
      type: 'DEPOSIT',
      units: 1000,
      total: 1000,
      fee_total: 0,
      currency: 'USD',
      created_at: new Date(Date.now() - 24 * 60 * 60 * 1000).toISOString(),
      asset: {
        id: 2,
        code: 'asset-002',
        user_id: 1,
        name: 'Cuenta de Ahorros',
        symbol: 'SAVINGS',
        currency: 'USD',
        category_id: 2,
        total_units: 15000,
        invested_total: 15000,
        auto_pricing_enabled: false,
        created_at: new Date().toISOString(),
        updated_at: new Date().toISOString()
      }
    }
  ],
  topAssets: [
    {
      id: 1,
      code: 'asset-001',
      user_id: 1,
      name: 'Apple Inc.',
      symbol: 'AAPL',
      ticker: 'AAPL',
      currency: 'USD',
      category_id: 4,
      total_units: 50,
      invested_total: 7500,
      auto_pricing_enabled: true,
      price_source: 'yahoo',
      created_at: new Date().toISOString(),
      updated_at: new Date().toISOString(),
      category: { id: 4, name: 'STOCK', created_at: '', updated_at: '' }
    },
    {
      id: 3,
      code: 'asset-003',
      user_id: 1,
      name: 'Bitcoin',
      symbol: 'BTC',
      ticker: 'BTC-USD',
      currency: 'USD',
      category_id: 6,
      total_units: 0.5,
      invested_total: 20000,
      auto_pricing_enabled: true,
      price_source: 'coingecko',
      created_at: new Date().toISOString(),
      updated_at: new Date().toISOString(),
      category: { id: 6, name: 'CRYPTO', created_at: '', updated_at: '' }
    }
  ]
};

export default function DashboardPage() {
  const [dashboardData, setDashboardData] = useState<DashboardData | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const { setError } = useAppStore();

  const loadDashboardData = async () => {
    try {
      setIsLoading(true);
      // En desarrollo, usar datos mock
      // const response = await dashboardService.getDashboardData();
      // setDashboardData(response.data);
      
      // Simular delay de red
      await new Promise(resolve => setTimeout(resolve, 1000));
      setDashboardData(mockDashboardData);
    } catch (error) {
      console.error('Error loading dashboard data:', error);
      setError('Error al cargar los datos del dashboard');
      // En caso de error, usar datos mock como fallback
      setDashboardData(mockDashboardData);
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    loadDashboardData();
  }, []);

  if (isLoading) {
    return (
      <MainLayout>
        <div className="space-y-6">
          <div className="flex items-center justify-between">
            <h1 className="text-3xl font-bold">Dashboard</h1>
          </div>
          <div className="flex items-center justify-center h-96">
            <div className="text-center">
              <RefreshCw className="h-8 w-8 animate-spin mx-auto mb-4 text-muted-foreground" />
              <p className="text-muted-foreground">Cargando datos del dashboard...</p>
            </div>
          </div>
        </div>
      </MainLayout>
    );
  }

  if (!dashboardData) {
    return (
      <MainLayout>
        <div className="space-y-6">
          <div className="flex items-center justify-between">
            <h1 className="text-3xl font-bold">Dashboard</h1>
          </div>
          <div className="text-center py-12">
            <p className="text-muted-foreground mb-4">No se pudieron cargar los datos</p>
            <Button onClick={loadDashboardData}>
              <RefreshCw className="h-4 w-4 mr-2" />
              Reintentar
            </Button>
          </div>
        </div>
      </MainLayout>
    );
  }

  return (
    <MainLayout>
      <div className="space-y-6">
        {/* Header */}
        <div className="flex items-center justify-between">
          <div>
            <h1 className="text-3xl font-bold">Dashboard</h1>
            <p className="text-muted-foreground">
              Resumen de tu portafolio de inversiones
            </p>
          </div>
          <Button variant="outline" onClick={loadDashboardData} disabled={isLoading}>
            <RefreshCw className={`h-4 w-4 mr-2 ${isLoading ? 'animate-spin' : ''}`} />
            Actualizar
          </Button>
        </div>

        {/* Summary Cards */}
        <SummaryCards summary={dashboardData.portfolioSummary} />

        {/* Charts and Lists Grid */}
        <div className="grid gap-6 lg:grid-cols-2">
          <AssetAllocationChart allocations={dashboardData.assetAllocations} />
          <RecentTransactions transactions={dashboardData.recentTransactions} />
        </div>

        {/* Top Assets */}
        <div className="grid gap-6 lg:grid-cols-2">
          <TopAssets assets={dashboardData.topAssets} />
          {/* Espacio para futuros componentes como gráficos de rendimiento temporal */}
          <div className="hidden lg:block" />
        </div>
      </div>
    </MainLayout>
  );
}
