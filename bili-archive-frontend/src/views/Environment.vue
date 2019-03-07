<template>
  <div id="environment">
    <md-table md-card>
      <md-table-toolbar>
        <div class="md-title">环境测试</div>
        <div class="md-subhead">测试前后端及服务器的各项状态。</div>
      </md-table-toolbar>
      <md-table-row>
        <md-table-head md-numeric>ID</md-table-head>
        <md-table-head>测试项</md-table-head>
        <md-table-head>测试状态</md-table-head>
        <md-table-head>测试结果</md-table-head>
      </md-table-row>
      <md-table-row>
        <md-table-cell md-numeric>1</md-table-cell>
        <md-table-cell>主站网络环境</md-table-cell>
        <md-table-cell>{{this.pending ? '测试中' : '测试完成'}}</md-table-cell>
        <md-table-cell>{{check(this.main)}}</md-table-cell>
      </md-table-row>
      <md-table-row>
        <md-table-cell md-numeric>2</md-table-cell>
        <md-table-cell>登录服务初始化</md-table-cell>
        <md-table-cell>{{this.pending ? '测试中' : '测试完成'}}</md-table-cell>
        <md-table-cell>{{check(this.login_qr)}}</md-table-cell>
      </md-table-row>
      <md-table-row>
        <md-table-cell md-numeric>3</md-table-cell>
        <md-table-cell>登陆服务校验</md-table-cell>
        <md-table-cell>{{this.pending ? '测试中' : '测试完成'}}</md-table-cell>
        <md-table-cell>{{check(this.login_info)}}</md-table-cell>
      </md-table-row>
      <md-table-row>
        <md-table-cell md-numeric>4</md-table-cell>
        <md-table-cell>用户空间</md-table-cell>
        <md-table-cell>{{this.pending ? '测试中' : '测试完成'}}</md-table-cell>
        <md-table-cell>{{check(this.space)}}</md-table-cell>
      </md-table-row>
      <md-table-row>
        <md-table-cell md-numeric>5</md-table-cell>
        <md-table-cell>视频分P</md-table-cell>
        <md-table-cell>{{this.pending ? '测试中' : '测试完成'}}</md-table-cell>
        <md-table-cell>{{check(this.video_page)}}</md-table-cell>
      </md-table-row>
      <md-table-row>
        <md-table-cell md-numeric>6</md-table-cell>
        <md-table-cell>用户收藏列表</md-table-cell>
        <md-table-cell>{{this.pending ? '测试中' : '测试完成'}}</md-table-cell>
        <md-table-cell>{{check(this.favorite_list)}}</md-table-cell>
      </md-table-row>
      <md-table-row>
        <md-table-cell md-numeric>7</md-table-cell>
        <md-table-cell>用户收藏内容</md-table-cell>
        <md-table-cell>{{this.pending ? '测试中' : '测试完成'}}</md-table-cell>
        <md-table-cell>{{check(this.favorite_list_item)}}</md-table-cell>
      </md-table-row>
    </md-table>
    <div v-if="!this.pending" class="center">
      <p>{{this.all_ok ? '准备就绪，请进入下一步。' : '您存在未通过的校验项，请确认网络环境！'}}</p>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      pending: true,
      main: false,
      login_qr: false,
      login_info: false,
      space: false,
      video_page: false,
      favorite_list: false,
      favorite_list_item: false,
      all_ok: true
    };
  },
  methods: {
    check(status) {
      if (typeof status !== "boolean") return "变量类型不为 boolean!";
      if (this.pending) return "测试中";
      return status ? "通过" : "未通过";
    }
  },
  mounted() {
    fetch(
      `//${
        window.port
          ? window.location.hostname + ":" + window.port
          : window.location.host
      }/api/test`
    )
      .then(data => data.json())
      .then(json => {
        this.pending = false;
        this.main = json.main;
        this.login_qr = json.login_qr;
        this.login_info = json.login_info;
        this.space = json.space;
        this.video_page = json.video_page;
        this.favorite_list = json.favorite_list;
        this.favorite_list_item = json.favorite_list_item;

        for (const key in json) {
          if (json.hasOwnProperty(key)) {
            this.all_ok |= json[key];
          }
        }

        if (this.all_ok) {
          this.$router.push({
            path: this.$router.path,
            query: {
              next: true
            }
          });
        }
      });
  }
};
</script>

<style lang="scss" scoped>
#environment {
  width: 100%;
}
</style>
