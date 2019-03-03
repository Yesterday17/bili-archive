<template>
  <div id="qr-box">
    <transition name="fade" mode="out-in">
      <span v-if="this.qrCode === ''">加载登录二维码中……</span>
      <img v-else :src="this.qrCode" alt="qrCode">
    </transition>
  </div>
</template>

<script>
export default {
  name: "Login",
  data() {
    return {
      qrCode: ""
    };
  },
  beforeRouteEnter(to, from, next) {
    localStorage.setItem("display", "1");
    next();
  },
  mounted() {
    localStorage.setItem("step", "1");
  },
  created() {
    fetch("//localhost:8080/login-qr")
      .then(data => data.json())
      .then(json => (this.qrCode = json.image))
      .then(() => {
        const status = setInterval(() => {
          fetch("//localhost:8080/login-status")
            .then(data => data.json())
            .then(json => {
              if (json.ok) {
                clearInterval(status);
                this.$router.push("step-02");
              }
            });
        }, 3000);
      });
  }
};
</script>

<style lang="scss" scoped>
#qr-box {
  width: 256px;
  height: 256px;

  display: flex;
  display: -webkit-flex;
  align-items: center;
  justify-content: center;
}
</style>
