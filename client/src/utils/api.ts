import { useEmailsStore } from '@/stores/emails';

export const fetchEmails = async (
  pageNumber: number,
  amountEmailsByPage: number,
  searchTerm: string
) => {
  const emailsStore = useEmailsStore();

  try {
    emailsStore.restoreServerErrorResponse()

    const response = await fetch(
      `http://${import.meta.env.VITE_API_URL}/emails?page=${pageNumber}&max=${amountEmailsByPage}&term=${searchTerm}`
    );

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
    emailsStore.ServerErrorResponse = {
      errorStatus: true,
      errorCode: 0,
      errorMessage: 'CONNECTION_ERROR'
    }
    }
  }

