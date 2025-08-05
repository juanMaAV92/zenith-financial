// Zenith Financial - TypeScript Types

export interface User {
  id: number;
  code: string;
  username: string;
  email: string;
  currency: string;
  created_at: string;
  updated_at: string;
}

export interface Category {
  id: number;
  name: string;
  created_at: string;
  updated_at: string;
}

export interface Asset {
  id: number;
  code: string;
  user_id: number;
  name: string;
  symbol: string;
  ticker?: string;
  currency: string;
  category_id: number;
  total_units: number;
  current_value?: number;
  invested_total: number;
  auto_pricing_enabled: boolean;
  price_source?: string;
  created_at: string;
  updated_at: string;
  category?: Category;
}

export interface Transaction {
  id: number;
  code: string;
  asset_id: number;
  type: TransactionType;
  units: number;
  total: number;
  fee_total: number;
  currency: string;
  note?: string;
  created_at: string;
  asset?: Asset;
}

export type TransactionType = 'BUY' | 'SELL' | 'DEPOSIT' | 'WITHDRAW';

export type CategoryType = 
  | 'CASH'
  | 'SAVINGS_ACCOUNT'
  | 'FIXED_INCOME'
  | 'STOCK'
  | 'ETF'
  | 'CRYPTO'
  | 'MUTUAL_FUND'
  | 'COMMODITY'
  | 'CURRENCIES'
  | 'OTHER';

// Dashboard types
export interface PortfolioSummary {
  totalValue: number;
  totalInvested: number;
  totalGainLoss: number;
  gainLossPercentage: number;
  currency: string;
}

export interface AssetAllocation {
  category: string;
  value: number;
  percentage: number;
  count: number;
}

export interface DashboardData {
  portfolioSummary: PortfolioSummary;
  assetAllocations: AssetAllocation[];
  recentTransactions: Transaction[];
  topAssets: Asset[];
}

// API Response types
export interface ApiResponse<T> {
  data: T;
  message?: string;
  success: boolean;
}

export interface PaginatedResponse<T> {
  data: T[];
  pagination: {
    page: number;
    per_page: number;
    total: number;
    total_pages: number;
  };
}

// Form types
export interface CreateAssetForm {
  name: string;
  symbol: string;
  ticker?: string;
  currency: string;
  category_id: number;
  auto_pricing_enabled: boolean;
  price_source?: string;
}

export interface CreateTransactionForm {
  asset_id: number;
  type: TransactionType;
  units: number;
  total: number;
  fee_total?: number;
  currency: string;
  note?: string;
}

// Chart data types
export interface ChartDataPoint {
  label: string;
  value: number;
  color?: string;
}

export interface TimeSeriesDataPoint {
  date: string;
  value: number;
}
