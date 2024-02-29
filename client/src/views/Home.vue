<!-- eslint-disable vue/multi-word-component-names -->
<script setup lang="ts">
import { ref,computed, onBeforeMount } from 'vue'
import { useRouter } from 'vue-router'
const router = useRouter()
import { SearchWithFuse } from '@/utils/emails.utilities';
import Header from '@/components/Header.vue'
import  type {Email, Hit} from '@/types/index'
import EmailDetails from '@/components/EmailDetails.vue'
import TotalResults from '@/components/TotalResults.vue'
import EmailsTable from '@/components/EmailsTable.vue'
import dictionary from '@/assets/words_dictionary.json'; 

let open = ref<boolean>(false)
let selectedEmail = ref<Email | undefined>()
let emails = ref<Hit[]>([])
let selectedEmailIndex = ref<number>(0)
let currentSearchTerm = ref<string>('')
let totalResults = ref<number>(0)
let currentPage = ref<number>(1)
let isLoading = ref<boolean>(false)
let conectionError = ref<boolean>(false)
let wordsInDictionary: string[] = Object.keys(dictionary);
let suggestion= ref()
let amountEmailsByPage = ref<number>(8)
let totalPages = computed(() => Math.ceil(totalResults.value / amountEmailsByPage.value))

/**
 * getData fetches email data from the server based on the specified page number and search term.
 * @param {number} pageNumber - The page number of the data to fetch.
 * @param {string} [searchTerm] - The optional search term to filter the data.
 */
const getData = async (pageNumber: number, searchTerm?: string) => {
    // Set the search term to an empty string if it's undefined.
  searchTerm = typeof searchTerm === 'undefined' ? '' : searchTerm
  let amountEmailsByPage = 8


  let strPageNumber = pageNumber.toString()
  currentSearchTerm.value = searchTerm
  isLoading.value = true
  console.log("testing334d")

  try {
    const response = await fetch(
      `http://${import.meta.env.VITE_API_URL}/emails?page=${strPageNumber}&max=${amountEmailsByPage}&term=${searchTerm}`
    )

    const data = await response.json()
    emails.value = data.hits.hits
    totalResults.value = data.hits.Total.value

    // If there are no results, generate search suggestions.
    if (totalResults.value===0){
      const suggestions = SearchWithFuse(wordsInDictionary, searchTerm);
      suggestion.value = suggestions;
    }else if (response.status === 404) {
    router.push('/404')
  }

  } catch (error) {
    console.error('Error fetching data:', error)
    conectionError.value=true
  }

  isLoading.value = false
  currentPage.value = pageNumber
}

const asigneSelectedContent = (index: number) => {
  selectedEmail.value = emails.value[index]._source
  selectedEmailIndex.value = index
  open.value=true
}

onBeforeMount(() => {
  getData(1)
})
</script>

<template>

  <Header 
  :getData="getData"
 />

  <EmailsTable
    :getData="getData" 
    :emails ="emails"
    :asigneSelectedContent = "asigneSelectedContent"
    :currentSearchTerm="currentSearchTerm"
    :isLoading="isLoading"
    :totalResults="totalResults"
    :conectionError="conectionError"
    :totalPages="totalPages"
    :currentPage="currentPage"
    :suggestions="suggestion"
    />

    <TotalResults v-if="!conectionError && totalResults !=0"
    :totalResults="totalResults"
    :isLoading="isLoading"
    :currentSearchTerm="currentSearchTerm"
    :currentPage="currentPage"
    :sugerencias="suggestion"
    :getData="getData"
    />


  <EmailDetails v-if="open"
    :selectedEmail="selectedEmail"
    :asigneSelectedContent="asigneSelectedContent"
    :currentSearchTerm="currentSearchTerm"
    :selectedEmailIndex="selectedEmailIndex"
    @close="open=false"
    />

</template>

<style>

::-webkit-scrollbar {
  width: 10px;
  height: 10px;
}


::-webkit-scrollbar-thumb {
  background-color: #9292FF;
  border-radius: 5px;
}

::-webkit-scrollbar-thumb:hover {
  background-color: #555;
}

::-webkit-scrollbar-track {
  background-color: #f1f1f1;
}
</style>
