<template>
  <div class="container mx-auto h-screen">
    <div class="grid grid-cols-1 gap-4">
      <div class="h-16"></div>
    </div>
    <div class="grid grid-cols-1 gap-4">
      <h1 class="text-5xl text-center">GitLab Wrap 🦊</h1>
      <h2 class="text-2xl text-center">Get your 2021 GitLab Stats</h2>
    </div>
    <div class="grid grid-cols-1 gap-4">
      <div class="h-16"></div>
      <GitLabWrapResult v-if="username" :username="username"></GitLabWrapResult>
      <GitLabWrapSearch v-else/>
      <div
          id="tweet"
          class="flex mt-3 md:mt-6 justify-center items-center h-2/3 flex-col"
      ></div>
      <router-link
          v-if="username"
          class="text-3xl flex mt-3 md:mt-6 justify-center items-center h-2/3 flex-col"
          :to="{ name: 'WrapIt' }"
      >I want a Wrap too!
      </router-link>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    username: {required: false},
  },
};
</script>

<script setup>
import {onUpdated, onMounted} from "vue";
import GitLabWrapSearch from "../components/GitLabWrapSearch.vue";
import GitLabWrapResult from "../components/GitLabWrapResult.vue";
import {useRouter} from "vue-router";

function createTwitterButton() {
  let router = useRouter();
  let text =
      "Want to know the impact you made to open-source? 👀\nClick the link below to find out!\n";
  let url = "https://gitlabwrap.fly.dev/";
  if (router.currentRoute.value.params.username != null) {
    text =
        "🚨 Do you know the impact you made to open-source? I do!\nClick the link below to find out!";
    url =
        "https://gitlabwrap.fly.dev/" + router.currentRoute.value.params.username;
  }

  let script = document.createElement("script");
  script.setAttribute("async", "");
  script.setAttribute("src", "https://platform.twitter.com/widgets.js");
  script.setAttribute("charset", "utf-8");
  let parent = document.getElementById("tweet");
  let wrapper = document.createElement("a");
  wrapper.innerHTML = "Tweet";
  wrapper.appendChild(script);
  wrapper.setAttribute("href", "https://twitter.com/share?ref_src=twsrc%5Etfw");
  wrapper.setAttribute("class", "twitter-share-button");
  wrapper.setAttribute("data-text", text);
  wrapper.setAttribute("data-size", "large");
  wrapper.setAttribute("data-url", url);
  wrapper.setAttribute("data-hashtags", "GitLabWrap");
  wrapper.setAttribute("data-show-count", "false");
  parent.appendChild(wrapper);
}

onUpdated(() => {
  let parent = document.getElementById("tweet");
  parent.removeChild(parent.childNodes[0]);
  createTwitterButton()
})

onMounted(() => {
  createTwitterButton()
});
</script>
