<template>
  <div>
    <p>欢迎使用
      <code>bili-archive</code> ！在使用之前，请详细阅读以下指南~
    </p>
    <h2 id="前言">前言</h2>
    <p>一切的起因是某天我去看我
      <a href="https://space.bilibili.com/123817257/favlist?fid=283019357">百万</a>收藏夹的时候发现半数以上视频惨遭爆破，后来又去找
      <a href="https://www.biliplus.com/video/av4931739/">M@STERPIECE</a>又发现被爆破。看来今后这样的事情会越来越多的吧，不如缓存下来省事。
    </p>
    <h2 id="如何使用">如何使用</h2>
    <h3 id="步骤-01---登录">步骤 01 - 登录</h3>
    <p>看完这页欢迎页面，你就需要进入下一步的登录。登录
      <strong>只支持二维码登录</strong>。
      <br>（因为这是最简单的实现方式了
      <br>（而且也不用让你们担心收集用户名密码什么的（逃
    </p>
    <p>注意在这里建议选择开通了
      <s>带</s>大会员的账号登录，以下载到最高画质(1080P60)的视频。
    </p>
    <h3 id="步骤-02---选择收藏">步骤 02 - 选择收藏</h3>
    <p>这一步的选择收藏其实是选择需要下载的收藏夹所属用户。比如我现在的运行模式是大号大会员投稿，小号日用
      <s>天天 awsl</s>，画风不同，不在 Timeline 打扰别人。这时候我想要下载我小号的收藏视频，又想要最高画质（大会员），就需要在这一步输入我小号的
      <code>UID</code>。相信了解到这个工具的用户不至于连 UID 怎么查都不知道吧，这里就不再赘述了。
    </p>
    <h3 id="步骤-03---开始下载">步骤 03 - 开始下载</h3>
    <p>这一步就是纯粹的等待了，这时候在控制台会显示详细的视频下载状况（调用了
      <a href="https://github.com/iawia002/annie">annie</a>），网页端也会有详细的信息显示(
      <code>WIP</code>)。
      <br>值得注意的是在这一步下载的过程中，如果发现之前有下载过某视频，会自动跳过该视频的下载（以分 P 为准）。因此，你可以把它当作日常备份工具来用。
      <br>
      <code>// 持久化存储cookies (WIP)</code>
    </p>
    <h2 id="faq">FAQ</h2>
  </div>
</template>

<script>
export default {
  name: "Welcome",
  data() {
    return {};
  },
  mounted() {
    if (!localStorage.getItem("step")) localStorage.setItem("step", "0");
  }
};
</script>

<style lang="scss" scoped>
</style>
