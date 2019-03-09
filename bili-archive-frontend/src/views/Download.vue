<template>
  <div class="full-height center">
    <md-card v-if="this.start" id="download">
      <md-card-header>
        <div class="md-title">下载中！</div>
        <div class="md-subhead">该页面正在完善，请至控制台查看完整信息。</div>
      </md-card-header>
      <md-card-content>
        <p>目前正在下载：</p>
        <p>{{this.downloadTitle}}</p>
      </md-card-content>

      <md-card-expand>
        <md-card-actions md-alignment="right">
          <md-card-expand-trigger>
            <md-button>详细信息(WIP)</md-button>
          </md-card-expand-trigger>
        </md-card-actions>

        <md-card-expand-content>
          <md-card-content>
            <md-progress-bar
              md-mode="determinate"
              :md-value="this.downloadNow / this.downloadTotal * 100"
            ></md-progress-bar>
          </md-card-content>
        </md-card-expand-content>
      </md-card-expand>
    </md-card>
    <div v-else class="page-container full-height" id="select">
      <md-app class="full-height" md-waterfall md-mode="fixed">
        <md-app-toolbar class="md-primary" id="app-toolbar">
          <span class="md-title">收藏</span>
          <md-button class="md-icon-button" @click="startDownload">
            <md-icon>check</md-icon>
          </md-button>
        </md-app-toolbar>

        <md-app-drawer md-permanent="full">
          <md-toolbar class="md-transparent" md-elevation="0">收藏列表</md-toolbar>

          <md-list class="md-scrollbar">
            <md-list-item v-for="(fav, index) in this.favlist" v-bind:key="index">
              <md-button class="md-list-item-text" @click="changeView(index)">{{fav.name}}</md-button>
            </md-list-item>
          </md-list>
        </md-app-drawer>

        <md-app-content>
          <div
            v-if="this.currentView >=0 && this.favlist.length > 0 && this.favorite.has(this.favlist[this.currentView].fid)"
          >
            <md-card
              class="video-card"
              v-for="(fav, index) in this.favorite.get(this.favlist[this.currentView].fid)"
              v-bind:key="index"
            >
              <md-card-media-cover md-solid>
                <md-card-media md-ratio="16:9">
                  <img :src="getPicture(fav.pic)">
                </md-card-media>

                <md-card-area>
                  <md-card-header>
                    <span>{{fav.title}}</span>
                  </md-card-header>
                </md-card-area>
                <md-checkbox
                  class="videocheck md-primary"
                  v-model="downloadCheck"
                  :value="fav.aid"
                  :disabled="fav.pic === 'http://i0.hdslb.com/bfs/archive/be27fd62c99036dce67efface486fb0a88ffed06.jpg'"
                ></md-checkbox>
              </md-card-media-cover>
            </md-card>
          </div>

          <div v-else style="margin-top: 50px;">
            <md-empty-state
              md-rounded
              md-icon="access_time"
              md-label="请耐心等待收藏加载完成"
              md-description="在左侧出现收藏夹列表后，你可以点击收藏夹名称将视频加入下载列表。"
            ></md-empty-state>
          </div>
        </md-app-content>
      </md-app>
    </div>
  </div>
</template>

<script>
import WSP from "websocket-as-promised";

