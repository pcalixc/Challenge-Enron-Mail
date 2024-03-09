import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { IEmail, IHit } from '@/types/index'
import dictionary from '@/assets/Dictionary.json'
import { SearchWithFuse } from '@/utils/emails.utilities'
import { fetchEmails } from '@/utils/api'
import type { IServerErrorResponse } from '@/types/index'

export const useEmailsStore = defineStore('emails', () => {
  const modalOpen = ref<boolean>(false)
  const selectedEmail = ref<IEmail | undefined>()
  const emails = ref<IHit[]>([])
  const selectedEmailIndex = ref<number>(0)
  const currentSearchTerm = ref<string>('')
  const totalResults = ref<number>(0)
  const currentPage = ref<number>(1)
  const isLoading = ref<boolean>(false)
  const wordsInDictionary: string[] = Object.keys(dictionary)
  const searchSuggestion = ref()
  const amountEmailsByPage = 8
  const totalPages = computed(() => Math.ceil(totalResults.value / amountEmailsByPage))
  const response = ref<any | undefined>()

  const ServerErrorResponse = ref<IServerErrorResponse>({
    errorStatus: false,
    errorCode: 0,
    errorMessage: ''
  })

  const getData = async (pageNumber: number, searchTerm = '') => {
    console.log("pp", currentPage.value)
    currentSearchTerm.value = searchTerm
    isLoading.value = true

    response.value = await fetchEmails(pageNumber, amountEmailsByPage, searchTerm)

    emails.value = response.value.data.hits.hits
    totalResults.value = response.value.data.hits.Total.value

    if (totalResults.value === 0) {
      searchSuggestion.value = SearchWithFuse(wordsInDictionary, searchTerm)
    }

    isLoading.value = false
    currentPage.value = pageNumber
  }

  const asigneSelectedContent = (index: number) => {
    selectedEmail.value = emails.value[index]._source
    selectedEmailIndex.value = index
    modalOpen.value = true
  }

  const restoreServerErrorResponse = () => {
    ServerErrorResponse.value = {
      errorStatus: false,
      errorCode: 0,
      errorMessage: ''
    }
  }

  return {
    emails,
    modalOpen,
    selectedEmail,
    searchSuggestion,
    totalPages,
    amountEmailsByPage,
    wordsInDictionary,
    currentPage,
    isLoading,
    selectedEmailIndex,
    currentSearchTerm,
    totalResults,
    asigneSelectedContent,
    getData,
    ServerErrorResponse,
    restoreServerErrorResponse
  }
})
