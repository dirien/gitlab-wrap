<script setup>
import {ref} from "vue";
import router from "../router";

const userName = ref("");

function submitUserName() {
  router.push({name: "WrapItResult", params: {username: userName.value}});
}
</script>

<template>
  <div class="container flex justify-center items-center">
    <div class="relative">
      <div class="absolute top-4 left-3">
        <font-awesome-icon
            :icon="['fas', 'search']"
            class="text-gray-400 z-20 hover:text-gray-500"
        />
      </div>
      <input
          v-model="userName"
          type="text"
          class="h-14 w-96 pl-10 pr-20 rounded-lg z-0 text-black focus:shadow focus:outline-none"
          placeholder="GitLab username"
      />
      <div class="absolute top-0 right-0">
        <button
            class="h-14 w-20 text-white rounded-r-lg bg-yellow-600 hover:bg-yellow-500"
            @click="submitUserName"
        >
          Wrap It
        </button>
      </div>
    </div>
  </div>
</template>
