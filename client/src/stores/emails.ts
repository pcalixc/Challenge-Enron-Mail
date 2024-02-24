// import { defineStore } from 'pinia';


// export const useEmailsStore = defineStore('emails', () => {

//     const open = ref<boolean>(false)
// const selectedEmail = ref<Email | undefined>()
// const emails = ref<Hit[]>([])
// const selectedEmailIndex = ref<number>(0)
// const currentSearchTerm = ref<string>('')
// const totalResults = ref<number>(0)
// const currentPage = ref<number>(1)
// const amountEmailsByPage = ref<number>(7)
// const isLoading = ref<boolean>(false)
// const conectionError = ref<boolean>(false)
// const totalPages = computed(() => Math.ceil(totalResults.value / amountEmailsByPage.value))

// const getData = async (pageNumber: number, searchTerm?: string) => {
//     searchTerm = typeof searchTerm === 'undefined' ? '' : searchTerm
  
//     const strPageNumber = pageNumber.toString()
//     currentSearchTerm.value = searchTerm
//     isLoading.value = true
  
//     try {
//       const response = await fetch(
//         `http://${import.meta.env.VITE_API_URL}/emails?page=${strPageNumber}&max=${amountEmailsByPage.value}&term=${searchTerm}`
//       )
  
//       const data = await response.json()
//       emails.value = data.hits.hits
//       totalResults.value = data.hits.Total.value
//     } catch (error) {
//       console.error('Error fetching data:', error)
//       conectionError.value=true
//     }
//     isLoading.value = false
//     currentPage.value = pageNumber
//   }
  
//   const asigneSelectedContent = (index: number) => {
//     selectedEmail.value = emails.value[index]._source
//     selectedEmailIndex.value = index
//     open.value=true
//   }

// })