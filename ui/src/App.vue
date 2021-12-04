<template>
  <metainfo>
    <template #title="{ content }">{{
      content ? `${content} | SITE_NAME` : `SITE_NAME`
    }}</template>
  </metainfo>
  <router-view />
</template>

<style>
@import url("https://fonts.googleapis.com/css?family=Nunito:400,700&display=swap");
</style>