export default {
  name: "Download",
  data() {
    return {
      uid: 0, // 下载用户的 UID
      start: false, // 下载是否开始

      /**
       * 正在下载视频的信息
       */
      downloadTitle: "",
      downloadStatus: 0,
      downloadMessage: "",
      downloadNow: 0,
      downloadTotal: Infinity,

      favlist: [], // 表示收藏内容的数组 v-for 用
      favlistDetail: new Map(), // fid 到收藏列表信息
      videoToFav: new Map(), // 视频内容到 fid 的反向 Map
      videoDetail: new Map(), // 视频 aid 到视频内容的 Map
      favorite: new Map(), // 收藏内容的详细信息 fid 对应视频数组
      downloadCheck: [], // 需要下载视频 aid 的数组
      currentView: -1 // 当前查看的收藏夹视图 id
    };
  },
  methods: {
    getFavList() {
      return fetch(
        `//${
          window.port
            ? window.location.hostname + ":" + window.port
            : window.location.host
        }/api/favlist?uid=${this.uid}`
      )
        .then(data => data.json())
        .then(json => {
          // TODO: 检测 json.ok
          this.favlist.splice(0, this.favlist.length);
          Array.prototype.push.apply(this.favlist, json.data);
          this.favlist.forEach(fav => {
            this.favlistDetail.set(fav.fid, fav);
            this.favorite.set(fav.fid, []);
          });
        });
    },
    getFavItems(fid, pn) {
      return fetch(
        `//${
          window.port
            ? window.location.hostname + ":" + window.port
            : window.location.host
        }/api/fav?uid=${this.uid}&fid=${fid}&pn=${pn}`
      )
        .then(data => data.json())
        .then(json => {
          // TODO: 检测 json.ok
          Array.prototype.push.apply(this.favorite.get(fid), json.data);
        });
    },
    getVideoPages(aid) {
      return fetch(
        `//${
          window.port
            ? window.location.hostname + ":" + window.port
            : window.location.host
        }/api/pages?aid=${aid}`
      )
        .then(data => data.json())
        .then(json => {
          return json.ok ? json.data : [];
        });
    },
    getPicture(src) {
      return `//${
        window.port
          ? window.location.hostname + ":" + window.port
          : window.location.host
      }/api/pic?url=${encodeURIComponent(src)}`;
    },
    changeView(to) {
      this.currentView = to;

      if (this.downloadCheck.length === 0) {
        // 当下载链接的数组中不存在内容时
        // 判定未初始化
        // 此时初始化 将所有**可用**视频加入列表
        // 由于 3fd1824562cfafbbb0b3c49bdd9812535680280c
        // 不会对已失效视频进行任何处理
        // TODO: 峰回路转 把视频加入稍后再看列表后可以显示失效视频信息
        this.favorite.forEach((value, key) => {
          value.forEach(fav => {
            if (
              fav.pic !==
              "http://i0.hdslb.com/bfs/archive/be27fd62c99036dce67efface486fb0a88ffed06.jpg"
            ) {
              this.downloadCheck.push(fav.aid);
              if (this.videoToFav.has(fav.aid)) {
                this.videoToFav.get(fav.aid).push(key);
              } else {
                this.videoToFav.set(fav.aid, [key]);
              }

              if (!this.videoDetail.has(fav.aid)) {
                this.videoDetail.set(fav.aid, fav);
              }
            }
          });
        });
      }
    },
    downloadVideo(queue, index) {
      if (index === queue.length) return;

      // TODO: 使用通常的 WebSocket 减少依赖
      const wsp = new WSP(
        `ws://${
          window.port
            ? window.location.hostname + ":" + window.port
            : window.location.host
        }/download`,
        {
          packMessage: data => JSON.stringify(data),
          unpackMessage: message => JSON.parse(message)
        }
      );

      wsp.onUnpackedMessage.addListener(data => {
        this.downloadTitle = data.title;
        this.downloadStatus = data.status;
        this.downloadMessage = data.message;
        this.downloadNow = data.data.progress;
        this.downloadTotal = data.data.size;
      });

      wsp.onClose.addListener(() => {
        // TODO: 更详细的日志提示
        console.log(queue[index].title + " 下载完成！");
        this.downloadVideo(queue, index + 1);
      });
      wsp.open().then(() => wsp.sendPacked(queue[index]));
    },
    async startDownload() {
      this.start = true;
      const downloadQueue = [];

      for (let aid of this.downloadCheck) {
        const videos = [];
        const favList = this.videoToFav.get(aid);
        const pages = await this.getVideoPages(aid);

        pages.forEach(page => {
          const data = {
            // 只在 list 的第一个收藏夹中下载 避免重复
            title: this.videoDetail.get(aid).title,
            fav_title: this.favlistDetail.get(favList[0]).name,
            aid: aid.toString(),
            page: {
              page: page.page,
              page_name: page.pagename,
              cid: page.cid.toString()
            }
          };
          videos.push(data);
        });
        // 确保同一视频的多分P连续
        // FIXME: 这个可能没有必要了
        Array.prototype.push.apply(downloadQueue, videos);
      }
      this.downloadVideo(downloadQueue, 0);
    }
  },
  mounted() {
    if (!localStorage.getItem("uid")) {
      alert("未检测到UID，将自动跳转回开始！");
      localStorage.setItem("step", 0);
      this.$router.push("step-00");
    } else {
      this.uid = parseInt(localStorage.getItem("uid"));
      this.getFavList().then(() => {
        this.favlist.forEach(fav => {
          for (let i = 0; i < fav.cur_count / 20; i++) {
            this.getFavItems(fav.fid, i + 1);
          }
        });
      });
    }
  }
};
</script>

<style lang="scss" scoped>
#download {
  width: 100%;
}

#select {
  #app-toolbar {
    flex-direction: row;
    justify-content: space-between;
  }

  .md-button {
    text-transform: initial;
  }

  .video-card {
    display: inline-block;
    width: 250px;
    margin: 10px;

    .videocheck {
      position: absolute;
      padding: 0px;
      margin: 0px;
      top: 10px;
      left: 10px;
      background: white;
    }
  }
}

.md-app {
  border: 1px solid rgba(#000, 0.12);
}

.md-drawer {
  width: 20%;
  max-width: calc(100vw - 125px);
}
</style>
