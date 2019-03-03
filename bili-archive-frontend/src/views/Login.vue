<template>
  <div id="qr-box">
    <transition name="fade" mode="out-in">
      <div v-if="this.timeout">二维码超时</div>
      <div v-else-if="this.qrCode === ''">加载登录二维码中……</div>
      <img v-else :src="this.qrCode" alt="qrCode">
    </transition>
  </div>
</template>

<script>
export default {
  name: "Login",
  data() {
    return {
      timeout: false,
      qrCode: ""
    };
  },
  created() {
    this.getQRCode();
  },
  methods: {
    getQRCode() {
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

          const timeout = setInterval(() => {
            clearInterval(status);
            clearInterval(timeout);
          }, 300000);
        });
    }
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
  position: relative;

  div {
    position: absolute;
  }
}
</style>
