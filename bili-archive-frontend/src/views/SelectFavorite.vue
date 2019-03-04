<template>
  <div id="select">
    <md-card md-with-hover>
      <md-ripple>
        <md-card-header>
          <div class="md-title">下载用户？</div>
          <div class="md-subhead">选择是否下载当前登录用户的收藏视频</div>
        </md-card-header>

        <md-card-content>
          <div v-if="currentUser">
            <p>下载来源为当前用户的收藏夹，即使用本账号下载本账号的收藏夹。</p>
            <p>如果你想要下载其他用户（比如自己小号）的收藏夹，请取消勾选下方的“当前用户”。</p>
          </div>
          <div v-else>
            <md-field :class="uidClass">
              <label>UID</label>
              <md-input v-model="uid" type="string" @change="changeMode"></md-input>
              <span class="md-error">{{this.error}}</span>
            </md-field>
            <p>下载来源为上方UID用户的收藏夹，即使用已登录账号下载上方账号的收藏夹。</p>
            <p>上方账号需要开放收藏夹可见权限以正确识别收藏夹内容。</p>
            <p>如果你想要下载当前用户的的收藏夹，请勾选下方的“当前用户”。</p>
          </div>
        </md-card-content>

        <md-card-actions>
          <md-button disabled v-if="this.status === 'updating'">更新信息中……</md-button>
          <md-checkbox v-model="currentUser" class="md-primary" @change="changeMode">当前用户</md-checkbox>
        </md-card-actions>
      </md-ripple>
    </md-card>

    <md-snackbar
      md-position="center"
      :md-duration="Infinity"
      :md-active.sync="showStatus"
      md-persistent
    >
      <span>{{this.status}}</span>
      <md-button class="md-primary" @click="next">确认</md-button>
    </md-snackbar>
  </div>
</template>

<script>
export default {
  name: "SelectFavorite",
  data() {
    return {
      currentUser: false,
      uid: "",
      error: "",

      showStatus: false,
      status: "",

      uname: "",
      avatar: ""
    };
  },
  computed: {
    uidClass() {
      return {
        "md-invalid": this.error !== ""
      };
    }
  },
  methods: {
    clear() {
      this.currentUser = false;
      this.uid = "";
      this.error = "";

      this.showStatus = false;
      this.status = "";

      this.uname = "";
      this.avatar = "";
    },
    getUserData() {
      return fetch("//localhost:8080/info?uid=" + this.uid)
        .then(data => data.json())
        .then(json => {
          if (!json.ok || json.data.mid === 0) {
            alert("用户不存在！");
            this.clear();
          } else {
            this.uname = json.data.name;
            this.avatar = json.data.face;
          }
        });
    },
    changeMode() {
      // 修改后切换到更新状态
      this.status = "updating";

      if (this.currentUser) {
        fetch("//localhost:8080/current-user")
          .then(data => data.json())
          .then(json => {
            if (!json.ok) {
              this.currentUser = false;
              this.status = "";
              alert(json.message);
            } else {
              this.uid = json.uid;
              // 获取该 UID 的详细信息
              this.getUserData().then(() => {
                this.status = `用户名: ${this.uname}, UID: ${this.uid}`;
                this.showStatus = true;
              });
            }
          });
      } else if (!this.uid.match(/^\d*$/)) {
        this.error = "UID非法！应为纯数字！（匹配/^\\d*$/）";
        this.status = "";
      } else {
        // 判断到这里没有错误了
        this.error = "";
        this.getUserData().then(() => {
          this.status = `用户名: ${this.uname}, UID: ${this.uid}`;
          this.showStatus = true;
        });
      }
    },
    next() {
      this.showStatus = false;
      localStorage.setItem("uid", this.uid);
      this.$router.push({
        path: this.$router.path,
        query: {
          next: true
        }
      });
    }
  }
};
</script>

<style lang="scss" scoped>
#select {
  margin-top: 100px;
  margin-bottom: 100px;
}
</style>
