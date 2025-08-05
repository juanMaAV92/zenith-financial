import apiClient from './client';
import {
  User,
  Asset,
  Transaction,
  Category,
  DashboardData,
  ApiResponse,
  PaginatedResponse,
  CreateAssetForm,
  CreateTransactionForm
} from '@/types';

// Health check
export const healthCheck = async (): Promise<ApiResponse<{ status: string }>> => {
  const response = await apiClient.get('/health');
  return response.data;
};

// User services
export const userService = {
  getCurrentUser: async (): Promise<ApiResponse<User>> => {
    const response = await apiClient.get('/api/users/me');
    return response.data;
  },
  
  updateUser: async (userData: Partial<User>): Promise<ApiResponse<User>> => {
    const response = await apiClient.put('/api/users/me', userData);
    return response.data;
  }
};

// Asset services
export const assetService = {
  getAssets: async (page = 1, perPage = 20): Promise<PaginatedResponse<Asset>> => {
    const response = await apiClient.get(`/api/assets?page=${page}&per_page=${perPage}`);
    return response.data;
  },
  
  getAsset: async (id: number): Promise<ApiResponse<Asset>> => {
    const response = await apiClient.get(`/api/assets/${id}`);
    return response.data;
  },
  
  createAsset: async (assetData: CreateAssetForm): Promise<ApiResponse<Asset>> => {
    const response = await apiClient.post('/api/assets', assetData);
    return response.data;
  },
  
  updateAsset: async (id: number, assetData: Partial<CreateAssetForm>): Promise<ApiResponse<Asset>> => {
    const response = await apiClient.put(`/api/assets/${id}`, assetData);
    return response.data;
  },
  
  deleteAsset: async (id: number): Promise<ApiResponse<void>> => {
    const response = await apiClient.delete(`/api/assets/${id}`);
    return response.data;
  }
};

// Transaction services
export const transactionService = {
  getTransactions: async (
    page = 1, 
    perPage = 20, 
    assetId?: number
  ): Promise<PaginatedResponse<Transaction>> => {
    let url = `/api/transactions?page=${page}&per_page=${perPage}`;
    if (assetId) {
      url += `&asset_id=${assetId}`;
    }
    const response = await apiClient.get(url);
    return response.data;
  },
  
  getTransaction: async (id: number): Promise<ApiResponse<Transaction>> => {
    const response = await apiClient.get(`/api/transactions/${id}`);
    return response.data;
  },
  
  createTransaction: async (transactionData: CreateTransactionForm): Promise<ApiResponse<Transaction>> => {
    const response = await apiClient.post('/api/transactions', transactionData);
    return response.data;
  },
  
  updateTransaction: async (
    id: number, 
    transactionData: Partial<CreateTransactionForm>
  ): Promise<ApiResponse<Transaction>> => {
    const response = await apiClient.put(`/api/transactions/${id}`, transactionData);
    return response.data;
  },
  
  deleteTransaction: async (id: number): Promise<ApiResponse<void>> => {
    const response = await apiClient.delete(`/api/transactions/${id}`);
    return response.data;
  }
};

// Category services
export const categoryService = {
  getCategories: async (): Promise<ApiResponse<Category[]>> => {
    const response = await apiClient.get('/api/categories');
    return response.data;
  }
};

// Dashboard services
export const dashboardService = {
  getDashboardData: async (): Promise<ApiResponse<DashboardData>> => {
    const response = await apiClient.get('/api/dashboard');
    return response.data;
  }
};
