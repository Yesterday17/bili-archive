<template>
  <div id="download">
    <md-card>
      <md-card-header>
        <div class="md-title">下载中！</div>
        <div class="md-subhead">该页面正在完善，请至控制台查看完整信息。</div>
      </md-card-header>
      <md-card-content>
        <p>目前正在下载：</p>
        <p>{{this.status}}</p>
      </md-card-content>

      <md-card-expand>
        <md-card-actions md-alignment="right">
          <md-card-expand-trigger>
            <md-button>详细信息(WIP)</md-button>
          </md-card-expand-trigger>
        </md-card-actions>

        <md-card-expand-content>
          <md-card-content>
            <md-progress-bar md-mode="indeterminate"></md-progress-bar>
          </md-card-content>
        </md-card-expand-content>
      </md-card-expand>
    </md-card>
  </div>
</template>

<script>
export default {
  name: "Download",
  data() {
    return {
      uid: 0,
      status: ""
    };
  },
  mounted() {
    if (!localStorage.getItem("uid")) {
      alert("未检测到UID，将自动跳转回开始！");
      localStorage.setItem("step", 0);
      this.$router.push("step-00");
    } else {
      this.uid = parseInt(localStorage.getItem("uid"));
      const ws = new WebSocket("ws://localhost:8080/ws");

      ws.addEventListener("message", event => {
        this.status = event.data;
      });
      ws.addEventListener("open", () => {
        ws.send(this.uid);
      });
    }
  }
};
</script>

<style lang="scss" scoped>
#download {
  margin-top: 100px;
  width: 768px;
}
</style>
