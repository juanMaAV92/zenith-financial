import { create } from 'zustand';
import { devtools } from 'zustand/middleware';
import { User, Asset, Transaction, Category, DashboardData } from '@/types';

interface AppState {
  // User state
  user: User | null;
  isAuthenticated: boolean;
  
  // Data state
  assets: Asset[];
  transactions: Transaction[];
  categories: Category[];
  dashboardData: DashboardData | null;
  
  // UI state
  isLoading: boolean;
  error: string | null;
  
  // Actions
  setUser: (user: User | null) => void;
  setAssets: (assets: Asset[]) => void;
  addAsset: (asset: Asset) => void;
  updateAsset: (id: number, asset: Partial<Asset>) => void;
  removeAsset: (id: number) => void;
  setTransactions: (transactions: Transaction[]) => void;
  addTransaction: (transaction: Transaction) => void;
  updateTransaction: (id: number, transaction: Partial<Transaction>) => void;
  removeTransaction: (id: number) => void;
  setCategories: (categories: Category[]) => void;
  setDashboardData: (data: DashboardData) => void;
  setLoading: (loading: boolean) => void;
  setError: (error: string | null) => void;
  reset: () => void;
}

const useAppStore = create<AppState>()(
  devtools(
    (set, get) => ({
      // Initial state
      user: null,
      isAuthenticated: false,
      assets: [],
      transactions: [],
      categories: [],
      dashboardData: null,
      isLoading: false,
      error: null,

      // Actions
      setUser: (user) => set({ user, isAuthenticated: !!user }),
      
      setAssets: (assets) => set({ assets }),
      
      addAsset: (asset) => set((state) => ({
        assets: [...state.assets, asset]
      })),
      
      updateAsset: (id, updatedAsset) => set((state) => ({
        assets: state.assets.map(asset => 
          asset.id === id ? { ...asset, ...updatedAsset } : asset
        )
      })),
      
      removeAsset: (id) => set((state) => ({
        assets: state.assets.filter(asset => asset.id !== id)
      })),
      
      setTransactions: (transactions) => set({ transactions }),
      
      addTransaction: (transaction) => set((state) => ({
        transactions: [...state.transactions, transaction]
      })),
      
      updateTransaction: (id, updatedTransaction) => set((state) => ({
        transactions: state.transactions.map(transaction => 
          transaction.id === id ? { ...transaction, ...updatedTransaction } : transaction
        )
      })),
      
      removeTransaction: (id) => set((state) => ({
        transactions: state.transactions.filter(transaction => transaction.id !== id)
      })),
      
      setCategories: (categories) => set({ categories }),
      
      setDashboardData: (dashboardData) => set({ dashboardData }),
      
      setLoading: (isLoading) => set({ isLoading }),
      
      setError: (error) => set({ error }),
      
      reset: () => set({
        user: null,
        isAuthenticated: false,
        assets: [],
        transactions: [],
        categories: [],
        dashboardData: null,
        isLoading: false,
        error: null,
      })
    }),
    { name: 'zenith-financial-store' }
  )
);

export default useAppStore;
