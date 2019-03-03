<template>
  <div>Download</div>
</template>

<script>
export default {
  name: "Download",
  data() {
    return {};
  },
  beforeRouteEnter(to, from, next) {
    localStorage.setItem("display", "3");
    next();
  },
  mounted() {
    localStorage.setItem("step", "3");
  }
};
</script>

<style lang="scss" scoped>
</style>
