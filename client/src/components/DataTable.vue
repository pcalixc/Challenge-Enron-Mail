<!-- eslint-disable vue/multi-word-component-names -->
<script setup lang="ts">
import { ref, onMounted } from 'vue'

export interface dataI {
  ID: number
  Age: number
  Workclass: string
  Fnlwgt: number
  Education: string
  EducationNum: number
  MaritalStatus: string
  Occupation: string
  Relationship: string
  Race: string
  Gender: string
  CapitalGain: number
  CapitalLoss: number
  HoursPerWeek: number
  NativeCountry: string
  Income: string
}

const filtersData = {
  age: {
    label: "Age",
    values: ["21-40", "41-60", "61"],
  },
  workclass:{
    label: "Workclass"
  },
  education: {
    label: "Education",
    values: ["Bachelors", "Some college", "11th", "12th", "HS grad", "Assoc acdm", "Masters", "Preeschool", "Doctorate"],
  },
  marital: {
    label: "Marital Status",
    values: ["Married Civ", "Spouse", "Divorced", "Never married", "Widowed", "Married spouse absent"],
  },
  occupation: {
    label: "Occupation",
    values: ["Tech support", "Craft repair", "Sales", "Cleaners", "Other service"],
  },
  race:{
    label: "race"
  },
  gender:{
    label: "Gender"
  },
  country:{
    label: "Native country"
  },
  income: {
    label: "Income",
    values: ["More than 50K", "Less than 50k"],
  },
};

function handleChangeVisibility() {
  var filterSelect = document.getElementById("filter");
  selectedFilter.value = filterSelect.value
}

var totalResults = ref(0)
var page = ref(1)
var testData = ref('')

var totalPages = ref(0)
var baseURL = ref(`http://${import.meta.env.VITE_API_URL}/data/`)
const dataByPage = ref<dataI[]>([])
var selectedFilterOption = ref('')
var selectedFilter= ref('')


//localhost:3333/search?term=crazy&from=1&max=1
const getData1 = async () => {

  try {
    const response = await fetch(`http://localhost:3333/search?term=crazy&from=1&max=1`)
    const data = await response.json()
    testData.value = data.hits.hits[0]._source.content

  } catch (error) {
    console.error('Error fetching data:', error)
  }
 console.log("",testData.value)
}

onMounted(() => {
  //getData(1)
  getData1()

})
</script>

