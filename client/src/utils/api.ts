import { useEmailsStore } from '@/stores/emails';

export const fetchEmails = async (
  pageNumber: number,
  amountEmailsByPage: number,
  searchTerm: string
) => {
  const emailsStore = useEmailsStore();

  try {
    emailsStore.restoreServerErrorResponse()
    const controller = new AbortController();
    const signal = controller.signal;

    const timeoutId = setTimeout(() => {
      controller.abort(); // Aborta la solicitud después de un tiempo de espera
    }, 9000); 

    const response = await fetch(
      `http://${import.meta.env.VITE_API_URL}/emails?page=${pageNumber}&max=${amountEmailsByPage}&term=${searchTerm}`
      , { signal } 
    );

    clearTimeout(timeoutId);

    if (!response.ok) {
      emailsStore.ServerErrorResponse = {
        errorStatus: true,
        errorCode: response.status,
        errorMessage: response.statusText
      }
      return new Error(`Server response error: ${response.statusText}`);
    }

    const data = await response.json();
    return { data };

  } catch (error) {
    if (error instanceof DOMException && error.name === 'AbortError') {
      emailsStore.ServerErrorResponse = {
        errorStatus: true,
        errorCode: 0,
        errorMessage: 'TIMEOUT_ERROR'
      };
    } else {
      emailsStore.ServerErrorResponse = {
        errorStatus: true,
        errorCode: 0,
        errorMessage: 'CONNECTIONN_ERROR'
      };
    }
    return error;
  }
  }

