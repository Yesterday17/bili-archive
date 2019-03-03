<template>
  <div>Welcome!</div>
</template>

<script>
export default {
  name: "Welcome",
  data() {
    return {};
  },
  mounted() {
    localStorage.setItem("step", "0");
  }
};
</script>

<style lang="scss" scoped>
</style>
