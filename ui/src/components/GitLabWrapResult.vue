<template>
  <div class="container flex justify-center items-center">
    <img :src="card(username)"/>
  </div>
</template>

<script>
export default {
  name: "GitLabWrapResult",
  props:
      {
        username: String,
      },
  methods: {
    card: function (username) {
      return "https://gitlabwrap.fly.dev/card/" + username;
    },
  }
}
</script>
