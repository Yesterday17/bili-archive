<template>
  <div id="app">
    <md-steppers md-linear md-sync-route md-dynamic-height>
      <md-step
        id="step-00"
        class="center"
        md-label="欢迎"
        md-description="欢迎使用"
        :md-editable="true"
        :md-done="this.step > 0"
        :to="'/step-00'"
      >
        <keep-alive>
          <router-view v-if="this.display === 0"/>
        </keep-alive>
      </md-step>
      <md-step
        id="step-01"
        class="center"
        md-label="检查"
        md-description="环境测试"
        :md-editable="true"
        :md-done="this.step > 1"
        :to="'/step-01'"
      >
        <keep-alive>
          <router-view v-if="this.display === 1"/>
        </keep-alive>
      </md-step>
      <md-step
        id="step-02"
        class="center"
        md-label="登录"
        md-description="扫描二维码"
        :md-editable="false"
        :md-done="this.step > 2"
        :to="'/step-02'"
      >
        <keep-alive>
          <router-view v-if="this.display === 2"/>
        </keep-alive>
      </md-step>
      <md-step
        id="step-03"
        class="center"
        md-label="选择收藏夹"
        md-description="需要存档的视频列表"
        :md-editable="false"
        :md-done="this.step > 3"
        :to="'/step-03'"
      >
        <keep-alive>
          <router-view v-if="this.display === 3"/>
        </keep-alive>
      </md-step>
      <md-step
        id="step-04"
        class="center"
        md-label="下载"
        md-description="存档视频"
        :md-editable="false"
        :to="'/step-04'"
      >
        <keep-alive>
          <router-view v-if="this.display === 4"/>
        </keep-alive>
      </md-step>
    </md-steppers>
    <md-button
      id="step-next"
      class="md-fab md-primary"
      @click="next"
      :disabled="!this.canGoNext"
      v-if="this.display !== 4"
    >
      <md-icon>navigate_next</md-icon>
    </md-button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      display: 0,
      step: 0,
      goNextFlag: false
    };
  },
  methods: {
    next() {
      if (localStorage.getItem("step")) {
        // 更新 display 与 step
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
      localStorage.setItem("step", this.step);
    }
  },
  watch: {
    $route() {
      const path = this.$route.path;
      this.step = localStorage.getItem("step")
        ? parseInt(localStorage.getItem("step"))
        : 0;
      this.display = parseInt(path.substr(path.length - 2, 2));
      this.goNextFlag = this.$route.query.next ? this.$route.query.next : false;
    }
  },
  computed: {
    canGoNext() {
      return this.display === 0 || this.goNextFlag;
    }
  },
  created() {
    // 获取或初始化 step, display
    this.step = localStorage.getItem("uid")
      ? 0
      : localStorage.getItem("step")
      ? parseInt(localStorage.getItem("step"))
      : 0;
    this.display = this.step;

    // 更新 step
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

#app,
.md-stepper,
.md-stepper-content,
.md-steppers,
.md-steppers-container,
.full-height {
  height: 100%;
  width: 100%;
}

.md-steppers-wrapper {
  height: 90% !important;
  width: 100%;
}

.center,
.md-stepper-content {
  display: flex;
  display: -webkit-flex;
  align-items: center;
  justify-content: center;
}

.column {
  flex-direction: column;
}

// rotate animation
.rotate {
  animation: rotate 3s linear infinite;
  transform-origin: 50% 50%;
}

@keyframes rotate {
  0% {
    transform: translate(0px, 0px) rotate(0deg);
  }

  100% {
    transform: translate(0px, 0px) rotate(-50deg);
  }
}

// fade transition
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s;
}
.fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */ {
  opacity: 0;
}
</style>
