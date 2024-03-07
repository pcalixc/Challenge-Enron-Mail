import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import type { IEmail, Hit } from '@/types/index'
import dictionary from '@/assets/Dictionary.json'
import { SearchWithFuse } from '@/utils/emails.utilities'
import { fetchEmails } from '@/utils/api'
const router = useRouter()

export const useEmailsStore = defineStore('emails', () => {
  const modalOpen = ref<boolean>(false) 
  const selectedEmail = ref<IEmail | undefined>() 
  const emails = ref<Hit[]>([]) 
  const selectedEmailIndex = ref<number>(0)
  const currentSearchTerm = ref<string>('') 
  const totalResults = ref<number>(0)
  const currentPage = ref<number>(1) 
  const isLoading = ref<boolean>(false) 
  const conectionError = ref<boolean>(false) 
  const wordsInDictionary: string[] = Object.keys(dictionary)
  const searchSuggestion = ref()
  const amountEmailsByPage = 8
  const totalPages = computed(() => Math.ceil(totalResults.value / amountEmailsByPage))

  /**
 * getData fetches email data from the server based on the specified page number and search term.
 * pageNumber - The page number of the data to fetch.
 * [searchTerm] - The optional search term to filter the data.
 */
  const getData = async (pageNumber: number, searchTerm?: string) => {
    searchTerm = typeof searchTerm === 'undefined' ? '' : searchTerm
    const amountEmailsByPage = 8
    currentSearchTerm.value = searchTerm
    isLoading.value = true
    console.log('api url', import.meta.env.VITE_API_URL)

    try {
      const { response, data } = await fetchEmails(pageNumber, amountEmailsByPage, searchTerm)
      emails.value = data.hits.hits
      totalResults.value = data.hits.Total.value
      currentSearchTerm.value=searchTerm

      // If there are no results, generate search suggestions.
      if (totalResults.value === 0) {
        const suggestions = SearchWithFuse(wordsInDictionary, searchTerm)
        searchSuggestion.value = suggestions
      } else if (response.status === 404) {
        router.push('/404')
      }
    } catch (error) {
      conectionError.value = true
    }

    isLoading.value = false
    currentPage.value = pageNumber
  }

  const asigneSelectedContent = (index: number) => {
    selectedEmail.value = emails.value[index]._source
    selectedEmailIndex.value = index
    modalOpen.value = true
  }

  const setSearchWord = (term: string) => {
    currentSearchTerm.value = term
}

  return {
    emails,
    modalOpen,
    selectedEmail,
    suggestion: searchSuggestion,
    totalPages,
    amountEmailsByPage,
    wordsInDictionary,
    currentPage,
    conectionError,
    isLoading,
    selectedEmailIndex,
    currentSearchTerm,
    totalResults,
    asigneSelectedContent,
    getData,
    setSearchWord
  }
})
