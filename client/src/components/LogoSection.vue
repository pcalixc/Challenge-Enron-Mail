<template>
  <div class=" flex justify-center md:items-center lg:w-auto w-auto ">
    <div class="flex items-center w-12 h-13 justify-center">
      <img
        src="../../public/envelope.png"
      />
    </div>
    <div class="flex flex-col ml-3 hidden_ ">
      <a href="/" class="mb-0 flex flex-col">
        <h2
          class="text-lg dark:text-slate-100 font-semibold tracking-wider uppercase text-black focus:outline-none focus:shadow-outline"
          >Mail Searcher</h2>
        <h4
        class="text-slate-600 dark:text-slate-300 text-[0.70rem] font-xs"
          >Find Emails Effortlessly!</h4>
      </a>
    </div>
  </div>
</template>

<style scoped>
@media (max-width: 1024px) {
  .hidden_ {
    display: none;
  }
}

</style>

