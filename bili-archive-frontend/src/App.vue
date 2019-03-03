<template>
  <div id="app">
    <md-steppers md-linear md-sync-route md-dynamic-height @md-changed="toStep">
      <md-step
        id="step-00"
        md-label="欢迎"
        md-description="欢迎使用"
        :md-editable="true"
        :md-done="this.step > 0"
        :to="'/step-00'"
      >
        <keep-alive>
          <router-view/>
        </keep-alive>
      </md-step>
      <md-step
        id="step-01"
        md-label="登录"
        md-description="扫描二维码"
        :md-editable="false"
        :md-done="this.step > 1"
        :to="'/step-01'"
      >
        <keep-alive>
          <router-view/>
        </keep-alive>
      </md-step>
      <md-step
        id="step-02"
        md-label="选择收藏夹"
        md-description="需要存档的视频列表"
        :md-editable="false"
        :md-done="this.step > 2"
        :to="'/step-02'"
      >
        <keep-alive>
          <router-view/>
        </keep-alive>
      </md-step>
      <md-step
        id="step-03"
        md-label="下载"
        md-description="存档视频"
        :md-editable="false"
        :to="'/step-03'"
      >
        <keep-alive>
          <router-view/>
        </keep-alive>
      </md-step>
    </md-steppers>
    <md-button id="step-next" class="md-fab md-primary" @click="next">
      <md-icon>navigate_next</md-icon>
    </md-button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      display: 0,
      step: 0
    };
  },
  methods: {
    next() {
      if (localStorage.getItem("display") && localStorage.getItem("step")) {
        // 更新 display 与 step
        this.display = parseInt(localStorage.getItem("display"));
        this.step = parseInt(localStorage.getItem("step"));

        if (this.display !== this.step) {
          // 不相等 说明用户切换到了第一页
          // 此时跳转回当前 step 对应页面
          this.$router.push("step-0" + this.step);
        } else {
          // 当前 step 页面与显示页面相同
          this.$router.push("step-0" + ++this.step);
        }
      }
      // 点击之后显示页面与 step 必然相同
      this.display = this.step;

      // 更新 display 与 step
      localStorage.setItem("display", this.display);
      localStorage.setItem("step", this.step);
    },
    toStep(step) {
      alert(step);
    }
  },
  computed: {
    active() {
      return `step-0${this.step}`;
    }
  },
  created() {
    // 获取或初始化 display 与 step
    this.display = localStorage.getItem("display")
      ? parseInt(localStorage.getItem("display"))
      : 0;
    this.step = localStorage.getItem("step")
      ? parseInt(localStorage.getItem("step"))
      : 0;

    // 更新 display 与 step
    localStorage.setItem("display", this.display);
    localStorage.setItem("step", this.step);

    // 路由跳转
    this.$router.push("step-0" + this.step);
  }
};
</script>

<style lang="scss">
#step-next {
  position: fixed;
  bottom: 20px;
  right: 20px;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s;
}
.fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */ {
  opacity: 0;
}
</style>