<template>
  <!-- filter  -->
  <div class="flex flex-row justify-center items-center relative mt-4">
    <div class="my-2 flex sm:flex-row flex-col">
      <div class="flex flex-row mb-1 sm:mb-0">
        <div class="relative mx-2">
          <select
            id="filter"
            class="appearance-none h-full rounded border block w-full bg-white border-gray-400 text-gray-700 py-2 px-4 pr-8 leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
            @change="handleChangeVisibility()"
            >
            <option selected>Filter by</option>
            <!-- <option value="age">Age</option> -->
            <option value="marital_status">Marital Status</option>
            <option value="occupation">Occupation</option>
            <option value="education">Education</option>
            <option value="income">Income</option>
          </select>
          <div
            class="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700"
          >
            <svg
              class="fill-current h-4 w-4"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 20 20"
            >
              <path
                d="M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z"
              />
            </svg>
          </div>
        </div>
        <!-- <div class="relative"
        v-if="selectedFilter=='age'"
        >
          <select
            id="age"
            class="appearance-none h-full rounded border block w-full bg-white border-gray-400 text-gray-700 py-2 px-4 pr-8 leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
            @change="handleChange('age')"
            >
            <option selected>Options</option>
            <option>21-40</option>
            <option>41-60</option>
            <option>61</option>
          </select>
          <div
            class="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700"
          >
            <svg
              class="fill-current h-4 w-4"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 20 20"
            >
              <path
                d="M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z"
              />
            </svg>
          </div>
        </div> -->

        <div class="relative"
        v-if="selectedFilter=='education'"
        >
          <select
            id="education"
            class="appearance-none h-full rounded border block w-full bg-white border-gray-400 text-gray-700 py-2 px-4 pr-8 leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
            @change="handleChange('education')"
          >
            <option value="">Options</option>
            <option value="Bachelors">Bachelors</option>
            <option value="Some college">Some college</option>
            <option value="11th">11th</option>
            <option value="HS grad">HS-grad</option>
            <option value="Assoc acdm">Assoc acdm</option>
            <option value="Masters">Masters</option>
            <option value="Preschool">Preschool</option>
            <option value="Doctorate">Doctorate</option>
          </select>
          <div
            class="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700"
          >
            <svg
              class="fill-current h-4 w-4"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 20 20"
            >
              <path
                d="M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z"
              />
            </svg>
          </div>
        </div>

        <div class="relative"
        v-if="selectedFilter=='marital_status'"
        >
          <select id="marital_status"
          class="appearance-none h-full rounded border block w-full bg-white border-gray-400 text-gray-700 py-2 px-4 pr-8 leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
          @change="handleChange('marital_status')"
            >
            <option>Options</option>
            <option>Married civ spouse</option>
            <option>Divorced</option>
            <option>Never married</option>
            <option>Separated</option>
            <option>Widowed</option>
            <option>Married spouse absent</option>
            <option>Married AF spouse</option>

          </select>
          <div
            class="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700"
          >
            <svg
              class="fill-current h-4 w-4"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 20 20"
            >
              <path
                d="M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z"
              />
            </svg>
          </div>
        </div>

        <div class="relative"
        v-if="selectedFilter=='occupation'"
        >
          <select id="occupation"
            class="appearance-none h-full rounded border block w-full bg-white border-gray-400 text-gray-700 py-2 px-4 pr-8 leading-tight focus:outline-none focus:bg-white focus:border-gray-500"

            @change="handleChange('occupation')"
            >
            <option>Options</option>
            <option>Tech support</option>
            <option>Craft repair</option>
            <option>Sales</option>
            <option>Handlers cleaners</option>
            <option>Prof specialty</option>
            <option>Exec managerial</option>
            <option>Machine op inspct</option>
            <option>Adm clerical</option>
            <option>Farming fishing</option>
            <option>Transport moving</option>
            <option>Other service</option>
          </select>
          <div
            class="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700"
          >
            <svg
              class="fill-current h-4 w-4"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 20 20"
            >
              <path
                d="M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z"
              />
            </svg>
          </div>
        </div>

        <div class="relative"
        v-if="selectedFilter=='income'">
          <select id="income"
            class="appearance-none h-full rounded border block w-full bg-white border-gray-400 text-gray-700 py-2 px-4 pr-8 leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
            @change="handleChange('income')"
            >
            <option>Options</option>
            <option value="1">More than 50K</option>
            <option value="2">Less than 50k</option>

          </select>
          <div
            class="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700"
          >
            <svg
              class="fill-current h-4 w-4"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 20 20"
            >
              <path
                d="M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z"
              />
            </svg>
          </div>
        </div>
      </div>
    </div>
  </div>

  <div class="flex w-80 rounded-lg text-sm text-white mx-24" role="alert">
        <svg class="w-5 h-5 inline mr-3" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z" clip-rule="evenodd"></path></svg>
        <div>
            <span class="font-medium">{{ totalResults }}</span> registers founded.
        </div>
    </div>


  <!-- component -->
  <section class="container px-4 mx-auto mt-5 mb-0">
    <div class="flex flex-col h-[65vh]">
      <div class="-mx-4 -my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
        <div class="inline-block min-w-full py-2 align-middle md:px-6 lg:px-8">
          <div class="overflow-hidden border border-gray-200 dark:border-gray-700 md:rounded-lg">
            <table class="min-w-full divide-y divide-gray-200 dark:divide-gray-700">
              <thead class="bg-gray-50 dark:bg-gray-800">
                <tr>
                  <th
                    scope="col"
                    class="py-3.5 px-4 text-sm font-normal text-left rtl:text-right text-gray-500 dark:text-gray-400"
                  >
                    <div class="flex items-center gap-x-3">
                      <button class="flex items-center gap-x-2 ml-3">
                        <span>ID</span>
                      </button>
                    </div>
                  </th>
                  <th 
                  v-for="(filter, index) in filtersData" :key="index" 
                    scope="col"
                    class="px-4 py-3.5 text-sm font-normal text-left rtl:text-right text-gray-500 dark:text-gray-400"
                  >
                    {{filter.label}}
                  </th>

                </tr>
              </thead>
              <tbody
                class="bg-white divide-y h-1.5 divide-gray-200 dark:divide-gray-700 dark:bg-gray-900"
              >
                <tr v-for="(data, index) in dataByPage" :key="index" class="h-[40px]">
                  <td class="px-4 py-4 text-sm font-medium text-gray-700 whitespace-nowrap">
                    <div
                      class="inline-flex items-center px-3 py-1 rounded-full gap-x-2 text-blue-300"
                    >
                      <h2 class="text-sm font-bold">{{ data.ID }}</h2>
                    </div>
                  </td>
                  <td class="px-4 py-4 text-sm text-gray-500 dark:text-gray-300 whitespace-nowrap">
                    {{ data.Age }}
                  </td>
                  <td class="px-4 py-4 text-sm font-medium text-gray-700 whitespace-nowrap">
                    <div
                      class="inline-flex items-center px-3 py-1 text-gray-200 rounded-full gap-x-2 bg-gray-100/60 dark:bg-gray-800"
                    >
                      <h2 class="text-sm font-normal">{{ data.Workclass }}</h2>
                    </div>
                  </td>
                  <td class="px-4 py-4 text-sm text-gray-500 dark:text-gray-300 whitespace-nowrap">
                    {{ data.Education }}
                  </td>
                  <td class="px-4 py-4 text-sm text-gray-500 dark:text-gray-300 whitespace-nowrap">
                    {{ data.MaritalStatus }}
                  </td>
                  <td class="px-4 py-4 text-sm text-gray-500 dark:text-gray-300 whitespace-nowrap">
                    {{ data.Occupation }}
                  </td>
                  <td class="px-4 py-4 text-sm text-gray-500 dark:text-gray-300 whitespace-nowrap">
                    {{ data.Race }}
                  </td>
                  <td class="px-4 py-4 text-sm font-medium text-gray-700 whitespace-nowrap">
                    <div
                      :class="[
                        'inline-flex items-center px-3 py-1 rounded-full gap-x-2  dark:bg-gray-800',
                        data.Gender == 'Female'
                          ? ' text-purple-500 bg-purple-100/60'
                          : 'text-blue-500 bg-blue-100/60'
                      ]"
                    >
                      <h2 class="text-sm font-normal">{{ data.Gender }}</h2>
                    </div>
                  </td>
                  <td class="px-4 py-4 text-sm text-gray-500 dark:text-gray-300 whitespace-nowrap">
                    {{ data.NativeCountry }}
                  </td>
                  <td class="px-4 py-4 text-sm whitespace-nowrap">
                    <div class="flex items-center gap-x-6">
                      <div
                        :class="[
                          'inline-flex items-center px-3 py-1 rounded-full gap-x-2 bg-gray-800 ',
                          data.Income == '<=50K' ? ' text-red-500' : 'text-green-500 '
                        ]"
                      >
                        <h2 class="text-sm font-semibold">{{ data.Income }}</h2>
                      </div>
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>

    <div class="flex items-center justify-center gap-2 mt-0">
      <button
        v-if="page === 1"
        class="flex items-center w-32 pointer-events-none px-5 py-2 text-sm capitalize transition-colors duration-200 border rounded-md gap-x-2 bg-gray-800 dark:text-gray-300 dark:border-gray-700"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          stroke-width="1.5"
          stroke="currentColor"
          class="w-5 h-5 rtl:-scale-x-100"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            d="M6.75 15.75L3 12m0 0l3.75-3.75M3 12h18"
          />
        </svg>

        <span> previous </span>
      </button>

      <button
        v-else
        class="flex items-start w-32 px-5 py-2 text-sm text-gray-700 capitalize transition-colors duration-200 bg-white border rounded-md gap-x-2 dark:bg-gray-900 dark:text-gray-200 dark:border-gray-700 hover:border-blue-700"
        @click="getData(page - 1, selectedFilter, selectedFilterOption)"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          stroke-width="1.5"
          stroke="currentColor"
          class="w-5 h-5 rtl:-scale-x-100"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            d="M6.75 15.75L3 12m0 0l3.75-3.75M3 12h18"
          />
        </svg>

        <span> previous </span>
      </button>

      <button
        v-if="page === totalPages"
        class="flex items-end justify-end w-32 px-5 py-2 text-sm text-gray-700 capitalize transition-colors duration-200 bg-white border rounded-md gap-x-2 pointer-events-none dark:bg-gray-700 dark:text-gray-200 dark:border-gray-700"
      >
        <span> Next </span>

        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          stroke-width="1.5"
          stroke="currentColor"
          class="w-5 h-5 rtl:-scale-x-100"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            d="M17.25 8.25L21 12m0 0l-3.75 3.75M21 12H3"
          />
        </svg>
      </button>

      <button
        v-else
        class="flex items-end justify-end px-5 py-2 w-32 text-sm text-gray-700 capitalize transition-colors duration-200 bg-white border rounded-md gap-x-2 dark:bg-gray-900 dark:text-gray-200 dark:border-gray-700 hover:border-blue-700"
        @click="getData(page + 1, selectedFilter, selectedFilterOption)"
      >
        <span> Next </span>

        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          stroke-width="1.5"
          stroke="currentColor"
          class="w-5 h-5 rtl:-scale-x-100"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            d="M17.25 8.25L21 12m0 0l-3.75 3.75M21 12H3"
          />
        </svg>
      </button>
    </div>
  </section>
</template>
