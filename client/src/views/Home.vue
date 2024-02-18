<!-- eslint-disable vue/multi-word-component-names -->
<script setup lang="ts">
import { ref, onBeforeMount } from 'vue'
import Header from '@/components/Header.vue'
import  type {Email, Hit} from '@/types/index'
import EmailDetails from '@/components/EmailDetails.vue'
import TotalResults from '@/components/TotalResults.vue'
import Table from '@/components/Table.vue'

const open = ref(false)
const selectedEmail = ref<Email>();
const emails = ref<Hit[]>([]);
const selectedEmailIndex = ref(0)
const currentSearchTerm = ref('')
const totalResults = ref(0)
const currentPage = ref(1)
var totalPages = ref(0)
const amountEmailsByPage = ref(7)
const isLoading = ref(false)
const conectionError= ref(false)


const getData = async (pageNumber: number, searchTerm?: string) => {
  searchTerm = typeof searchTerm === 'undefined' ? '' : searchTerm
  let pn = pageNumber.toString()
  currentSearchTerm.value = searchTerm
  isLoading.value = true
  try {
    const response = await fetch(
      `http://${import.meta.env.VITE_API_URL}/emails?page=${pn}&max=${amountEmailsByPage.value}&term=${searchTerm}`
    )

    const data = await response.json()

    emails.value = data.hits.hits
    totalResults.value = data.hits.Total.value
  } catch (error) {
    console.error('Error fetching data:', error, Response)
    conectionError.value=true
  }
  isLoading.value = false
  currentPage.value = pageNumber
  totalPages.value = Math.ceil(totalResults.value / amountEmailsByPage.value)
}

const asigneSelectedContent = (index: number) => {
  selectedEmail.value = emails.value[index]._source
  selectedEmailIndex.value = index
  open.value=true
  console.log(index, selectedEmail.value)
}

onBeforeMount(() => {
  getData(1)
})
</script>

<template>

<Header :getData="getData" />

  <Table
    :getData="getData" 
    :emails ="emails"
    :asigneSelectedContent = "asigneSelectedContent"
    :currentSearchTerm="currentSearchTerm"
    :isLoading="isLoading"
    :totalResults="totalResults"
    :conectionError="conectionError"
    :totalPages="totalPages"
    :currentPage="currentPage"
    />

  <TotalResults v-if="!conectionError"
    :totalResults="totalResults"
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
