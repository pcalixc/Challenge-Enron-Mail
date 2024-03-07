export const fetchEmails = async (pageNumber : number, amountEmailsByPage: number, searchTerm: string) => {
    try {
      const response = await fetch(
        `http://${import.meta.env.VITE_API_URL}/emails?page=${pageNumber}&max=${amountEmailsByPage}&term=${searchTerm}`
      );
  
      const data = await response.json();
      return { response, data };
    } catch (error) {
      console.error('Error fetching data:', error);
      throw error;
    }
  };
