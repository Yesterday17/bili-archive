<template>
  <div id="environment">
    <md-table md-card>
      <md-table-toolbar>
        <div class="md-title">环境测试</div>
        <div class="md-subhead">测试目前前后端及服务器的各项状态。</div>
      </md-table-toolbar>
      <md-table-row>
        <md-table-head md-numeric>ID</md-table-head>
        <md-table-head>测试项</md-table-head>
        <md-table-head>测试状态</md-table-head>
        <md-table-head>测试结果</md-table-head>
      </md-table-row>
      <md-table-row>
        <md-table-cell md-numeric>1</md-table-cell>
        <md-table-cell>网络环境</md-table-cell>
        <md-table-cell>Shawna Dubbin</md-table-cell>
        <md-table-cell></md-table-cell>
      </md-table-row>
      <md-table-row>
        <md-table-cell md-numeric>2</md-table-cell>
        <md-table-cell>登录服务初始化</md-table-cell>
        <md-table-cell>Shawna Dubbin</md-table-cell>
        <md-table-cell></md-table-cell>
      </md-table-row>
      <md-table-row>
        <md-table-cell md-numeric>3</md-table-cell>
        <md-table-cell>登陆服务校验</md-table-cell>
        <md-table-cell>Shawna Dubbin</md-table-cell>
        <md-table-cell></md-table-cell>
      </md-table-row>
    </md-table>
  </div>
</template>

<script>
export default {
  mounted() {
    //
  }
};
</script>

<style lang="scss" scoped>
#environment {
  margin-top: 100px;
  width: 768px;
}
</style>
