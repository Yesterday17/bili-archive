<template>
  <div id="app">
    <md-steppers md-sync-route md-dynamic-height md-linear>
      <md-step
        id="step-welcome"
        md-label="欢迎"
        md-description="欢迎使用"
        :md-editable="true"
        to="/step-00"
      >
        <router-view/>
      </md-step>
      <md-step
        id="step-login"
        md-label="登录"
        md-description="扫描二维码"
        :md-editable="false"
        to="/step-01"
      >
        <router-view/>
      </md-step>
      <md-step
        id="step-select-favorite"
        md-label="选择收藏夹"
        md-description="需要存档的视频列表"
        :md-editable="false"
        to="/step-02"
      >
        <router-view/>
      </md-step>
      <md-step
        id="step-archive"
        md-label="下载"
        md-description="存档视频"
        :md-editable="false"
        to="/step-03"
      >
        <router-view/>
      </md-step>
    </md-steppers>
  </div>
</template>

<script>
export default {
  data() {
    return {
      step: 0
    };
  },
  created() {
    this.step = localStorage.getItem("step") ? localStorage.getItem("step") : 0;
    this.$router.push("step-0" + this.step);
  }
};
</script>

<style lang="scss">
*:visited {
  color: inherit;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s;
}
.fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */ {
  opacity: 0;
}
</style>
