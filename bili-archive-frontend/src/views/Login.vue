<template>
  <div id="login" class="center">
    <div id="qr-box" class="center">
      <transition name="fade" mode="out-in">
        <div v-if="this.timeout">二维码超时</div>
        <div v-else-if="this.qrCode === ''">加载登录二维码中……</div>
        <img v-else :src="this.qrCode" alt="qrCode">
      </transition>
    </div>
    <div>请扫描二维码以登录</div>
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
                  this.$router.push({
                    path: this.$router.path,
                    query: {
                      next: true
                    }
                  });
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
#login {
  margin-top: 100px;
  margin-bottom: 100px;
  padding: 25px;
  width: calc(256px + 2 * 25px);
  border: 1px solid grey;

  flex-direction: column;
}

#qr-box {
  width: 256px;
  height: 256px;
}
</style>
