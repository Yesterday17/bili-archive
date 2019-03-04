<template>
  <div id="login" class="center column">
    <div id="qr-box" class="center">
      <transition name="fade" mode="out-in">
        <div v-if="this.timeout" class="center column">
          <div>二维码已超时</div>
          <md-button class="md-dense md-raised md-primary" @click="refresh">点击刷新</md-button>
        </div>
        <div v-else-if="this.qrCode === ''">加载登录二维码中……</div>
        <img v-else :src="this.qrCode" alt="qrCode">
      </transition>
    </div>
    <div>
      <div id="login-icon" :class="this.logined ? '' : 'rotate'">
        <md-icon v-if="this.logined">verified_user'</md-icon>
        <md-icon v-else>sync</md-icon>
      </div>
      <span>{{ this.logined ? '登录成功，请点击下一步。' : '扫描二维码以登录' }}</span>
    </div>
    <div v-if="!logined">
      <p>登录检查每 3 秒进行一次</p>
      <p>二维码超时时间为 300 秒</p>
    </div>
  </div>
</template>

<script>
export default {
  name: "Login",
  data() {
    return {
      timeout: false,
      logined: false,
      qrCode: "",

      iStatus: undefined,
      iTimeout: undefined
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
          this.iStatus = setInterval(() => {
            fetch("//localhost:8080/login-status")
              .then(data => data.json())
              .then(json => {
                if (json.ok) {
                  clearInterval(this.iStatus);
                  clearInterval(this.iTimeout);
                  this.logined = true;
                  this.$router.push({
                    path: this.$router.path,
                    query: {
                      next: true
                    }
                  });
                }
              });
          }, 3000);

          this.iTimeout = setInterval(() => {
            clearInterval(this.iStatus);
            clearInterval(this.iTimeout);
          }, 300000);
        });
    },
    refresh() {
      this.qrCode = "";
      this.timeout = false;
      this.getQRCode();
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
  height: 404.6px;
  border: 1px solid grey;

  p {
    margin-bottom: 5px;
  }
}

#qr-box {
  width: 256px;
  height: 256px;
}

#login-icon {
  display: inline-block;
}
</style>
