import { ref } from 'vue';

export interface ToastMessage {
  message: string;
  type?: 'success' | 'error' | 'warning' | 'info';
  duration?: number;
}

export const toastMessage = ref<ToastMessage | null>(null);

export const showToast = (message: string, type: 'success' | 'error' | 'warning' | 'info' = 'info', duration = 3000) => {
  toastMessage.value = { message, type, duration };
  
  if (duration > 0) {
    setTimeout(() => {
      toastMessage.value = null;
    }, duration);
  }
};

export const showSuccess = (message: string, duration = 3000) => showToast(message, 'success', duration);
export const showError = (message: string, duration = 3000) => showToast(message, 'error', duration);
export const showWarning = (message: string, duration = 3000) => showToast(message, 'warning', duration);
export const showInfo = (message: string, duration = 3000) => showToast(message, 'info', duration);
